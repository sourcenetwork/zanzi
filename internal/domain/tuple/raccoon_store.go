// package raccon implements the store interfaces using raccoondb
package tuple

import (
    "crypto/sha256"

    rcdb "github.com/sourcenetwork/raccoondb"
    cosmos "github.com/cosmos/cosmos-sdk/store/types"
    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/pkg/utils"
    opt "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

const (
    sha256bytes int = 256 / 8
)

const relsIdx string = "relations" // raccoon index name for source node relation

var _ TupleStore[proto.Message] = (*RCDBTupleStore[proto.Message])(nil)
var _ rcdb.NodeKeyer[*ObjRelRecord] = (*tupleKeyer)(nil)
var _ rcdb.Edge[*ObjRelRecord] = (*TupleRecord)(nil)

// tupleKeyer implements the Keyer interface as defined in raccoondb
// Map an ObjRelRecord into []byte
// Uses a sha256 to generete the keys
type tupleKeyer struct {}

func (t *tupleKeyer) Key(objRel *ObjRelRecord) []byte {
    h := sha256.New()
    h.Write([]byte(objRel.Namespace))
    h.Write([]byte(objRel.Id))
    h.Write([]byte(objRel.Relation))
    return h.Sum(nil)
}

// Min possible key is a slice of sha256bytes min byte val (0)
func (t *tupleKeyer) MinKey() []byte {
    k := make([]byte, sha256bytes)
    return k
}

// Min possible key is a slice of sha256bytes max byte val (255)
func (t *tupleKeyer) MaxKey() []byte {
    k := make([]byte, sha256bytes)
    for i := range k {
        k[i] = byte(255)
    }
    return k
}


// GetSource implements Edge interface from RaccoonDb
func (r TupleRecord) GetSource() *ObjRelRecord {
    return r.ObjectRel
}

// GestDest implements Edge interface from RaccoonDb
func (r TupleRecord) GetDest() *ObjRelRecord {
    return r.Actor
}


// relMapper returns the Source Node relation from a tuple record.
// Mapper used by racoons relation secondary index
func relMapper(rec *TupleRecord) []byte {
    return []byte(rec.ObjectRel.Relation)
}


// Return schema for racooon's tuple store
func buildSchema(kv cosmos.KVStore, prefix []byte) rcdb.RaccoonSchema[*TupleRecord, *ObjRelRecord] {
    factory := func() *TupleRecord {
        return &TupleRecord{}
    }
    marshaler := rcdb.ProtoMarshaler[*TupleRecord](factory)
    return rcdb.RaccoonSchema[*TupleRecord, *ObjRelRecord] {
        Indexes: []rcdb.SecondaryIndex[*TupleRecord, *ObjRelRecord]{
            rcdb.SecondaryIndex[*TupleRecord, *ObjRelRecord] {
                Name: relsIdx,
                Mapper: relMapper,
            },
        },
        Store: kv,
        KeysPrefix: prefix,
        Keyer: &tupleKeyer{},
        Marshaler: marshaler,
    }
}


// RCDBTupleStore implements the TupleStore interface
// using RaccoonDB as the backend storage engine.
type RCDBTupleStore[D proto.Message] struct {
    rc rcdb.RaccoonStore[*TupleRecord, *ObjRelRecord]
}

// Return RCDBTupleStore from a cosmos KVStore and a global key prefix
func NewRaccoonStore[D proto.Message](kv cosmos.KVStore, prefix []byte) RCDBTupleStore[D] {
    schema := buildSchema(kv, prefix)
    raccoon := rcdb.NewRaccoonStore[*TupleRecord, *ObjRelRecord](schema)
    return RCDBTupleStore[D] {
        rc: raccoon,
    }
}


func (s *RCDBTupleStore[D]) SetTuple(tuple Tuple[D]) error {
    rec := tuple.ToRec()
    return s.rc.Set(&rec)
}

func (s *RCDBTupleStore[D]) GetTuple(source ObjRel, dest ObjRel) (opt.Option[Tuple[D]], error) {
    sourceRec := source.ToRec()
    destRec := dest.ToRec()

    rcOpt, err := s.rc.Get(&sourceRec, &destRec)
    if err != nil || rcOpt.IsEmpty() {
        return opt.None[Tuple[D]](), err
    }

    tuple := toTuple[D](rcOpt.Value())
    return opt.Some[Tuple[D]](tuple), nil
}

func (s *RCDBTupleStore[D]) DeleteTuple(source ObjRel, dest ObjRel) error {
    src := source.ToRec()
    dst := dest.ToRec()
    rec := TupleRecord {
        ObjectRel: &src,
        Actor: &dst,
    }
    return s.rc.Delete(&rec)
}

func (s *RCDBTupleStore[D]) GetSucessors(source ObjRel) ([]Tuple[D], error) {
    src := source.ToRec()
    records, err := s.rc.GetSucessors(&src)
    return s.toTuples(records, err)
}

func (s *RCDBTupleStore[D]) GetAncestors(source ObjRel) ([]Tuple[D], error) {
    src := source.ToRec()
    records, err := s.rc.GetAncestors(&src)
    return s.toTuples(records, err)
}

func (s *RCDBTupleStore[D]) GetTuplesFromRelationAndUserObject(relation string, objNamespace string, objectId string) ([]Tuple[D], error) {
    filter := func(rec *TupleRecord) bool {
        isNamespace := rec.Actor.Namespace == objNamespace
        isObj := rec.Actor.Id == objectId
        isRel := rec.ObjectRel.Relation == relation
        return isNamespace && isObj && isRel
    }

    records, err := s.rc.FilterByIdx(relsIdx, []byte(relation), filter)
    return s.toTuples(records, err)
}

func (s *RCDBTupleStore[D]) toTuples(records []*TupleRecord, err error) ([]Tuple[D], error) {
    if err != nil {
        return nil, err
    }

    tuples := utils.MapSlice(records, func(r *TupleRecord) Tuple[D] {return toTuple[D](r)})
    return tuples, nil
}
