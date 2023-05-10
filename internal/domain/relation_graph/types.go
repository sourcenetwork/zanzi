package relation_graph

import (
	"context"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/sourcenetwork/zanzi/internal/domain/tuple"
)

type RelationGraph interface {
	// Build a DFS Tree resulting from walking the RelationGraph,
	// starting from source.
	Walk(ctx context.Context, partition string, source tuple.TupleNode) (RelationNode, error)

	// Return shortest path between source and dest.
	GetPath(ctx context.Context, partition string, source tuple.TupleNode, dest tuple.TupleNode) ([]tuple.TupleNode, error)

	// expand quiet
	GetSucessors(ctx context.Context, partition string, source tuple.TupleNode) (mapset.Set[tuple.TupleNode], error)

	// reverse lookup
	GetAncestors(ctx context.Context, partition string, source tuple.TupleNode) ([]tuple.TupleNode, error)

	// Return true if dest can be reached from source
	IsReachable(ctx context.Context, partition string, source tuple.TupleNode, dest tuple.TupleNode) (bool, error)
}
