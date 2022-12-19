package source_zanzibar

import (
    "context"
    "fmt"

    rcdb "github.com/sourcenetwork/raccoondb"

    "github.com/sourcenetwork/source-zanzibar/types"
    rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph/simple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    o "github.com/sourcenetwork/source-zanzibar/pkg/option"
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
        relationshipService: newRelationshipService(tStore),
        policyService: newPolicyService(pStore),
        authorizer: newAuthorizer(relGraph),
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


type recordClient[T any, PT types.ProtoConstraint[*T]] struct {
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


func NewRecordClient[T any, PT types.ProtoConstraint[*T]]() types.RecordClient[T, PT] {
    return nil
}
