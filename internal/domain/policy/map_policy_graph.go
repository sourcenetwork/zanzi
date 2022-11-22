package policy

import (
    "strings"
)

var _ PolicyGraph = (*MapPolicyGraph)(nil)

struct MapPolicyGraph {
    nodes map[string]Rule
    edges map[string]map[string]struct{}
    idxMap map[string]int64
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
        nodes: make(map[string]Rule),
        edges: make(map[string]map[string]struct{}),
    }
    graph.buildNodes(policy)
    graph.buildEdges(policy)
    return &graph
}

func (g *MapPolicyGraph) buildNodes(policy Policy) {
    for _, resource := range policy.Resources {
        for _, rule := range resource.Rules {
            key := buildRuleKey(resource, rule)
            g.nodes[key] = rule
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
            key := buildRuleKey(resource, rule)
            leaves := rule.ExpressionTree.GetLeaves()
            for _, leaf := range leaves {
                switch rule := leaf.Rule.(type) {
                case *Rule_This:
                case *Rule_ComputedUserset:
                    target := buildRuleKey(resource, rule.relation)
                    g.edges[key][target] = struct{}{}
                case *Rule_TupleToUserset:
                    target := buildRuleKey(rule.tuplesetNamespace, rule.tuplesetRelation)
                    g.edges[key][target] = struct{}{}
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

func (g *MapPolicyGraph) GetRule(resource, name string) types.Option[Rule] {
    key := buildRuleKey(resource, name)
    rule, ok := g.nodes[key]
    if !ok {
        return utils.Empty[Rule]()
    }

    return utils.Some[Rule](rule)
}

func (g *MapPolicyGraph) GetAncestors(resource, name string) []Rule {
    sourceKey := buildRuleKey(resource, name)

    var ancestors []Rule
    ancestorMap, ok := g.edges[sourceKey]
    if !ok {
        return ancestors
    }

    for _, ruleKey := range ancestorMap {
        rule := g.nodes[ruleKey]
        ancestors = append(ancestors, rule)
    }

    return ancestors
}
