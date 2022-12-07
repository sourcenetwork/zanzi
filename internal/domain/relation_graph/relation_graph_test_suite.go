package relation_graph

import (
    "context"
    "testing"

     "github.com/stretchr/testify/assert"

    t "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    p "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    "github.com/sourcenetwork/source-zanzibar/internal/test_utils"
)

const partition = ""



func getFixutres() (t.TupleBuilder[*test_utils.Appdata], []t.Tuple[*test_utils.Appdata], p.Policy) {
    tb := t.TupleBuilder[*test_utils.Appdata]{
        ActorNamespace: "user",
        Partition: partition,
    }

    tuples := []t.Tuple[*test_utils.Appdata] {
        tb.Grant("group", "admin", "member", "alice"),
        tb.Grant("group", "staff", "member", "bob"),

        tb.Delegate("directory", "project", "owner", "group", "admin", "member"),
        tb.Delegate("file", "readme", "reader", "group", "staff", "member"),

        tb.Attribute("file", "foo", "parent", "directory", "project"),
        tb.Attribute("file", "readme", "parent", "directory", "project"),


    }

    policy := p.Policy{
        Id: "1",
        Name: "Test policy",
        Resources: []*p.Resource{
            p.BuildResource("file",
                p.ThisRelation("reader"),
                p.ThisRelation("parent"),
                p.BuildPerm("read", p.Union(p.CU("write"), p.CU("reader"))),
                p.BuildPerm("write", p.Union(p.CU("owner"), p.TTU("parent", "directory", "dir_owner"))),
            ),
            p.BuildResource("directory",
                p.ThisRelation("owner"),
            ),
            p.BuildResource("group",
                p.ThisRelation("member"),
            ),

            p.BuildResource("group",
                p.ThisRelation("member"),
            ),
        },
        Actors: []*p.Actor{
            p.BuildActor("user"),
        },
    }

    return tb, tuples, policy
}


type RelationGraphTestSuite struct {
    tb t.TupleBuilder[*test_utils.Appdata]
}

func (s *RelationGraphTestSuite) Run(t *testing.T) {

}

func (s *RelationGraphTestSuite) testWalk(t *testing.T, rg RelationGraph[*test_utils.Appdata]) {
    ctx := context.Background()
    src := s.tb.RelSource("file", "readme", "read")

    tree, err := rg.Walk(ctx, partition, src)

    assert.Nil(err)
    assert.NotNil(tree)
}
