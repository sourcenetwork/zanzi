package domain

import (
	"fmt"
)

func CUNode(targetRelation string) *RelationExpressionTree {
	return &RelationExpressionTree{
		Node: &RelationExpressionTree_Rule{
			Rule: &Rule{
				Rule: &Rule_Cu{
					Cu: &ComputedUserset{
						TargetRelation: targetRelation,
					},
				},
			},
		},
	}
}

func TTUNode(tuplesetRelation, computedUsersetRelation string) *RelationExpressionTree {
	return &RelationExpressionTree{
		Node: &RelationExpressionTree_Rule{
			Rule: &Rule{
				Rule: &Rule_Ttu{
					Ttu: &TupleToUserset{
						TuplesetRelation:        tuplesetRelation,
						ComputedUsersetRelation: computedUsersetRelation,
					},
				},
			},
		},
	}
}

func ThisNode() *RelationExpressionTree {
	return &RelationExpressionTree{
		Node: &RelationExpressionTree_Rule{
			Rule: &Rule{
				Rule: &Rule_This{
					This: &This{},
				},
			},
		},
	}
}

func NewOpNode(left *RelationExpressionTree, op Operator, right *RelationExpressionTree) *RelationExpressionTree {
	return &RelationExpressionTree{
		Node: &RelationExpressionTree_OpNode{
			OpNode: &OpNode{
				Left:     left,
				Operator: op,
				Right:    right,
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

// Convert a Relation Expression Tree to its string representation
// as defined in the Relation Expression mini-language
func (tree *RelationExpressionTree) RelationExpression() string {
	switch node := tree.Node.(type) {
	case *RelationExpressionTree_Rule:
		return node.Rule.RelationExpression()
	case *RelationExpressionTree_OpNode:
		left := node.OpNode.Left.RelationExpression()
		right := node.OpNode.Right.RelationExpression()
		return fmt.Sprintf("%v%v %v %v%v", GroupBeginLexeme, left, node.OpNode.Operator.RelationExpression(), right, GroupEndLexeme)
	}

	return ""
}

// Convert a Rule to its string representation
// as defined in the Relation Expression mini-language
func (r *Rule) RelationExpression() string {
	var ruleStr string
	switch rule := r.Rule.(type) {
	case *Rule_This:
		ruleStr = ThisLexeme
	case *Rule_Cu:
		ruleStr = rule.Cu.TargetRelation
	case *Rule_Ttu:
		ruleStr = rule.Ttu.TuplesetRelation + ArrowLexeme + rule.Ttu.ComputedUsersetRelation
	}
	return ruleStr
}

// Convert an Operator to its string representation
// as defined in the Relation Expression mini-language
func (op *Operator) RelationExpression() string {
	switch *op {
	case Operator_UNION:
		return UnionLexeme
	case Operator_INTERSECTION:
		return IntersectionLexeme
	case Operator_DIFFERENCE:
		return DifferenceLexeme
	}
	return ""
}

const (
	UnionLexeme        string = "+"
	IntersectionLexeme        = "&"
	DifferenceLexeme          = "-"
	ArrowLexeme               = "->"
	ThisLexeme                = "_this"
	GroupBeginLexeme          = "("
	GroupEndLexeme            = ")"
)
