package integration

import (
    "testing"
    "log"

	rcdb "github.com/sourcenetwork/raccoondb"

    "github.com/sourcenetwork/source-zanzibar/test"
    zanzi "github.com/sourcenetwork/source_zanzibar"
)

func setup() types.SimpleClient {
    kv := rcdb.NewMemKV()

    client := zanzi.NewSimpleFromKV(kv)

    policy, relationships := test.FilesystemFixture()
    
    relServ := client.GetRelationshipService()
    for _, relationship := range relationships {
        err := relServ.Set(relationship)
        if err != nil {
            log.Panicf("failed setting relationship: %v", err)
        }
    }

    polServ := client.GetPolicyService()
    err := polServ.Set(policy)
    if err != nil {
        log.Panicf("failed setting policy: %v", err)
    }
    return client
}
