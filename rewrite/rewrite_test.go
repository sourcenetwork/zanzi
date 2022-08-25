package rewrite

import (
    "testing"
    "reflect"

    "github.com/sourcenetwork/source-zanzibar/model"
)


func TestBuildExpressionTreeWithNestedOpNodes(t *testing.T) {
    // Given Relation tree with 3 rewrite rules: this, ttu and cu (2 op nodes, 3 leaves)
    tree := &model.RewriteNode {
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
        },
    }

    got := convertTree(tree, "name")

    want := &OpNode {
        Op: model.Operation_UNION,
        Left: &RuleNode {
            Children: nil,
            Rule: Rule {
                Type: RuleType_THIS,
                Args: []string{"name"},
            },
        },
        Right: &OpNode {
            Op: model.Operation_UNION,
            Left: &RuleNode{
                Children: nil,
                Rule: Rule{
                    Type: RuleType_CU,
                    Args: []string{"cu"},
                },
            },
            Right: &RuleNode{
                Children: nil,
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


func TestBuildExpressionTreeWithThisAndComputedUserset(t *testing.T) {
    tree := &model.RewriteNode {
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
    }

    got := convertTree(tree, "Reader")

    want := &OpNode {
        Op: model.Operation_UNION,
        Left: &RuleNode {
            Children: nil,
            Rule: Rule {
                Type: RuleType_THIS,
                Args: []string{"Reader"},
            },
        },
        Right: &RuleNode{
            Children: nil,
            Rule: Rule{
                Type: RuleType_CU,
                Args: []string{"owner"},
            },
        },
    }

    if reflect.DeepEqual(got, want) {
        t.Errorf("relationToTree = %#v; want %#v", got, want)
    }
}
