package services

import (
    _ "google.golang.org/protobuf/proto"
    "crypto/sha256"
    "log"

    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
    "github.com/sourcenetwork/source-zanzibar/types"
    "github.com/sourcenetwork/source-zanzibar/pkg/utils"
)

var _ types.RecordService = (*recordService)(nil)

func RecordServiceFromStores[T any, PT types.ProtoConstraint[*T]](kv rcdb.KVStore, recordDataPrefix []byte, tupleStore tuple.TupleStore) types.RecordService {
    marshaler := rcdb.ProtoMarshaler[T](func() PT {return new(T)})

    objKV := rcdb.NewObjKV[T](kv, recordDataPrefix, marshaler)

    return &recordService {
        tuples: tupleStore,
        objKV: objKV,
    }
}

// recordService provides a RecordService implementation
// Records are broken up and the relationship is stored in a tuple storage
// while the satellite data is stored in a raccoon ObjKV store instance
// FIXME All of this *must* be made atomic
type recordService[T any, PT types.ProtoConstraint[*T]] struct {
    tuples: tuple.TupleStore
    objKV: rcdb.ObjKV[T]
    ider rcdb.Ider[tuple.Tuple]
    mapper RelationshipMapper
}

func (s *recordService[T, PT]) Set(rel types.Relationship, data T) error {
    tuple := s.mapper.FromRelationship(rel)
    key := s.ider.Id(tuple)

    err := s.objKv.Set(key, data)
    if err != nil {
        return fmt.Errorf("failed setting record data for rel %v: %w", rel, err)
    }

    err = s.tuples.SetTuple(tuple)
    if err != nil {
        // FIXME this should have a rollback policy
        return fmt.Errorf("failed setting relationship %v: %w", rel, err)
    }

    return nil
}

func (s *recordService[T, PT]) Delete(rel types.Relationship) error {
    tuple := s.mapper.FromRelationship(rel)
    key := s.ider.Id(tuple)

    err := s.objKv.Delete(key)
    if err != nil {
        return fmt.Errorf("failed deleting record data for rel %v: %w", rel, err)
    }

    err = s.tuples.DeleteTuple(tuple)
    if err != nil {
        // FIXME this should have a rollback policy
        return fmt.Errorf("failed deleting relationship %v: %w", rel, err)
    }

    return nil
}

func (s *recordService[T, PT]) Get(rel types.Relationship) (o.Option[Record[T]], error) {
    tuple := s.mapper.FromRelationship(rel)
    key := s.ider.Id(tuple)

    // FIXME this should have a lock or version check in order to
    // fetch the matching record data and tuple

    dataOpt, err := s.objKv.Get(key)
    if err != nil {
        return o.None[Record[T]](), fmt.Errorf("failed fetching record data for rel %v: %w", rel, err)
    }

    tupleOpt, err = s.tuples.GetTuple(tuple)
    if err != nil {
        return o.None[Record[T]](), fmt.Errorf("failed deleting relationship %v: %w", rel, err)
    }

    if tupleOpt.IsEmpty() == dataOpt.IsEmpty() {
        return o.None[Record[T]](), nil
    } 
    if tupleOpt.IsEmpty() && !dataOpt.IsEmpty() {
        // NOTE this is bad because it could be a sign of data corruption / failed transaction
        msg := fmt.Sprintf("failed to get record: inconsistency detected: relationship %v does not exist while associated data does", rel)
        log.Println(msg)
        return o.None[Record[T]](), fmt.Errorf(msg)
    }
    // TODO think over the else case
    // is it ok for a record to have a relationship but no data?

    record := types.Record {
        Relationship: s.mapper.ToRelationship(tuple),
        Data: data,
    }
    return o.Some[Record[T]](record), nil
}

func (s *recordService[T, PT]) Has(rel Relationship) (bool, error) {
    tuple := s.mapper.FromRelationship(rel)
    opt, err = s.tuples.GetTuple(tuple)
    if err != nil {
        return false, fmt.Errorf("failed fetching relationship %v: %w", rel, err)
    }
    return !opt.IsEmpty(), nil
}

// tupleKeyer implements the NodeKeyer interface as defined in raccoondb
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
