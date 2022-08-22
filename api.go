// Package source-zanzibar exposes a zanzibar-like public API
package source-zanzibar

import (
    "github.com/sourcenetwork/source-zanzibar/model"
    "github.com/sourcenetwork/source-zanzibar/db"
)


// Service manages resources needed to perform operations
struct Service {
    db db.DB
}

func NewService() *Serivce {

    var name string
    var path string

    // Initialize / load instance of DB.
    // Use Goleveldb as that is supposedly the standard according to the [docs].
    // [docs]: https://pkg.go.dev/github.com/tendermint/tm-db#section-readme
    levelDb, err := tm-db.NewGoLevelDB(name, path)
    if err != nil {
        return nil, err
    }

    // build repository

    // build service

    return nil
}


// Check verifies whether a user has relation with object.
// Returns when any valid path from the (object, relation) pair is found to the user 
// or until all possibilities are exhausted.
//
// Check traverses through userset rewrites.
func (s *Service) Check(object model.Object, relation string, user model.User) bool { 
    return true
}

// Expand builds a complete representation of all users that have the given relation with object.
//
// The representation is given in the shape of a tree which represents the underlying userset rewrite rules.
// The resulting tree preserves the rules which lead to a given set of leaves to be included.
func (s *Service) Expand(object model.Object, relation string) *Tree {
    return nil
}

// Lookup filters store for tuples with the given parameters.
// Lookup follows usersets but does not perform userset rewrite expansion.
func (s *Service) Lookup(obj model.Object, relation string) []model.TupleRecord {
    return nil
}

// ReverseLookup returns all tuples a user has access to.
// Accounts for userset rewrites and follows userset references
func (s *Service) ReverseLookup(user model.User) []model.TupleRecord {
    return nil
}


func (s *Service) WriteTuple(tuple model.Tuple) error {
    return nil
}

func (s *Service) DeleteTuple(tuple model.Tuple) error { 
    return nil
}

func (s *Service) WriteNamespace(namespace model.Namespace) error {
    // before storing there should be some check ensuring relationships aren't cyclic
    return nil
}

func (s *Service) GetNamespace(name model.Namespace) *model.Namespace {
    return nil
}

func (s *Service) DeleteNamespace(name string) {
    // should namespace removal cascade over stored tuples?
}
