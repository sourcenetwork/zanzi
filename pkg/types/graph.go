package types

import (
	"strings"

	"github.com/awalterschulze/gographviz"
)

func NewMapDiGraph() *MapDiGraph {
	return &MapDiGraph{
		Nodes:        make(map[string]*GraphNode),
		ForwardEdges: make(map[string]*EdgeMap),
		BackEdges:    make(map[string]*EdgeMap),
	}
}

func (g *MapDiGraph) AddNode(node *GraphNode) {
	g.Nodes[node.Id] = node
}

func (g *MapDiGraph) AddEdge(edg *GraphEdge) {
	_, ok := g.ForwardEdges[edg.Source]
	if !ok {
		g.ForwardEdges[edg.Source] = &EdgeMap{
			Edges: make(map[string]*GraphEdge),
		}
	}

	_, ok = g.BackEdges[edg.Dest]
	if !ok {
		g.BackEdges[edg.Dest] = &EdgeMap{
			Edges: make(map[string]*GraphEdge),
		}
	}

	g.ForwardEdges[edg.Source].Edges[edg.Dest] = edg
	g.BackEdges[edg.Dest].Edges[edg.Source] = edg
}

func (g *MapDiGraph) GetNode(id string) *GraphNode {
	node, ok := g.Nodes[id]
	if !ok {
		return nil
	}
	return node
}

// GraphToDOT converts the given graph into a DOT representation
func GraphToDOT(graph *MapDiGraph) string {
	g := gographviz.NewGraph()
	g.SetDir(true) //directed graph true
	g.SetName("G")

	for _, node := range graph.Nodes {
		attrs := map[string]string{
			"label": wrap(node.Data),
		}
		g.AddNode("G", mapId(node.Id), attrs)
	}

	for src, edges := range graph.ForwardEdges {
		for dst, _ := range edges.Edges {
			g.AddEdge(mapId(src), mapId(dst), true, nil)
		}
	}

	return g.String()
}

func mapId(id string) string {
	id = strings.ReplaceAll(id, "\"", "\\\"")
	return wrap(id)
}

func wrap(txt string) string {
	return "\"" + txt + "\""
}
