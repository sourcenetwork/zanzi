// Package tree define data types which expresses the result of Userset Rewrite expansions
package tree 


// Rule represents an userset rewrite rule.
// Each Rule has a type (a fixed set of variantes)
// and a map of arguments specified at namespace definition
type Rule struct {
    Type RuleType
    Args map[string]string
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

// Node type represents a variant of the defined tree
type Node interface {
    GetChildren() []Node
}

// Expression nodes represent the subset of the nodes
// which deal with userset rewrite rules
//
// ExpressionNode can model a single Userset rewrite rule
// or a deeply nested complex relationship.
type ExpressionNode interface {
    Node
    isExpressionNode()
}
