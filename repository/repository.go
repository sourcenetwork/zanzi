// Package repository provides an interface to interact with the stored relation tuples.
//
// Internally repository manages database indexes, keeps indexes and provides queries.
package repository

// Repository leverages the "tm-db" DB interface as a storage backend abstraction.
// This choice was tentatively made to ease future integration with a dApp powered by Cosmos SDK,
// potentially through [github.com/cosmos/cosmos-sdk/store/types.CommitMultiStore.MountStoreWithDB]

import (
    tmdb "github.com/tendermint/tm-db"

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
    GetTuple(namespace string, object string, relation string, userNamespace string, uid string) (model.TupleRecord, error)

    // Return tuples related to given userset (specified through namespace, objectId and relation)
    // Does not follows userset indirection.
    GetTuplesFromObjRel(namespace string, objectId string, relation string) ([]model.TupleRecord, error)

    // Return tuples where their users are specified by the given userset
    // aka reverse lookup
    // Does not perform userset chasing
    GetTuplesFromUserset(namespace string, id string, relation string) ([]model.TupleRecord, error)

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
type RepositoryChaser interface {

    // Recursively fetches tuples matching namespace, objectId and relation
    ChaseUsersets(userset model.Userset) ([]model.TupleRecord, error)

    // Perform an inverse lookup of all tuples whose user match the given userset.
    // ie. Starting from an user node, walks up the tuple graph by following usersets
    ReverseChaseUserset(namespace, id, relation string) ([]model.TupleRecord, error)
}

/*
type TMDBRepository struct {
    // NOTE Single DB for all indexes? hashicorps mem-db has a schema feature for a single db.
    // The tendermint abstraction provides nothing of the sort, forcing partition to be done through keys.
    // Should we use multiple instances of a db? I am not sure of the tradeoffs.
    Db tmdb.DB
}

// Minor c# idiom until i figure out what to do with repochaser
type RepositoryChaserImpl struct {
    repo Repository
}

func (r *RepositoryChaserImpl) ChaseUsersets(userset model.Userset) ([]model.TupleRecord, error) {
    records, err := r.repo.GetTuplesFromUserset(userset.namespace, userset.objectId, userset.relation)
    if err != nil {
        // TODO wrap err
        return nil, err
    }

    usersets := make(model.User, 0, len(records))
    for record := range records {
        user := record.Tuple.User
        if user.Type == model.UserType_USER_SET {
            append(usersets, user)
        }
    }

    for userset := range usersets {
        subRecords, err := chaseUsersets(ctx, userset.Namespace, userset.Identifier, userset.Relation)
        if err != nil {
            // TODO wrap err
            return nil, err
        }
        records = append(records, subRecords)
    }

    return records, nil
}
*/
var _ *tmdb.DB = nil
