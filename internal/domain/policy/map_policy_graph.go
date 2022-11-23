package policy

import (
    "strings"

    "github.com/sourcenetwork/sourcezanzibar/pkg/mapgraph"
    opt "github.com/sourcenetwork/sourcezanzibar/pkg/option"
)

var _ PolicyGraph = (*MapPolicyGraph)(nil)

struct MapPolicyGraph {
    graph mapgraph.MapGraph[Rule]
    resources []Resource
    actors []Actor
}

func NewMapPolicyGraph(policy Policy) PolicyGraph {
    var resources []Resource
    var actors []Actor

    for _, resource := range policy.Resources {
        resources = append(resouces, *resource)
    }

    for _, actor := range policy.Actors {
        actors = append(actors, *actor)
    }

    graph := MapPolicyGraph{
        resources: resources,
        actors: actors,
        graph: mapgraph.New[Rule](),
    }
    graph.buildNodes(policy)
    graph.buildEdges(policy)
    return &graph
}

func (g *MapPolicyGraph) buildNodes(policy Policy) {
    for _, resource := range policy.Resources {
        for _, rule := range resource.Rules {
            key := buildRuleKey(resource, rule)
            g.graph.SetNode(key, rule)
        }
    }
}

func buildRuleKey(resource *Resource, rule *Rule) string {
    var builder strings.Builder
    builder.WriteString(resource.Name)
    builder.WriteString('/')
    builder.WriteString(rule.Name)
    return builder.String()
}

func (g *MapPolicyGraph) buildEdges(policy Policy) {
    for _, resource := range policy.Resources {
        for _, rule := range resource.Rules {
            source := buildRuleKey(resource, rule)
            leaves := rule.ExpressionTree.GetLeaves()
            for _, leaf := range leaves {
                switch rule := leaf.Rule.(type) {
                case *Rule_This:
                case *Rule_ComputedUserset:
                    dest := buildRuleKey(resource, rule.relation)
                    g.graph.SetEdge(source, dest)
                case *Rule_TupleToUserset:
                    dest := buildRuleKey(rule.tuplesetNamespace, rule.tuplesetRelation)
                    g.graph.SetEdge(source, dest)
                default:
                    panic("invalid rule type")
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

func (g *MapPolicyGraph) GetRule(resource, name string) opt.Option[Rule] {
    key := buildRuleKey(resource, name)
    return g.graph.GetNode(key)
}

func (g *MapPolicyGraph) GetAncestors(resource, name string) []Rule {
    key := buildRuleKey(resource, name)
    return g.graph.GetAncestors(key)
}
