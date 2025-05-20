package policy

import (
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/errors"
)

type ValidPolicySpec struct{}

func (s *ValidPolicySpec) Verify(policy *domain.Policy) error {
	err := s.verify(policy)
	if err != nil {
		return errors.Wrap("invalid policy", err,
			errors.Pair(errors.AttrPolicy, policy.Id))
	}
	return nil
}

func (s *ValidPolicySpec) verify(policy *domain.Policy) error {
	lut := NewPolicyLookUpTable(policy)

	if err := s.UniqueResourceAndRelationsNames(policy); err != nil {
		return err
	}

	if err := s.RequiredFieldsPresent(policy); err != nil {
		return err
	}

	if err := s.SubjectRestrictionsConsistent(policy, lut); err != nil {
		return err
	}

	if err := s.RelationExpressionWellDefined(policy, lut); err != nil {
		return err
	}

	return nil
}

func (s *ValidPolicySpec) RequiredFieldsPresent(policy *domain.Policy) error {
	return policy.Validate()
}

func (s *ValidPolicySpec) UniqueResourceAndRelationsNames(policy *domain.Policy) error {
	names := make(map[string]struct{})
	for _, resource := range policy.Resources {
		name := resource.Name
		_, ok := names[name]
		if ok {
			return errors.Wrap("duplicated resource definition", errors.BadInput,
				errors.Pair(errors.AttrResource, name))
		}
		names[name] = struct{}{}

		err := s.UniqueRelationNameForResource(resource)
		if err != nil {
			return errors.Attrs(err, errors.Pair(errors.AttrResource, name))
		}
	}
	return nil
}

func (s *ValidPolicySpec) UniqueRelationNameForResource(resource *domain.Resource) error {
	names := make(map[string]struct{})

	for _, relation := range resource.Relations {
		name := relation.Name
		_, ok := names[name]
		if ok {
			return errors.Wrap("duplicated relation definition",
				errors.BadInput, errors.Pair("relation", name))
		}
		names[name] = struct{}{}
	}
	return nil
}

func (s *ValidPolicySpec) RelationExpressionWellDefined(policy *domain.Policy, lut PolicyLookUpTable) error {
	for _, resource := range policy.Resources {
		for _, relation := range resource.Relations {
			tree, err := GetExpressionTree(relation)
			if err != nil {
				return errors.Attrs(err,
					errors.Pair(errors.AttrResource, resource.Name),
					errors.Pair(errors.AttrRelation, relation.Name),
				)
			}
			rules := tree.GetRules()
			for _, rule := range rules {
				err := s.policyContainsRelationsInRule(resource.Name, relation.Name, lut, rule)
				if err != nil {
					return errors.Attrs(err, errors.Pair("rule", rule.String()))
				}
			}
		}
	}
	return nil
}

func (s *ValidPolicySpec) policyContainsRelationsInRule(resource, relation string, lut PolicyLookUpTable, rule *domain.Rule) error {
	switch r := rule.Rule.(type) {
	case *domain.Rule_Cu:
		target := r.Cu.TargetRelation
		rel := lut.GetRelation(resource, target)
		if rel == nil {
			return errors.Wrap("resource does not have relation", ErrRelExpTree,
				errors.Pair(errors.AttrRelation, target),
				errors.Pair(errors.AttrResource, resource),
			)
		}
	case *domain.Rule_This:
	case *domain.Rule_Ttu:
		target := r.Ttu.TuplesetRelation
		tuplesetRel := lut.GetRelation(resource, target)
		if tuplesetRel == nil {
			return errors.Wrap("resource does not have relation", ErrRelExpTree,
				errors.Pair(errors.AttrRelation, target),
				errors.Pair(errors.AttrResource, resource),
			)
		}

		// TODO this check is trickier, as I have to consider
		// every allowed resource type in the subject restriction graph.
		// tbh it's pretty simple.
		//cuRel := lut.GetRelation(resource, r.Ttu.ComputedUsersetRelation)
	default:
		return errors.Wrap("invalid rule variant", errors.ErrInvalidVariant)
	}
	return nil
}

