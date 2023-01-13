package cli

import (
	rcdb "github.com/sourcenetwork/raccoondb"

	zanzi "github.com/sourcenetwork/source-zanzibar"
	"github.com/sourcenetwork/source-zanzibar/types"
)

var client types.SimpleClient

func init() {
	store := rcdb.NewMemKV()
	tuplePrefix := []byte("/tuples")
	policyPrefix := []byte("/policy")

	client = zanzi.NewSimpleFromKVWithPrefixes(store, tuplePrefix, policyPrefix)

        policyService := client.GetPolicyService()
        err := policyService.Set(buildPolicy())
        if err != nil {
            panic(err)
        }

        relationshipService := client.GetRelationshipService()

	for _, relationship := range buildRelationships() {
		err = relationshipService.Set(relationship)
                if err != nil {
                    panic(err)
                }
	}
}

const POLICY_ID string = "1"
const ACTOR_NAMESPACE = "user"

func buildPolicy() types.Policy {
    rb := types.ResourceBuilder{}
    rb.Name("file")
    rb.Relations("owner", "reader", "parent")
    rb.Perm("read", "write+reader")
    rb.Perm("write", "owner + parent->dir_owner")
    file := rb.Build()

    rb.Name("group")
    rb.Relations("member")
    group := rb.Build()

    rb.Name("directory")
    rb.Relations("dir_owner")
    directory := rb.Build()

    pb := types.PolicyBuilder()
    pb.IdNameDescription(POLICY_ID,  "Policy", "Description")
    pb.Actors(types.NewActor(ACTOR_NAMESPACE))
    pb.Resources(file, group, directory)
    pb.Attr("author", "source")
    pol := pb.Build()

    return pol
}

func buildRelationships() []types.Relationship {
	b := types.RelationshipBuilder(POLICY_ID)

	return []types.Relationship{
            b.Grant(b.En("group", "admin"), "member", b.En(ACTOR_NAMESPACE, "alice")),
            b.Grant(b.En("group", "staff"), "member", b.En(ACTOR_NAMESPACE, "bob")),

            b.Delegate(b.En("directory", "project"), "dir_owner", b.En("group", "admin"), "member"),
            b.Delegate(b.En("file", "readme"), "reader", b.En("group", "staff"), "member"),

            b.Attribute(b.En("file", "foo"), "parent", b.En("directory", "project")),
            b.Attribute(b.En("file", "readme"), "parent", b.En("directory", "project")),
	}
}
