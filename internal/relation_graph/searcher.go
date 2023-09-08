package relation_graph

import (
    "context"
    "fmt"

    "github.com/sourcenetwork/zanzi/pkg/domain"
    "github.com/sourcenetwork/zanzi/pkg/types"
    "github.com/sourcenetwork/zanzi/internal/utils"
)

func NewSearcher(builder goalTreeBuilder, logger types.Logger) Searcher{
    return Searcher{
        builder: builder,
        logger: logger,
        seenNodes: make(map[string]struct{}),
    }
}

type Searcher struct {
    builder goalTreeBuilder
    logger types.Logger
    seenNodes map[string]struct{}
    resolvedTrees map[string]GoalTree
}

// Search receives an Origin and a Goal and searches for a Path in the RelationGraph between
// the two Nodes.
// Returns a GoalTree of the search, a bool indicating whether it was found or an error
func (s *Searcher) Search(ctx context.Context, policy *domain.Policy, origin *domain.RelationNode, goal *Goal) (GoalTree, error) {
    s.logger.Infof("Searching for %v in policy %v", origin, policy.Id)

    if goal.Target != nil {
        _, isWildcard := goal.Target.Node.(*domain.RelationNode_Wildcard)
        if isWildcard {
            return nil, fmt.Errorf("search failed: goal %v: %w", goal, ErrWildcardGoal)
        }
    }

    path := &PathNode{
        Parent: nil,
        RelationNode: origin,
        Result: SearchResult_UNKNOWN,
        Path: nil,
        Reason: "Request",
    }

    tree, err := s.searchPath(ctx, policy, path, goal)
    if err != nil {
        return nil, fmt.Errorf("search failed: %w", err)
    }

    return tree, nil
}

func (s *Searcher) search(ctx context.Context, policy *domain.Policy, tree GoalTree, goal *Goal) (GoalTree, error) {
    s.logger.Debugf("searching tree: %#v ", tree)

    var newTree GoalTree
    var err error

    switch treeType := tree.(type) {
    case *ORNode:
        newTree, err = s.searchOR(ctx, policy, treeType, goal)
    case *ANDNode:
        newTree, err = s.searchAND(ctx, policy, treeType, goal)
    case *DifferenceNode:
        newTree, err = s.searchDifference(ctx, policy, treeType, goal)
    case *PathNode:
        newTree, err = s.searchPath(ctx, policy, treeType, goal)
    default:
        err = fmt.Errorf("RewriteGoalTree %v: %w", treeType, domain.ErrInvalidVariant)
    }

    if err != nil {
        s.logger.Debugf("%v", err)
        return nil, fmt.Errorf("search failed: %w", err)
    }

    return newTree, nil
}


func (s *Searcher) searchAND(ctx context.Context, policy *domain.Policy, tree *ANDNode, goal *Goal) (GoalTree, error) { 
    paths := make([]GoalTree, len(tree.Paths))
    copy(paths, tree.Paths)
    foundAll := true

    for i, path := range tree.Paths{
        newPath, err := s.search(ctx, policy, path, goal)
        paths[i] = newPath
        if err != nil {
            return nil, err
        }

        // ANDNode requires that all Paths to the Goal are found
        if newPath.GetResult() == SearchResult_FAILURE {
            foundAll = false
            s.logger.Debugf("AND Node terminated: node %v did not reach path", newPath)
            break
        }
    }

    result := utils.Conditional(foundAll, SearchResult_SUCCESS, SearchResult_FAILURE)
    and := &ANDNode{
        Paths: paths,
        Result: result,
        Parent: nil,
    }
    for _, node := range paths {
        node.SetParent(and)
    }
    return and, nil
}

func (s *Searcher) searchOR(ctx context.Context, policy *domain.Policy, tree *ORNode, goal *Goal) (GoalTree, error) { 
    paths := make([]GoalTree, len(tree.Paths))
    for i, p := range tree.Paths  {
        paths[i] = p
    }
    foundAny := false

    for i, path := range tree.Paths{
        newPath, err := s.search(ctx, policy, path, goal)
        paths[i] = newPath
        if err != nil {
            return nil, err
        }

        // ORNode requires that any Path reaches the Goal
        if newPath.GetResult() == SearchResult_SUCCESS {
            foundAny = true
            s.logger.Debugf("OR Node terminated: node %v reached goal", newPath)
            break
        }
    }

    result := utils.Conditional(foundAny, SearchResult_SUCCESS, SearchResult_FAILURE)
    or := &ORNode{
        Paths: paths,
        Result: result,
        Parent: nil,
    }

    for _, node := range paths {
        node.SetParent(or)
    }

    return or, nil
}

