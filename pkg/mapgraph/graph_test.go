package mapgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddingANodeReturnsIt(t *testing.T) {
	g := New[string]()

	g.SetNode("a", "node a")

	opt := g.GetNode("a")
	assert.False(t, opt.IsEmpty())
	assert.Equal(t, "node a", opt.Value())
}

func TestGetAncestorsReturnAncestorNodes(t *testing.T) {
	g := New[string]()

	g.SetNode("a", "a")
	g.SetNode("b", "b")
	g.SetNode("c", "c")
	g.SetNode("d", "d")

	g.SetEdge("a", "b")
	g.SetEdge("a", "c")
	g.SetEdge("b", "d")
	g.SetEdge("c", "d")

	edges := g.GetAncestors("b")

	assert.Contains(t, edges, "a")
	assert.Equal(t, 1, len(edges))
}

func TestGetSucessorsReturnDirectChild(t *testing.T) {
	g := New[string]()

	g.SetNode("a", "a")
	g.SetNode("b", "b")
	g.SetNode("c", "c")
	g.SetNode("d", "d")

	g.SetEdge("a", "b")
	g.SetEdge("a", "c")
	g.SetEdge("b", "d")
	g.SetEdge("c", "d")

	println("sucessors")
	edges := g.GetSucessors("a")

	assert.Contains(t, edges, "b")
	assert.Contains(t, edges, "c")
	assert.Equal(t, 2, len(edges))
}
