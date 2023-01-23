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
    rb.Perm("read", "write - reader")
    rb.Perm("write", "owner + parent->dir_owner")
    file := rb.Build()

    rb.Name("group")
    rb.Relations("member")
    group := rb.Build()

    rb.Name("directory")
    rb.Relations("dir_owner")
    directory := rb.Build()

    pb := types.PolicyBuilder()
    pb.IdNameDescription(POLICY_ID,  "Filesystem Policy", "Auth CLI example filesystem policy")
    pb.Actors(types.NewActor(ACTOR_NAMESPACE))
    pb.Resources(file, group, directory)
    pb.Attr("author", "Source Network")
    pol := pb.Build()

    return pol
}

func buildRelationships() []types.Relationship {
	b := types.RelationshipBuilder(POLICY_ID)

	return []types.Relationship{
            b.Grant("group", "admin", "member", ACTOR_NAMESPACE, "alice"),
            b.Grant("group", "staff", "member", ACTOR_NAMESPACE, "bob"),

            b.Delegate("directory", "project", "dir_owner", "group", "admin", "member"),
            b.Delegate("file", "readme", "reader", "group", "staff", "member"),

            b.Attribute("file", "foo", "parent", "directory", "project"),
            b.Attribute("file", "readme", "parent", "directory", "project"),
	}
}
