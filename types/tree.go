package types

var _ ExpressionTree = (*Factor)(nil)
var _ ExpressionTree = (*Expression)(nil)

type Operator rune

const (
	UNION        Operator = '+'
	DIFFERENCE   Operator = '-'
	INTERSECTION Operator = '&'
)

type ExpandTree struct {
	Entity             Entity
	RelOrPerm          string
	RelationExpression ExpressionTree
}

type ExpressionTree interface {
	isExpressionTree()
}

type Factor struct {
	Operand string
	Result  []ExpandTree
}

func (f Factor) isExpressionTree() {}

type Expression struct {
	Operator Operator
	Left     ExpressionTree
	Right    ExpressionTree
}

func (e Expression) isExpressionTree() {}
