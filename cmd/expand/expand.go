package main

import (
    "context"
    "log"
    "strings"

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

    doc := model.Namespace{
        Name: "Doc",
        Relations: []*model.Relation{
            owner,
            reader,
        },
    }

    return []model.Namespace{
        doc,
    }
}

func tuplesFixture() []model.Tuple {
    tuples := []model.Tuple{
        model.Tuple { //(group:admin, member, bob)
            ObjectRel: &model.Userset{
                Namespace: "Group",
                ObjectId: "Admin",
                Relation: "Member",
            },
            User: &model.User{
                Type: model.UserType_USER,
                Userset: &model.Userset{
                    Namespace: "Users",
                    ObjectId: "Bob",
                    Relation: "...",
                },
            },
        },

        model.Tuple { //(doc:readme, owner, (group:admin, member))
            ObjectRel: &model.Userset{
                Namespace: "Doc",
                ObjectId: "Readme",
                Relation: "Owner",
            },
            User: &model.User{
                Type: model.UserType_USER_SET,
                Userset: &model.Userset{
                    Namespace: "Group",
                    ObjectId: "Admin",
                    Relation: "Member",
                },
            },
        },

        model.Tuple { //(doc:readme, reader, alice)
            ObjectRel: &model.Userset{
                Namespace: "Doc",
                ObjectId: "Readme",
                Relation: "Reader",
            },
            User: &model.User{
                Type: model.UserType_USER,
                Userset: &model.Userset{
                    Namespace: "Users",
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
        Namespace: "Doc",
        ObjectId: "Readme",
        Relation: "Reader",
    }

    tree, err := querier.Expand(ctx, uset)
    if err != nil {
        log.Fatal(err)
    }

    builder := strings.Builder {}
    buildTree(tree, &builder, 1)
    output := builder.String()
    print(output)
}

func buildTree(root rewrite.Node, builder *strings.Builder, level int) {
    header := strings.Repeat("#", level)
    builder.WriteString(header)
    builder.WriteString(" ")
    builder.WriteString(root.Display())
    builder.WriteString("\n")

    children := root.GetChildren()
    for _, child := range children {
        buildTree(child, builder, level+1)
    }
}
