package relation_graph

import (
	"github.com/awalterschulze/gographviz"

	"github.com/sourcenetwork/zanzi/pkg/domain"
)

const relationTreeGraphName string = "Relationships"
const directed bool = true

type RelationTreeDOTSerializer struct{}

func (s *RelationTreeDOTSerializer) Serialize(tree *domain.RelationTree) (string, error) {
	graph := gographviz.NewGraph()
	graph.SetDir(directed) //directed graph true
	graph.SetName(relationTreeGraphName)

	err := s.handleNode(nil, tree, graph)
	if err != nil {
		return "", err
	}

	return graph.String(), nil
}

func (s *RelationTreeDOTSerializer) handleNode(parent *domain.RelationTree, node *domain.RelationTree, graph *gographviz.Graph) error {
	id := s.nodeId(node.Node)
	label := s.nodeLabel(node.Node)

	err := graph.AddNode(relationTreeGraphName, id, map[string]string{
		"label": label,
	})
	if err != nil {
		return err // TODO wrap
	}

	if parent != nil {
		parentId := s.nodeId(parent.Node)
		err := graph.AddEdge(parentId, id, directed, nil)
		if err != nil {
			return err
		}
	}

	for _, child := range node.Children {
		err := s.handleNode(node, child, graph)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *RelationTreeDOTSerializer) nodeId(node *domain.RelationNode) string {
	if node == nil {
		return s.wrap("root")
	}
	return s.wrap(node.Id())
}

func (s *RelationTreeDOTSerializer) nodeLabel(node *domain.RelationNode) string {
	if node == nil {
		return s.wrap("root")
	}
	return s.wrap(node.PrettyString())
}

func (s *RelationTreeDOTSerializer) wrap(str string) string {
	return "\"" + str + "\""
}
