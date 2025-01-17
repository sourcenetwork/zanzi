package kv_store

import (
	"context"

	rcdb "github.com/sourcenetwork/raccoondb"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/utils"
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

// newRelationshipFetcher returns a fetcher for a given raccon store
func newRelationshipFetcher(store rcdb.RaccoonStore[*Relationship, *RelationNode], mapper relationshipMapper) *relationshipFetcher {
	return &relationshipFetcher{
		store:  store,
		mapper: mapper,
	}
}

// relationshipFetcher is a helper which returns relationships from the relation graph, given a selector
// fetcher inspects a selector and optimizes well-known access patterns
type relationshipFetcher struct {
	store  rcdb.RaccoonStore[*Relationship, *RelationNode]
	mapper relationshipMapper
}

// Fetch returns a set of relationships for a given Selector
func (f *relationshipFetcher) Fetch(ctx context.Context, selector *domain.RelationshipSelector) ([]*domain.Relationship, error) {
	// with a given object/rel pair, we can directly lookup successors
	if selector.ObjectSelector.GetObjectSpec() != nil && selector.RelationSelector.GetRelationName() != "" {
		entity := selector.ObjectSelector.GetObjectSpec()
		relation := selector.RelationSelector.GetRelationName()
		return f.fetchSucessor(ctx, entity, relation, selector.SubjectSelector)
	} else {
		return f.scan(ctx, selector)
	}
}

// fetchSucessors optmizes a scan by scanning only a node's sucessors and applying the subject predicate
func (f *relationshipFetcher) fetchSucessor(ctx context.Context, obj *domain.Entity, relation string, subjectSelector *domain.SubjectSelector) ([]*domain.Relationship, error) {
	node := RelationNode{
		Resource: obj.Resource,
		Id:       obj.Id,
		Relation: relation,
		Type:     RelationType_OBJECT_SET,
	}
	sucessors, err := f.store.GetSucessors(&node)
	if err != nil {
		return nil, err
	}

	// after narrowing down the possible sucessors
	// apply a selector spec to filter out subjects
	builder := domain.SelectorBuilder{}
	selector := builder.AnyObject().AnyRelation().AnySubject().Build()
	selector.SubjectSelector = subjectSelector

	spec, err := policy.NewSelectorSpec(&selector)
	if err != nil {
		return nil, err
	}

	mapped := utils.MapSlice(sucessors, func(relationship *Relationship) *domain.Relationship {
		mapped := f.mapper.FromInternal(relationship)
		return &mapped
	})
	filtered := utils.Filter(mapped, func(relationship *domain.Relationship) bool { return spec.Satisfies(relationship) })

	return filtered, nil
}

// scan performs a full table scan and applies the spec mapped from the selector
// to filter out relationships
func (f *relationshipFetcher) scan(ctx context.Context, selector *domain.RelationshipSelector) ([]*domain.Relationship, error) {
	rels, err := f.store.List()
	if err != nil {
		return nil, err
	}

	spec, err := policy.NewSelectorSpec(selector)
	if err != nil {
		return nil, err
	}

	mapped := utils.MapSlice(rels, func(relationship *Relationship) *domain.Relationship {
		mapped := f.mapper.FromInternal(relationship)
		return &mapped
	})
	filtered := utils.Filter(mapped, func(relationship *domain.Relationship) bool { return spec.Satisfies(relationship) })
	return filtered, nil
}
