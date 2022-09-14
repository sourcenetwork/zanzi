package simple

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
        "github.com/davecgh/go-spew/spew"


	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/tree"
	"github.com/sourcenetwork/source-zanzibar/repository/maprepo"
	"github.com/sourcenetwork/source-zanzibar/utils"
	"github.com/sourcenetwork/source-zanzibar/model/builder"
)

func TestExpandOnAVirtualNode(t *testing.T) {

    // given test namespace where owner implies reader
    ns := []model.Namespace{
        *builder.Namespace("test", 
            builder.ThisRelation("owner"),
            builder.ThisRelation(""),
            builder.Relation("reader", builder.Union(builder.This(), builder.CU("owner"))),
        ),
    }

    tb := builder.TupleBuilder{}
    tuples := []model.Tuple {
        tb.ObjRel("test", "obj", "owner").Userset("test", "bob", "").Build(),
    }
    

    expander := NewExpander()

    uset := model.Userset{
        Namespace: "test",
        ObjectId: "obj",
        Relation: "reader",
    }
    ctx := setupCtx(tuples, ns)
    got, err := expander.Expand(ctx, uset)

    assert.Nil(t, err)

    want := tree.UsersetNode{
        Userset: uset,
        Child: &tree.OpNode{
            JoinOp: model.Operation_UNION,
            Left: &tree.RuleNode{
                Rule: tree.Rule{
                    Type: tree.RuleType_THIS,
                },
                Children: []*tree.UsersetNode{},
            },
            Right: &tree.RuleNode{
                Rule: tree.Rule{
                    Type: tree.RuleType_CU,
                    Args: map[string]string {
                        "Relation": "owner",
                    },
                },
                Children: []*tree.UsersetNode{
                    &tree.UsersetNode{
                        Userset: model.Userset{
                            Namespace: "test",
                            ObjectId: "obj",
                            Relation: "owner",
                        },
                        Child: &tree.RuleNode{
                            Rule: tree.Rule{
                                Type: tree.RuleType_THIS,
                            },
                            Children: []*tree.UsersetNode{
                                &tree.UsersetNode{
                                    Userset: model.Userset{
                                        Namespace: "test",
                                        ObjectId: "bob",
                                        Relation: "",
                                    },
                                    Child: &tree.RuleNode{
                                        Rule: tree.Rule{
                                            Type: tree.RuleType_THIS,
                                        },
                                        Children: []*tree.UsersetNode{},
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

func setupCtx(tuples []model.Tuple, ns []model.Namespace) context.Context {
	tr := maprepo.NewTupleRepo(tuples)
	nr := maprepo.NewNamespaceRepo(ns)

	ctx := context.Background()
	ctx = utils.WithTupleRepository(ctx, tr)
	ctx = utils.WithNamespaceRepository(ctx, nr)
        return ctx
}
