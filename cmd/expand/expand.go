package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/model/builder"
	"github.com/sourcenetwork/source-zanzibar/querier"
	"github.com/sourcenetwork/source-zanzibar/repository/maprepo"
	"github.com/sourcenetwork/source-zanzibar/tree"
	"github.com/sourcenetwork/source-zanzibar/utils"
)

func namespacesFixture() []model.Namespace {
	owner := &model.Relation{
		Name: "Owner",
		Rewrite: &model.UsersetRewrite{
			ExpressionTree: builder.This(),
		},
	}

	reader := &model.Relation{
		Name: "Reader",
		Rewrite: &model.UsersetRewrite{
			ExpressionTree: builder.Union(builder.This(),
				builder.Union(builder.CU("Owner"), builder.TTU("Parent", "Owner")),
			),
		},
	}

	member := &model.Relation{
		Name: "Member",
		Rewrite: &model.UsersetRewrite{
			ExpressionTree: builder.This(),
		},
	}

	empty := &model.Relation{
		Name: "...",
		Rewrite: &model.UsersetRewrite{
			ExpressionTree: builder.This(),
		},
	}

	parent := &model.Relation{
		Name: "Parent",
		Rewrite: &model.UsersetRewrite{
			ExpressionTree: builder.This(),
		},
	}

	namespace := model.Namespace{
		Name:      "Test",
		Relations: []*model.Relation{owner, reader, member, empty, parent},
	}

	return []model.Namespace{
		namespace,
	}
}

func tuplesFixture() []model.Tuple {
	tuples := []model.Tuple{
		model.Tuple{ //(object, owner, bob)
			ObjectRel: &model.Userset{
				Namespace: "Test",
				ObjectId:  "Object",
				Relation:  "Owner",
			},
			User: &model.User{
				Type: model.UserType_USER,
				Userset: &model.Userset{
					Namespace: "Test",
					ObjectId:  "Bob",
					Relation:  "...",
				},
			},
		},

		model.Tuple{ //(object, reader, (group, member))
			ObjectRel: &model.Userset{
				Namespace: "Test",
				ObjectId:  "Object",
				Relation:  "Reader",
			},
			User: &model.User{
				Type: model.UserType_USER_SET,
				Userset: &model.Userset{
					Namespace: "Test",
					ObjectId:  "Group",
					Relation:  "Member",
				},
			},
		},

		model.Tuple{ //(object, reader, (group, member))
			ObjectRel: &model.Userset{
				Namespace: "Test",
				ObjectId:  "Object",
				Relation:  "Reader",
			},
			User: &model.User{
				Type: model.UserType_USER,
				Userset: &model.Userset{
					Namespace: "Test",
					ObjectId:  "Charlie",
					Relation:  "...",
				},
			},
		},

		model.Tuple{ //(group, member, alice)
			ObjectRel: &model.Userset{
				Namespace: "Test",
				ObjectId:  "Group",
				Relation:  "Member",
			},
			User: &model.User{
				Type: model.UserType_USER,
				Userset: &model.Userset{
					Namespace: "Test",
					ObjectId:  "Alice",
					Relation:  "...",
				},
			},
		},

		model.Tuple{ //(object, parent, (directory, ...))
			ObjectRel: &model.Userset{
				Namespace: "Test",
				ObjectId:  "Object",
				Relation:  "Parent",
			},
			User: &model.User{
				Type: model.UserType_USER,
				Userset: &model.Userset{
					Namespace: "Test",
					ObjectId:  "Directory",
					Relation:  "...",
				},
			},
		},

		model.Tuple{ //(directory, owner, charlie)
			ObjectRel: &model.Userset{
				Namespace: "Test",
				ObjectId:  "Directory",
				Relation:  "Owner",
			},
			User: &model.User{
				Type: model.UserType_USER,
				Userset: &model.Userset{
					Namespace: "Test",
					ObjectId:  "Steve",
					Relation:  "...",
				},
			},
		},
	}
	return tuples
}

func main() {
	tr := maprepo.NewTupleRepo(tuplesFixture())
	nr := maprepo.NewNamespaceRepo(namespacesFixture())

	ctx := context.Background()
	ctx = utils.WithTupleRepository(ctx, tr)
	ctx = utils.WithNamespaceRepository(ctx, nr)

	uset := model.Userset{
		Namespace: "Test",
		ObjectId:  "Object",
		Relation:  "Reader",
	}

	usetNode, err := querier.Expand(ctx, uset)
	if err != nil {
		log.Fatal(err)
	}

	usets := tree.Eval(usetNode)

	for _, uset := range usets {
		println(uset.String())
	}

	printNode(1, usetNode)
}

func printNode(lvl int, node tree.Node) {
	header := strings.Repeat(" ", lvl)

	switch n := node.(type) {

	case *tree.RuleNode:
		fmt.Printf("%v RuleNode rule=%v\n", header, n.Rule)
		for _, child := range n.Children {
			printNode(lvl+1, child)
		}

	case *tree.UsersetNode:
		uset := fmt.Sprintf("{%v, %v, %v}", n.Userset.Namespace, n.Userset.ObjectId, n.Userset.Relation)
		fmt.Printf("%v UsersetNode userset=%v\n", header, uset)
		printNode(lvl+1, n.Child)

	case *tree.OpNode:
		fmt.Printf("%v OpNode OP=%v\n", header, n.JoinOp)
		printNode(lvl+1, n.Left)
		printNode(lvl+1, n.Right)
	}
}
