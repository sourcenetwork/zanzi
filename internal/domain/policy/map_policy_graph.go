package policy

import (
	"strings"

	"github.com/sourcenetwork/zanzi/pkg/mapgraph"
	opt "github.com/sourcenetwork/zanzi/pkg/option"
)

var _ PolicyGraph = (*MapPolicyGraph)(nil)

type MapPolicyGraph struct {
	graph     mapgraph.MapGraph[PolicyNode]
	resources []Resource
	actors    []Actor
}

func NewMapPolicyGraph(policy Policy) PolicyGraph {
	var resources []Resource
	var actors []Actor

	for _, resource := range policy.Resources {
		resources = append(resources, *resource)
	}

	for _, actor := range policy.Actors {
		actors = append(actors, *actor)
	}

	graph := MapPolicyGraph{
		resources: resources,
		actors:    actors,
		graph:     mapgraph.New[PolicyNode](),
	}
	graph.buildNodes(policy)
	graph.buildEdges(policy)
	return &graph
}

func (g *MapPolicyGraph) buildNodes(policy Policy) {
	for _, resource := range policy.Resources {
		for _, rule := range resource.Relations {
			key := buildRelationKey(resource.Name, rule.Name)
			node := PolicyNode{
				Resource: resource.Name,
				Relation: *rule,
			}
			g.graph.SetNode(key, node)
		}
	}
}

func buildRelationKey(resource string, rule string) string {
	var builder strings.Builder
	builder.WriteString(resource)
	builder.WriteString("/")
	builder.WriteString(rule)
	return builder.String()
}

func (g *MapPolicyGraph) buildEdges(policy Policy) {
	for _, resource := range policy.Resources {
		for _, relation := range resource.Relations {
			source := buildRelationKey(resource.Name, relation.Name)
			leaves := relation.ExpressionTree.GetLeaves()
			for _, leaf := range leaves {
				switch relation := leaf.Rule.RewriteRule.(type) {
				case *RewriteRule_This:
				case *RewriteRule_ComputedUserset:
					cu := relation.ComputedUserset
					dest := buildRelationKey(resource.Name, cu.Relation)
					g.graph.SetEdge(source, dest)
				case *RewriteRule_TupleToUserset:
					ttu := relation.TupleToUserset
					dest := buildRelationKey(ttu.CuRelationNamespace, ttu.CuRelation)
					g.graph.SetEdge(source, dest)
				default:
					panic("invalid relation type")
				}
			}
		}
	}
}

func (g *MapPolicyGraph) GetResources() []Resource {
	return g.resources
}

func (g *MapPolicyGraph) GetActors() []Actor {
	return g.actors
}

func (g *MapPolicyGraph) GetRelation(resource, name string) opt.Option[Relation] {
	key := buildRelationKey(resource, name)
	o := g.graph.GetNode(key)
	if o.IsEmpty() {
		return opt.None[Relation]()
	}
	relation := o.Value().Relation
	return opt.Some[Relation](relation)
}

func (g *MapPolicyGraph) GetAncestors(resource, name string) []PolicyNode {
	key := buildRelationKey(resource, name)
	return g.graph.GetAncestors(key)
}
