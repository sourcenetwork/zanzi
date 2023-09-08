package relation_graph

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

func newGoalTreeBuilder(evaluator evaluator, logger types.Logger) goalTreeBuilder {
	return goalTreeBuilder{
		evaluator: evaluator,
		logger:    logger,
	}
}

type goalTreeBuilder struct {
	evaluator evaluator
	logger    types.Logger
}

func (b *goalTreeBuilder) getEvaluator() evaluator {
	return b.evaluator
}

// Build a GoalTree for node.
// Builder evaluates the RelationNode Relation (if it exists) and the resulting GoalTree
// contains the dynamically evaluated RelationNodes
func (b *goalTreeBuilder) Build(ctx context.Context, pol *domain.Policy, node *domain.RelationNode) (GoalTree, error) {
	var err error
	var tree GoalTree

	switch nodeType := node.Node.(type) {
	case *domain.RelationNode_EntitySet:
		tree, err = b.handleEntitySet(ctx, pol, nodeType.EntitySet)
	case *domain.RelationNode_Entity, *domain.RelationNode_Wildcard:
		tree = b.handleTerminalNodes(node)
	default:
		err = fmt.Errorf("RelationNode %v: %w", nodeType, domain.ErrInvalidVariant)
	}

	if err != nil {
		return nil, fmt.Errorf("build GoalTree failed: policy %v node %v: %w", pol.Id, node, err)
	}

	b.logger.Debugf("goal tree for node %v: tree: %#v", node, tree)
	return tree, nil
}

// handleTerminalNodes handle the base case for Goal Trees,
// meaning leaves / terminal nodes where there is no further possible expansion.
func (b *goalTreeBuilder) handleTerminalNodes(node *domain.RelationNode) GoalTree {
	return &PathNode{
		Result:       SearchResult_UNKNOWN,
		RelationNode: node,
		Path:         nil,
	}
}

// handleEntitySet builds the GoalTree for an EntitySetNode
func (b *goalTreeBuilder) handleEntitySet(ctx context.Context, pol *domain.Policy, node *domain.EntitySetNode) (GoalTree, error) {
	table := policy.NewPolicyLookUpTable(pol)

	relation := table.GetRelation(node.Object.Resource, node.Relation)
	if relation == nil {
		return nil, policy.ErrRelationNotFound
	}

	expressionTree, err := policy.GetExpressionTree(relation)
	if err != nil {
		return nil, err
	}

	tree, err := b.buildGoalTree(ctx, pol.Id, expressionTree, node)
	if err != nil {
		return nil, err
	}

	return tree, nil
}

func (b *goalTreeBuilder) buildGoalTree(ctx context.Context, policyId string, exprTree *domain.RelationExpressionTree, node *domain.EntitySetNode) (GoalTree, error) {
	var goalTree GoalTree
	var err error

	switch exprTreeNode := exprTree.Node.(type) {
	case *domain.RelationExpressionTree_OpNode:
		goalTree, err = b.buildGoalForOpNode(ctx, policyId, exprTreeNode.OpNode, node)
	case *domain.RelationExpressionTree_Rule:
		goalTree, err = b.buildGoalForRule(ctx, policyId, exprTreeNode.Rule, node)
	default:
		err = fmt.Errorf("relation expression tree node %v: %w", exprTreeNode, domain.ErrInvalidVariant)
	}

	return goalTree, err
}

func (b *goalTreeBuilder) buildGoalForRule(ctx context.Context, policyId string, rule *domain.Rule, node *domain.EntitySetNode) (GoalTree, error) {
	evaluator := b.getEvaluator()

	sucessors, err := evaluator.Evaluate(ctx, policyId, rule, node)
	if err != nil {
		return nil, err
	}

	reason := reasonBuilder(rule)
	paths := utils.MapSlice(sucessors, func(node *domain.RelationNode) GoalTree {
		return b.relationNodeToPathNode(node, reason)
	})

	// Wrap all paths in an ORNode as the result of evaluating a rule
	// results in a Set of nodes where find the goal in any of them
	// is enough to fulfill the goal.
	return &ORNode{
		Paths: paths,
	}, nil
}

func (b *goalTreeBuilder) relationNodeToPathNode(node *domain.RelationNode, reason string) GoalTree {
	return &PathNode{
		RelationNode: node,
		Path:         nil,
		Parent:       nil,
		Result:       SearchResult_UNKNOWN,
		Reason:       reason,
	}
}

func (b *goalTreeBuilder) buildGoalForOpNode(ctx context.Context, polId string, opNode *domain.OpNode, node *domain.EntitySetNode) (GoalTree, error) {
	var tree GoalTree

	left, err := b.buildGoalTree(ctx, polId, opNode.Left, node)
	if err != nil {
		return nil, err
	}

	right, err := b.buildGoalTree(ctx, polId, opNode.Right, node)
	if err != nil {
		return nil, err
	}

	switch opNode.Operator {
	case domain.Operator_UNION:
		// A set Union is equivalent to
		// an OR Node since any set which contains
		// the goal would fulfill the Goal
		tree = &ORNode{
			Paths: []GoalTree{
				left,
				right,
			},
			Parent: nil,
			Result: SearchResult_UNKNOWN,
		}
	case domain.Operator_INTERSECTION:
		// A set Intersection is equivalent to
		// an AND Node since the goal is only
		// fulfilled if the left and right paths
		// contains the Goal.
		tree = &ANDNode{
			Paths: []GoalTree{
				left,
				right,
			},
			Parent: nil,
			Result: SearchResult_UNKNOWN,
		}
	case domain.Operator_DIFFERENCE:
		tree = &DifferenceNode{
			Left:   left,
			Right:  right,
			Parent: nil,
			Result: SearchResult_UNKNOWN,
		}
	default:
		return nil, fmt.Errorf("operator %v: %w", opNode.Operator, domain.ErrInvalidVariant)
	}

	return tree, nil
}

/*
type goalTreeFolder struct {
}

func (f *goalTreeFolder) Fold(tree GoalTree) []*domain.RelationNode { }

func (f *goalTreeFolder) fold(tree GoalTree) map[string]*domain.RelationNode {
    switch node := tree.(type) {
    case *ANDNode:
    case *ORNode:
    case *DifferenceNode:
    case *PathNode:
    case nil:
    default:
    }
}

func (f *goalTreeFolder) foldORNode(tree *ORNode) map[string]*domain.RelationNode {
    nodes := make(map[string]*RelationNode)

    for _, path := range tree.Paths{
        f.fold(path, nodes)
    }
}

func (f *goalTreeFolder) foldANDNode(tree GoalTree) []*domain.RelationNode { }

func (f *goalTreeFolder) foldDifferenceNode(tree GoalTree) []*domain.RelationNode { }

func (f *goalTreeFolder) foldPathNode(tree GoalTree) []*domain.RelationNode {

}
*/

func reasonBuilder(rule *domain.Rule) string {
	var reason string
	switch ruleType := rule.Rule.(type) {
	case *domain.Rule_This:
		reason = "Rule This"
	case *domain.Rule_Ttu:
		reason = fmt.Sprintf("Rule Tuple To Userset: %v %v", ruleType.Ttu.TuplesetRelation, ruleType.Ttu.ComputedUsersetRelation)
	case *domain.Rule_Cu:
		reason = fmt.Sprintf("Rule Computed Userset: %v", ruleType.Cu.TargetRelation)
	}
	return reason
}
