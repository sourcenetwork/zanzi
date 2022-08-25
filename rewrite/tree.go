package rewrite

import (
    "fmt"
    "strings"

	"github.com/sourcenetwork/source-zanzibar/model"
)

var (
    _ Node = (*OpNode)(nil)
    _ Node = (*Leaf)(nil)
    _ Node = (*RuleNode)(nil)
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

// Type Rule represents a Userset Rewrite Rule.
// Rule contains the rule type and a list of arguments passed to the rule definition.
//
// Userset Rewrite Rules specifiy a function which controls how a relation will return a userset during a query.
// Quoting the paper:
// "Each rule specifies a function that takes an object ID as input and outputs a userset expression tree."
//
// Contents of Args is not enforced by the type system but follows the convetion:
// THIS Args: Name of the relation which defined the This rule
// CU Args: Name of the target relation
// TTU Args: Contains two arguments -> Args[0] = Tupleset lookup Relation name; Args[1] Computed Userset Relation name
type Rule struct {
    Type RuleType
    Args []string
}

type OpNode struct {
	Left  Node
	Right Node
	Op    model.Operation
}

// Type Leaf models leaves in the userset expansion tree
// Leaves contains a list of usersets
type Leaf struct {
	Users []model.Userset
}

// RuleNode represents an Userset Rewrite Rule
// Its child contains the userset expression tree for the related relation
type RuleNode struct {
	Rule Rule
	Children []Node
}


type Node interface {
	GetNodeType() NodeType
        Display() string
        GetChildren() []Node
}

func (n *RuleNode) GetNodeType() NodeType {
	return Node_RuleNode
}

func (n *RuleNode) Display() string {
    return fmt.Sprintf("Rulenode: Rule=%v, Args=%v", MapRuleName(n.Rule.Type), n.Rule.Args)
}

func (n *RuleNode) GetChildren() []Node {
	return n.Children
}

func (n *Leaf) GetNodeType() NodeType {
	return Node_Leaf
}

func (n *Leaf) Display() string {
    builder := strings.Builder{}
    builder.WriteString("Leaf \n")
    for _, userset := range n.Users {
        builder.WriteString("- ")
        builder.WriteString(userset.String())
        builder.WriteString("\n")
    }
    return builder.String()
}

func (n *Leaf) GetChildren() []Node {
	return nil
}

func (n *OpNode) GetNodeType() NodeType {
	return Node_OpNode
}

func (n *OpNode) Display() string {
    return fmt.Sprintf("Opnode: Operation=%v", n.Op.String())
}

func (n *OpNode) GetChildren() []Node {
    return []Node{n.Left, n.Right}
}

func MapRuleName(t RuleType) string {
    switch t {
    case RuleType_THIS:
        return "This"
    case RuleType_CU:
        return "Computed Userset"
    case RuleType_TTU:
        return "Tuple To Userset"
    default:
        return ""
    }
}
