package source_zanzibar

import (
    "google.golang.org/protobuf/proto"
    cosmos "github.com/cosmos/cosmos-sdk/store/types"

    "github.com/sourcenetwork/source-zanzibar/types"
)

// Exposes Zanzibar-like functionality to client applications
type Service[T proto.Message] struct {
    relationService types.RelationService[T]
    policyService types.PolicyService
    authorizer types.Authorizer
}

func (s *Service[T]) GetAuthorizer() types.Authorizer {
    return s.authorizer
}

func (s *Service[T]) GetPolicyService() types.PolicyService {
    return s.policyService
}

func (s *Service[T]) GetRelationService() types.RelationService[T] {
    return s.relationService
}

// Build a Zanzibar service from a cosmos-sdk kv store
func InitFromCosmosKV[T proto.Message](
    policyStore cosmos.KVStore,
    policyPrefix string,
    relationStore cosmos.KVStore,
    relationPrefix string) Service[T] {
    return Service {
    }
}
