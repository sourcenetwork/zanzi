// Package repository provides an interface to interact with the stored relation tuples.
//
// Internally repository manages database indexes, keeps indexes and provides queries.
package repository

// Repository leverages the "tm-db" DB interface as a storage backend abstraction.
// This choice was tentatively made to ease future integration with a dApp powered by Cosmos SDK,
// potentially through [github.com/cosmos/cosmos-sdk/store/types.CommitMultiStore.MountStoreWithDB]

import (
    "fmt"

    _ "github.com/tendermint/tm-db"

    "github.com/sourcenetwork/source-zanzibar/model"
)

const (
    tuplePath      = "tuples/"
    namespacePath  = "namespaces/"
    usersetIdxPath = "usersets/"
)

// possible indexes:
//
// tuples:
// tuples/{namespace}/{obj_id}/{relation}/{namespace}/{id}/{relation} -> TupleRecord
// tuples/reversed/{user_namespace}/{id}/{relation}/{namespace}/{object_id}/{relation} -> TupleRecord  (reverse lookup)
//
// namespaces:
// namespaces/{namespace-name} -> Namespace
// namespaces/{namespace-name}/{relation-name} -> Relation
// namespace/{namespace_name}/{relation_name}/referenced_by/{relation_name} -> Relation (reverse relation lookup used in reverse lookup)

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
    GetParentTuples(userset model.Userset) ([]model.TupleRecord, error)

    // Purge tuple from storage.
    RemoveTuple(tuple model.Tuple) error
}

type NamespaceRepository interface {
    GetNamespace(namespace string) (model.Namespace, error)

    SetNamespace(namespace model.Namespace) error

    RemoveNamespace(namespace string) error

    // Return a Relation definition from a namespace
    // NOTE use relation index?
    GetRelation(namespace, relation string) (model.Relation, error)
}

type EntityNotFound struct {
    Entity string
    Args []any
}

func NewEntityNotFound(entity string, args ...any) error {
    return &EntityNotFound {
        Entity: entity,
        Args: args,
    }
}

func (e *EntityNotFound) Error() string {
    return fmt.Sprintf("Entity not found: entity=%v, args=%v", e.Entity, e.Args)
}

var _ error = (*EntityNotFound)(nil)
