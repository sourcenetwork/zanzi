// Package repository provides an interface to interact with the stored relation tuples.
//
// Internally repository manages database indexes, keeps indexes and provides queries.
package repository

// Repository leverages the "tm-db" DB interface as a storage backend abstraction.
// This choice was tentatively made to ease future integration with a dApp powered by Cosmos SDK,
// potentially through [github.com/cosmos/cosmos-sdk/store/types.CommitMultiStore.MountStoreWithDB]

import (
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

// Define methods to recursively search a repository by following usersets indirection
// NOTE the idea behind this type is that it could work with any repository.
// so perhaps turn this into a struct with a concrete implementation
type Chaser interface {

    // Recursively fetches tuples matching namespace, objectId and relation
    ChaseUsersets(userset model.Userset) ([]model.TupleRecord, error)

    // Perform an inverse lookup of all tuples whose user match the given userset.
    // ie. Starting from an user node, walks up the tuple graph by following usersets
    //ReverseChaseUserset(namespace, id, relation string) ([]model.TupleRecord, error)
}


// Minor c# idiom until i figure out what to do with repochaser
type ChaserImpl struct {
    repo TupleRepository
}

func (r *ChaserImpl) ChaseUsersets(userset model.Userset) ([]model.TupleRecord, error) {
    records, err := r.repo.GetRelatedUsersets(userset)
    if err != nil {
        // TODO wrap err
        return nil, err
    }

    usersets := make([]*model.Userset, 0, len(records))
    for _, record := range records {
        user := record.Tuple.User
        if user.Type == model.UserType_USER_SET {
            usersets = append(usersets, user.Userset)
        }
    }

    for _, userset := range usersets {
        subRecords, err := r.ChaseUsersets(*userset)
        if err != nil {
            // TODO wrap err
            return nil, err
        }
        records = append(records, subRecords...)
    }

    return records, nil
}

func NewChaser(repo TupleRepository) Chaser {
    return &ChaserImpl {
        repo: repo,
    }
}

var _ Chaser = (*ChaserImpl)(nil)
