// Package authorizer provides abstractions to perform authorization verification operations.
// The primary operations are part of Zanzibar's spec, eg Check, Expand and Reverse Lookup
package authorizer

import (
	"context"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/tree"
)

type Checker interface {
	// Verify whether object has relation with user.
	// Check calls can be short circuited,
	// considering that once a path evaluates to true other checks become redundant.
	Check(ctx context.Context, objRel model.Userset, user model.Userset) (bool, error)
}

type Reverser interface {
	// Reverse lookup return all tuples user is related to
	ReverseLookup(ctx context.Context, user model.User) ([]model.Userset, error)

	// Return all objects an user is related to expressed through a Tree.
	// Node children are annotated with the source of the relation,
	// either through direct lookup or userset rewrites
	//ExplainedReverseLookup(ctx context.Context, user model.User) (*tree.UsersetNode, error)
}

type Expander interface {
	// Build Tree of all Usersets related to the given Userset
	Expand(ctx context.Context, userset model.Userset) (tree.UsersetNode, error)
}
