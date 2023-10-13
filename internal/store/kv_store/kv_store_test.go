package kv_store

import (
	"testing"

	rcdb "github.com/sourcenetwork/raccoondb"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/store"
	_testing "github.com/sourcenetwork/zanzi/internal/testing"
)

func factory() policy.Repository {
	memKV := rcdb.NewMemKV()
	kv, err := NewKVStore(memKV)
	if err != nil {
		panic(err)
	}
	return kv.GetPolicyRepository()
}

func TestKVStore(t *testing.T) {
	suite := store.NewPolicyRepositoryTestSuite(factory)
	_testing.RunSuite(t, &suite)
}
