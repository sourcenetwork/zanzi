package simple

import (
    "testing"

    "github.com/cosmos/cosmos-sdk/store/mem"

    t "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    p "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    "github.com/sourcenetwork/source-zanzibar/internal/test_utils"
    rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
)

func TestSimpleRelationGraph(test *testing.T) {
    tKv := mem.NewStore()
    pKv := mem.NewStore()

    pStore := p.NewPolicyKVStore(nil, pKv)
    tStore := t.NewRaccoonStore[*test_utils.Appdata](tKv, nil)

    simple := NewSimple[*test_utils.Appdata](tStore, pStore)

    suite := rg.NewTestSuite(tStore, pStore, simple)
    suite.Run(test)
}
