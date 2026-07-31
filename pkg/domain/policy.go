package domain

import "github.com/sourcenetwork/zanzi/internal/utils"

func NewPolicyRecord(policy *Policy, data []byte) *PolicyRecord {
	return &PolicyRecord{
		Policy:  policy,
		AppData: data,
	}
}

// GetResourceByName returns the named resource. If not found returns nil
func (p *Policy) GetResourceByName(name string) *Resource {
	for _, resource := range p.Resources {
		if resource.Name == name {
			return resource
		}
	}
	return nil
}

// GetResourcesNames returns a slice of resource names contained in the policy
func (p *Policy) GetResourcesNames() []string {
	return utils.MapSlice(p.Resources, func(r *Resource) string { return r.Name })
}

// GetRelationsNames returns a slice of relation names contained in the resource
func (r *Resource) GetRelationsNames() []string {
	return utils.MapSlice(r.Relations, func(r *Relation) string { return r.Name })
}
