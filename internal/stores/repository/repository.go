// Package repository provides an interface to interact with the stored relation tuples.
package repository

import (
	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/utils"
)

// TupleRepository abstract interfacing with tuple storage.
type TupleRepository[T model.Relationship] interface {
	// Store a new tuple. Update all indexes
	SetTuple(tuple T) error

	// Looks up a relation tuple from the backend storage
	GetTuple(object, subject model.AuthNode) (utils.Option[T], error) // TODO this may be kinda annoying, having two auth nodes without an intermediary type

	// Return tuples directly connected to userset
	// Does not follows userset indirection.
	GetRelatedUsersets(node model.AuthNode) ([]T, error)

	// Return tuples that have an outgoing edge to userset (aka reverse lookup)
	// Does not perform userset chasing
	GetIncomingUsersets(node model.AuthNode) ([]T, error)

	// Purge tuple from storage.
	RemoveTuple(object, subject model.AuthNode) error

	// Return all tuples that match the filter:
	// tuple.Relation == relation
	// tuple.User.ObjectId == objectId
	//
	// This peculiar query is used to reverse tuple to userset rules
	// during reverse lookup
	GetTuplesFromRelationAndUserObject(relation string, objNamespace string, objectId string) ([]T, error)
}

// NamespaceRepository abstract interfacing with namespace storage.
//
// Contract: methods should return instance of `EntityNotFound`
// when some operation failed because the lookup resulted in an empty set.
type NamespaceRepository interface {
	GetNamespace(namespace string) (utils.Option[model.Namespace], error)

	SetNamespace(namespace model.Namespace) (model.Namespace, error)

	RemoveNamespace(namespace string) error

	// Return a Relation definition from a namespace
	GetRelation(namespace, relation string) (model.Relation, error)

	// Return all relations which reference the given `relation`.
	GetReferrers(namespace, relation string) ([]model.Relation, error)
}