func (s *Searcher) searchDifference(ctx context.Context, policy *domain.Policy, tree *DifferenceNode, goal *Goal) (GoalTree, error) {
    leftTree, err := s.search(ctx, policy, tree.Left, goal)
    if err != nil {
        return nil, err
    }

    rightTree, err := s.search(ctx, policy, tree.Right, goal)
    if err != nil {
        return nil, err
    }

    found := leftTree.GetResult() == SearchResult_SUCCESS && rightTree.GetResult() == SearchResult_FAILURE
    result := utils.Conditional(found, SearchResult_SUCCESS, SearchResult_FAILURE)

    diff := &DifferenceNode{
        Left: leftTree,
        Right: rightTree,
        Result: result,
        Parent: nil,
    }
    leftTree.SetParent(diff)
    rightTree.SetParent(diff)
    return diff, nil

}

func (s *Searcher) searchPath(ctx context.Context, policy *domain.Policy, node *PathNode, goal *Goal) (GoalTree, error) {
    pathNode := &PathNode{
        Parent: nil,
        RelationNode: node.RelationNode,
        Result: SearchResult_UNKNOWN,
        Path: nil,
        Reason: node.Reason,
    }

    // add trail
    // use cached value
    nodeId := node.RelationNode.Id()
    if _, ok := s.seenNodes[nodeId]; ok {
        s.logger.Debugf("duplicated node - terminating brach: %v", node.RelationNode)
        return pathNode, nil
    } else {
        s.seenNodes[nodeId] = struct{}{}
    }

    spec := GoalFoundSpec{}
    found := spec.Found(node, goal)
    terminal := node.RelationNode.IsTerminalNode()

    if found || terminal {
        s.logger.Debugf("Path terminated: Node %v Goal Reached %v", node.RelationNode, found)
        pathNode.Result = utils.Conditional(found, SearchResult_SUCCESS, SearchResult_FAILURE)
        return pathNode, nil
    }

    goalTree, err := s.builder.Build(ctx, policy, pathNode.RelationNode)
    if err != nil {
        return nil, err
    }

    goalTree, err = s.search(ctx, policy, goalTree, goal)
    if err != nil {
        return nil, err
    }

    goalTree.SetParent(pathNode)
    pathNode.Path = goalTree
    pathNode.Result = goalTree.GetResult()

    return pathNode, nil
}


type GoalFoundSpec struct{}

func (s *GoalFoundSpec) Found(pathNode *PathNode, goal *Goal) bool {
    if goal.Target == nil {
        // Target is nil during Expand calls, such that the goal is never reached
        // and the Searcher walks through the entire Graph
        return false
    }

    switch node := goal.Target.Node.(type) {
    case *domain.RelationNode_EntitySet:
        return s.isEntitySetGoalFound(pathNode, node.EntitySet)
    case *domain.RelationNode_Entity:
        return s.isEntityGoalFound(pathNode, node.Entity)
    case *domain.RelationNode_Wildcard:
        // Note this is invalid condition and should probably cause a panic
        return false
    default:
        return false
    }
}

func (s *GoalFoundSpec) isEntityGoalFound(pathNode *PathNode, goal *domain.EntityNode) bool {
    switch node := pathNode.RelationNode.Node.(type){
    case *domain.RelationNode_Entity:
        goalEntity := goal.Object
        pathEntity := node.Entity.Object
        return goalEntity.Resource == pathEntity.Resource && goalEntity.Id == pathEntity.Id
    case *domain.RelationNode_EntitySet:
        return false
    case *domain.RelationNode_Wildcard:
        return node.Wildcard.Resource == goal.Object.Resource
    default:
        return false
    }
}

func (s *GoalFoundSpec) isEntitySetGoalFound(pathNode *PathNode, goal *domain.EntitySetNode) bool {
    switch node := pathNode.RelationNode.Node.(type){
    case *domain.RelationNode_EntitySet:
        return node.EntitySet == goal
    case *domain.RelationNode_Entity, *domain.RelationNode_Wildcard:
        return false
    default:
        return false
    }
}
