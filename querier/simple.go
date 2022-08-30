
package querier

import (
	"context"
        "fmt"

	"github.com/sourcenetwork/source-zanzibar/tree"
	"github.com/sourcenetwork/source-zanzibar/repository"
	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/utils"
)


func Expand(ctx context.Context, userset model.Userset) (*tree.UsersetNode, error) {
    exp := newExpander()

    expTree, err := exp.getExpTree(ctx, userset.Namespace, userset.Relation)
    if err != nil {
        // wrap
        return nil, err
    }

    tree, err := exp.expandTree(ctx, expTree, userset)
    if err != nil {
        return nil, err
    }
    
    return tree, nil
}

type expander struct {
    trail map[key]struct{}
    trees map[key]tree.Node
    rewriteTrees map[key]*model.RewriteNode
}

func newExpander() expander {
    return expander {
        trail: make(map[key]struct{}),
        trees: make(map[key]tree.Node),
        rewriteTrees: make(map[key]*model.RewriteNode),
    }
}

// Expand an Userset Rewrite Expression Tree
// Uses local cache if an expand call has already been evaluated
// keeps a trail through the depth search to avoid cyclic expands
func (e *expander) expandTree(ctx context.Context, root *model.RewriteNode, uset model.Userset) (*tree.UsersetNode, error) {
    // expandTree is the recusion entrypoint for Expand subcalls
    // Therefore it's a centralized point to check whether ctx is still valid before continuing
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
    }

    //key = toKey(uset)

    // handles a backtrail expand call such as: A -> B -> A
    // return a non-expanded leaf with the uset in it
    //_, ok := e.trail[key]
    //if ok {
        //node := &tree.UsersetNode {
            //Userset: uset,
        //}
        //return node, nil
    //}

    // check whether expand was already evaluated, return if it was
    //tree, ok := e.trees[key]
    //if ok {
        //return tree, nil
    //}

    //e.trail[key] = struct{}{}

    exprNode, err := e.expandExprNode(ctx, root, uset)
    if err != nil {
        // TODO wrap
        return nil, err
    }

    node := &tree.UsersetNode {
        Userset: uset,
        Child: exprNode,
    }

    // remove from trail and cache result
    //delete(e.trail, key)
    //e.trees[key] = node

    return node, nil
}

func (e *expander) expandExprNode(ctx context.Context, root *model.RewriteNode, uset model.Userset) (tree.ExpressionNode, error) {
    switch n := root.Node.(type) {
    case *model.RewriteNode_Opnode:
        opnode := n.Opnode
        return e.expandOpNode(ctx, opnode, uset)
    case *model.RewriteNode_Leaf:
        leaf := n.Leaf
        return e.expandRuleNode(ctx, leaf, uset)
    default:
        err := fmt.Errorf("Rewrite Node type unknown: %#v", root)
        return nil, err
    }
}


// Expand OpNode by expand its left and right children
func (e *expander) expandOpNode(ctx context.Context, root *model.OpNode, uset model.Userset) (*tree.OpNode, error) {
    left, err := e.expandExprNode(ctx, root.Left, uset)
    if err != nil {
        return nil, err
    }

    right, err := e.expandExprNode(ctx, root.Right, uset)
    if err != nil {
        return nil, err
    }

    node := &tree.OpNode {
        Left: left,
        Right: right,
        JoinOp: root.Op,
    }
    return node, nil
}