func (s *ValidPolicySpec) SubjectRestrictionsConsistent(policy *domain.Policy, lut PolicyLookUpTable) error {
	for _, resource := range policy.Resources {
		for _, relation := range resource.Relations {
			var err error
			if relation.SubjectRestriction != nil {
				err = s.ValidSubjectRestriction(relation.SubjectRestriction, lut)
			}
			if err != nil {
				return errors.Attrs(err,
					errors.Pair(errors.AttrResource, resource.Name),
					errors.Pair(errors.AttrRelation, relation.Name),
				)
			}
		}
	}
	return nil
}

func (s *ValidPolicySpec) ValidSubjectRestriction(restriction *domain.SubjectRestriction, lut PolicyLookUpTable) error {
	switch restrictionSet := restriction.SubjectRestriction.(type) {
	case *domain.SubjectRestriction_RestrictionSet:
		return s.validSubjectRestrictionSet(restrictionSet.RestrictionSet, lut)
	case *domain.SubjectRestriction_UniversalSet:
		return nil
	default:
		return errors.Wrap("invalid subject resitrction", errors.ErrInvalidVariant)
	}
}

func (s *ValidPolicySpec) validSubjectRestrictionSet(restrictionSet *domain.SubjectRestrictionSet, lut PolicyLookUpTable) error {
	for _, elem := range restrictionSet.Restrictions {
		switch restriction := elem.Entry.(type) {
		case *domain.SubjectRestrictionSet_Restriction_Entity:
			if lut.GetResource(restriction.Entity.ResourceName) == nil {
				return errors.Wrap("resource not found", errors.BadInput,
					errors.Pair(errors.AttrResource, restriction.Entity.ResourceName))
			}
		case *domain.SubjectRestrictionSet_Restriction_EntitySet:
			resource, relation := restriction.EntitySet.ResourceName, restriction.EntitySet.RelationName
			if lut.GetRelation(resource, relation) == nil {
				return errors.Wrap("relation not found", errors.BadInput,
					errors.Pair(errors.AttrResource, resource),
					errors.Pair(errors.AttrRelation, relation))
			}
		default:
			return errors.Wrap("invalid subject resitrction", errors.ErrInvalidVariant)
		}
	}
	return nil
}

// models predicates that a valid relationship must satisfy in order
// to be added to a policy
type AllowedRelationshipSpec struct {
}

// ValidateRelationship checks the relationship contains the rquried fields
func (s *AllowedRelationshipSpec) requiredFields(relationship *domain.Relationship) error {
	err := relationship.Validate()
	if err != nil {
		return errors.Wrap(err.Error(), errors.BadInput)
	}

	if relationship.Object.Id == "" {
		return errors.Wrap("invalid relationship: object id cannot be empty", errors.BadInput)
	}
	return nil
}

// checkSymbols verifies that the symbols referred to in a Relationship are valid within a Policy
// eg. validates the resources and relation names
func (s *AllowedRelationshipSpec) checkSymbols(relationship *domain.Relationship, lut PolicyLookUpTable) error {
	resName, relName := relationship.Object.Resource, relationship.Relation

	res := lut.GetResource(resName)
	if res == nil {
		return errors.Wrap("resource not found", errors.BadInput, errors.Pair(errors.AttrResource, resName))
	}

	rel := lut.GetRelation(resName, relName)
	if rel == nil {
		return errors.Wrap("relation not found", errors.BadInput,
			errors.Pair(errors.AttrResource, resName),
			errors.Pair(errors.AttrRelation, relName))
	}

	switch s := relationship.Subject.Subject.(type) {
	case *domain.Subject_Entity:
		res = lut.GetResource(s.Entity.Resource)
		if res == nil {
			return errors.Wrap("resource not found", errors.BadInput, errors.Pair(errors.AttrResource, s.Entity.Resource))
		}
	case *domain.Subject_EntitySet:
		res, rel := s.EntitySet.Entity.Resource, s.EntitySet.Relation
		resource := lut.GetRelation(res, rel)
		if resource == nil {
			return errors.Wrap("relation not found", errors.BadInput,
				errors.Pair(errors.AttrResource, res),
				errors.Pair(errors.AttrRelation, rel))
		}
	case *domain.Subject_ResourceSet:
		resource := lut.GetResource(s.ResourceSet.ResourceName)
		if resource == nil {
			return errors.Wrap("resource not found", errors.BadInput, errors.Pair(errors.AttrResource, resource.Name))
		}
	default:
		return errors.Wrap("subject", errors.ErrInvalidVariant)
	}

	return nil
}

