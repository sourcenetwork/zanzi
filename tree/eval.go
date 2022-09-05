package tree

import (
	"github.com/sourcenetwork/source-zanzibar/model"

	mapset "github.com/deckarep/golang-set/v2"
)

func Eval(node *UsersetNode) mapset.Set[model.KeyableUset] {
	return evalExprNode(node.Child)
}

func apply(left, right mapset.Set[model.KeyableUset], op model.Operation) mapset.Set[model.KeyableUset] {
	switch op {
	case model.Operation_UNION:
		return left.Union(right)
	case model.Operation_INTERSECTION:
		return left.Intersect(right)
	case model.Operation_DIFFERENCE:
		return left.Difference(right)
	default:
		panic("invalid Operation")
	}
}

func evalExprNode(exprNode ExpressionNode) mapset.Set[model.KeyableUset] {
	switch node := exprNode.(type) {
	case *OpNode:
		left := evalExprNode(node.Left)
		right := evalExprNode(node.Right)
		return apply(left, right, node.JoinOp)

	case *RuleNode:
		usets := mapset.NewSet[model.KeyableUset]()
		// This is bad
		for _, child := range node.Children {
			key := child.Userset.ToKey()
			usets.Add(key)
			result := evalExprNode(child.Child)
			usets = usets.Union(result)
		}
		return usets
	default:
		panic("invalid ExpressionNode type")
	}
}
