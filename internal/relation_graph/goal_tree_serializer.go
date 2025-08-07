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

type ExplainCheckTreeMapper struct{}

func (m *ExplainCheckTreeMapper) Map(tree GoalTree) *types.CheckExplainTree {
	switch node := tree.(type) {
	case *PathNode:
		return m.handlePathNode(node)
	case *ORNode:
		return m.handleOPNode(node, node.Paths, "+")
	case *ANDNode:
		return m.handleOPNode(node, node.Paths, "&")
	case *DifferenceNode:
		return m.handleOPNode(node, []GoalTree{node.Left, node.Right}, "-")
	}
	return nil
}

func (m *ExplainCheckTreeMapper) handlePathNode(tree *PathNode) *types.CheckExplainTree {
	t := &types.CheckExplainTree{
		Text:   tree.RelationNode.PrettyString(),
		Detail: "reason: " + tree.Reason,
	}

	if tree.Path == nil {
		return t
	}

	child := m.Map(tree.Path)
	children := []*types.CheckExplainTree{child}
	if tree.Result.Authorized {
		t.Ok = children
	} else if tree.Result.Completed && !tree.Result.Authorized {
		t.Failed = children
	} else {
		t.Unknown = children
	}

	return t
}

func (m *ExplainCheckTreeMapper) handleOPNode(tree GoalTree, children []GoalTree, nodeText string) *types.CheckExplainTree {
	t := &types.CheckExplainTree{
		Text:   nodeText,
		Detail: "",
	}
	for _, subPath := range children {
		child := m.Map(subPath)
		if subPath.GetResult().Authorized {
			t.Ok = append(t.Ok, child)
		} else if subPath.GetResult().Completed && !subPath.GetResult().Authorized {
			t.Failed = append(t.Failed, child)
		} else {
			t.Unknown = append(t.Unknown, child)
		}
	}
	return t
}

const dotGraphName string = "GoalTree"

// DotSerializer maps a CheckExplainTree into a DOT Graph
type DotSerializer struct {
	counter int
}

// Serialize converts a CheckExplainTree into a DOT representation
// omitUknonwn can be used to skip nodes which were not fully explored
// from the final graph
func (s *DotSerializer) Serialize(tree *types.CheckExplainTree, omitUknown bool) (string, error) {
	graph := gographviz.NewGraph()
	graph.SetDir(true) //directed graph true
	graph.SetName(dotGraphName)

	err := graph.AddNode(dotGraphName, "root", nil)
	if err != nil {
		return "", fmt.Errorf("dot serialize: failed to set root: %v", err)
	}

	err = s.step("root", tree, graph, RootColor, omitUknown)
	if err != nil {
		return "", fmt.Errorf("dot serialize: %v", err)
	}

	return graph.String(), nil
}

func (s *DotSerializer) step(parentId string, node *types.CheckExplainTree, graph *gographviz.Graph, color string, omitUknown bool) error {
	id := s.nextId()
	attrs := map[string]string{
		"label":     s.sprintf("%v\\n%v", node.Text, node.Detail),
		"color":     s.sprintf(color),
		"fontcolor": s.sprintf(color),
	}

	err := graph.AddNode(dotGraphName, id, attrs)
	if err != nil {
		return fmt.Errorf("node %v: %w", node.Text, err)
	}

	directed := true
	edgeAttrs := map[string]string{
		"color": s.sprintf(color),
	}
	err = graph.AddEdge(parentId, id, directed, edgeAttrs)
	if err != nil {
		return fmt.Errorf("edg %v->%v: %w", parentId, id, err)
	}

	for _, child := range node.Failed {
		err = s.step(id, child, graph, FailedColor, omitUknown)
		if err != nil {
			return err
		}
	}
	for _, child := range node.Ok {
		err = s.step(id, child, graph, OkColor, omitUknown)
		if err != nil {
			return err
		}
	}
	if !omitUknown {
		for _, child := range node.Unknown {
			err = s.step(id, child, graph, UnknownColor, omitUknown)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *DotSerializer) nextId() string {
	id := s.sprintf("%v", s.counter)
	s.counter++
	return id
}

// sprintf formats a string and wraps it with quotes
func (s *DotSerializer) sprintf(format string, args ...any) string {
	format = "\"" + format + "\""
	return fmt.Sprintf(format, args...)
}
