package mappers

import (
    "fmt"

	"github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
	"github.com/sourcenetwork/source-zanzibar/pkg/utils"
	"github.com/sourcenetwork/source-zanzibar/types"
	parser "github.com/sourcenetwork/source-zanzibar/internal/permission_parser"
)

type PolicyMapper struct{}

func (m *PolicyMapper) ToInternal(p types.Policy) (policy.Policy, error) {
    var resources []*policy.Resource = make([]*policy.Resource, 0, len(p.Resources))

    for _, resource := range p.Resources {
        mapped, err := toInternalResource(resource)
        if err != nil {
            return policy.Policy{}, fmt.Errorf("failed mapping policy %v: %w", p.Id, err)
        }
        resources = append(resources, mapped)
    }

	return policy.Policy{
		Id:         p.Id,
		Name:       p.Name,
                Description: p.Description,
		Resources:  resources,
		Actors:     utils.MapSlice(p.Actors, toInternalActor),
		Attributes: p.Attributes,
	}, nil
}


func  relationToRule(rel types.Relation) *policy.Rule {
	return &policy.Rule{
		Type: policy.RuleType_RELATION,
		Name: rel.Name,
                RewriteExpr: "_this",
                ExpressionTree: policy.ThisTree(),
		// TODO constraints
	}
}

func  toValidator(t policy.ActorIdType) types.Validator {
    switch t {
    case policy.ActorIdType_STRING:
        return types.Validator_STRING
    case policy.ActorIdType_NUMBER:
        return types.Validator_NUMBER
    default:
        panic("invalid ActorIdType" )
    }
}

func  fromValidator(validator types.Validator) policy.ActorIdType {
    switch validator {
    case types.Validator_STRING:
        return policy.ActorIdType_STRING
    case types.Validator_NUMBER:
        return policy.ActorIdType_NUMBER
    default:
        panic("invalid Validator" )
    }
}

func  permissionToRule(perm types.Permission) (*policy.Rule, error) {
    parseTree, err := parser.Parse(perm.Expression)
    if err != nil {
        return nil, fmt.Errorf("failed mapping permission %v: %w", perm.Name, err)
    }
    tree := ToRewriteTree(parseTree)

	return &policy.Rule{
		Type:        policy.RuleType_PERMISSION,
		Name:        perm.Name,
		RewriteExpr: perm.Expression,
                ExpressionTree: tree,
	}, nil
}

func  toInternalActor(actor types.Actor) *policy.Actor {
	return &policy.Actor{
		Name: actor.Name,
                Constraints: utils.MapSlice(actor.Validators, fromValidator),
	}
}

func  fromInternalActor(actor *policy.Actor) types.Actor {
	return types.Actor{
		Name: actor.Name,
                Validators: utils.MapSlice(actor.Constraints, toValidator),
	}
}


func  toInternalResource(res types.Resource) (*policy.Resource, error) {
	var rules []*policy.Rule

	for _, rel := range res.Relations {
		rule := relationToRule(rel)
		rules = append(rules, rule)
	}

	for _, perm := range res.Permissions {
		rule, err := permissionToRule(perm)
                if err != nil {
                    return nil, fmt.Errorf("failed mapping relation %v: %w", res.Name, err)
                }
		rules = append(rules, rule)
	}

	return &policy.Resource{
		Name:  res.Name,
		Rules: rules,
	}, nil
}

func  fromInternalResource(res *policy.Resource) types.Resource {
	return types.Resource{}
}

func  (m *PolicyMapper) FromInternal(p *policy.Policy) types.Policy {
    panic("TODO")
	return types.Policy{
		// TODO
	}
}

func  mapInternalRule(rule *policy.Rule) any {
	// TODO
	return nil
}
