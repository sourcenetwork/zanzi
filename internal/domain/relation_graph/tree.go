package relation_graph

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	"github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
)

var (
	_ RewriteNode = (*RuleNode)(nil)
	_ RewriteNode = (*OpNode)(nil)
	_ Node        = (*OpNode)(nil)
	_ Node        = (*RuleNode)(nil)
	_ Node        = (*RelationNode)(nil)
)

// Node type represents a node variant for the Walk Tree
type Node interface {
	GetChildren() []Node
}

// RewriteNode expresses a Zanzibar "Userset Rewrite" tree
// Rewrite Node can either be a single Rewrite Rule
// or a tree of Rewrite Rules
type RewriteNode interface {
	Node
	isRewriteNode()
}

// OpNode is a RewriteNode which combines RuleNodes.
// It is composed of two RuleNodes and a set operation.
type OpNode struct {
	Left   RewriteNode
	Right  RewriteNode
	JoinOp policy.Operation
}

func (n *OpNode) GetChildren() []Node {
	return []Node{n.Left, n.Right}
}

func (n *OpNode) isRewriteNode() {}

// RuleNode models an userset rule expansion operation.
// Each userset rule produces a collection of Nodes (usersets).
type RuleNode struct {
	RuleData RuleData
	Children []*RelationNode
}

func (n *RuleNode) GetChildren() []Node {
	children := make([]Node, len(n.Children))
	for i := range n.Children {
		children[i] = n.Children[i]
	}
	return children
}

func (n *RuleNode) isRewriteNode() {}

// RelationNode models a RelationGraph AuthNode.
// RelationNodes children are determined after evaluating its rewrite node.
type RelationNode struct {
	ObjRel tuple.TupleNode
	Child  RewriteNode
}

func (n *RelationNode) GetChildren() []Node {
	return []Node{n.Child}
}

// Rule represents an userset rewrite rule.
// Each Rule has a type (a fixed set of variantes)
// and a map of arguments specified at namespace definition
type RuleData struct {
	Type RuleType
	Args map[string]string
}

// Return new RuleData from a policy RewriteRule
func NewRuleData(rule *policy.RewriteRule) RuleData {
	switch r := rule.RewriteRule.(type) {
	case *policy.RewriteRule_This:
		return RuleData{
			Type: RuleType_THIS,
		}

	case *policy.RewriteRule_TupleToUserset:
		ttu := r.TupleToUserset
		return RuleData{
			Type: RuleType_TTU,
			Args: map[string]string{
				"TuplesetRelation":        ttu.TuplesetRelation,
				"ComputedUsersetRelation": ttu.CuRelation,
			},
		}

	case *policy.RewriteRule_ComputedUserset:
		cu := r.ComputedUserset
		return RuleData{
			Type: RuleType_CU,
			Args: map[string]string{
				"Relation": cu.Relation,
			},
		}

	default:
		err := fmt.Errorf("unknown rule type: %#v", r)
		panic(err)
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

func EvalTree(node Node) mapset.Set[tuple.TupleNode] {
	switch n := node.(type) {
	case *RuleNode:
		set := mapset.NewSet[tuple.TupleNode]()
		for _, child := range n.Children {
			set = set.Union(EvalTree(child))
		}
		return set
	case *OpNode:
		l := EvalTree(n.Left)
		r := EvalTree(n.Right)
		return applyOp(l, r, n.JoinOp)
	case *RelationNode:
		relations := EvalTree(n.Child)
		relations.Add(n.ObjRel)
		return relations
	case nil:
		return mapset.NewSet[tuple.TupleNode]()
	default:
		panic("invalid Node type")
	}
}

func applyOp(left, right mapset.Set[tuple.TupleNode], op policy.Operation) mapset.Set[tuple.TupleNode] {
	switch op {
	case policy.Operation_UNION:
		return left.Union(right)
	case policy.Operation_INTERSECTION:
		return left.Intersect(right)
	case policy.Operation_DIFFERENCE:
		return left.Difference(right)
	default:
		panic("invalid Operation")
	}
}
