package rewrite

import (
    "github.com/sourcenetwork/source-zanzibar/model"
)

// Type definies possible Operations for a Node in the Userset Expression Tree
type Operation uint8
const (
    Union Operation = iota
    Difference
    Intersection
    NoOp // NoOp indicates a Leaf node which should not perform an operation
)


type JoinableNodeType uint8
var (
    JoinableNode_OpNode JoinableNode = iota
    JoinableNode_Leaf
)

type NodeType uint8
var (
    Node_OpNode NodeType = iota
    Node_Leaf
    Node_RuleNode
)

type JoinableNode interface {
    GetJoinableNodeType() JoinableNodeType
}

type Node interface {
    GetNodeType() NodeType
}


type OpNode struct {
    Left Node
    Right Node
    Op Operation
}

func (n *OpNode) GetJoinableNodeType() { 
    return JoinableNode_OpNode 
}

func (n *OpNode) GetNodeType() { 
    return Node_OpNode 
}

type Leaf struct {
    Users []model.User
}

func (n *Leaf) GetJoinableNodeType() { 
    return JoinableNode_Leaf 
}

func (n *Leaf) GetNodeType() { 
    return Node_Leaf 
}

type RuleJoinNode struct {
    Rule Rule
    Node JoinableNode
}

func (n *RuleJoinNode) GetNodeType() { 
    return Node_RuleNode
}
