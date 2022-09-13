package model

import (
	"fmt"
)

func (rw *RewriteNode) GetLeaves() []*Leaf {
	switch node := rw.Node.(type) {
	case *RewriteNode_Opnode:
		opnode := node.Opnode
		left := opnode.Left.GetLeaves()
		right := opnode.Right.GetLeaves()
		nLeaves := len(left) + len(right)
		leaves := make([]*Leaf, 0, nLeaves)
		leaves = append(leaves, left...)
		leaves = append(leaves, right...)
		return leaves
	case *RewriteNode_Leaf:
		leaf := node.Leaf
		return []*Leaf{leaf}
	default:
		panic(fmt.Sprintf("Invalid RewriteNode type: %v", rw))
	}
}