func (s *AllowedRelationshipSpec) Satisfies(relationship *domain.Relationship, lut PolicyLookUpTable) error {
	err := s.satisfies(relationship, lut)
	if err != nil {
		return errors.Wrap("invalid relationship", err, errors.Pair(errors.AttrPolicy, lut.GetId()))
	}

	return nil
}

func (s *AllowedRelationshipSpec) satisfies(relationship *domain.Relationship, lut PolicyLookUpTable) error {
	if err := s.requiredFields(relationship); err != nil {
		return err
	}
	if err := s.checkSymbols(relationship, lut); err != nil {
		return err
	}

	subjectRestrictionSpec := SubjectRestrictionSpec{}
	if err := subjectRestrictionSpec.Satisfies(relationship, lut); err != nil {
		return err
	}

	return nil
}

// SubjectRestrictionSpec is a Specification specialized to validate whether a Relationship's Subject
// is allowed by a Policy's SubjectRestriction rules
type SubjectRestrictionSpec struct{}

// Satisfies checks whether relationship.Subject is allowed by the Policy's SubjectRestriction rules
func (s *SubjectRestrictionSpec) Satisfies(relationship *domain.Relationship, table PolicyLookUpTable) error {
	relation := table.GetRelation(relationship.Object.Resource, relationship.Relation)
	if relation == nil {
		return errors.Wrap("relation not found in policy", errors.BadInput,
			errors.Pair(errors.AttrPolicy, table.GetId()),
			errors.Pair(errors.AttrRelation, relation.Name))
	}

	switch restriction := relation.SubjectRestriction.SubjectRestriction.(type) {
	case *domain.SubjectRestriction_UniversalSet:
		return nil
	case *domain.SubjectRestriction_RestrictionSet:
		ok := s.satisfiesRestrictionSet(relationship.Subject, restriction.RestrictionSet)
		if !ok {
			return errors.Wrap("invalid relationship: subject violates subject restriction in relation: double check allowed subjects for relation the policy definition",
				errors.BadInput,
				errors.Pair(errors.AttrPolicy, table.GetId()),
				errors.Pair(errors.AttrResource, relationship.Object.Resource),
				errors.Pair(errors.AttrRelation, relationship.Relation),
			)
		}
		return nil
	default:
		return errors.Wrap("invalid relation subject restriction", errors.ErrInvalidVariant)
	}
}

func (s *SubjectRestrictionSpec) satisfiesRestrictionSet(subject *domain.Subject, set *domain.SubjectRestrictionSet) bool {
	var subjResource, subjRelation string
	var validator func(*domain.SubjectRestrictionSet_Restriction, string, string) bool = s.satisfiesRestriction

	switch subjectType := subject.Subject.(type) {
	case *domain.Subject_Entity:
		subjResource = subjectType.Entity.Resource
	case *domain.Subject_EntitySet:
		subjResource = subjectType.EntitySet.Entity.Resource
		subjRelation = subjectType.EntitySet.Relation
		validator = s.satisfiesRestrictionForEntitySet
	case *domain.Subject_ResourceSet:
		subjResource = subjectType.ResourceSet.ResourceName
	}

	for _, restriction := range set.Restrictions {
		ok := validator(restriction, subjResource, subjRelation)
		if ok {
			return true
		}
	}
	return false
}

