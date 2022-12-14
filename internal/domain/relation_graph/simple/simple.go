// Package simple provides a simplified implementation of authorizers methods
//
// The primary implementation is done through the Expand call.
// Expand call is used by the Check call.
// Check is used by ReverseLookup
package simple

import (
    "fmt"
    "context"

    "google.golang.org/protobuf/proto"
    mapset "github.com/deckarep/golang-set/v2"

    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
)

var _ rg.RelationGraph = (*RelationGraph[proto.Message])(nil)

func NewSimple[T proto.Message](tStore tuple.TupleStore[T], pStore policy.PolicyStore) rg.RelationGraph {
    return &RelationGraph[T]{
        tupleStore: tStore,
        policyStore: pStore,
        walker: newWalker[T](tStore, pStore),
    }
}

type RelationGraph[T proto.Message] struct {
    tupleStore tuple.TupleStore[T]
    policyStore policy.PolicyStore
    walker walker[T]
}

func (g *RelationGraph[T]) Walk(ctx context.Context, policyId string, source tuple.TupleNode) (rg.RelationNode, error) {
    return g.walker.Walk(ctx, policyId, source)
}

func (g *RelationGraph[T]) GetPath(ctx context.Context, policyId string, source tuple.TupleNode, dest tuple.TupleNode) ([]tuple.TupleNode, error) {
    return nil, fmt.Errorf("Not impemented")
}

func (g *RelationGraph[T]) GetSucessors(ctx context.Context, policyId string, source tuple.TupleNode) (mapset.Set[tuple.TupleNode], error) {
    tree, err := g.Walk(ctx, policyId, source)
    if err != nil {
        return nil, err
    }

    return rg.EvalTree(&tree), nil
}

func (g *RelationGraph[T]) GetAncestors(ctx context.Context, policyId string, source tuple.TupleNode) ([]tuple.TupleNode, error) {
    return nil, fmt.Errorf("Not impemented")
}

func (g *RelationGraph[T]) IsReachable(ctx context.Context, policyId string, source, dest tuple.TupleNode) (bool, error) {
    nodes, err := g.GetSucessors(ctx, policyId, source)
    if err != nil {
        return false, err
    }

    return nodes.Contains(dest), nil
}
