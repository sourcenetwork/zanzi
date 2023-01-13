package services

import (
	"fmt"

	"github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	o "github.com/sourcenetwork/source-zanzibar/pkg/option"
	"github.com/sourcenetwork/source-zanzibar/types"
	"github.com/sourcenetwork/source-zanzibar/internal/mappers"
)

var _ types.PolicyService = (*policyService)(nil)

// Return an PolicyService implementation from a PolicyStore
func PolicyServiceFromPolicyStore(store policy.PolicyStore) types.PolicyService {
	return &policyService{
		pStore: store,
	}
}

// policyService wraps a PolicyStore in order to implement PolicyService
type policyService struct {
	pStore policy.PolicyStore
        m mappers.PolicyMapper
}

func (s *policyService) Set(p types.Policy) error {
	mapped, err := s.m.ToInternal(p)
        if err != nil {
            return fmt.Errorf("error storing policy %v: %w", p.Name, err)
        }

	return s.pStore.SetPolicy(&mapped)
}

func (s *policyService) Get(id string) (o.Option[types.Policy], error) {
	polOpt, err := s.pStore.GetPolicy(id)
	if err != nil || polOpt.IsEmpty() {
		return o.None[types.Policy](), err
	}

	pol := s.m.FromInternal(polOpt.Value())

	return o.Some[types.Policy](pol), nil
}

func (s *policyService) Delete(id string) error {
	return s.pStore.DeletePolicy(id)
}
