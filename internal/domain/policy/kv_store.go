package policy

import (
    cosmos "github.com/cosmos/cosmos-sdk/store/types"
)

// PolicyStore abstract interfacing with namespace storage.
type PolicyKVStore struct {
    store raccoon.ObjectStore[Policy]
}

func NewPolicyKVStore(prefix []byte, store cosmos.KVStore) PolicyStore {
    return &PolicyKVStore {

    }
}

func (s *PolicyKVStore) GetPolicy(policyId string) (utils.Option[Policy], error) {
    return s.store.Get([]byte(policyId))
}

func (s *PolicyKVStore) SetPolicy(policy Policy) error {
    return s.store.Set(policy)
}

func (s *PolicyKVStore) DeletePolicy(policyId string) error {
    return s.store.DeleteById(policy)
}

func (s *PolicyKVStore) GetPolicyGraph(policyId string) (utils.Option[PolicyGraph], error) {
    opt, err := s.store.Get([]byte(policyId))
    if err != nil || opt.IsEmpty() {
        return ultils.None(), err
    }

    policy := opt.Value()
    graph := NewMapPolicyGraph(policy)
    return graph
}
