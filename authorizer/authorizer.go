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
	Check(ctx context.Context, objRel model.AuthNode, user model.AuthNode) (bool, error)
}

type Reverser interface {
	// Reverse lookup return all tuples user is related to
	ReverseLookup(ctx context.Context, user model.AuthNode) ([]model.AuthNode, error)

	// Return all objects an user is related to expressed through a Tree.
	// Node children are annotated with the source of the relation,
	// either through direct lookup or userset rewrites
	//ExplainedReverseLookup(ctx context.Context, user model.User) (*tree.AuthNodeNode, error)
}

type Expander interface {
	// Build Tree of all AuthNodes related to the given AuthNode
	Expand(ctx context.Context, userset model.AuthNode) (tree.AuthNodeNode, error)
}

type Authorizer {
    Expander,
    Reverser,
    Checker
}
