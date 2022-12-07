package relation_graph

import (
    "context"

    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
)

type RelationGraph[T proto.Message] interface {
    // Build a DFS Tree resulting from walking the RelationGraph,
    // starting from source.
    Walk(ctx context.Context, partition string, source tuple.TupleNode) (RelationNode, error)

    // Return shortest path between source and dest.
    GetPath(ctx context.Context, partition string, source tuple.TupleNode, dest tuple.TupleNode) ([]tuple.TupleNode, error)

    // expand quiet
    GetSucessors(ctx context.Context, partition string, source tuple.TupleNode) ([]tuple.TupleNode, error)

    // reverse lookup
    GetAncestors(ctx context.Context, partition string, source tuple.TupleNode) ([]tuple.TupleNode, error)
}
