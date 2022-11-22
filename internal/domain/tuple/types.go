import tuple

import (
    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/internal/tuples"
    "github.com/sourcenetwork/source-zanzibar/types"
)

type TupleStore[T proto.Message] interface {
    // Store a new tuple
    SetTuple(tuple tuples.Tuple[T]) error

    // Looks up a relation tuple from the backend storage
    GetTuple(source tuples.ObjRel, dest tuples.ObjRel) (types.Option[tuples.Tuple[T]], error)

    // Purge tuple from storage.
    DeleteTuple(source tuples.ObjRel, dest tuples.ObjRel) error

    // maybe separate these
    GetSucessors(source tuples.ObjRel) ([]tuples.Tuple[T], error)
    GetAncestors(source tuples.ObjRel) ([]tuples.Tuple[T], error)
    
    // Return all tuples that match the filter:
    // tuple.Relation == relation
    // tuple.User.ObjectId == objectId
    //
    // This peculiar query is used to reverse tuple to userset rules
    // during reverse lookup
    GetTuplesFromRelationAndUserObject(relation string, objNamespace string, objectId string) ([]tuples.Tuple[T], error)
    // FIXME rename this method
    // TODO need some method to deal with relations
}
