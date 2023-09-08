package relation_graph

import (
    "encoding/json"
    "fmt"

        "github.com/davecgh/go-spew/spew"
        "github.com/awalterschulze/gographviz"

	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

type spewSerializer struct{}

func (s *spewSerializer) Serialize(goalTree GoalTree) (string, error) {
    return spew.Sdump(goalTree), nil
}

type jsonSerializer struct{}

func (s *jsonSerializer) Serialize(goalTree GoalTree) (string, error) {
    bytes, err := json.MarshalIndent(goalTree, "", "  ")
    return string(bytes), err
}

const dotGraphName string = "GoalTree"
type dotSerializer struct{
    counter int
}

func (s *dotSerializer) Serialize(goalTree GoalTree) (string, error) {
    graph := gographviz.NewGraph()
    graph.SetDir(true) //directed graph true
    graph.SetName(dotGraphName)

    err := graph.AddNode(dotGraphName, "root", nil)
    if err != nil{
        return "", fmt.Errorf("dot serialize: failed to set root: %v", err)
    }

    err = s.handleGoalTree("root", goalTree, graph)
    if err != nil {
        return "", fmt.Errorf("dot serialize: %v")
    }

    return graph.String(), nil
}


func (s *dotSerializer) handleGoalTree(parentId string, tree GoalTree, graph *gographviz.Graph) error { 
    switch node := tree.(type) {
    case *PathNode:
        s.handlePathNode(parentId, node, graph)
    case *ORNode:
        s.handleORNode(parentId, node, graph)
    case *ANDNode:
        s.handleANDNode(parentId, node, graph)
    case *DifferenceNode:
        s.handleDifferenceNode(parentId, node, graph)
    case nil:
        break
    default:
        return fmt.Errorf("GoalTree %v: %w", node, domain.ErrInvalidVariant)
    }
    return nil
}

func (s *dotSerializer) handlePathNode(parentId string, node *PathNode, graph *gographviz.Graph) error { 
    id := s.nextId()
    label := s.sprintf("PathNode\nrelation node: %v\nreason: %v\nresult: %v", node.RelationNode.PrettyString(), node.Reason, node.Result.String())
    attrs := map[string]string{
        "label": label,
    }

    err := graph.AddNode(dotGraphName, id, attrs)
    if err != nil {
        return fmt.Errorf("node %v: %w", label, err)
    }

    directed := true
    err = graph.AddEdge(parentId, id, directed, nil)
    if err != nil {
        return fmt.Errorf("edg %v->%v: %w", parentId, id, err)
    }

    return s.handleGoalTree(id, node.Path, graph)
}

func (s *dotSerializer) handleORNode(parentId string, node *ORNode, graph *gographviz.Graph) error { 
    id := s.nextId()
    label := s.sprintf("ORNode\nresult: %v", node.Result.String())
    attrs := map[string]string{
        "label": label,
    }

    err := graph.AddNode(dotGraphName, id, attrs)
    if err != nil {
        return fmt.Errorf("node %v: %w", label, err)
    }

    directed := true
    err = graph.AddEdge(parentId, id, directed, nil)
    if err != nil {
        return fmt.Errorf("edg %v->%v: %w", parentId, id, err)
    }

    for _, child := range node.Paths{
        err = s.handleGoalTree(id, child, graph)
        if err != nil {
            return err
        }
    }
    return nil
}

func (s *dotSerializer) handleANDNode(parentId string, node *ANDNode, graph *gographviz.Graph) error {
    id := s.nextId()
    label := s.sprintf("ANDNode\nresult: %v", node.Result.String())
    attrs := map[string]string{
        "label": label,
    }

    err := graph.AddNode(dotGraphName, id, attrs)
    if err != nil {
        return fmt.Errorf("node %v: %w", label, err)
    }

    directed := true
    err = graph.AddEdge(parentId, id, directed, nil)
    if err != nil {
        return fmt.Errorf("edg %v->%v: %w", parentId, id, err)
    }

    for _, child := range node.Paths{
        err = s.handleGoalTree(id, child, graph)
        if err != nil {
            return err
        }
    }
    return nil
}

func (s *dotSerializer) handleDifferenceNode(parentId string, node *DifferenceNode, graph *gographviz.Graph) error { 
    id := s.nextId()
    label := s.sprintf("DifferenceNode\nresult: %v", node.Result.String())
    attrs := map[string]string{
        "label": label,
    }

    err := graph.AddNode(dotGraphName, id, attrs)
    if err != nil {
        return fmt.Errorf("node %v: %w", label, err)
    }

    directed := true
    err = graph.AddEdge(parentId, id, directed, nil)
    if err != nil {
        return fmt.Errorf("edg %v->%v: %w", parentId, id, err)
    }

    err = s.handleGoalTree(id, node.Left, graph)
    if err != nil {
        return err
    }

    err = s.handleGoalTree(id, node.Right, graph)
    if err != nil {
        return err
    }

    return nil
}

func (s *dotSerializer) nextId() string {
    id := s.sprintf("%v", s.counter)
    s.counter++
    return id
}

// sprintf formats a string and wraps it with quotes
func (s *dotSerializer) sprintf(format string, args ...any) string {
    format = "\"" + format + "\""
    return fmt.Sprintf(format, args...)
}


func SerializerFactory(model api.ExplainFormat) (GoalTreeSerializer, error) {
    switch model {
    case api.ExplainFormat_SPEW:
        return &spewSerializer{}, nil
    case api.ExplainFormat_JSON:
        return &jsonSerializer{}, nil
    case api.ExplainFormat_DOT:
        return &dotSerializer{}, nil
    default:
        return nil, fmt.Errorf("ExplainFormat %v: %w", model, domain.ErrInvalidVariant)
    }
}
