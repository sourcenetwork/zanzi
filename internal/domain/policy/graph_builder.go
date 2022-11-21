package policy

import (
    "strings"
)

func Edge(from, to) { }

func Node(node N) { }

func BuildGraph(policy Policy) {
    for _, resource := range policy.Resources {
        for _, rule := range resource.Rules {
            var builder strings.Builder
            builder.WriteString(resource.Name)
            builder.WriteString('/')
            builder.WriteString(rule.Name)
            key := builder.String()
        }
    }

}

func buildNodes(policy Policy) {
    nodes := make(map[string]any)
    for _, resource := range policy.Resources {
        for _, rule := range resource.Rules {
            key := buildRuleKey(resource, rule)
            nodes[key] = rule
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

func buildEdges(policy Policy) {
    edges := [](string, string)
    for _, resource := range policy.Resources {
        for _, rule := range resource.Rules {
            key := buildRuleKey(resource, rule)
            leaves := rule.ExpressionTree.GetLeaves()
            for _, leaf := range leaves {
                switch rule := leaf.Rule.(type) {
                case *Rule_This:
                case *Rule_ComputedUserset:
                    target := buildRuleKey(resource, rule.relation)
                    edges = append(edges, (key, target))
                case *Rule_TupleToUserset:
                    target := buildRuleKey(rule.tuplesetNamespace, rule.tuplesetRelation)
                    edges = append(edges, (key, target))
                default:
                    panic("invalid rule type")
                }
            }
        }
    }
}
