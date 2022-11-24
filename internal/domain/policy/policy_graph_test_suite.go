package policy

import (
    "testing"
    "fmt"
    "reflect"
    "strings"

    "github.com/stretchr/testify/assert"

    "github.com/sourcenetwork/source-zanzibar/pkg/tuple"
    "github.com/sourcenetwork/source-zanzibar/pkg/utils"
)


// test represents a test function
type test func(*testing.T) 

// policyGraphBuilder represents a producer function 
// that builds an a Policy Graph from a Policy
type policyGraphBuilder func(Policy) PolicyGraph

// policyGraphTestSuite represents a set of tests used to
// assert correctness of PolicyGraph implementations
type policyGraphTestSuite struct {
    g PolicyGraph
    p Policy
}

// Build a policyGraphTestSuite object from a builder
func buildTestSuite(builder policyGraphBuilder) policyGraphTestSuite {
    return policyGraphTestSuite {
        g: builder(policyFixture),
        p: policyFixture,
    }
}

// Run test suite
func (s *policyGraphTestSuite) Run(t *testing.T) {
    selfVal := reflect.ValueOf(s)
    typeT := reflect.TypeOf(s)
    for i := 0; i  < typeT.NumMethod(); i++ {
        method := typeT.Method(i)
        if strings.HasPrefix(method.Name, "Test") {
            f := method.Func
            t.Run(method.Name, func(t *testing.T) {
                tVal := reflect.ValueOf(t)
                in := []reflect.Value {selfVal, tVal}
                f.Call(in)
            })
        }
    }
}

// Standardized policy fixture used to test policyGraph
var policyFixture Policy = Policy{
    Id: "1",
    Name: "Test policy",
    Resources: []*Resource{
        BuildResource("file",
            ThisRelation("owner"),
            ThisRelation("reader"),
            BuildPerm("read", Union(CU("write"), CU("reader"))),
            BuildPerm("write", Union(CU("owner"), TTU("directory", "owner", "write"))),
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

func (s *policyGraphTestSuite) TestGetResourcesReturnAllResources(t *testing.T) {
    resources := s.g.GetResources()
    assert.Equal(t, len(s.p.Resources), len(resources))

    names := utils.MapSlice(resources, func(r Resource) string {return r.Name})
    for _, resource := range s.p.Resources {
        assert.Contains(t, names, resource.Name)
    }
}

func (s *policyGraphTestSuite) TestGetActorReturnAllActors(t *testing.T) {
    actors := s.g.GetActors()
    assert.Equal(t, len(s.p.Actors), len(actors))

    names := utils.MapSlice(actors, func(a Actor) string {return a.Name})
    for _, actor := range s.p.Actors {
        assert.Contains(t, names, actor.Name)
    }
}


func (s *policyGraphTestSuite) TestRulesAreFetchable(t *testing.T) {

    rules := []tuple.Pair[string, string]{
        tuple.NewPair("directory", "parent"),
        tuple.NewPair("directory", "owner"),
        tuple.NewPair("file", "owner"),
        tuple.NewPair("file", "reader"),
        tuple.NewPair("file", "read"),
        tuple.NewPair("file", "write"),
    }

    for _, pair := range rules {
        resource, ruleName := pair.Fst(), pair.Snd()
        testName := fmt.Sprintf("res=%s,rule=%s", resource, ruleName)
        t.Run(testName, func(t *testing.T) {
            opt := s.g.GetRule(resource, ruleName)
            assert.False(t, opt.IsEmpty())
            rule := opt.Value()
            assert.Equal(t, ruleName, rule.Name)
        })
    }
}

func (s *policyGraphTestSuite) TestAncestorsAreReachable(t *testing.T) {
    // FIXME Add check for cross resource references

    var ancestors map[string][]string = map[string][]string{
        "owner": []string{"write"},
        "reader": []string{"read"},
        "write": []string{"read"},
    }

    for rule, expected := range ancestors {
        t.Run(rule, func(t *testing.T) {
            ancestors := s.g.GetAncestors("file", rule)
            names := getNames(ancestors)
            for _, val := range expected {
                assert.Contains(t, names, val)
            }
        })
    }
}

func (s *policyGraphTestSuite) TestAncestorAcrossResourceIsReachable(t *testing.T) {
    ancestors := s.g.GetAncestors("directory", "owner")

    found := false
    for _, rule := range ancestors {
        if rule.Name == "write" {
            found = true
        }
    }
    assert.True(t, found)
}

func getNames(rules []Rule) []string {
    return utils.MapSlice(rules, func(rule Rule) string {return rule.Name})
}
