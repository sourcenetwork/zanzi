package tree

import (
    "github.com/sourcenetwork/source-zanzibar/model"
)

type Node interface {
	GetNodeType() NodeType
        GetChildren() []Node
}

// Enum which identifies the Node variant Node variant
type NodeType uint8
const (
    NodeType_USERSET NodeType= iota
    NodeType_OP
    NodeType_RULE
)

func (n NodeType) String() string {
    switch n {
    case NodeType_USERSET:
        return "Userset Node"
    case NodeType_RULE:
        return "Rule Node"
    case NodeType_OP:
        return "Op Node"
    default:
        return ""
    }
}

type ExpressionNode interface {
    Node
    isExpressionNode()
}

var (
    _ ExpressionNode = (*RuleNode)(nil)
    _ ExpressionNode = (*OpNode)(nil)
)

func (n *OpNode) isExpressionNode() { }
func (n *RuleNode) isExpressionNode() { }

type RuleType uint8
const (
    RuleType_THIS RuleType = iota
    RuleType_CU
    RuleType_TTU
)

func (n RuleType) String() string {
    switch n {
    case RuleType_THIS:
        return "This"
    case RuleType_CU:
        return "ComputedUserset"
    case RuleType_TTU:
        return "TupleToUserset"
    default:
        return ""
    }
}


type Rule struct {
    Type RuleType
    Args map[string]string
}


type OpNode struct {
    Left ExpressionNode
    Right ExpressionNode
    JoinOp model.Operation
}

type RuleNode struct {
    Rule Rule
    Children []*UsersetNode
}

type UsersetNode struct {
    Userset model.Userset
    Child ExpressionNode
}

func (n *OpNode) GetNodeType() NodeType {
	return NodeType_OP
}

func (n *OpNode) GetChildren() []Node {
    return []Node{n.Left, n.Right}
}

func (n *UsersetNode) GetChildren() []Node {
    return []Node{n.Child}
}

func (n *UsersetNode) GetNodeType() NodeType {
    return NodeType_USERSET
}

func (n *RuleNode) GetChildren() []Node {
    return nil
}

func (n *RuleNode) GetNodeType() NodeType {
    return NodeType_RULE
}
