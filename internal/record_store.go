package internal

/*
import (
    raccoon "github.com/sourcenetwork/raccoondb"

    "github.com/sourcenetwork/source-zanzibar/types"
    "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

type RecordStore[T any, PT types.ProtoConstraint[T, *T]] interface {
    Get(*types.Relationship) (option.Option[types.Record[T]], error)
    Set(*types.Record[T]) error
    Has(rel *types.Relationship) (bool, error)
    Delete(*types.Relationship) error
}

var _ raccoon.Ider[T any] = (*identifier)(nil)

type identifier struct {}

func NewKVRecordStore[T any, PT types.ProtoConstraint[*T]](prefix []byte, store raccoon.KVStore) RecordStore[T, PT] {
    factory := func() PT {return &T{}}
    protoMarshaler := raccoon.ProtoMarshaler[PT](factory)
    ider := &identifier{}
    store := raccoon.NewObjStore[PT](store, prefix, protoMarshaler, ider)
    return &kVRecordStore{
        store: store,
    }

}

type kVRecordStore[T any, PT types.ProtoConstraint[T, *T]] struct {
    store raccoon.ObjectStore[*Policy]
}
*/
