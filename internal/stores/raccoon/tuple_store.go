// package raccon implements the store interfaces using raccoondb
package raccoon

import (
    "github.com/sourcenetwork/source-zanzibar/internal/stores"
    "github.com/sourcenetwork/source-zanzibar/internal/model/tuple"
    "github.com/sourcenetwork/source-zanzibar/lib/types"


    rcdb "github.com/sourcenetwork/raccoon"
    cosmos "github.com/cosmos/cosmos-sdk/store/types"
    "google.golang.org/protobuf/proto"
)

var _ stores.TupleStore = (*TupleStore)(nil)


var schema rcdb.RaccoonSchema = rcdb.RaccoonSchema{
    Indexes []SecondaryIndex[Edg, N]
    Store cosmos.KVStore
    KeysPrefix []byte
    Keyer NodeKeyer[N]
    Marshaler Marshaler[Edg]
}

type TupleStore[T proto.Message] struct {
    rc rcdb.RaccoonStore[tuple.TupleRecord, tuple.ObjRelRecord]
}


func (s *TupleStore[T]) SetTuple(tuple tuple.Tuple[T]) error {
    rec := tuple.ToRec()
    return s.rc.Set(rec)
}

func (s *TupleStore[T]) GetTuple(source tuple.ObjRel, dest tuple.ObjRel) (types.Option[tuple.Tuple[T]], error) {
    sourceRec := source.Rec()
    destRec := dest.Rec()

    rcOpt, err := s.rc.Get(sourceRec, destRec)
    if err != nil {
        return types.None(), err
    }

    opt := types.None()

    if !rcOpt.IsEmpty() {
        rec := rcOpt.Value()
        tuple := rec.ToTuple()
        opt = types.Some(tuple)
    }

    return opt, nil
}

func (s *TupleStore[T]) DeleteTuple(source tuple.ObjRel, dest tuple.ObjRel) error {
    rec := types.TupleRecord {
        ObjectRel: source.Rec(),
        Actor: dest.Rec(),
    }

    return s.rc.Delete(rec)
}

func (s *TupleStore[T]) GetSucessors(source tuple.ObjRel) ([]tuple.Tuple[T], error) {
    rec := source.ToRec()
    return s.rc.GetSucessors(rec)
}

func (s *TupleStore[T]) GetAncestors(source tuple.ObjRel) ([]tuple.Tuple[T], error) {
    rec := source.ToRec()
    return s.rc.GetAncestors(rec)
}

func (s *TupleStore[T]) GetTuplesFromRelationAndUserObject(relation string, objNamespace string, objectId string) ([]tuple.Tuple[T], error) {
    filter := func(rec tuple.TupleRecord) bool {
        isNamespace := tuple.Actor.Namespace == objNamespace
        isObj := tuple.Actor.Id == objectId
        isRel := tuple.ObjectRel.Relation == relation
        return isNamespace && isObj && isRel
    }

    return r.raccoon.FilterByIdx(relsIdx, relation, filter)
}
