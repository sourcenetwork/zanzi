package tree

import (
	"github.com/sourcenetwork/source-zanzibar/model"
)

var (
	_ ExpressionNode = (*RuleNode)(nil)
	_ ExpressionNode = (*OpNode)(nil)
	_ Node           = (*OpNode)(nil)
	_ Node           = (*RuleNode)(nil)
	_ Node           = (*UsersetNode)(nil)
)

// OpNode represents an inner node which combines leave nodes.
// Opnode contains two children and a set operation.
//
// Since not all operations are commutative (eg difference),
// OpNode operand semantic is: `LEFT OP RIGHT`.
//
// OpNodes may be nested to build expression trees.
// Expression Tree leaves are `RuleNode`s
type OpNode struct {
	Left   ExpressionNode
	Right  ExpressionNode
	JoinOp model.Operation
}

func (n *OpNode) GetNodeType() NodeType {
	return NodeType_OP
}

func (n *OpNode) GetChildren() []Node {
	return []Node{n.Left, n.Right}
}

func (n *OpNode) isExpressionNode() {}

// RuleNode models an userset rule expansion operation.
// Each userset rule produces a set of Usersets,
// which are RuleNode's children.
type RuleNode struct {
	Rule     Rule
	Children []*UsersetNode
}

func (n *RuleNode) GetChildren() []Node {
	return nil
}

func (n *RuleNode) GetNodeType() NodeType {
	return NodeType_RULE
}

func (n *RuleNode) isExpressionNode() {}

// UsersetNode models the result of evaluating userset rewrite rules for an userset.
// It contains an Userset indicating the root Node and an ExpressionNode child.
// By evaluating the Rule(s) in ExpressionNode
// the UsersetNode may form multiple DFS trees which shall be combined
// through operations defined in OpNode
type UsersetNode struct {
	Userset model.AuthNode
	Child   ExpressionNode
}

func (n *UsersetNode) GetChildren() []Node {
	return []Node{n.Child}
}

func (n *UsersetNode) GetNodeType() NodeType {
	return NodeType_USERSET
}
