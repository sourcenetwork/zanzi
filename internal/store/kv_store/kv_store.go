package kv_store

import (
	rcdb "github.com/sourcenetwork/raccoondb"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/relation_graph"
	"github.com/sourcenetwork/zanzi/internal/store"
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

var _ store.Store = (*KVStore)(nil)

var (
	policyStorePrefix           []byte = []byte("1")
	relationshipStorePrefix     []byte = []byte("2")
	relationshipDataStorePrefix []byte = []byte("3")
	storeMetadataPrefix         []byte = []byte("4")
)

func NewKVStore(store rcdb.KVStore) (KVStore, error) {
	factory := func() *domain.PolicyRecord { return &domain.PolicyRecord{} }
	ider := &policyIDer{}
	protoMarshaler := rcdb.ProtoMarshaler[*domain.PolicyRecord](factory)
	policyPrefixWrapper := rcdb.NewWrapperKV(store, policyStorePrefix)
	policyStore := rcdb.NewObjStore[*domain.PolicyRecord](policyPrefixWrapper, nil, protoMarshaler, ider)

	return KVStore{
		baseKV:                   store,
		policyStore:              policyStore,
		relationshipPrefixKV:     rcdb.NewWrapperKV(store, relationshipStorePrefix),
		relationshipDataPrefixKV: rcdb.NewWrapperKV(store, relationshipDataStorePrefix),
	}, nil
}

type KVStore struct {
	baseKV                   rcdb.KVStore
	policyStore              rcdb.ObjectStore[*domain.PolicyRecord]
	relationshipPrefixKV     rcdb.KVStore
	relationshipDataPrefixKV rcdb.KVStore
}

func (kv *KVStore) getPolicyStore() *rcdb.ObjectStore[*domain.PolicyRecord] {
	return &kv.policyStore
}

func (kv *KVStore) getRelationshipStore(policyId string) rcdb.RaccoonStore[*Relationship, *RelationNode] {
	factory := func() *Relationship { return &Relationship{} }
	marshaler := rcdb.ProtoMarshaler[*Relationship](factory)
	store := rcdb.NewWrapperKV(kv.relationshipPrefixKV, []byte(policyId))
	schema := rcdb.RaccoonSchema[*Relationship, *RelationNode]{
		Store:     store,
		Keyer:     &relationNodeKeyer{},
		Marshaler: marshaler,
	}
	return rcdb.NewRaccoonStore(schema)
}

func (kv *KVStore) getRelationshipDataStore(policyId string) *rcdb.ObjectStore[*RelationshipData] {
	factory := func() *RelationshipData { return &RelationshipData{} }
	marshaler := rcdb.ProtoMarshaler[*RelationshipData](factory)
	ider := relationshipDataIDer{}
	store := rcdb.NewObjStore[*RelationshipData](kv.relationshipDataPrefixKV, nil, marshaler, &ider)
	return &store
}

func (kv *KVStore) GetPolicyRepository() policy.Repository {
	repo := newPolicyRepository(kv)
	return &repo
}

func (kv *KVStore) GetRelationNodeRepository() relation_graph.NodeRepository {
	repo := newNodeRepository(kv)
	return &repo
}
