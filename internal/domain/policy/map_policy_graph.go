package policy

import (
    "strings"

    "github.com/sourcenetwork/source-zanzibar/pkg/mapgraph"
    opt "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

var _ PolicyGraph = (*MapPolicyGraph)(nil)

type MapPolicyGraph struct {
    graph mapgraph.MapGraph[PolicyNode]
    resources []Resource
    actors []Actor
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
        actors: actors,
        graph: mapgraph.New[PolicyNode](),
    }
    graph.buildNodes(policy)
    graph.buildEdges(policy)
    return &graph
}

func (g *MapPolicyGraph) buildNodes(policy Policy) {
    for _, resource := range policy.Resources {
        for _, rule := range resource.Rules {
            key := buildRuleKey(resource.Name, rule.Name)
            node := PolicyNode {
                Resource: resource.Name,
                Rule: *rule,
            }
            g.graph.SetNode(key, node)
        }
    }
}

func buildRuleKey(resource string, rule string) string {
    var builder strings.Builder
    builder.WriteString(resource)
    builder.WriteString("/")
    builder.WriteString(rule)
    return builder.String()
}

func (g *MapPolicyGraph) buildEdges(policy Policy) {
    for _, resource := range policy.Resources {
        for _, rule := range resource.Rules {
            source := buildRuleKey(resource.Name, rule.Name)
            leaves := rule.ExpressionTree.GetLeaves()
            for _, leaf := range leaves {
                switch rule := leaf.Rule.Rule.(type) {
                case *RewriteRule_This:
                case *RewriteRule_ComputedUserset:
                    cu := rule.ComputedUserset
                    dest := buildRuleKey(resource.Name, cu.Relation)
                    g.graph.SetEdge(source, dest)
                case *RewriteRule_TupleToUserset:
                    ttu := rule.TupleToUserset
                    dest := buildRuleKey(ttu.TuplesetNamespace, ttu.TuplesetRelation)
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
    o := g.graph.GetNode(key)
    if o.IsEmpty() {
        return opt.None[Rule]()
    }
    rule := o.Value().Rule
    return opt.Some[Rule](rule)
}

func (g *MapPolicyGraph) GetAncestors(resource, name string) []PolicyNode {
    key := buildRuleKey(resource, name)
    return g.graph.GetAncestors(key)
}
