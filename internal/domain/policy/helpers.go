package policy

import (
	"fmt"
)

func (t *Tree) GetLeaves() []Leaf {
	switch node := t.Node.(type) {
	case *Tree_Opnode:
		opnode := node.Opnode
		left := opnode.Left.GetLeaves()
		right := opnode.Right.GetLeaves()
		nLeaves := len(left) + len(right)
		leaves := make([]Leaf, 0, nLeaves)
		leaves = append(leaves, left...)
		leaves = append(leaves, right...)
		return leaves
	case *Tree_Leaf:
		leaf := node.Leaf
		return []Leaf{*leaf}
	default:
		panic(fmt.Sprintf("Invalid RewriteNode type: %v", rw))
	}
}
