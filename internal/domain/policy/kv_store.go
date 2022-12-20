package policy

import (
	raccoon "github.com/sourcenetwork/raccoondb"

	opt "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

// PolicyStore abstract interfacing with namespace storage.
type PolicyKVStore struct {
	store raccoon.ObjectStore[*Policy]
}

type polIder struct{}

func (id *polIder) Id(policy *Policy) []byte {
	return []byte(policy.Id)
}

var _ raccoon.Ider[*Policy] = (*polIder)(nil)

func NewPolicyKVStore(prefix []byte, store raccoon.KVStore) PolicyStore {
	factory := func() *Policy { return &Policy{} }
	protoMarshaler := raccoon.ProtoMarshaler[*Policy](factory)
	ider := &polIder{}
	objStore := raccoon.NewObjStore[*Policy](store, prefix, protoMarshaler, ider)
	return &PolicyKVStore{
		store: objStore,
	}
}

func (s *PolicyKVStore) GetPolicy(policyId string) (opt.Option[*Policy], error) {
	opt, err := s.store.GetObject([]byte(policyId))
	return toOpt(opt), err
}

func (s *PolicyKVStore) SetPolicy(policy *Policy) error {
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
	graph := NewMapPolicyGraph(*policy)
	return opt.Some(graph), nil
}

func toOpt[T any](ropt raccoon.Option[T]) opt.Option[T] {
	if ropt.IsEmpty() {
		return opt.None[T]()
	}
	return opt.Some[T](ropt.Value())
}
