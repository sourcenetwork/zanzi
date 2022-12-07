package simple

import (
	"context"
	"testing"

	_ "github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/model/builder"
	"github.com/sourcenetwork/source-zanzibar/repository/maprepo"
	"github.com/sourcenetwork/source-zanzibar/tree"
)

func TestExpandOnAVirtualNode(t *testing.T) {

	// given test namespace where owner implies reader
	nr := maprepo.NewNamespaceRepo(
		builder.Namespace("test",
			builder.ThisRelation("owner"),
			builder.ThisRelation(""),
			builder.Relation("reader", builder.Union(builder.This(), builder.CU("owner"))),
		),
	)

	tb := builder.RelationshipBuilder{}
	tr := maprepo.NewTupleRepo(
		tb.ObjRel("test", "obj", "owner").Userset("test", "bob", "").Build(),
	)

	ctx := context.Background()
	expander := NewExpander(nr, tr)

	uset := rg.ObjRelNode{
		Namespace: "test",
		ObjectId:  "obj",
		Relation:  "reader",
	}

	got, err := expander.Expand(ctx, uset)

	assert.Nil(t, err)

	want := rg.RelationNode{
		ObjectRelNode uset,
		Child: &rg.OpNode{
			JoinOp: model.Operation_UNION,
			Left: &rg.RuleNode{
				Rule: rg.Rule{
					Type: rg.RuleType_THIS,
				},
				Children: []*rg.RelationNode{},
			},
			Right: &rg.RuleNode{
				Rule: rg.Rule{
					Type: rg.RuleType_CU,
					Args: map[string]string{
						"Relation": "owner",
					},
				},
				Children: []*rg.RelationNode{
					&rg.RelationNode{
						ObjectRelNode rg.ObjRelNode{
							Namespace: "test",
							ObjectId:  "obj",
							Relation:  "owner",
						},
						Child: &rg.RuleNode{
							Rule: rg.Rule{
								Type: rg.RuleType_THIS,
							},
							Children: []*rg.RelationNode{
								&rg.RelationNode{
									ObjectRelNode rg.ObjRelNode{
										Namespace: "test",
										ObjectId:  "bob",
										Relation:  "",
									},
									Child: &rg.RuleNode{
										Rule: rg.Rule{
											Type: rg.RuleType_THIS,
										},
										Children: []*rg.RelationNode{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	assert.Equal(t, got, want)
}