// satisfiesRestrictionForEntitySet verifies whether the given restriction
// allows a Subject of type EntitySet
func (s *SubjectRestrictionSpec) satisfiesRestrictionForEntitySet(restriction *domain.SubjectRestrictionSet_Restriction, resource, relation string) bool {
	switch restrictionType := restriction.Entry.(type) {
	case *domain.SubjectRestrictionSet_Restriction_Entity:
		// An Entity Restriction cannot satisfies an EntitySet restriction
		// since the relation is not empty
		return false
	case *domain.SubjectRestrictionSet_Restriction_EntitySet:
		setRestriction := restrictionType.EntitySet
		if setRestriction.ResourceName == resource && setRestriction.RelationName == relation {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

// satisfiesRestrictions verifies whether a restriction allows a Subject
// of type Entity or ResourceSet.
func (s *SubjectRestrictionSpec) satisfiesRestriction(restriction *domain.SubjectRestrictionSet_Restriction, resource, _ string) bool {
	switch restrictionType := restriction.Entry.(type) {
	case *domain.SubjectRestrictionSet_Restriction_Entity:
		if restrictionType.Entity.ResourceName == resource {
			return true
		} else {
			return false
		}
	case *domain.SubjectRestrictionSet_Restriction_EntitySet:
		// Entity or ResourceSet Subjects does not satisfy a EntitySet Restriction
		return false
	}
	return false
}

// ValidSelectorSpec verifies whether a RelationshipSelector is valid for a Policy
type ValidSelectorSpec struct{}

func (s *ValidSelectorSpec) Satisfies(selector *domain.RelationshipSelector, policy *domain.Policy) error {
	err := s.validObjectSelector(selector.ObjectSelector, policy)
	if err != nil {
		return errors.Wrap("invalid object selector", err, errors.Pair(errors.AttrPolicy, policy.Id))
	}

	err = s.validRelationSelector(selector.ObjectSelector, selector.RelationSelector, policy)
	if err != nil {
		return errors.Wrap("invalid relation selector", err, errors.Pair(errors.AttrPolicy, policy.Id))
	}

	err = s.validSubjectSelector(selector.SubjectSelector, policy)
	if err != nil {
		return errors.Wrap("invalid subject selector", err, errors.Pair(errors.AttrPolicy, policy.Id))
	}

	return nil
}

func (s *ValidSelectorSpec) validObjectSelector(selector *domain.ObjectSelector, policy *domain.Policy) error {
	switch s := selector.Selector.(type) {
	case *domain.ObjectSelector_ObjectSpec:
		resource := policy.GetResourceByName(s.ObjectSpec.Resource)
		if resource == nil {
			return errors.Wrap("resource not defined in policy", errors.BadInput,
				errors.Pair(errors.AttrResource, s.ObjectSpec.Resource))
		}
	case *domain.ObjectSelector_ResourceSpec:
		resource := policy.GetResourceByName(s.ResourceSpec)
		if resource == nil {
			return errors.Wrap("resource not defined in policy", errors.BadInput,
				errors.Pair(errors.AttrResource, s.ResourceSpec))
		}
	case *domain.ObjectSelector_Wildcard:
		break
	default:
		return errors.Wrap("object selector", errors.ErrInvalidVariant)
	}
	return nil
}

func (s *ValidSelectorSpec) validRelationSelector(objSelector *domain.ObjectSelector, relSelector *domain.RelationSelector, policy *domain.Policy) error {
	if relSelector.GetWildcard() != nil {
		return nil
	}

	relName := relSelector.GetRelationName()

	if objSelector.GetWildcard() != nil {
		return nil
	}

	resourceName := s.getObjectSelectorResourceName(objSelector)
	resource := policy.GetResourceByName(resourceName)
	relation := resource.GetRelationByName(relName)
	if relation == nil {
		return errors.Wrap("relation resource not defined in resource", errors.BadInput,
			errors.Pair(errors.AttrResource, resourceName),
			errors.Pair(errors.AttrRelation, relName),
		)
	}

	return nil
}

func (s *ValidSelectorSpec) getObjectSelectorResourceName(objSelector *domain.ObjectSelector) string {
	switch s := objSelector.Selector.(type) {
	case *domain.ObjectSelector_ObjectSpec:
		return s.ObjectSpec.Resource
	case *domain.ObjectSelector_ResourceSpec:
		return s.ResourceSpec
	}
	return ""
}

func (s *ValidSelectorSpec) validSubjectSelector(selector *domain.SubjectSelector, policy *domain.Policy) error {
	switch s := selector.Selector.(type) {
	case *domain.SubjectSelector_ResourceSpec:
		resource := policy.GetResourceByName(s.ResourceSpec)
		if resource == nil {
			return errors.Wrap("resource not defined in policy", errors.BadInput,
				errors.Pair(errors.AttrResource, s.ResourceSpec))
		}
	case *domain.SubjectSelector_SubjectSpec:
		name := s.SubjectSpec.GetResourceName()
		resource := policy.GetResourceByName(name)
		if resource == nil {
			return errors.Wrap("resource not defined in policy", errors.BadInput,
				errors.Pair(errors.AttrResource, name))
		}
	case *domain.SubjectSelector_Wildcard:
		break
	default:
		return errors.Wrap("subject selector", errors.ErrInvalidVariant)
	}
	return nil
}
