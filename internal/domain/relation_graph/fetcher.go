package relation_graph

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/zanzi/internal/domain/policy"
	"github.com/sourcenetwork/zanzi/internal/domain/tuple"
	"github.com/sourcenetwork/zanzi/pkg/utils"
)

// Sucessor fetcher fetches the logical sucessors of a node in a relation graph.
//
// Relation Graph can be thought of as a "view" over the graph defined by relation tuples.
// Each Tuple defines two nodes in the concrete Relation Graph.
// Other nodes are added to the Relation Graph through a Policy definition.
//
// Policies defines rules, which are relationships between tuple relations.
//
// In order to fetch the sucessors of a node, it must combine the tuple storage with the rules in a policy.
type RuleSucessorFetcher struct {
	tStore tuple.TupleStore
}

func NewRuleSucessorFetcher(tStore tuple.TupleStore) RuleSucessorFetcher {
	return RuleSucessorFetcher{
		tStore: tStore,
	}
}

// Apply rewrite rule to node and return the resulting node sucessors
func (f *RuleSucessorFetcher) Fetch(ctx context.Context, rule *policy.RewriteRule, policyId string, node tuple.TupleNode) ([]tuple.TupleNode, error) {
	var sucessors []tuple.TupleNode
	var err error
	switch r := rule.RewriteRule.(type) {

	case *policy.RewriteRule_This:
		sucessors, err = f.getThisSucessors(ctx, policyId, node)
	case *policy.RewriteRule_TupleToUserset:
		ttu := r.TupleToUserset
		sucessors, err = f.getTTUSucessors(ctx, ttu, policyId, node)
	case *policy.RewriteRule_ComputedUserset:
		cu := r.ComputedUserset
		sucessors = f.getCUSucessors(ctx, cu, policyId, node)
	default:
		// This should not happen unless library clients implement their own rule types
		// which should not be possible as rules are part of the internal pkg.
		err = fmt.Errorf("invalid rule type: %#v", r)
		panic(err)
	}

	if err != nil {
		err = fmt.Errorf("error fetching sucessors: rule %v, node %v: %v", rule.RewriteRule, node, err)
	}
	return sucessors, err
}

// Return direct descendents of uset
func (f *RuleSucessorFetcher) getThisSucessors(ctx context.Context, policyId string, node tuple.TupleNode) ([]tuple.TupleNode, error) {
	sucessors, err := f.tStore.GetSucessors(policyId, node)

	if len(sucessors) == 0 {
		return nil, nil
	}

	if err != nil {
		err = fmt.Errorf("failed fetching This sucessors for: node=%v: %v", node, err)
		return nil, err
	}

	nodes := utils.MapSlice(sucessors, getDestNode)

	return nodes, nil
}

// Return logical descendent made by evaluating a Computed Userset rule
func (f *RuleSucessorFetcher) getCUSucessors(ctx context.Context, cu *policy.ComputedUserset, policyId string, node tuple.TupleNode) []tuple.TupleNode {
	node.Relation = cu.Relation
	node.Type = tuple.NodeType_RELATION_SOURCE
	return []tuple.TupleNode{
		node,
	}
}

// Return logical sucessors reachable from node by applying a Tuple To Userset rule.
func (f *RuleSucessorFetcher) getTTUSucessors(ctx context.Context, ttu *policy.TupleToUserset, policyId string, node tuple.TupleNode) ([]tuple.TupleNode, error) {
	// TTU is a complex rule to implement.
	// The steps are:
	// 1. Use node to build a new node with its relation replaced - call it filter
	// 2. Fetch filter's sucessors from tuple store
	// 3. Map fetched sucessors into a new node by replacing its relation with the cu_relation given in
	// the TupleToUserset rule
	// The mapped sucessors are the resulting nodes from evaluating TTU
	tuplesetFilter := tuple.TupleNode{
		Namespace: node.Namespace,
		Id:        node.Id,
		Relation:  ttu.TuplesetRelation,
	}

	sucessors, err := f.tStore.GetSucessors(policyId, tuplesetFilter)
	if len(sucessors) == 0 {
		// An empty result set from a TTU call cannot be considered
		// an error, as there is no guarantee the target will exist
		return nil, nil
	}

	if err != nil {
		err = fmt.Errorf("failed to produce TTU neighbors for node %v: %v", node, err)
		return nil, err
	}

	nodes := utils.MapSlice(sucessors, getDestNode)
	for i := range nodes {
		nodes[i].Relation = ttu.CuRelation
		nodes[i].Type = tuple.NodeType_RELATION_SOURCE
	}

	return nodes, nil
}

