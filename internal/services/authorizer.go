package services

import (
	"context"
	"fmt"

	rg "github.com/sourcenetwork/zanzi/internal/domain/relation_graph"
	tuple "github.com/sourcenetwork/zanzi/internal/domain/tuple"
	"github.com/sourcenetwork/zanzi/internal/mappers"
	"github.com/sourcenetwork/zanzi/types"
)

// Return an Authorizer implementation from a relation graph
func AuthorizerFromRelationGraph(relGraph rg.RelationGraph) types.Authorizer {
	return &authorizer{
		rg: relGraph,
	}
}

// authorizer implements the Authorizer interface by wrapping a relation graph
type authorizer struct {
	rg      rg.RelationGraph
	builder tuple.TupleBuilder
	mapper  mappers.ExpandTreeMapper
}

func (a *authorizer) Check(policyId string, obj types.Entity, relation string, actor types.Entity) (bool, error) {
	ctx := context.Background()
	src := a.builder.RelSource(obj.Namespace, obj.Id, relation)
	dst := a.builder.ActorWithNamespace(actor.Namespace, actor.Id)
	reachable, err := a.rg.IsReachable(ctx, policyId, src, dst)
	if err != nil {
		return false, fmt.Errorf("check failed: %w", err)
	}

	return reachable, nil
}

func (a *authorizer) Reverse(policyId string, actor types.Entity) ([]types.EntityRelPair, error) {
	return nil, nil
}

func (a *authorizer) Expand(policyId string, obj types.Entity, relation string) (types.ExpandTree, error) {
	ctx := context.Background()
	src := a.builder.RelSource(obj.Namespace, obj.Id, relation)

	tree, err := a.rg.Walk(ctx, policyId, src)
	if err != nil {
		return types.ExpandTree{}, fmt.Errorf("expand failed: %w", err)
	}

	return a.mapper.ToExpandTree(&tree), nil
}
