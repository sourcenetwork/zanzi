package simple

import (
	"context"
	"errors"
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
	trail     map[model.AuthNode]struct{}
	tupleRepo repository.TupleRepository
	nsRepo    repository.NamespaceRepository
}

func (e *expander) Expand(ctx context.Context, userset model.AuthNode) (tree.UsersetNode, error) {
	e.trail = make(map[model.AuthNode]struct{})

	expTree, err := e.getExpTree(ctx, userset.Namespace, userset.Relation)
	if err != nil {
		err = fmt.Errorf("Expand failed for %v: %v", userset, err)
		return tree.UsersetNode{}, err
	}

	node, err := e.expandTree(ctx, expTree, userset)
	if err != nil {
		err = fmt.Errorf("Expand failed for %v: %v", userset, err)
		return tree.UsersetNode{}, err
	}

	return *node, nil
}

// Expand an Userset Rewrite Expression Tree
// keeps a trail through the depth search to avoid cyclic expands
func (e *expander) expandTree(ctx context.Context, root *model.RewriteNode, uset model.AuthNode) (*tree.UsersetNode, error) {
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
	_, ok := e.trail[uset]
	if ok {
		return node, nil
	}

	// Handle an empty expression tree
	// Should only happen when uset.Relation is the empty relation constant
	if root == nil {
		return node, nil
	}

	e.trail[uset] = struct{}{}
	defer delete(e.trail, uset)

	exprNode, err := e.expandExprNode(ctx, root, uset)
	if err != nil {
		err = fmt.Errorf("expandTree failed for userset %v: %v", uset, err)
		return nil, err
	}

	node.Child = exprNode
	return node, nil
}

func (e *expander) expandExprNode(ctx context.Context, root *model.RewriteNode, uset model.AuthNode) (tree.ExpressionNode, error) {
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
func (e *expander) expandOpNode(ctx context.Context, root *model.OpNode, uset model.AuthNode) (*tree.OpNode, error) {
	left, err := e.expandExprNode(ctx, root.Left, uset)
	if err != nil {
		err = fmt.Errorf("failed expanding opnode %v: %v", root, err)
		return nil, err
	}

	right, err := e.expandExprNode(ctx, root.Right, uset)
	if err != nil {
		err = fmt.Errorf("failed expanding opnode %v: %v", root, err)
		return nil, err
	}

	node := &tree.OpNode{
		Left:   left,
		Right:  right,
		JoinOp: root.Op,
	}
	return node, nil
}

func (e *expander) expandRuleNode(ctx context.Context, root *model.Leaf, uset model.AuthNode) (*tree.RuleNode, error) {

	var neighbors []model.AuthNode
	var rule tree.Rule
	var err error

	switch r := root.Rule.GetRule().(type) {

	case *model.Rule_This:
		neighbors, err = e.produceThis(uset)
		rule = tree.Rule{
			Type: tree.RuleType_THIS,
		}

	case *model.Rule_TupleToUserset:
		ttu := r.TupleToUserset
		neighbors, err = e.produceTTU(uset, ttu.TuplesetRelation, ttu.ComputedUsersetRelation)
		rule = tree.Rule{
			Type: tree.RuleType_TTU,
			Args: map[string]string{
				"TuplesetRelation":        ttu.TuplesetRelation,
				"ComputedUsersetRelation": ttu.ComputedUsersetRelation,
			},
		}

	case *model.Rule_ComputedUserset:
		cu := r.ComputedUserset
		neighbors = e.produceCU(uset, cu.Relation)
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
		err = fmt.Errorf("Failed expanding rule %v for node %v: %v", root.Rule.GetRule(), uset, err)
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

// Recurses and builds an expand tree for all neighbors.
// Return a RuleNode with the built trees as children.
func (e *expander) expandRule(ctx context.Context, uset model.AuthNode, neighbors []model.AuthNode, rule tree.Rule) (*tree.RuleNode, error) {
	children := make([]*tree.UsersetNode, 0, len(neighbors))
	for _, neigh := range neighbors {
		expTree, err := e.getExpTree(ctx, neigh.Namespace, neigh.Relation)
		if err != nil {
			err = fmt.Errorf("failed fetching exp tree: %v", err)
			return nil, err
		}

		// bfs call
		child, err := e.expandTree(ctx, expTree, neigh)
		if err != nil {
			err = fmt.Errorf("failed building subrtree: %v", err)
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

// Return direct descendents of uset
func (e *expander) produceThis(uset model.AuthNode) ([]model.AuthNode, error) {
	tuples, err := e.tupleRepo.GetRelatedUsersets(uset)

	if errors.Is(err, &repository.EntityNotFound{}) {
		return nil, nil
	}

	if err != nil {
		err = fmt.Errorf("failed fetching descendents for node: %v", err)
		return nil, err
	}

	usets := utils.MapSlice(tuples, tupleToUserset)

	return usets, nil
}

// Return logical descendent made by evaluating a Computed Userset rule
func (e *expander) produceCU(uset model.AuthNode, relation string) []model.AuthNode {
	return []model.AuthNode{
		model.AuthNode{
			Namespace: uset.Namespace,
			ObjectId:  uset.ObjectId,
			Relation:  relation,
		},
	}
}

// Return logical descendents reachable from uset by computing a TTU rule.
// TTU rule is defined by tsetRel and cuRel
func (e *expander) produceTTU(uset model.AuthNode, tsetRel string, cuRel string) ([]model.AuthNode, error) {
	tuplesetFilter := model.AuthNode{
		Namespace: uset.Namespace,
		ObjectId:  uset.ObjectId,
		Relation:  tsetRel,
	}

	records, err := e.tupleRepo.GetRelatedUsersets(tuplesetFilter)
	if errors.Is(err, &repository.EntityNotFound{}) {
		// An empty result set from a TTU call cannot be considered
		// an error, as there is no guarantee the target will exist
		return nil, nil
	}
	if err != nil {
		err = fmt.Errorf("failed to produce TTU neighbors for userset %v: %v", uset, err)
		return nil, err
	}

	usets := utils.MapSlice(records, tupleToUserset)
	for i := range usets {
		usets[i].Relation = cuRel
	}

	return usets, nil
}

// Map a record's User to an Userset
func tupleToUserset(tuple model.Relationship) model.AuthNode {
    return tuple.GetSubject()
}
