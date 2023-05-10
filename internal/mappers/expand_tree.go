package mappers

import (
	_ "google.golang.org/protobuf/proto"

	"github.com/sourcenetwork/zanzi/internal/domain/policy"
	rg "github.com/sourcenetwork/zanzi/internal/domain/relation_graph"
	"github.com/sourcenetwork/zanzi/pkg/utils"
	"github.com/sourcenetwork/zanzi/types"
)

type ExpandTreeMapper struct{}

func (m *ExpandTreeMapper) ToExpandTree(tree *rg.RelationNode) types.ExpandTree {
	return types.ExpandTree{
		Entity: types.Entity{
			Namespace: tree.ObjRel.Namespace,
			Id:        tree.ObjRel.Id,
		},
		RelOrPerm:          tree.ObjRel.Relation,
		RelationExpression: m.toExpressionTree(tree.Child),
	}
}

func (m *ExpandTreeMapper) mapOp(op policy.Operation) types.Operator {
	var operator types.Operator
	switch op {
	case policy.Operation_UNION:
		operator = types.UNION
	case policy.Operation_DIFFERENCE:
		operator = types.DIFFERENCE
	case policy.Operation_INTERSECTION:
		operator = types.INTERSECTION
	}
	return operator
}

func (m *ExpandTreeMapper) toFactor(node *rg.RuleNode) *types.Factor {
	f := func(relNode *rg.RelationNode) types.ExpandTree { return m.ToExpandTree(relNode) }
	return &types.Factor{
		Operand: m.ruleToOperand(&node.RuleData),
		Result:  utils.MapSlice(node.Children, f),
	}
}

func (m *ExpandTreeMapper) toExpression(node *rg.OpNode) *types.Expression {
	return &types.Expression{
		Operator: m.mapOp(node.JoinOp),
		Left:     m.toExpressionTree(node.Left),
		Right:    m.toExpressionTree(node.Right),
	}
}

func (m *ExpandTreeMapper) toExpressionTree(node rg.RewriteNode) types.ExpressionTree {
	switch n := node.(type) {
	case *rg.OpNode:
		return m.toExpression(n)
	case *rg.RuleNode:
		return m.toFactor(n)
	default:
		panic("unknown RewriteNode")
	}
}

func (m *ExpandTreeMapper) ruleToOperand(rule *rg.RuleData) string {
	switch rule.Type {
	case rg.RuleType_THIS:
		return "this"
	case rg.RuleType_CU:
		operand := rule.Args["Relation"]
		return operand
	case rg.RuleType_TTU:
		operand := rule.Args["TuplesetRelation"] + "->" + rule.Args["ComputedUsersetRelation"]
		return operand
	default:
		panic("unknown rule type")
	}
}
