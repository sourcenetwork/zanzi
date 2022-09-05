// Package repository provides an interface to interact with the stored relation tuples.
package repository

import (
	"fmt"

	"github.com/sourcenetwork/source-zanzibar/model"
)

// TupleRepository abstract interfacing with tuple storage.
//
// Contract: methods should return instance of `EntityNotFound`
// when some operation failed because the lookup resulted in an empty set.
type TupleRepository interface {
	// Store a new tuple. Update all indexes
	SetTuple(tuple model.Tuple) error

	// Looks up a relation tuple from the backend storage
	GetTuple(tuple model.Tuple) (model.TupleRecord, error)

	// Return tuples directly connected to userset
	// Does not follows userset indirection.
	GetRelatedUsersets(userset model.Userset) ([]model.TupleRecord, error)

	// Return tuples that have an outgoing edge to userset (aka reverse lookup)
	// Does not perform userset chasing
	GetIncomingUsersets(userset model.Userset) ([]model.TupleRecord, error)

	// Purge tuple from storage.
	RemoveTuple(tuple model.Tuple) error
}

// NamespaceRepository abstract interfacing with namespace storage.
//
// Contract: methods should return instance of `EntityNotFound`
// when some operation failed because the lookup resulted in an empty set.
type NamespaceRepository interface {
	GetNamespace(namespace string) (model.Namespace, error)

	SetNamespace(namespace model.Namespace) error

	RemoveNamespace(namespace string) error

	// Return a Relation definition from a namespace
	GetRelation(namespace, relation string) (model.Relation, error)

	// Return all relations which reference the given `relation`.
        //GetReferrers(namespace, relation string) ([]model.Relation, error)
}

// EntityNotFound type implements error interface.
// Caller code may compare against it to verify whether
// returned error was a storage error or a not found error
type EntityNotFound struct {
	Entity string
	Args   []any
}

// Build instance of EntityNotFound.
// Caller may specify entity name and any filters used during
// lookup
// Supplied information will be used to generate final error message.
func NewEntityNotFound(entity string, args ...any) error {
	return &EntityNotFound{
		Entity: entity,
		Args:   args,
	}
}

func (e *EntityNotFound) Error() string {
	return fmt.Sprintf("Entity not found: entity=%v, args=%v", e.Entity, e.Args)
}

var _ error = (*EntityNotFound)(nil)
