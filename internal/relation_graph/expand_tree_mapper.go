package relation_graph

import (
	"github.com/sourcenetwork/zanzi/internal/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

func ToExpandTree(t *PathNode) *types.VerboseExpandTree {
	return mapPathNodeToExpandTree(t)
}

func mapOr(tree *ORNode) *types.ExpressionNode {
	// an OPNode is terminal if its children
	// are PathNodes
	if len(tree.Paths) == 0 {
		return nil
	}
	first, isTerminalOp := tree.Paths[0].(*PathNode)
	if isTerminalOp {
		children := utils.MapSlice(tree.Paths, mapPathNodeToExpandTree)
		return &types.ExpressionNode{
			Node: &types.ExpressionNode_FactorNode{
				FactorNode: &types.FactorNode{
					RewriteRule: first.Reason,
					Children:    children,
				},
			},
		}
	} else {
		return &types.ExpressionNode{
			Node: &types.ExpressionNode_OpNode{
				OpNode: &types.OpNode{
					Left:     mapOpNodes(tree.Paths[0]),
					Operator: types.Operator_DIFFERENCE,
					Right:    mapOpNodes(tree.Paths[1]),
				},
			},
		}
	}
}

func mapAnd(tree *ANDNode) *types.ExpressionNode {
	// an OPNode is terminal if its children
	// are PathNodes
	if len(tree.Paths) == 0 {
		return nil
	}
	first, isTerminalOp := tree.Paths[0].(*PathNode)
	if isTerminalOp {
		children := utils.MapSlice(tree.Paths, mapPathNodeToExpandTree)
		return &types.ExpressionNode{
			Node: &types.ExpressionNode_FactorNode{
				FactorNode: &types.FactorNode{
					RewriteRule: first.Reason,
					Children:    children,
				},
			},
		}
	} else {
		return &types.ExpressionNode{
			Node: &types.ExpressionNode_OpNode{
				OpNode: &types.OpNode{
					Left:     mapOpNodes(tree.Paths[0]),
					Operator: types.Operator_INTERSECTION,
					Right:    mapOpNodes(tree.Paths[1]),
				},
			},
		}
	}
}

func mapDiff(tree *DifferenceNode) *types.ExpressionNode {
	// an OPNode is terminal if its children
	// are PathNodes
	if (tree.Left == nil) && (tree.Right == nil) {
		return nil
	}
	first, isTerminalOp := tree.Left.(*PathNode)
	if isTerminalOp {
		return &types.ExpressionNode{
			Node: &types.ExpressionNode_FactorNode{
				FactorNode: &types.FactorNode{
					RewriteRule: first.Reason,
					Children:    []*types.VerboseExpandTree{mapPathNodeToExpandTree(tree.Left), mapPathNodeToExpandTree(tree.Right)},
				},
			},
		}
	} else {
		return &types.ExpressionNode{
			Node: &types.ExpressionNode_OpNode{
				OpNode: &types.OpNode{
					Left:     mapOpNodes(tree.Left),
					Operator: types.Operator_INTERSECTION,
					Right:    mapOpNodes(tree.Right),
				},
			},
		}
	}
}

func mapOpNodes(tree GoalTree) *types.ExpressionNode {
	switch node := tree.(type) {
	case *ORNode:
		return mapOr(node)
	case *ANDNode:
		return mapAnd(node)
	case *DifferenceNode:
		return mapDiff(node)
	default:
		return nil
	}
}

func mapPathNodeToExpandTree(pathNode GoalTree) *types.VerboseExpandTree {
	node := pathNode.(*PathNode)
	var obj *domain.Entity
	relation := ""
	if entityNode := node.RelationNode.GetEntity(); entityNode != nil {
		obj = entityNode.Object
	} else if entitySet := node.RelationNode.GetEntitySet(); entitySet != nil {
		obj = entitySet.Object
		relation = entitySet.Relation
	}

	var next *types.ExpressionNode
	if node.Path != nil {
		next = mapOpNodes(node.Path)
	}

	return &types.VerboseExpandTree{
		Entity:   obj,
		Relation: relation,
		Node:     next,
	}
}
