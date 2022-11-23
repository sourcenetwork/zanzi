package policy

import (
    opt "github.com/sourcenetwork/sourcezanzibar/pkg/option"
)

// PolicyGraph models the definitinos in a policy as a graph
// meaning, the hieararchy between different permissions or relations
type PolicyGraph interface {
    // policy graph nodes are identified by the resource name + relation/perm name
    // policy graph nodes store the relation / perm definition
    // there should be a single internal type for a relation permission as they are equivalent
    GetResources() []Resource
    GetActors() []Actor
    GetRule(resource, name string) opt.Option[Rule]
    GetAncestors(resource, name string) []Rule
}

// PolicyStore abstract interfacing with namespace storage.
type PolicyStore interface {
	GetPolicy(policyId string) (opt.Option[model.Policy], error)

	SetPolicy(policy model.Policy) error

	DeletePolicy(policyId string) error

        GetPolicyGraph(policyId string) (opt.Option[PolicyGraph], error)
}
