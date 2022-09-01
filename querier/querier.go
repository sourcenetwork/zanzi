// Package querier provides an abstraction to perform complex lookups over relationships.
//
// Querier defines an interface for the possible operations and exposes implementations
// which differ in performance and resource requirements
package querier

import (
	"context"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/tree"
)

// Querier abstracts graph traversal and search operations
type Querier interface {
	// Build Tree of all Usersets related to the given Userset
	Expand(ctx context.Context, userset model.Userset) (tree.UsersetNode, error)

	// Verify whether object has relation with user.
	// Check calls can be short circuited,
	// considering that once a path evaluates to true other checks become redundant.
	Check(ctx context.Context, namespace, objectId string) (bool, error)

	// Reverse lookup return all tuples user is related to
	ReverseLookup(ctx context.Context, user model.User) ([]model.TupleRecord, error)

	// Return all objects an user is related to expressed through a Tree.
	// Node children are annotated with the source of the relation,
	// either through direct lookup or userset rewrites
	ExplainedReverseLookup(ctx context.Context, user model.User) (tree.UsersetNode, error)
}
