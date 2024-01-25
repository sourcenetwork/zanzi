package domain

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewPolicyRecord(policy *Policy, data []byte) *PolicyRecord {
	return &PolicyRecord{
		Policy:    policy,
		AppData:   data,
		CreatedAt: timestamppb.Now(),
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
