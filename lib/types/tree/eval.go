package tree

import (
	"github.com/sourcenetwork/source-zanzibar/model"

	mapset "github.com/deckarep/golang-set/v2"
)

func Eval(node *UsersetNode) mapset.Set[model.AuthNode] {
	return evalExprNode(node.Child)
}

func apply(left, right mapset.Set[model.AuthNode], op model.Operation) mapset.Set[model.AuthNode] {
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

func evalExprNode(exprNode ExpressionNode) mapset.Set[model.AuthNode] {
	switch node := exprNode.(type) {
	case *OpNode:
		left := evalExprNode(node.Left)
		right := evalExprNode(node.Right)
		return apply(left, right, node.JoinOp)

	case *RuleNode:
		usets := mapset.NewSet[model.AuthNode]()
		// This is bad
		for _, child := range node.Children {
			uset := child.Userset
			usets.Add(uset)
			result := evalExprNode(child.Child)
			usets = usets.Union(result)
		}
		return usets
	case nil:
		return mapset.NewSet[model.AuthNode]()
	default:
		panic("invalid ExpressionNode type")
	}
}
