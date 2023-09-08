package policy

import (
	"fmt"
	"reflect"

	"github.com/sourcenetwork/zanzi/pkg/domain"
)

type predicate func(*domain.Relationship) bool

type RelationshipSelectorSpec struct {
	selector          domain.RelationshipSelector
	objectPredicate   predicate
	relationPredicate predicate
	subjectPredicate  predicate
}

func NewSelectorSpec(selector *domain.RelationshipSelector) (RelationshipSelectorSpec, error) {
	spec := RelationshipSelectorSpec{}

	switch objectSelector := selector.ObjectSelector.Selector.(type) {
	case *domain.ObjectSelector_ObjectSpec:
		spec.objectPredicate = spec.objectSpec(objectSelector.ObjectSpec)
	case *domain.ObjectSelector_Wildcard:
		spec.objectPredicate = spec.objectWildcardSpec()
	case *domain.ObjectSelector_ResourceSpec:
		spec.objectPredicate = spec.objectResourceSpec(objectSelector.ResourceSpec)
	default:
		return spec, fmt.Errorf("ObjectSelector %v: %w", objectSelector, domain.ErrInvalidVariant)
	}

	switch relationSelector := selector.RelationSelector.Selector.(type) {
	case *domain.RelationSelector_RelationName:
		spec.relationPredicate = spec.relationNameSpec(relationSelector.RelationName)
	case *domain.RelationSelector_Wildcard:
		spec.relationPredicate = spec.relationWildcardSpec()
	default:
		return spec, fmt.Errorf("RelationSelector %v: %w", relationSelector, domain.ErrInvalidVariant)
	}

	switch subjectSelector := selector.SubjectSelector.Selector.(type) {
	case *domain.SubjectSelector_SubjectSpec:
		spec.subjectPredicate = spec.subjectSpec(subjectSelector.SubjectSpec)
	case *domain.SubjectSelector_Wildcard:
		spec.subjectPredicate = spec.subjectWildcardSpec()
	case *domain.SubjectSelector_ResourceSpec:
		spec.subjectPredicate = spec.subjectResourceSpec(subjectSelector.ResourceSpec)
	default:
		return spec, fmt.Errorf("SubjectSelector %v: %w", subjectSelector, domain.ErrInvalidVariant)
	}

	return spec, nil
}

func (s *RelationshipSelectorSpec) Satisfies(relationship *domain.Relationship) bool {
	return s.objectPredicate(relationship) && s.relationPredicate(relationship) && s.subjectPredicate(relationship)
}

func (s *RelationshipSelectorSpec) objectSpec(entity *domain.Entity) predicate {
	return func(r *domain.Relationship) bool {
		return reflect.DeepEqual(entity, r.Object)
	}
}

func (s *RelationshipSelectorSpec) objectWildcardSpec() predicate {
	return func(r *domain.Relationship) bool {
		return true
	}
}

func (s *RelationshipSelectorSpec) objectResourceSpec(resourceName string) predicate {
	return func(r *domain.Relationship) bool {
		return r.Object.Resource == resourceName
	}
}

func (s *RelationshipSelectorSpec) relationNameSpec(relation string) predicate {
	return func(r *domain.Relationship) bool {
		return r.Relation == relation
	}
}

func (s *RelationshipSelectorSpec) relationWildcardSpec() predicate {
	return func(r *domain.Relationship) bool {
		return true
	}
}

func (s *RelationshipSelectorSpec) subjectSpec(subject *domain.Subject) predicate {
	return func(r *domain.Relationship) bool {
		return reflect.DeepEqual(r.Subject, subject)
	}
}

func (s *RelationshipSelectorSpec) subjectWildcardSpec() predicate {
	return func(r *domain.Relationship) bool {
		return true
	}
}

func (s *RelationshipSelectorSpec) subjectResourceSpec(resourceName string) predicate {
	return func(r *domain.Relationship) bool {
		switch subject := r.Subject.Subject.(type) {
		case *domain.Subject_Entity:
			return subject.Entity.Resource == resourceName
		case *domain.Subject_EntitySet:
			return subject.EntitySet.Entity.Resource == resourceName
		case *domain.Subject_ResourceSet:
			return subject.ResourceSet.ResourceName == resourceName
		default:
			return false
		}
	}
}
