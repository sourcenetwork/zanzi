package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThisToRelationExpression(t *testing.T) {
	tree := ThisNode()

	relExpr := tree.RelationExpression()

	require.Equal(t, relExpr, "_this")
}

func TestComputedUsersetToRelationExpression(t *testing.T) {
	tree := CUNode("relation")

	relExpr := tree.RelationExpression()

	require.Equal(t, relExpr, "relation")
}

func TestTupleToUsersetToRelationExpression(t *testing.T) {
	tree := TTUNode("parent", "owner")

	relExpr := tree.RelationExpression()

	require.Equal(t, relExpr, "parent->owner")
}

func TestUnionToRelationExpression(t *testing.T) {
	tree := UnionNode(CUNode("left"), CUNode("right"))

	relExpr := tree.RelationExpression()

	require.Equal(t, relExpr, "(left + right)")
}

func TestIntersectionToRelationExpression(t *testing.T) {
	tree := IntersectionNode(CUNode("left"), CUNode("right"))

	relExpr := tree.RelationExpression()

	require.Equal(t, relExpr, "(left & right)")
}

func TestDifferenceToRelationExpression(t *testing.T) {
	tree := DifferenceNode(CUNode("left"), CUNode("right"))

	relExpr := tree.RelationExpression()

	require.Equal(t, relExpr, "(left - right)")
}
