// package raccon implements the store interfaces using raccoondb
package tuple

import (
    "crypto/sha256"

    rcdb "github.com/sourcenetwork/raccoondb"
    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/pkg/utils"
    opt "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

const (
    sha256bytes int = 256 / 8
)

const relsIdx string = "relations" // raccoon index name for source node relation

var _ TupleStore[proto.Message] = (*RCDBTupleStore[proto.Message])(nil)
var _ rcdb.NodeKeyer[*TupleNodeRecord] = (*tupleKeyer)(nil)
var _ rcdb.Edge[*TupleNodeRecord] = (*TupleRecord)(nil)

// tupleKeyer implements the Keyer interface as defined in raccoondb
// Map an TupleNodeRecord into []byte
// Uses a sha256 to generete the keys
type tupleKeyer struct {}

func (t *tupleKeyer) Key(objRel *TupleNodeRecord) []byte {
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


// relMapper returns the Source Node relation from a tuple record.
// Mapper used by racoons relation secondary index
func relMapper(rec *TupleRecord) []byte {
    return []byte(rec.Source.Relation)
}


// Return schema for racooon's tuple store
func buildSchema(kv rcdb.KVStore, prefix []byte) rcdb.RaccoonSchema[*TupleRecord, *TupleNodeRecord] {
    factory := func() *TupleRecord {
        return &TupleRecord{}
    }
    marshaler := rcdb.ProtoMarshaler[*TupleRecord](factory)
    return rcdb.RaccoonSchema[*TupleRecord, *TupleNodeRecord] {
        Indexes: []rcdb.SecondaryIndex[*TupleRecord, *TupleNodeRecord]{
            rcdb.SecondaryIndex[*TupleRecord, *TupleNodeRecord] {
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
    globalPrefix []byte
    kvStore rcdb.KVStore
    stores map[string]rcdb.RaccoonStore[*TupleRecord, *TupleNodeRecord]
}

// Return RCDBTupleStore from a raccoon KVStore and a global key prefix
func NewRaccoonStore[D proto.Message](kv rcdb.KVStore, prefix []byte) *RCDBTupleStore[D] {
    return &RCDBTupleStore[D] {
        globalPrefix: prefix,
        stores: make(map[string]rcdb.RaccoonStore[*TupleRecord, *TupleNodeRecord]),
        kvStore: kv,
    }
}

func (s *RCDBTupleStore[D]) getStore(partition string) rcdb.RaccoonStore[*TupleRecord, *TupleNodeRecord] {
    store, ok := s.stores[partition]
    if !ok {
        prefix := make([]byte, 0, len(partition) + len(s.globalPrefix) + 1)
        prefix = append(prefix, s.globalPrefix...)
        prefix = append(prefix, byte('/'))
        prefix = append(prefix, []byte(partition)...)
        schema := buildSchema(s.kvStore, prefix)
        store = rcdb.NewRaccoonStore[*TupleRecord, *TupleNodeRecord](schema)
        s.stores[partition] = store
    }
    return store
}


func (s *RCDBTupleStore[D]) SetTuple(tuple Tuple[D]) error {
    store := s.getStore(tuple.Partition)
    rec := tuple.ToRec()
    return store.Set(&rec)
}

func (s *RCDBTupleStore[D]) GetTuple(partition string, source TupleNode, dest TupleNode) (opt.Option[Tuple[D]], error) {
    store := s.getStore(partition)

    sourceRec := source.ToRec()
    destRec := dest.ToRec()
    rcOpt, err := store.Get(&sourceRec, &destRec)
    if err != nil || rcOpt.IsEmpty() {
        return opt.None[Tuple[D]](), err
    }

    tuple := toTuple[D](rcOpt.Value())
    return opt.Some[Tuple[D]](tuple), nil
}

func (s *RCDBTupleStore[D]) DeleteTuple(partition string, source TupleNode, dest TupleNode) error {
    store := s.getStore(partition)

    src := source.ToRec()
    dst := dest.ToRec()
    rec := TupleRecord {
        Source: &src,
        Dest: &dst,
    }
    return store.Delete(&rec)
}

func (s *RCDBTupleStore[D]) GetSucessors(partition string, source TupleNode) ([]Tuple[D], error) {
    store := s.getStore(partition)

    src := source.ToRec()
    records, err := store.GetSucessors(&src)
    return s.toTuples(records, err)
}

func (s *RCDBTupleStore[D]) GetAncestors(partition string, source TupleNode) ([]Tuple[D], error) {
    store := s.getStore(partition)

    src := source.ToRec()
    records, err := store.GetAncestors(&src)
    return s.toTuples(records, err)
}

func (s *RCDBTupleStore[D]) GetGrantingTuples(partition string, relation string, objNamespace string, objectId string) ([]Tuple[D], error) {
    store := s.getStore(partition)

    filter := func(rec *TupleRecord) bool {
        isNamespace := rec.Dest.Namespace == objNamespace
        isObj := rec.Dest.Id == objectId
        isRel := rec.Source.Relation == relation
        return isNamespace && isObj && isRel
    }

    records, err := store.FilterByIdx(relsIdx, []byte(relation), filter)
    return s.toTuples(records, err)
}

func (s *RCDBTupleStore[D]) toTuples(records []*TupleRecord, err error) ([]Tuple[D], error) {
    if err != nil {
        return nil, err
    }

    tuples := utils.MapSlice(records, func(r *TupleRecord) Tuple[D] {return toTuple[D](r)})
    return tuples, nil
}
