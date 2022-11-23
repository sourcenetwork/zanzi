package policy

import (
    "testing"
    "fmt"

    "github.com/stretchr/testify/assert"
)

// General test suite for Policy Graph implementations

type policyGraphBuilder func(Policy) PolicyGraph

var policyFixture Policy = Policy{
    Id: "1",
    Name: "Test policy",
    Descripton: "test",
    Resources: []*Resource{
        BuildResource("file",
            ThisRelation("owner"),
            ThisRelation("reader"),
            BuildPerm("read", Union(CU("write"), CU("reader"))),
            BuildPerm("write", Union(CU("owner"), TTU("directory", "owner", "write")),
        ),

        BuildResource("directory",
            ThisRelation("owner"),
            ThisRelation("parent"),
        ),
    },
    Actors: []*Actor{
        BuildActor("user"),
    },
}

func testGetResourcesReturnAllResources(g PolicyGraph, t *testing.T) {
    resources := g.GetResources()

    assert.Equal(len(policy.Resources), len(resources)
    for _, resouce := range policy.Resources {
        assert.Contains(t, resources, resource)
    }
}

func testGetActorReturnAllActors(g PolicyGraph, t *testing.T) {
    actors := g.GetActors()

    assert.Equal(len(policy.Actors), len(actors)
    for _, actor := range policy.Actors {
        assert.Contains(t, actors, actor)
    }
}


func testRulesAreFetchable(g PolicyGraph, t *testing.T) {

    rules := [](string, string) {
        ("directory", "parent"),
        ("directory", "owner"),
        ("file", "owner"),
        ("file", "reader"),
        ("file", "read"),
        ("file", "write"),
    }

    for _, (resource, ruleName) := range rules {
        testName = fmt.Sprintf("res: %s rule: %s", resource, ruleName)
        t.Run(testName, func(t *testing.T) {
            opt := g.GetRule(resource, ruleName)
            assert.False(opt.IsEmpty())
            rule := opt.Value()
            assert.Equal(ruleName, rule.Name)
        }
    }
}

func testAncestorsAreReachable(g PolicyGraph, t *testing.T) {
    // FIXME Add check for cross resource references

    var ancestors map[string][]string {
        "owner": []string{"write"},
        "reader": []string{"read"},
        "write": []string{"read"},
    }

    for rule, expected := range ancestors {
        t.Run(rule, func(t *testing.T) {
            ancestors := g.GetAncestors("file", rule)
            names := getNames(ancestors)
            for _, val := range expected {
                assert.Contains(t, names, val)
            }
        }
    }
}

func getNames(rules []Rule) []string {
    return utils.MapSlice(rules, func(rule Rule) string {return rule.Name})
}

func policyGraphTestSuite(builder policyGraphBuilder, t *testing.T) {
    g := builder(policyFixture)

    t.Run("testGetResourcesReturnAllResources", func(t *testing.T) {testGetResourcesReturnAllResources(g, t)})
    t.Run("testGetActorReturnAllActors", func(t *testing.T) {testGetActorReturnAllActors(g, t)})
    t.Run("testRulesAreFetchable", func(t *testing.T) {testRulesAreFetchable(g, t)})
    t.Run("testAncestorsAreReachable", func(t *testing.T) {testAncestorsAreReachable(g, t)})
}
