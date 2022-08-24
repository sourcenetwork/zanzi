package rewrite

import (
    "testing"

    "github.com/sourcenetwork/source-zanzibar/model"
)


func TestBuildExpressionTree(t *testing.T) {
    relation := getTestRelation()

    got := relationToTree(relation, "name")

    want := OpNode {
        Op: Operation_UNION,
        Left: RuleNode {
            Child: nil,
            Rule: Rule {
                Type: RuleType_THIS,
                Args: []string{"name"},
            },
        },
        Right: OpNode {
            Op: Operation_UNION,
            Left: RuleNode{
                Child: nil,
                Rule: Rule{
                    Type: RuleType_CU,
                    Args: []string{"cu"},
                },
            },
            Right: RuleNode{
                Child: nil,
                Rule: Rule{
                    Type: RuleType_TTU,
                    Args: []string{"ttu1", "ttu2"},
                },
            },
        },
    }

    if tree != want {
        t.Errorf("relationToTree = %d; want %d", got, want)
    }
}

// Build Relation tree with 3 rewrite rules: this, ttu and cu
// The tree contains 2 op nodes and 3 leaves
func getTestRelation() model.Relation {
    tree := rewrite.OpNode {
        Op: rewrite.Operation_UNION,
        Left: rewrite.RewriteNode {
            Node: rewrite.Leaf {
                Rule: Rewrite.This {},
            },
        },
        Right: rewrite.RewriteNode {
            Node: rewrite.OpNode {
                Op: rewrite.Operation_UNION,
                Left: rewrite.RewriteNode {
                    Node: rewrite.Leaf {
                        Rule: Rewrite.ComputedUserset {
                            Relation: "cu",
                        },
                    },
                },
                Right: rewrite.RewriteNode {
                    Node: rewrite.Leaf {
                        Rule: Rewrite.TupleToUserset {
                            TuplesetRelation: "ttu1",
                            ComputedUsersetRelation: "ttu2",
                        },
                    },
                },
            },
        },
    }

    relation := model.Relation {
        Name: "name",
        Rewrite: rewrite.UsersetRewrite {
            ExpressionTree: rewrite.RewriteNode {
                Node: tree,
            },
        },
    }
    return relation
}
