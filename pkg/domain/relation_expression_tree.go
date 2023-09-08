package domain

func CUNode(targetRelation string) *RelationExpressionTree{
    return &RelationExpressionTree{
        Node: &RelationExpressionTree_Rule{
            Rule: &Rule{
                Rule: &Rule_Cu{
                    Cu: &ComputedUserset {
                        TargetRelation: targetRelation,
                    },
                },
            },
        },
    }
}

func TTUNode(tuplesetRelation, computedUsersetRelation string) *RelationExpressionTree{
    return &RelationExpressionTree{
        Node: &RelationExpressionTree_Rule{
            Rule: &Rule{
                Rule: &Rule_Ttu{
                    Ttu: &TupleToUserset {
                        TuplesetRelation: tuplesetRelation,
                        ComputedUsersetRelation: computedUsersetRelation,
                    },
                },
            },
        },
    }
}

func ThisNode() *RelationExpressionTree{
        return &RelationExpressionTree{
            Node: &RelationExpressionTree_Rule{
                Rule: &Rule{
                    Rule: &Rule_This{
                        This: &This { },
                    },
                },
            },
        }
}

func NewOpNode(left *RelationExpressionTree, op Operator, right *RelationExpressionTree) *RelationExpressionTree {
    return &RelationExpressionTree{
        Node: &RelationExpressionTree_OpNode{
            OpNode: &OpNode{
                Left: left,
                Operator: op,
                Right: right,
            },
        },
    }
}

func UnionNode(left, right *RelationExpressionTree) *RelationExpressionTree {
    return NewOpNode(left, Operator_UNION, right)
}

func IntersectionNode(left, right *RelationExpressionTree) *RelationExpressionTree {
    return NewOpNode(left, Operator_INTERSECTION, right)
}

func DifferenceNode(left, right *RelationExpressionTree) *RelationExpressionTree {
    return NewOpNode(left, Operator_DIFFERENCE, right)
}

// GetRules returns all the rules in a RelationExpressionTree.
func (tree *RelationExpressionTree) GetRules() []*Rule {
    var rules []*Rule
    return tree.getRules(rules)
}

func (tree *RelationExpressionTree) getRules(acc []*Rule) []*Rule {
    switch n := tree.Node.(type) {
    case *RelationExpressionTree_Rule:
        acc = append(acc, n.Rule)
    case *RelationExpressionTree_OpNode:
        acc = n.OpNode.Left.getRules(acc)
        acc = n.OpNode.Right.getRules(acc)
    }
    return acc
}
