package cli

import (
	"log"

	rcdb "github.com/sourcenetwork/raccoondb"

	zanzi "github.com/sourcenetwork/zanzi"
	"github.com/sourcenetwork/zanzi/types"
)

var client types.SimpleClient

// Initialize Zanzibar client from Fixture.
// The initialized store is a volatile in memory store
func initService(fixture Fixture) {
	store := rcdb.NewMemKV()
	tuplePrefix := []byte("/tuples")
	policyPrefix := []byte("/policy")

	client = zanzi.NewSimpleFromKVWithPrefixes(store, tuplePrefix, policyPrefix)

	policyService := client.GetPolicyService()
	for _, policy := range fixture.Policies {
		err := policyService.Set(policy)
		if err != nil {
			log.Fatalf("error initializing policy: %v", err)
		}
	}

	relationshipService := client.GetRelationshipService()

	for _, relationship := range fixture.Relationships {
		err := relationshipService.Set(relationship)
		if err != nil {
			log.Fatalf("error initializing relationship: %v", err)
		}
	}
}
