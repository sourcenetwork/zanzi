package kv_store

import (
	"testing"

	rcdb "github.com/sourcenetwork/raccoondb"
	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/store"
	_testing "github.com/sourcenetwork/zanzi/internal/testing"
	"github.com/sourcenetwork/zanzi/pkg/domain"
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

func TestKVStore_RaccoonMarshaler_Deterministic(t *testing.T) {
	factory := func() *domain.PolicyRecord {
		return &domain.PolicyRecord{}
	}
	marshaler := rcdb.ProtoMarshaler[*domain.PolicyRecord](factory)

	recFactory := func() *domain.PolicyRecord {
		return &domain.PolicyRecord{
			Policy: &domain.Policy{
				Attributes: map[string]string{
					"att1": "att1",
					"att2": "att2",
				},
			},
		}
	}

	rec := recFactory()
	baseline, err := marshaler.Marshal(&rec)
	require.NoError(t, err)

	iterations := 200
	for i := 0; i < iterations; i++ {
		r := recFactory()
		got, err := marshaler.Marshal(&r)
		require.NoError(t, err)
		require.Equal(t, baseline, got, "marshaling diverged at iteration %d", i)
	}
}
