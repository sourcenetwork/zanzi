package services

import (
    _ "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
    "github.com/sourcenetwork/source-zanzibar/types"
    "github.com/sourcenetwork/source-zanzibar/pkg/utils"
)

// TODO finish
type PolicyMapper struct {}

func (m *PolicyMapper) ToInternal(p types.Policy) policy.Policy {
    resMapper := func(res types.Resource) *policy.Resource {return m.toInternalResource(res)}
    actorMapper := func(a types.Actor) *policy.Actor {return m.toInternalActor(a)}
    return policy.Policy{
        Id: p.Id,
        Name: p.Name,
        Resources: utils.MapSlice(p.Resources, resMapper),
        Actors: utils.MapSlice(p.Actors, actorMapper),
        Attributes: p.Attributes,
    }
}

func (m *PolicyMapper) FromInternal(p *policy.Policy) types.Policy {
    return types.Policy {
        // TODO
    }
}


func (m *PolicyMapper) mapInternalRule(rule *policy.Rule) any {
    // TODO
    return nil
}

func (m *PolicyMapper) relationToRule(rel types.Relation) *policy.Rule {
    return &policy.Rule {
        Type: policy.RuleType_RELATION,
        Name: rel.Name,
        // TODO map constraint
    }
}

func (m *PolicyMapper) permissionToRule(perm types.Permission) *policy.Rule {
    return &policy.Rule {
        Type: policy.RuleType_PERMISSION,
        Name: perm.Name,
        RewriteExpr: perm.Expression,
        // TODO call parser in order to build tree
    }
}

func (m *PolicyMapper) toInternalActor(actor types.Actor) *policy.Actor {
    return &policy.Actor {
        Name: actor.Name,
        // TODO constraints
    }
}

func (m *PolicyMapper) fromInternalActor(actor *policy.Actor) types.Actor {
    return types.Actor {
        Name: actor.Name,
        // TODO map kinds
    }
}

func (m *PolicyMapper) addRules() {

}

func (m *PolicyMapper) toInternalResource(res types.Resource) *policy.Resource {
    var rules []*policy.Rule

    for _, rel := range res.Relations {
        rule := m.relationToRule(rel)
        rules = append(rules, rule)
    }

    for _, perm := range res.Permissions {
        rule := m.permissionToRule(perm)
        rules = append(rules, rule)
    }

    return &policy.Resource {
        Name: res.Name,
        Rules: rules,
    }
}

func (m *PolicyMapper) fromInternalResource(res *policy.Resource) types.Resource {
    return types.Resource {

    }
}
