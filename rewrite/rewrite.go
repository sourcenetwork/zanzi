// Package rewrite defines types for Userset Rewrite Expression tree handling
package rewrite

import (
	"context"
)

// Build Expression Tree for a relation.
// Leaves are rewrite rules, inner nodes have 2 children and indicate a set operation
// (ie union, intersection, difference).
// Return reference to root node.
func BuildExpressionTree(ctx context.Context, namespace string, relation string) (Node, error) {
	// NOTE should I use a ctx or pass the Repository instance directly here?
	// Not sure which would be a better approach
}
