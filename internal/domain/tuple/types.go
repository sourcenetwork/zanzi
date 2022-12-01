package tuple

import (
    "google.golang.org/protobuf/proto"

    opt "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

type TupleStore[T proto.Message] interface {
    // Store a new tuple
    SetTuple(tuple Tuple[T]) error

    // Looks up a relation tuple from the backend storage
    GetTuple(source ObjRel, dest ObjRel) (opt.Option[Tuple[T]], error)

    // Purge tuple from storage.
    DeleteTuple(source ObjRel, dest ObjRel) error

    // maybe separate these
    GetSucessors(source ObjRel) ([]Tuple[T], error)
    GetAncestors(source ObjRel) ([]Tuple[T], error)
    
    // Return all tuples that match the filter:
    // tuple.ObjRel.Relation == relation
    // tuple.User.ObjectId == objectId
    //
    // This peculiar query is used to reverse tuple to userset rules
    // during reverse lookup
    GetTuplesFromRelationAndUserObject(relation string, objNamespace string, objectId string) ([]Tuple[T], error)
    // FIXME rename this method
    // TODO need some method to deal with relations
}
