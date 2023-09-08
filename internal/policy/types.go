package policy

import (
	"context"

	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

func (r *RelationFQN) Id() string {
	return r.ResourceName + "/" + r.RelationName
}

func NewRelationFQN(resourceName, relationName string) RelationFQN {
	return RelationFQN{
		ResourceName: resourceName,
		RelationName: relationName,
	}
}

type Repository interface {
	SetPolicy(context.Context, *domain.PolicyRecord) (types.RecordFound, error)
	GetPolicy(context.Context, string) (*domain.PolicyRecord, error)
	DeletePolicy(context.Context, string) (types.RecordFound, error)
	ListPolicyIds(context.Context) ([]string, error)

	SetRelationship(ctx context.Context, record *domain.RelationshipRecord) (bool, error)
	DeleteRelationship(ctx context.Context, policyId string, spec *domain.Relationship) (types.RecordFound, error)
	FindRelationships(ctx context.Context, policyId string, selector *domain.RelationshipSelector) ([]*domain.Relationship, error)
	FindRelationshipRecords(ctx context.Context, policyId string, selector *domain.RelationshipSelector) ([]*domain.RelationshipRecord, error)
	DeleteRelationships(ctx context.Context, policyId string, selector *domain.RelationshipSelector) (uint64, error)
	GetRelationship(ctx context.Context, policyId string, relationship *domain.Relationship) (*domain.RelationshipRecord, error)
}

// policy lookup table
type PolicyLookUpTable struct {
	relations map[string]*domain.Relation
	resources map[string]*domain.Resource
}

func NewPolicyLookUpTable(policy *domain.Policy) PolicyLookUpTable {
	lut := PolicyLookUpTable{
		relations: make(map[string]*domain.Relation),
		resources: make(map[string]*domain.Resource),
	}

	for _, resource := range policy.Resources {
		lut.resources[resource.Name] = resource
		for _, relation := range resource.Relations {
			key := lut.buildRelationKey(resource.Name, relation.Name)
			lut.relations[key] = relation
		}
	}
	return lut
}

func (l *PolicyLookUpTable) buildRelationKey(resource, relation string) string {
	fqn := NewRelationFQN(resource, relation)
	return fqn.String()
}

func (l *PolicyLookUpTable) GetRelation(resource, relation string) *domain.Relation {
	key := l.buildRelationKey(resource, relation)
	rel, ok := l.relations[key]
	if !ok {
		return nil
	}
	return rel
}

func (l *PolicyLookUpTable) GetResource(resourceName string) *domain.Resource {
	resource, ok := l.resources[resourceName]
	if !ok {
		return nil
	}
	return resource
}
