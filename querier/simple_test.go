package querier

import (
    "testing"
    "context"
 
    "github.com/stretchr/testify/assert"

    "github.com/sourcenetwork/source-zanzibar/repository/maprepo"
    "github.com/sourcenetwork/source-zanzibar/repository"
    "github.com/sourcenetwork/source-zanzibar/utils"
    "github.com/sourcenetwork/source-zanzibar/rewrite"
    "github.com/sourcenetwork/source-zanzibar/model"
)

func produceNamespaceRepo() repository.NamespaceRepository {
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
    return maprepo.NewNamespaceRepo([]model.Namespace{namespace})
}

func produceTupleRepo() repository.TupleRepository {

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
                Type: model.UserType_USER,
                Userset: &model.Userset{
                    Namespace: "Test",
                    ObjectId: "Group",
                    Relation: "Member",
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

    return maprepo.NewTupleRepo(tuples)
}


func TestExpandWithThisRelations(t *testing.T) {

    tr := produceTupleRepo()
    nr := produceNamespaceRepo()

    ctx := context.Background()
    ctx = utils.WithTupleRepository(ctx, tr)
    ctx = utils.WithNamespaceRepository(ctx, nr)

    uset := model.Userset {
        Namespace: "Test",
        ObjectId: "Object",
        Relation: "Reader",
    }

    got, err := Expand(ctx, uset)

    assert.Nil(t, err)

    want := &rewrite.UsersetNode {
        Userset: model.Userset {
            Namespace: "Test",
            ObjectId: "Object",
            Relation: "Reader",
        },
        Child: &rewrite.RuleNode {
            Rule: rewrite.Rule {
                Type: rewrite.RuleType_THIS,
            },
            Children: []rewrite.Node{

                &rewrite.UsersetNode {
                    Userset: model.Userset {
                        Namespace: "Test",
                        ObjectId: "Group",
                        Relation: "Member",
                    },
                    Child: &rewrite.RuleNode {
                        Rule: rewrite.Rule {
                            Type: rewrite.RuleType_THIS,
                        },
                        Children: []rewrite.Node{
                            &rewrite.UsersetNode {
                                Userset: model.Userset {
                                    Namespace: "Test",
                                    ObjectId: "Alice",
                                    Relation: "...",
                                },
                                Child: &rewrite.RuleNode {
                                    Rule: rewrite.Rule {
                                        Type: rewrite.RuleType_THIS,
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
