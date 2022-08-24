package rewrite

import (
    "testing"
    "reflect"

    "github.com/sourcenetwork/source-zanzibar/model"
)


func TestBuildExpressionTree(t *testing.T) {
    relation := getTestRelation()

    got := convertTree(relation.Rewrite.ExpressionTree, "name")

    want := &OpNode {
        Op: model.Operation_UNION,
        Left: &RuleNode {
            Child: nil,
            Rule: Rule {
                Type: RuleType_THIS,
                Args: []string{"name"},
            },
        },
        Right: &OpNode {
            Op: model.Operation_UNION,
            Left: &RuleNode{
                Child: nil,
                Rule: Rule{
                    Type: RuleType_CU,
                    Args: []string{"cu"},
                },
            },
            Right: &RuleNode{
                Child: nil,
                Rule: Rule{
                    Type: RuleType_TTU,
                    Args: []string{"ttu1", "ttu2"},
                },
            },
        },
    }

    if reflect.DeepEqual(got, want) {
        t.Errorf("relationToTree = %#v; want %#v", got, want)
    }
}

// Build Relation tree with 3 rewrite rules: this, ttu and cu
// The tree contains 2 op nodes and 3 leaves
func getTestRelation() model.Relation {
    tree := &model.RewriteNode_Opnode {
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
                Node: &model.RewriteNode_Opnode {
                    Opnode: &model.OpNode {
                        Op: model.Operation_UNION,
                        Left: &model.RewriteNode {
                            Node: &model.RewriteNode_Leaf {
                                Leaf: &model.Leaf {
                                    Rule: &model.Rule {
                                        Rule: &model.Rule_ComputedUserset {
                                            ComputedUserset: &model.ComputedUserset {
                                                Relation: "cu",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        Right: &model.RewriteNode {
                            Node: &model.RewriteNode_Leaf {
                                Leaf: &model.Leaf {
                                    Rule: &model.Rule {
                                        Rule: &model.Rule_TupleToUserset {
                                            TupleToUserset: &model.TupleToUserset {
                                                TuplesetRelation: "ttu1",
                                                ComputedUsersetRelation: "ttu2",
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

    relation := model.Relation {
        Name: "name",
        Rewrite: &model.UsersetRewrite {
            ExpressionTree: &model.RewriteNode {
                Node: tree,
            },
        },
    }
    return relation
}
