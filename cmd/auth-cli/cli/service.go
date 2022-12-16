package cli

import (
    "google.golang.org/protobuf/proto"
    rcdb "github.com/sourcenetwork/raccoondb"

    zanzi "github.com/sourcenetwork/source-zanzibar"
    tu "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    pol "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
)

var Service zanzi.Service[proto.Message]

func init() {
    tStore := tu.NewRaccoonStore[proto.Message](rcdb.NewMemKV(), nil)
    pStore := pol.NewPolicyKVStore(nil, rcdb.NewMemKV())

    policy := buildPolicy()
    pStore.SetPolicy(&policy)

    tuples := buildTuples()
    for _, tuple := range tuples {
        tStore.SetTuple(tuple)
    }
}

const POL_ID string = "1"
const ACTOR_NAMESPACE = "user"

func buildPolicy() pol.Policy {
    return pol.Policy{
        Id: "1",
        Name: "Test policy",
        Resources: []*pol.Resource{
        pol.BuildResource("file",
            pol.ThisRelation("owner"),
            pol.ThisRelation("reader"),
            pol.BuildPerm("read", pol.Union(pol.CU("write"), pol.CU("reader"))),
            pol.BuildPerm("write", pol.Union(pol.CU("owner"), pol.TTU("parent", "directory", "dir_owner"))),
        ),
        pol.BuildResource("directory",
            pol.ThisRelation("dir_owner"),
            pol.ThisRelation("parent"),
        ),
    },
    Actors: []*pol.Actor{
        pol.BuildActor(ACTOR_NAMESPACE),
    },
}
}

func buildTuples() []tu.Tuple[proto.Message] {
    tb := tu.TupleBuilder[proto.Message]{}
    tb.ActorNamespace = ACTOR_NAMESPACE
    tb.Partition = POL_ID

    return []tu.Tuple[proto.Message] {
        tb.Grant("group", "admin", "member", "alice"),
        tb.Grant("group", "staff", "member", "bob"),

        tb.Delegate("directory", "project", "owner", "group", "admin", "member"),
        tb.Delegate("file", "readme", "reader", "group", "staff", "member"),

        tb.Attribute("file", "foo", "parent", "directory", "project"),
        tb.Attribute("file", "readme", "parent", "directory", "project"),
    }
}
