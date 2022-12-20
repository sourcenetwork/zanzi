package source_zanzibar

import (
    rcdb "github.com/sourcenetwork/raccoondb"

    "github.com/sourcenetwork/source-zanzibar/types"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph/simple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    "github.com/sourcenetwork/source-zanzibar/internal/services"
)

//var _ types.RecordService[any, proto.Message] = (*relService[any, proto.Message])(nil)
var _ types.SimpleClient = (*simpleClient)(nil)

const (
    tuplePrefix string = "/tuples"
    policyPrefix string = "/policy"
)

func NewSimpleFromKV(store rcdb.KVStore) types.SimpleClient {
    return NewSimpleFromKVWithPrefixes(store, []byte(tuplePrefix), []byte(policyPrefix))
}

func NewSimpleFromKVWithPrefixes(store rcdb.KVStore, tuplesPrefix []byte, policyPrefix []byte) types.SimpleClient {
    pStore := policy.NewPolicyKVStore(policyPrefix, store)
    tStore := tuple.NewRaccoonStore(store, tuplesPrefix)
    relGraph := simple.NewSimple(tStore, pStore)

    return &simpleClient {
        relationshipService: services.RelationshipServiceFromTupleStore(tStore),
        policyService: services.PolicyServiceFromPolicyStore(pStore),
        authorizer: services.AuthorizerFromRelationGraph(relGraph),
    }
}

type simpleClient struct {
    relationshipService types.RelationshipService
    policyService types.PolicyService
    authorizer types.Authorizer
}

func (s *simpleClient) GetAuthorizer() types.Authorizer {
    return s.authorizer
}

func (s *simpleClient) GetPolicyService() types.PolicyService {
    return s.policyService
}

func (s *simpleClient) GetRelationshipService() types.RelationshipService {
    return s.relationshipService
}


type recordClient[T any, PT types.ProtoConstraint[T]] struct {
    recordService types.RecordService[T, PT]
    policyService types.PolicyService
    authorizer types.Authorizer
}

func (s *recordClient[T, PT]) GetAuthorizer() types.Authorizer {
    return s.authorizer
}

func (s *recordClient[T, PT]) GetPolicyService() types.PolicyService {
    return s.policyService
}

func (s *recordClient[T, PT]) GetRecordService() types.RecordService[T, PT] {
    return s.recordService
}


func NewRecordClient[T any, PT types.ProtoConstraint[T]](store rcdb.KVStore, recordPrefix []byte, relationshipPrefix []byte, policyPrefix []byte) types.RecordClient[T, PT] {
    pStore := policy.NewPolicyKVStore(policyPrefix, store)
    tStore := tuple.NewRaccoonStore(store, relationshipPrefix)
    relGraph := simple.NewSimple(tStore, pStore)

    return &recordClient[T, PT] {
        recordService: services.RecordServiceFromStores[T, PT](store, recordPrefix, tStore),
        policyService: services.PolicyServiceFromPolicyStore(pStore),
        authorizer: services.AuthorizerFromRelationGraph(relGraph),
    }
}