func (e *expander) expandRuleNode(ctx context.Context, root *model.Leaf, uset model.Userset) (*tree.RuleNode, error) {

    var neighbors []model.Userset
    var rule tree.Rule
    var err error

    switch r := root.Rule.GetRule().(type) {

    case *model.Rule_This:
        neighbors, err = produceThis(ctx, uset)
        rule = tree.Rule {
            Type: tree.RuleType_THIS,
        }

    case *model.Rule_TupleToUserset:
        ttu := r.TupleToUserset
        neighbors, err = produceTTU(ctx, uset, ttu.TuplesetRelation, ttu.ComputedUsersetRelation)
        rule = tree.Rule {
            Type: tree.RuleType_TTU,
            Args: map[string]string{
                "TuplesetRelation": ttu.TuplesetRelation,
                "ComputedUsersetRelation": ttu.ComputedUsersetRelation,
            },
        }

    case *model.Rule_ComputedUserset:
        cu := r.ComputedUserset
        neighbors = produceCU(ctx, uset, cu.Relation)
        rule = tree.Rule {
            Type: tree.RuleType_CU,
            Args: map[string]string{
                "Relation": cu.Relation,
            },
        }

    default:
        err = fmt.Errorf("Unknown rule type: %#v", r)
    }

    if err != nil {
        return nil, err
    }

    return e.expandRule(ctx, uset, neighbors, rule)
}

func (e *expander) getExpTree(ctx context.Context, namespace, relation string) (*model.RewriteNode, error) {
    repo := utils.GetNamespaceRepo(ctx)

    rel, err := repo.GetRelation(namespace, relation)
    if err != nil {
        // wrap err?
        return nil, err
    }

    return rel.Rewrite.ExpressionTree, nil
}


func (e *expander) expandRule(ctx context.Context, uset model.Userset, neighbors []model.Userset, rule tree.Rule) (*tree.RuleNode, error) {
    children := make([]*tree.UsersetNode, 0, len(neighbors))
    for _, neigh := range neighbors {
        expTree, err := e.getExpTree(ctx, neigh.Namespace, neigh.Relation)
        if err != nil {
            return nil, err
        }

        // bfs call
        child, err := e.expandTree(ctx, expTree, neigh)
        if err != nil {
            return nil, err
        }

        children = append(children, child)
    }

    node := &tree.RuleNode {
        Rule: rule,
        Children: children,
    }
    return node, nil
}

// Receive releation and object, builds userset and performs a bfs search on graph
// for all reachable nodes.
func produceThis(ctx context.Context, uset model.Userset) ([]model.Userset, error) {
        repo := utils.GetTupleRepo(ctx)

        tuples, err := repo.GetRelatedUsersets(uset)

        if _, ok := err.(*repository.EntityNotFound); ok {
            return nil, nil
        }

        if err != nil {
            // wrap err
            return nil, err
        }

        usets := utils.MapSlice(tuples, tupleToUserset)

        return usets, nil
}

// return a joinable node because the result might be a leaf or a opnode
func produceCU(ctx context.Context, uset model.Userset, relation string) []model.Userset {
    return []model.Userset{
        model.Userset {
            Namespace: uset.Namespace,
            ObjectId: uset.ObjectId,
            Relation: relation,
        },
    }
}

func produceTTU(ctx context.Context, uset model.Userset, tsetRel string, cuRel string) ([]model.Userset, error) {
    repo := utils.GetTupleRepo(ctx)

    tuplesetFilter := model.Userset{
        Namespace: uset.Namespace,
        ObjectId: uset.ObjectId,
        Relation: tsetRel,
    }

    records, err := repo.GetRelatedUsersets(tuplesetFilter)
    if _, ok := err.(*repository.EntityNotFound); ok {
        return nil, nil
    }
    if err != nil {
        // wrap
        return nil, err
    }

    usets := utils.MapSlice(records, tupleToUserset)
    
    return usets, nil
}

func tupleToUserset(tuple model.TupleRecord) model.Userset {
    return model.Userset{
        Namespace: tuple.Tuple.User.Userset.Namespace,
        ObjectId: tuple.Tuple.User.Userset.ObjectId,
        Relation: tuple.Tuple.User.Userset.Relation,
    }
}

// Key represents a stripped down version of Userset,
// such that it can be used as a map key
type key struct  {
    Namespace string
    ObjectId string
    Relation string
}

func toKey(userset model.Userset) key {
    return key {
        Namespace: userset.Namespace,
        ObjectId: userset.ObjectId,
        Relation: userset.Relation,
    }
}
