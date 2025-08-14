package relation_graph

import (
	"fmt"

	"github.com/awalterschulze/gographviz"

	"github.com/sourcenetwork/zanzi/pkg/types"
)

const (
	OkColor      string = "#299033"
	FailedColor         = "#BF2528"
	UnknownColor        = "#D0D0D0"
	RootColor           = "black"
)

type ExplainCheckTreeMapper struct {
	counter int
	nodes   []*types.CheckExplainNode
	edges   []*types.CheckExplainEdge
}

func (s *ExplainCheckTreeMapper) nextId() string {
	id := fmt.Sprintf("%v", s.counter)
	s.counter++
	return id
}

func (m *ExplainCheckTreeMapper) Map(tree GoalTree) *types.CheckExplainGraph {
	m.mapTree(tree, "")
	return &types.CheckExplainGraph{
		RootNodeId: "1",
		Edges:      m.edges,
		Nodes:      m.nodes,
	}
}

func (m *ExplainCheckTreeMapper) mapTree(tree GoalTree, parent_id string) {
	id := m.nextId()
	switch node := tree.(type) {
	case *PathNode:
		m.handlePathNode(id, parent_id, node)
	case *ORNode:
		m.handleOPNode(id, parent_id, node, node.Paths, types.CheckExplainNodeType_UNION_NODE, "+", "")
	case *ANDNode:
		m.handleOPNode(id, parent_id, node, node.Paths, types.CheckExplainNodeType_INTERSECTION_NODE, "&", "")
	case *DifferenceNode:
		m.handleOPNode(id, parent_id, node, []GoalTree{node.Left, node.Right}, types.CheckExplainNodeType_DIFF_NODE, "-", "")
	}
}

func (m *ExplainCheckTreeMapper) handlePathNode(id string, parent_id string, tree *PathNode) {
	// only a top level node may be a root node
	if parent_id != "" {
		edge := &types.CheckExplainEdge{
			SourceNodeId: parent_id,
			DestNodeId:   id,
			Message:      "",
		}
		m.edges = append(m.edges, edge)
	}

	node := &types.CheckExplainNode{
		Id:       id,
		NodeType: types.CheckExplainNodeType_USERSET_NODE,
		Text:     tree.RelationNode.PrettyString(),
		Detail:   "reason: " + tree.Reason,
		Result: &types.SearchResult{
			Authorized: tree.Result.Authorized,
			Explored:   tree.Result.Explored,
			Exhausted:  tree.Result.Completed,
		},
	}
	m.nodes = append(m.nodes, node)

	if tree.Path == nil {
		return
	}

	m.mapTree(tree.Path, id)
}

func (m *ExplainCheckTreeMapper) handleOPNode(id string, parent_id string, tree GoalTree, children []GoalTree, nodeType types.CheckExplainNodeType, text string, detail string) {
	edg := &types.CheckExplainEdge{
		SourceNodeId: parent_id,
		DestNodeId:   id,
		Message:      "",
	}
	m.edges = append(m.edges, edg)

	node := &types.CheckExplainNode{
		Id:       id,
		NodeType: nodeType,
		Text:     text,
		Detail:   detail,
		Result: &types.SearchResult{
			Authorized: tree.GetResult().Authorized,
			Explored:   tree.GetResult().Explored,
			Exhausted:  tree.GetResult().Completed,
		},
	}
	m.nodes = append(m.nodes, node)

	for _, subPath := range children {
		m.mapTree(subPath, id)
	}
}

const dotGraphName string = "GoalTree"

// DOTMapper maps a CheckExplainGraph into a DOT Graph
type DOTMapper struct {
	counter int
}

// Map converts a CheckExplainGraph into a DOT representation
// omitUknonwn can be used to skip nodes which were not fully explored
// from the final graph
func (s *DOTMapper) Map(g *types.CheckExplainGraph, omitUknown bool) (string, error) {
	graph := gographviz.NewGraph()
	graph.SetDir(true) //directed graph true
	graph.SetName(dotGraphName)

	nodeMap := make(map[string]*types.CheckExplainNode)

	for _, node := range g.Nodes {
		// skip nodes not explored
		if node.Result.Explored == false {
			continue
		}

		txt := s.sprintf("%v\\n%v", node.Text, node.Detail)
		color := s.getColor(node.Result)
		attrs := map[string]string{
			"label":     txt,
			"color":     s.sprintf(color),
			"fontcolor": s.sprintf(color),
		}
		err := graph.AddNode(dotGraphName, s.sprintf(node.Id), attrs)
		if err != nil {
			return "", fmt.Errorf("dot serialize: failed to set node: %v: %v", node.Id, err)
		}
		nodeMap[node.Id] = node
	}

	for _, edg := range g.Edges {
		srcNode, ok := nodeMap[edg.SourceNodeId]
		// if node is not found, it means we are omitting unknown
		// and the parent wasn't explored
		if !ok {
			continue
		}

		// color of edge is defined by the soruce node
		color := s.getColor(srcNode.Result)

		directed := true
		edgeAttrs := map[string]string{
			"color": s.sprintf(color),
		}
		err := graph.AddEdge(edg.SourceNodeId, edg.DestNodeId, directed, edgeAttrs)
		if err != nil {
			return "", fmt.Errorf("edg %v->%v: %w", edg.SourceNodeId, edg.DestNodeId, err)
		}
	}

	return graph.String(), nil
}

func (s *DOTMapper) getColor(result *types.SearchResult) string {
	if result.Explored == false {
		return UnknownColor
	}
	if result.Authorized {
		return OkColor
	}
	return FailedColor
}

// fmt formats a string and wraps it with quotes
func (s *DOTMapper) sprintf(str string, args ...any) string {
	return "\"" + fmt.Sprintf(str, args...) + "\""
}
