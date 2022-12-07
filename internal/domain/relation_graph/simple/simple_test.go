package simple

import (
    "context"
    "testing"

     "github.com/stretchr/testify/assert"
    "github.com/cosmos/cosmos-sdk/store/mem"

    t "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    p "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    "github.com/sourcenetwork/source-zanzibar/internal/test_utils"
    rg "github.com/sourcenetwork/source-zanzibar/internal/relation_graph"
)

func TestSimpleRelationGraph(t *testing.T) {
    tKv := mem.NewStore()
    pKv := mem.NewStore()

    pStore := p.NewPolicyKVStore(nil, pKv)
    tStore := t.NewRaccoonStore[*test_utils.Appdata](tKv, nil)

    suite := rg.NewSuite[*test_utils.Appdata](pStore, tStore)
    suite.Run(t)
}
