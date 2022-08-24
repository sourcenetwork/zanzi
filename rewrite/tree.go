package rewrite

import (
	"github.com/sourcenetwork/source-zanzibar/model"
)

var (
    _ Node = (*OpNode)(nil)
    _ Node = (*Leaf)(nil)
    _ Node = (*RuleNode)(nil)
    _ JoinableNode = (*OpNode)(nil)
    _ JoinableNode = (*Leaf)(nil)
)

// Enum which identifies the Node variant Node variant
type NodeType uint8

const (
	Node_OpNode NodeType = iota
	Node_Leaf
	Node_RuleNode
)

// Enum defines the Rule
type RuleType uint8

const (
    RuleType_THIS RuleType = iota
    RuleType_TTU
    RuleType_CU
)

// Type Rule represents a Userset Rewrite Rule
// Rule contains the rule type and a list of arguments
// which was passed to the matching rule during userset rewrite definition.
//
// Contents of Args is not enforced by the type system but follow a convetion as follows:
// THIS Args: Name of the relation which defined the This rule
// CU Args: Name of the target relation
// TTU Args: Contains two arguments -> Args[0] = Tupleset lookup Relation name; Args[1] Computed Userset Relation name
type Rule struct {
    Type RuleType
    Args []string
}

// Interface type to restrict the variants on RuleNode child
// This is done to avoid the semantically meaningless scenario where
// a RuleNode has another RuleNode as a child.
type JoinableNode interface {
	isJoinableNode()
}

type Node interface {
	GetNodeType() NodeType
}

type OpNode struct {
	Left  Node
	Right Node
	Op    model.Operation
}

func (n *OpNode) isJoinableNode() { }

func (n *OpNode) GetNodeType() NodeType {
	return Node_OpNode
}

// Type Leaf models leaves in the userset expansion tree
// Leaves contains a list of usersets
type Leaf struct {
	Users []model.Userset
}

func (n *Leaf) isJoinableNode() { }

func (n *Leaf) GetNodeType() NodeType {
	return Node_Leaf
}

// RuleNode represents an Userset Rewrite Rule
// Its child contains the userset expression tree for the related relation
type RuleNode struct {
	Rule Rule
	Child JoinableNode
}

func (n *RuleNode) GetNodeType() NodeType {
	return Node_RuleNode
}
