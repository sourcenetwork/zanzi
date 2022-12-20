package simple

import (
	"context"
	"fmt"
	"log"

	"github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
	"github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
)

func newWalker(tStore tuple.TupleStore, pStore policy.PolicyStore) walker {
	return walker{
		tStore:  tStore,
		pStore:  pStore,
		fetcher: rg.NewRuleSucessorFetcher(tStore),
	}
}

// walker performs a walk through the relation graph
type walker struct {
	tStore  tuple.TupleStore
	pStore  policy.PolicyStore
	fetcher rg.RuleSucessorFetcher

	trail    map[tuple.TupleNode]struct{}
	pg       policy.PolicyGraph
	policyId string
}

func (w *walker) Walk(ctx context.Context, policyId string, source tuple.TupleNode) (rg.RelationNode, error) {
	w.trail = make(map[tuple.TupleNode]struct{})
	w.policyId = policyId

	opt, err := w.pStore.GetPolicyGraph(policyId)

	if err != nil {
		return rg.RelationNode{}, fmt.Errorf("walk failed: %w", err)
	}

	if opt.IsEmpty() {
		return rg.RelationNode{}, fmt.Errorf("walk failed: policy not found id=%v", policyId)
	}

	w.pg = opt.Value()

	ruleOpt := w.pg.GetRule(source.Namespace, source.Relation)
	if ruleOpt.IsEmpty() {
		// checking that policy contains the given relation is an important invariant.
		// during the entry point it can simply mean an invalid input.
		// if this condition happens during the fetch however, it signals an inconsistency
		// either a policy udpate with orphaned tuples or a schema enforcer bug
		return rg.RelationNode{}, fmt.Errorf("walk failed: policy %v does not have rule (%v,%v)", w.policyId, source.Namespace, source.Relation)
	}

	node, err := w.walk(ctx, source)
	if err != nil {
		err = fmt.Errorf("walk failed for %v: %v", source, err)
		return rg.RelationNode{}, err
	}

	return *node, nil
}

// Walk through the relation graph starting at source.
//
// keeps a trail through the depth search to avoid cyclic expands
func (w *walker) walk(ctx context.Context, source tuple.TupleNode) (*rg.RelationNode, error) {
	// TODO maybe set max recursion depth

	// expandTree is the recusion entrypoint for Expand subcalls
	// Therefore it's a centralized point to check whether ctx is still valid before continuing
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	// handles a backtrail expand call such as: A -> B -> A
	// return a non-expanded leaf with the uset in it
	_, isCycle := w.trail[source]

	// return if cycle or node is actor / object (recursion base cases)
	if isCycle || source.Type != tuple.NodeType_RELATION_SOURCE {
		node := &rg.RelationNode{
			ObjRel: source,
		}
		return node, nil
	}

	w.trail[source] = struct{}{}
	defer delete(w.trail, source)

	ruleOpt := w.pg.GetRule(source.Namespace, source.Relation)
	if ruleOpt.IsEmpty() {
		// This is bad:
		// As an invariant the Walk method checks whether the relaion exists within the policy.
		// Therefore, the only scenario where this can happen is if a previously fetched Tuple
		// references a relation / permission not contained in the policy.
		// The root causes might be due to a Policy mutation which led to orphaned tuples
		// or a faulty policy / schema enforcing mechanism while adding tuples.
		// Either way, this is an error condition and must be reported
		msg := fmt.Sprintf("inconsistent relation graph: tuplenode references unknown relation name: policy=%s tupleNode=%v", w.policyId, source)
		log.Println(msg)

		// NOTE: Not sure what I should do here.
		// Should I ignore this and skip?
		// Technically I could return nil and recover from this situation.
		// Nevertheless, the inconsistency must be logged / tracked as it indicates a bug.
		return nil, fmt.Errorf(msg)
	}
	rule := ruleOpt.Value()
	rewriteTree := rule.ExpressionTree

	rewriteNode, err := w.evalRewriteTree(ctx, rewriteTree, source)
	if err != nil {
		err = fmt.Errorf("failed evaluating rewrite tree for %v: %v", source, err)
		return nil, err
	}

	node := &rg.RelationNode{
		ObjRel: source,
		Child:  rewriteNode,
	}
	return node, nil
}

func (w *walker) evalRewriteTree(ctx context.Context, rewriteTree *policy.Tree, node tuple.TupleNode) (rg.RewriteNode, error) {
	switch n := rewriteTree.Node.(type) {
	case *policy.Tree_Opnode:
		opnode := n.Opnode
		return w.expandOpNode(ctx, opnode, node)
	case *policy.Tree_Leaf:
		leaf := n.Leaf
		return w.expandRuleNode(ctx, leaf, node)
	default:
		err := fmt.Errorf("Rewrite Node type unknown: %#v", node)
		panic(err)
	}
}

// Expand OpNode by expand its left and right children
func (w *walker) expandOpNode(ctx context.Context, root *policy.OpNode, node tuple.TupleNode) (*rg.OpNode, error) {
	left, err := w.evalRewriteTree(ctx, root.Left, node)
	if err != nil {
		err = fmt.Errorf("failed expanding opnode %v: %v", node, err)
		return nil, err
	}

	right, err := w.evalRewriteTree(ctx, root.Right, node)
	if err != nil {
		err = fmt.Errorf("failed expanding opnode %v: %v", node, err)
		return nil, err
	}

	opnode := &rg.OpNode{
		Left:   left,
		Right:  right,
		JoinOp: root.Op,
	}
	return opnode, nil
}

// Recurses and builds an expand tree for node's logical neighbors
// Return a RuleNode with the built trees as children.
func (w *walker) expandRuleNode(ctx context.Context, leaf *policy.Leaf, node tuple.TupleNode) (*rg.RuleNode, error) {
	rule := leaf.Rule
	sucessors, err := w.fetcher.Fetch(ctx, rule, w.policyId, node)
	if err != nil {
		return nil, err
	}

	data := rg.NewRuleData(rule)

	children := make([]*rg.RelationNode, 0, len(sucessors))
	for _, node := range sucessors {
		child, err := w.walk(ctx, node)
		if err != nil {
			err = fmt.Errorf("failed building subrtree: %v", err)
			return nil, err
		}
		children = append(children, child)
	}

	ruleNode := &rg.RuleNode{
		RuleData: data,
		Children: children,
	}
	return ruleNode, nil
}
