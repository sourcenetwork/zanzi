package simple

import (
	"testing"

	rcdb "github.com/sourcenetwork/raccoondb"

	p "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
	t "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
)

func TestSimpleRelationGraph(test *testing.T) {
	tKv := rcdb.NewMemKV()
	pKv := rcdb.NewMemKV()

	pStore := p.NewPolicyKVStore(nil, pKv)
	tStore := t.NewRaccoonStore(tKv, nil)

	simple := NewSimple(tStore, pStore)

	suite := rg.NewTestSuite(tStore, pStore, simple)
	suite.Run(test)
}
