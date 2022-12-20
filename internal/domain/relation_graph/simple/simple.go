// Package simple provides a simplified implementation of authorizers methods
//
// The primary implementation is done through the Expand call.
// Expand call is used by the Check call.
// Check is used by ReverseLookup
package simple

import (
	"context"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
	"github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
)

var _ rg.RelationGraph = (*RelationGraph)(nil)

func NewSimple(tStore tuple.TupleStore, pStore policy.PolicyStore) rg.RelationGraph {
	return &RelationGraph{
		tupleStore:  tStore,
		policyStore: pStore,
		walker:      newWalker(tStore, pStore),
	}
}

type RelationGraph struct {
	tupleStore  tuple.TupleStore
	policyStore policy.PolicyStore
	walker      walker
}

func (g *RelationGraph) Walk(ctx context.Context, policyId string, source tuple.TupleNode) (rg.RelationNode, error) {
	return g.walker.Walk(ctx, policyId, source)
}

func (g *RelationGraph) GetPath(ctx context.Context, policyId string, source tuple.TupleNode, dest tuple.TupleNode) ([]tuple.TupleNode, error) {
	return nil, fmt.Errorf("Not impemented")
}

func (g *RelationGraph) GetSucessors(ctx context.Context, policyId string, source tuple.TupleNode) (mapset.Set[tuple.TupleNode], error) {
	tree, err := g.Walk(ctx, policyId, source)
	if err != nil {
		return nil, err
	}

	return rg.EvalTree(&tree), nil
}

func (g *RelationGraph) GetAncestors(ctx context.Context, policyId string, source tuple.TupleNode) ([]tuple.TupleNode, error) {
	return nil, fmt.Errorf("Not impemented")
}

func (g *RelationGraph) IsReachable(ctx context.Context, policyId string, source, dest tuple.TupleNode) (bool, error) {
	nodes, err := g.GetSucessors(ctx, policyId, source)
	if err != nil {
		return false, err
	}

	return nodes.Contains(dest), nil
}
