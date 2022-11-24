package policy

import (
    cosmos "github.com/cosmos/cosmos-sdk/store/types"
    raccoon "github.com/sourcenetwork/raccoondb"

    opt "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

// PolicyStore abstract interfacing with namespace storage.
type PolicyKVStore struct {
    store raccoon.ObjectStore[Policy]
}

func NewPolicyKVStore(prefix []byte, store cosmos.KVStore) PolicyStore {
    return &PolicyKVStore {

    }
}

func (s *PolicyKVStore) GetPolicy(policyId string) (opt.Option[Policy], error) {
    opt, err := s.store.GetObject([]byte(policyId))
    return toOpt(opt), err
}

func (s *PolicyKVStore) SetPolicy(policy Policy) error {
    return s.store.SetObject(policy)
}

func (s *PolicyKVStore) DeletePolicy(policyId string) error {
    return s.store.DeleteById([]byte(policyId))
}

func (s *PolicyKVStore) GetPolicyGraph(policyId string) (opt.Option[PolicyGraph], error) {
    ropt, err := s.store.GetObject([]byte(policyId))
    o := toOpt(ropt)
    if err != nil || o.IsEmpty() {
        return opt.None[PolicyGraph](), err
    }

    policy := o.Value()
    graph := NewMapPolicyGraph(policy)
    return opt.Some(graph), nil
}

func toOpt[T any](ropt raccoon.Option[T]) opt.Option[T] {
    if ropt.IsEmpty() {
        return opt.None[T]()
    }
    return opt.Some[T](ropt.Value())
}
