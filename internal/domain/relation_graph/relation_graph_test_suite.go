package relation_graph

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	p "github.com/sourcenetwork/zanzi/internal/domain/policy"
	t "github.com/sourcenetwork/zanzi/internal/domain/tuple"
	"github.com/sourcenetwork/zanzi/internal/test_utils"
)

const partition = "1"

func getFixutres() (t.TupleBuilder, []t.Tuple, p.Policy) {
	tb := t.TupleBuilder{
		ActorNamespace: "user",
		Partition:      partition,
	}

	tuples := []t.Tuple{
		tb.Grant("group", "admin", "member", "alice"),
		tb.Grant("group", "staff", "member", "bob"),

		tb.Delegate("directory", "project", "owner", "group", "admin", "member"),
		tb.Delegate("file", "readme", "reader", "group", "staff", "member"),

		tb.Attribute("file", "foo", "parent", "directory", "project"),
		tb.Attribute("file", "readme", "parent", "directory", "project"),
	}

	policy := p.Policy{
		Id:   "1",
		Name: "Test policy",
		Resources: []*p.Resource{
			p.NewResource("file",
				p.ThisRelation("reader"),
				p.ThisRelation("writer"),
				p.ThisRelation("parent"),
				p.NewRelation("read", p.Union(p.CU("write"), p.CU("reader")), ""),
				p.NewRelation("write", p.Union(p.CU("writer"), p.TTU("parent", "directory", "owner")), ""),
			),
			p.NewResource("directory",
				p.ThisRelation("owner"),
			),
			p.NewResource("group",
				p.ThisRelation("member"),
			),
		},
		Actors: []*p.Actor{
			p.NewActor("user"),
		},
	}

	return tb, tuples, policy
}

func NewTestSuite(ts t.TupleStore, ps p.PolicyStore, rg RelationGraph) RelationGraphTestSuite {
	return RelationGraphTestSuite{
		ts: ts,
		ps: ps,
		rg: rg,
	}
}

type RelationGraphTestSuite struct {
	tb t.TupleBuilder
	ts t.TupleStore
	ps p.PolicyStore
	rg RelationGraph
}

func (s *RelationGraphTestSuite) Run(t *testing.T) {
	s.setup()
	test_utils.RunSuite(s, t)
}

func (s *RelationGraphTestSuite) setup() {
	tb, tuples, policy := getFixutres()
	s.tb = tb

	for _, tuple := range tuples {
		err := s.ts.SetTuple(tuple)
		if err != nil {
			panic(err)
		}
	}

	err := s.ps.SetPolicy(&policy)
	if err != nil {
		panic(err)
	}
}

func (s *RelationGraphTestSuite) TestWalk(t *testing.T) {
	ctx := context.Background()
	src := s.tb.RelSource("file", "readme", "read")

	tree, err := s.rg.Walk(ctx, partition, src)

	assert.Nil(t, err)
	assert.NotNil(t, tree)
	spew.Printf("%v", tree)
}

var _ test_utils.TestSuite = (*RelationGraphTestSuite)(nil)
