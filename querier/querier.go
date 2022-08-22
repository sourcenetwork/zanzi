// Package querier provides an abstraction to perform complex lookups over relationships.
//
// Querier defines an interface for the possible operations and exposes implementations
// which differ in performance and resource requirements
package querier

// Querier
type Querier interface {
    Check() ()
    Expand() ()
    Lookup() ()
    ReverseLookup() ()

    // receives object and performs lookup according to rewrite rule
    Expand(ctx context.Context, namespace, objectId string) ([]model.User, error)

    // Verify whether object has relation with user.
    // Check calls can be short circuited,
    // considering that once a path evaluates to true other checks become redundant.
    Check(ctx context.Context, namespace, objectId string) (bool, error)

    // Reverse lookup return all tuples user is related to
    Reverse(ctx context.Context, user model.User) ([]model.TupleRecord, error)
}