/*

// Ancestor Fetcher retrieves ancestor tuples from a node
// Performs reverse rewrite rule evaluation
type AncestorFetcher[T proto.Message] struct {
	logicalAncestors []tuple.TupleNode
        tStore tuple.TupleStore
}

func NewAncestorFetcher[T proto.Message](tupleStore tuple.TupleStore) AncestorFetcher {
	return AncestorFetcher{
		tStore: tupleRepo,
	}
}

// Return all ancestors nodes
// Includes direct ancestors and logical ones (obtained by reverting rewrite rules)
func (f *AncestorFetcher) FetchAll(ctx context.Context, pg policy.PolicyGraph, root tuple.TupleNode) ([]tuple.TupleNode, error) {
	directAncestors, err := f.FetchAncestors(ctx, pg, root)
	if err != nil {
		return nil, err
	}

	logicalAncestors, err := f.FetchLogicalAncestors(ctx, pg, root)
	if err != nil {
		return nil, err
	}

	return append(directAncestors, logicalAncestors...), nil
}

// Fetch all directly accessible Ancestors of uset
func (f *AncestorFetcher) FetchDirectAncestors(ctx context.Context, pg policy.PolicyGraph, root tuple.TupleNode) ([]tuple.TupleNode, error) {
	ancestors, err := f.tStore.GetAncestors(pg.GetPolicyId(), root)
	if err != nil {
		err = fmt.Errorf("fetch ancestors failed for %v: %v", root, err)
		return nil, err
	}
	return utils.MapSlice(ancestors, getSourceNode), nil
}

// fetchLogicalAncestors return Relationships which are "logical ancestors" of uset.
// Logical Ancestors are edges reachable through userset rewrite rules.
func (f *AncestorFetcher) FetchLogicalAncestors(ctx context.Context, pg policy.PolicyGraph, root tuple.TupleNode) ([]tuple.TupleNode, error) {
	f.logicalAncestors = nil

        id := pg.GetPolicyId()
        opt := pg.GetAncestors(root.Namespace, root.Relation)
        if opt.IsEmpty() {
            err := fmt.Errorf("failed fetching ancestors for %v: policy %v does not have resource relation pair (%v, %v)", root, id, root.Namespace, root.Relation)
            return nil, err
        }

        ancestorRels := opt.Value()
	for _, relation := range ancestorRels {
		err := f.buildAncestorsFromRel(ctx, pg, root, relation)
		if err != nil {
			err = fmt.Errorf("failed building ancestors for %v: %v", root, err)
			return nil, err
		}
	}

	return f.logicalAncestors, nil
}


// Extracts all rules from relation which reference uset.Relation
// for each matching rule, build possible ancestors and
// append results to f.logicalAncestors
func (f *AncestorFetcher) buildAncestorsFromRel(ctx context.Context, uset tuple.TupleNode, relation model.Relation) error {
	rules := f.getReferrers(relation, uset)

	for _, rule := range rules {
		err := f.buildAncestorsFromRule(ctx, uset, relation, rule)
		if err != nil {
			fmt.Errorf("failed building ancestors for relation %v: %v", relation.Name, err)
			return err
		}
	}
	return nil
}


// Append potential logical ancestors of uset for a relation, rule pair
// Assume rules were previously filtered to only contain TTU and CU rules
// which applies to the given uset
func (f *AncestorFetcher) fetchAncestorsForAncestorRule(ctx context.Context, node tuple.TupleNode, ancestor PolicyNode) error {
	switch rule := rule.Rule.(type) {
	case *model.Rule_TupleToUserset:
		ttu := rule.TupleToUserset

		usets, err := f.revertTTU(ctx, uset, relation.Name, *ttu)
		if err != nil {
			return err
		}

		f.logicalAncestors = append(f.logicalAncestors, usets...)

	case *model.Rule_ComputedAuthNode:
		uset := f.revertCU(uset, relation.Name)
		f.logicalAncestors = append(f.logicalAncestors, uset)

	default:
		panic("invalid rule type")
	}
	return nil
}

// Apply a CU rule backwards and return the original node
// Returned node is not guaranteed to exist
func (f *AncestorFetcher) revertCU(uset tuple.TupleNode, referer string) tuple.TupleNode {
	return tuple.TupleNode{
		Namespace: uset.Namespace,
		ObjectId:  uset.ObjectId,
		Relation:  referer,
                Type: uset.Type,
	}
}

// Apply a TTU rule backwards and return nodes that may reach uset
// Returned nodes are not guaranteed to exist
func (f *AncestorFetcher) ancestorsTTU(ctx context.Context, pg policy.PolicyGraph, ttu policy.TupleToUserset, node tuple.TupleNode) ([]tuple.TupleNode, error) {
	tuples, err := f.tStore.GetGrantingTuples(pg.GetPolicyId(), ttu.TuplesetRelation, node.Namespace, node.ObjectId)
	if err != nil {
		err = fmt.Errorf("failed reverting TTU rule %v: %v", ttu, err)
		return nil, err
	}

	ancestors := utils.MapSlice(tuples, getSourceNode)

	for i, _ := range ancestors {
		ancestors[i].Relation = referer
	}

	return ancestors, nil
}

*/

func getSourceNode(tuple tuple.Tuple) tuple.TupleNode {
	return tuple.Source
}

func getDestNode(tuple tuple.Tuple) tuple.TupleNode {
	return tuple.Dest
}
