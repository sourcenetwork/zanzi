// package mapgraph implements a minimal volatile graph datastructure
package mapgraph

import (
	"fmt"

	opt "github.com/sourcenetwork/zanzi/pkg/option"
)

// MapGraph implements a Graph data structure using Go maps.
type MapGraph[T any] struct {
	nodes         map[string]T
	forwardEdges  map[string]map[string]struct{}
	backwardEdges map[string]map[string]struct{}
}

func New[T any]() MapGraph[T] {
	return MapGraph[T]{
		nodes:         make(map[string]T),
		forwardEdges:  make(map[string]map[string]struct{}),
		backwardEdges: make(map[string]map[string]struct{}),
	}
}

// Add or update node to graph with id and data node
func (g *MapGraph[T]) SetNode(id string, node T) {
	g.nodes[id] = node
}

// Return node given by id
func (g *MapGraph[T]) GetNode(id string) opt.Option[T] {
	node, ok := g.nodes[id]
	if !ok {
		return opt.None[T]()
	}
	return opt.Some[T](node)
}

// Get ancestors for node given by id.
// If node does not exist, return empty slice
func (g *MapGraph[T]) GetAncestors(id string) []T {
	edgeMap, ok := g.backwardEdges[id]
	if !ok {
		return nil
	}
	return g.getNodesFromEdges(id, edgeMap)
}

func (g *MapGraph[T]) getNodesFromEdges(id string, edges map[string]struct{}) []T {
	nodes := make([]T, 0, len(edges))
	for key := range edges {
		node := g.nodes[key]
		nodes = append(nodes, node)
	}
	return nodes
}

// Get sucessors for node given by id.
// If node does not exist, return empty slice
func (g *MapGraph[T]) GetSucessors(id string) []T {
	edgeMap, ok := g.forwardEdges[id]
	if !ok {
		return nil
	}
	return g.getNodesFromEdges(id, edgeMap)
}

// Define an edge between two nodes.
// if sourceId or destId does not exist, return error
func (g *MapGraph[T]) SetEdge(sourceId, destId string) error {
	srcOpt := g.GetNode(sourceId)
	if srcOpt.IsEmpty() {
		return fmt.Errorf("node not found: %s", sourceId)
	}

	dstOpt := g.GetNode(destId)
	if dstOpt.IsEmpty() {
		return fmt.Errorf("node not found: %s", destId)
	}

	sourceEdgMap, ok := g.forwardEdges[sourceId]
	if !ok {
		sourceEdgMap = make(map[string]struct{})
		g.forwardEdges[sourceId] = sourceEdgMap
	}
	sourceEdgMap[destId] = struct{}{}

	destEdgMap, ok := g.backwardEdges[destId]
	if !ok {
		destEdgMap = make(map[string]struct{})
		g.backwardEdges[destId] = destEdgMap
	}
	destEdgMap[sourceId] = struct{}{}

	return nil
}
