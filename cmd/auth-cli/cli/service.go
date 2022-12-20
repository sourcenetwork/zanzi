package cli

import (
	rcdb "github.com/sourcenetwork/raccoondb"

	zanzi "github.com/sourcenetwork/source-zanzibar"
	pol "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	tu "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
	"github.com/sourcenetwork/source-zanzibar/types"
)

var Client types.SimpleClient

func init() {
	store := rcdb.NewMemKV()
	tuplePrefix := []byte("/tuples")
	policyPrefix := []byte("/policy")
	tStore := tu.NewRaccoonStore(store, tuplePrefix)
	pStore := pol.NewPolicyKVStore(policyPrefix, store)

	// FIXME this should ideally use the *public* types
	// such relationship and the public policy.
	// Since policy has no parser yet, we are using the
	// internal types for now
	policy := buildPolicy()
	pStore.SetPolicy(&policy)

	tuples := buildTuples()
	for _, tuple := range tuples {
		tStore.SetTuple(tuple)
	}
	Client = zanzi.NewSimpleFromKVWithPrefixes(store, tuplePrefix, policyPrefix)
}

const POL_ID string = "1"
const ACTOR_NAMESPACE = "user"

func buildPolicy() pol.Policy {
	return pol.Policy{
		Id:   "1",
		Name: "Test policy",
		Resources: []*pol.Resource{
			pol.BuildResource("file",
				pol.ThisRelation("owner"),
				pol.ThisRelation("reader"),
				pol.ThisRelation("parent"),
				pol.BuildPerm("read", pol.Union(pol.CU("write"), pol.CU("reader"))),
				pol.BuildPerm("write", pol.Union(pol.CU("owner"), pol.TTU("parent", "directory", "dir_owner"))),
			),
			pol.BuildResource("directory",
				pol.ThisRelation("dir_owner"),
			),
			pol.BuildResource("group",
				pol.ThisRelation("member"),
			),
		},
		Actors: []*pol.Actor{
			pol.BuildActor(ACTOR_NAMESPACE),
		},
	}
}

func buildTuples() []tu.Tuple {
	tb := tu.TupleBuilder{}
	tb.ActorNamespace = ACTOR_NAMESPACE
	tb.Partition = POL_ID

	return []tu.Tuple{
		tb.Grant("group", "admin", "member", "alice"),
		tb.Grant("group", "staff", "member", "bob"),

		tb.Delegate("directory", "project", "dir_owner", "group", "admin", "member"),
		tb.Delegate("file", "readme", "reader", "group", "staff", "member"),

		tb.Attribute("file", "foo", "parent", "directory", "project"),
		tb.Attribute("file", "readme", "parent", "directory", "project"),
	}
}
