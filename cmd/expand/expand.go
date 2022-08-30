package main

import (
     "context"
     "log"
    "strings"
    "fmt"

     "github.com/sourcenetwork/source-zanzibar/repository/maprepo"
     "github.com/sourcenetwork/source-zanzibar/querier"
     "github.com/sourcenetwork/source-zanzibar/rewrite"
     "github.com/sourcenetwork/source-zanzibar/model"
     "github.com/sourcenetwork/source-zanzibar/utils"
)

func namespacesFixture() []model.Namespace {
    owner := &model.Relation {
        Name: "Owner",
        Rewrite: &model.UsersetRewrite {
            ExpressionTree: &model.RewriteNode {
                Node: &model.RewriteNode_Leaf {
                    Leaf: &model.Leaf {
                        Rule: &model.Rule {
                            Rule: &model.Rule_This {
                                This: &model.This {},
                            },
                        },
                    },
                },
            },
        },
    }

    reader := &model.Relation {
        Name: "Reader",
        Rewrite: &model.UsersetRewrite {
            ExpressionTree: &model.RewriteNode {
                Node: &model.RewriteNode_Opnode {
                    Opnode: &model.OpNode {
                        Op: model.Operation_UNION,
                        Left: &model.RewriteNode {
                            Node: &model.RewriteNode_Leaf {
                                Leaf: &model.Leaf {
                                    Rule: &model.Rule {
                                        Rule: &model.Rule_This {
                                            This: &model.This {},
                                        },
                                    },
                                },
                            },
                        },
                        Right: &model.RewriteNode {
                            Node: &model.RewriteNode_Leaf {
                                Leaf: &model.Leaf {
                                    Rule: &model.Rule {
                                        Rule: &model.Rule_ComputedUserset {
                                            ComputedUserset: &model.ComputedUserset {
                                                Relation: "Owner",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    }

    member := &model.Relation {
        Name: "Member",
        Rewrite: &model.UsersetRewrite {
            ExpressionTree: &model.RewriteNode {
                Node: &model.RewriteNode_Leaf {
                    Leaf: &model.Leaf {
                        Rule: &model.Rule {
                            Rule: &model.Rule_This {
                                This: &model.This {},
                            },
                        },
                    },
                },
            },
        },
    }

    empty := &model.Relation {
        Name: "...",
        Rewrite: &model.UsersetRewrite {
            ExpressionTree: &model.RewriteNode {
                Node: &model.RewriteNode_Leaf {
                    Leaf: &model.Leaf {
                        Rule: &model.Rule {
                            Rule: &model.Rule_This {
                                This: &model.This {},
                            },
                        },
                    },
                },
            },
        },
    }

    namespace := model.Namespace {
        Name: "Test",
        Relations: []*model.Relation{owner, reader, member, empty},
    }

    return []model.Namespace{
        namespace,
    }
}

func tuplesFixture() []model.Tuple {
    tuples := []model.Tuple{
        model.Tuple { //(object, owner, bob)
            ObjectRel: &model.Userset{
                Namespace: "Test",
                ObjectId: "Object",
                Relation: "Owner",
            },
            User: &model.User{
                Type: model.UserType_USER,
                Userset: &model.Userset{
                    Namespace: "Test",
                    ObjectId: "Bob",
                    Relation: "...",
                },
            },
        },

        model.Tuple { //(object, reader, (group, member))
            ObjectRel: &model.Userset{
                Namespace: "Test",
                ObjectId: "Object",
                Relation: "Reader",
            },
            User: &model.User{
                Type: model.UserType_USER_SET,
                Userset: &model.Userset{
                    Namespace: "Test",
                    ObjectId: "Group",
                    Relation: "Member",
                },
            },
        },

        model.Tuple { //(object, reader, (group, member))
            ObjectRel: &model.Userset{
                Namespace: "Test",
                ObjectId: "Object",
                Relation: "Reader",
            },
            User: &model.User{
                Type: model.UserType_USER,
                Userset: &model.Userset{
                    Namespace: "Test",
                    ObjectId: "Charlie",
                    Relation: "...",
                },
            },
        },

        model.Tuple { //(group, member, alice)
            ObjectRel: &model.Userset{
                Namespace: "Test",
                ObjectId: "Group",
                Relation: "Member",
            },
            User: &model.User{
                Type: model.UserType_USER,
                Userset: &model.Userset{
                    Namespace: "Test",
                    ObjectId: "Alice",
                    Relation: "...",
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
        ObjectId: "Object",
        Relation: "Reader",
    }

    tree, err := querier.Expand(ctx, uset)
    if err != nil {
        log.Fatal(err)
    }

    printNode(1, tree)
}


func printNode(lvl int, node rewrite.Node) {
    header := strings.Repeat(" ", lvl)

    switch n := node.(type) {

    case *rewrite.RuleNode:
        fmt.Printf("%v RuleNode rule=%v\n", header, n.Rule)
        for _, child := range n.Children {
            printNode(lvl+1, child)
        }

    case *rewrite.UsersetNode:
        uset := fmt.Sprintf("{%v, %v, %v}", n.Userset.Namespace, n.Userset.ObjectId, n.Userset.Relation)
        fmt.Printf("%v UsersetNode userset=%v\n", header, uset)
        printNode(lvl+1, n.Child)

    case *rewrite.OpNode:
        fmt.Printf("%v OpNode OP=%v\n", header, n.JoinOp)
        printNode(lvl+1, n.Left)
        printNode(lvl+1, n.Right)
    }
}
