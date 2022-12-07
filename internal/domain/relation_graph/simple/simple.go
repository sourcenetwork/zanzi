// Package simple provides a simplified implementation of authorizers methods
//
// The primary implementation is done through the Expand call.
// Expand call is used by the Check call.
// Check is used by ReverseLookup
package simple

import (
    "context"

    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    rg "github.com/sourcenetwork/source-zanzibar/relation_graph"
)

var _ rg.RelationGraph[proto.Message] = (*RelationGraph[proto.Message])(nil)

type RelationGraph[T proto.Message] struct {
    tupleStore tuple.TupleStore[T]
    policyStore policy.PolicyStore
    walker walker[T]
}

func (g *RelationGraph[T]) Walk(ctx context.Context, policyId string, source tuple.TupleNode) (rg.RelationNode, error) {
    return g.walker.Walk(ctx, policyId, source)
}

func (g *RelationGraph[T]) GetPath(ctx context.Context, policyId string, source tuple.TupleNode, dest tuple.TupleNode) ([]tuple.TupleNode, error) {
    return nil, nil
}

func (g *RelationGraph[T]) GetSucessors(ctx context.Context, policyId string, source tuple.TupleNode) ([]tuple.TupleNode, error) {
    return nil, nil
}

func (g *RelationGraph[T]) GetAncestors(ctx context.Context, policyId string, source tuple.TupleNode) ([]tuple.TupleNode, error) {
    return nil, nil
}
