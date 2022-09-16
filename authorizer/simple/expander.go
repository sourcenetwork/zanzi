package simple

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/source-zanzibar/authorizer"
	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/repository"
	"github.com/sourcenetwork/source-zanzibar/tree"
	"github.com/sourcenetwork/source-zanzibar/utils"
)

var _ authorizer.Expander = (*expander)(nil)

// Expander implements the authorizer Expander interface
type expander struct {
	trail map[model.KeyableUset]struct{}
        tupleRepo repository.TupleRepository
        nsRepo repository.NamespaceRepository
}

func (e *expander) Expand(ctx context.Context, userset model.Userset) (tree.UsersetNode, error) {
	e.trail = make(map[model.KeyableUset]struct{})

	expTree, err := e.getExpTree(ctx, userset.Namespace, userset.Relation)
	if err != nil {
                err = fmt.Errorf("Expand failed for %v: %w", userset, err)
		return tree.UsersetNode{}, err
	}

	node, err := e.expandTree(ctx, expTree, userset)
	if err != nil {
                err = fmt.Errorf("Expand failed for %v: %w", userset, err)
		return tree.UsersetNode{}, err
	}

	return *node, nil
}

// Expand an Userset Rewrite Expression Tree
// keeps a trail through the depth search to avoid cyclic expands
func (e *expander) expandTree(ctx context.Context, root *model.RewriteNode, uset model.Userset) (*tree.UsersetNode, error) {
	// expandTree is the recusion entrypoint for Expand subcalls
	// Therefore it's a centralized point to check whether ctx is still valid before continuing
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

        node := &tree.UsersetNode{
            Userset: uset,
        }

	// handles a backtrail expand call such as: A -> B -> A
	// return a non-expanded leaf with the uset in it
	key := uset.ToKey()
	_, ok := e.trail[key]
	if ok {
            return node, nil
	}

        // Handle an empty expression tree
        // Should only happen when uset.Relation is the empty relation constant
        if root == nil {
            return node, nil
        }

	e.trail[key] = struct{}{}
	defer delete(e.trail, key)


	exprNode, err := e.expandExprNode(ctx, root, uset)
	if err != nil {
                err = fmt.Errorf("expandTree failed for userset: %v: %v", uset, err)
		return nil, err
	}

        node.Child = exprNode
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
            panic(err)
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

	node := &tree.OpNode{
		Left:   left,
		Right:  right,
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
		neighbors, err = e.produceThis(ctx, uset)
		rule = tree.Rule{
			Type: tree.RuleType_THIS,
		}

	case *model.Rule_TupleToUserset:
		ttu := r.TupleToUserset
		neighbors, err = e.produceTTU(ctx, uset, ttu.TuplesetRelation, ttu.ComputedUsersetRelation)
		rule = tree.Rule{
			Type: tree.RuleType_TTU,
			Args: map[string]string{
				"TuplesetRelation":        ttu.TuplesetRelation,
				"ComputedUsersetRelation": ttu.ComputedUsersetRelation,
			},
		}

	case *model.Rule_ComputedUserset:
		cu := r.ComputedUserset
		neighbors = e.produceCU(ctx, uset, cu.Relation)
		rule = tree.Rule{
			Type: tree.RuleType_CU,
			Args: map[string]string{
				"Relation": cu.Relation,
			},
		}

	default:
		err = fmt.Errorf("Unknown rule type: %#v", r)
                panic(err)
	}

	if err != nil {
		return nil, err
	}

	return e.expandRule(ctx, uset, neighbors, rule)
}

func (e *expander) getExpTree(ctx context.Context, namespace, relation string) (*model.RewriteNode, error) {
        // Empty relation shouldn't be explcitly defined per namespace
        // therefore we skip it
        if relation == model.EMPTY_REL {
            return nil, nil
        }

	rel, err := e.nsRepo.GetRelation(namespace, relation)
	if err != nil {
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

	node := &tree.RuleNode{
		Rule:     rule,
		Children: children,
	}
	return node, nil
}

// Receive releation and object, builds userset and performs a bfs search on graph
// for all reachable nodes.
func (e *expander) produceThis(ctx context.Context, uset model.Userset) ([]model.Userset, error) {
	tuples, err := e.tupleRepo.GetRelatedUsersets(uset)

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


func (e *expander) produceCU(ctx context.Context, uset model.Userset, relation string) []model.Userset {
	return []model.Userset{
		model.Userset{
			Namespace: uset.Namespace,
			ObjectId:  uset.ObjectId,
			Relation:  relation,
		},
	}
}

// Return all Nodes reachable from uset by following a TTU.
// TTU rule is defined by tsetRel and cuRel
func (e *expander) produceTTU(ctx context.Context, uset model.Userset, tsetRel string, cuRel string) ([]model.Userset, error) {
	tuplesetFilter := model.Userset{
		Namespace: uset.Namespace,
		ObjectId:  uset.ObjectId,
		Relation:  tsetRel,
	}

	records, err := e.tupleRepo.GetRelatedUsersets(tuplesetFilter)
	if _, ok := err.(*repository.EntityNotFound); ok {
            // An empty result set from a TTU call cannot be considered
            // an error, as there is no guarantee the target will exist
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	usets := utils.MapSlice(records, tupleToUserset)
	for i := range usets {
		usets[i].Relation = cuRel
	}

	return usets, nil
}

// Map a record's User to an Userset
func tupleToUserset(tuple model.TupleRecord) model.Userset {
	return model.Userset{
		Namespace: tuple.Tuple.User.Userset.Namespace,
		ObjectId:  tuple.Tuple.User.Userset.ObjectId,
		Relation:  tuple.Tuple.User.Userset.Relation,
	}
}
