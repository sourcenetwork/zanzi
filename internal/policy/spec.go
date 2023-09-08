package policy

import (
    "fmt"

    "github.com/sourcenetwork/zanzi/pkg/domain"
)

type ValidPolicySpec struct {}


func (s *ValidPolicySpec) Verify(policy *domain.Policy) error {
    err := s.verify(policy)
    if err != nil {
        return fmt.Errorf("policy %v: %w", policy.Id, err)
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

    for _, resource := range policy.Resources{
        name := resource.Name
        _, ok := names[name]
        if ok {
            return fmt.Errorf("resource %v: %w", name, ErrDuplicateDefinition)
        }
        names[name] = struct{}{}

        err := s.UniqueRelationNameForResource(resource)
        if err != nil {
            return fmt.Errorf("resource %v: %w", name, err)
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
                return fmt.Errorf("relation %v: %w", name, ErrDuplicateDefinition)
            }
            names[name] = struct{}{}
        }
        return nil
}


func (s *ValidPolicySpec) RelationExpressionWellDefined(policy *domain.Policy, lut PolicyLookUpTable) error {
    for _, resource := range policy.Resources{
        for _, relation := range resource.Relations {
            tree, err := GetExpressionTree(relation)
            if err != nil {
                return fmt.Errorf("resoruce %v: relation %v: %w", resource.Name, relation.Name, err)
            }
            rules := tree.GetRules()
            for _, rule := range rules {
                err := s.policyContainsRelationsInRule(resource.Name, relation.Name, lut, rule)
                if err != nil {
                    return fmt.Errorf("resoruce %v: relation %v: rule %v: %w", resource.Name, relation.Name, rule, err)
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
            return fmt.Errorf("resource %v missing relation %v: %w", resource, target, ErrRelExpTree)
        }
    case *domain.Rule_This:
    case *domain.Rule_Ttu:
        target := r.Ttu.TuplesetRelation
        tuplesetRel := lut.GetRelation(resource, target)
        if tuplesetRel == nil {
            return fmt.Errorf("resource %v missing relation %v: %w", resource, target, ErrRelExpTree)
        }

        // TODO this check is trickier, as I have to consider
        // every allowed resource type in the subject restriction graph.
        // tbh it's pretty simple.
        //cuRel := lut.GetRelation(resource, r.Ttu.ComputedUsersetRelation)
    default:
        return fmt.Errorf("rule %v: %w", r, domain.ErrInvalidVariant)
    }
    return nil
}

func (s *ValidPolicySpec) SubjectRestrictionsConsistent(policy *domain.Policy, lut PolicyLookUpTable) error {
    for _, resource := range policy.Resources{
        for _, relation := range resource.Relations {
            var err error
            if relation.SubjectRestriction != nil {
                err = s.ValidSubjectRestriction(relation.SubjectRestriction, lut)
            }
            if err != nil {
                return fmt.Errorf("%w: resource %v, relation %v: %v", ErrSubjectRestriction, resource.Name, relation.Name, err)
            }
        }
    }
    return nil
}

func (s *ValidPolicySpec) ValidSubjectRestriction(restriction *domain.SubjectRestriction, lut PolicyLookUpTable) error {
    switch restrictionSet := restriction.SubjectRestriction.(type){
    case *domain.SubjectRestriction_RestrictionSet:
        return s.validSubjectRestrictionSet(restrictionSet.RestrictionSet, lut)
    case *domain.SubjectRestriction_UniversalSet:
        return nil
    default:
        return fmt.Errorf("invalid subject restriction type %v", restriction)
    }
}

func (s *ValidPolicySpec) validSubjectRestrictionSet(restrictionSet *domain.SubjectRestrictionSet, lut PolicyLookUpTable) error {

    for _, elem := range restrictionSet.Restrictions {
        switch restriction := elem.Entry.(type) {
        case *domain.SubjectRestrictionSet_Restriction_Entity:
            if lut.GetResource(restriction.Entity.ResourceName) == nil {
                return fmt.Errorf("no resource %v", restriction.Entity.ResourceName)
            }
        case *domain.SubjectRestrictionSet_Restriction_EntitySet:
            resource, relation := restriction.EntitySet.ResourceName, restriction.EntitySet.RelationName
            if lut.GetRelation(resource, relation) == nil {
                return fmt.Errorf("no relation %v for resource %v", relation, resource)
            }
        default:
            return fmt.Errorf("invalid type %v", restrictionSet)
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
        return fmt.Errorf("%v: %w", ErrInvalidRelationship)
    }
    return nil
}

// checkSymbols verifies that the symbols referred to in a Relationship are valid within a Policy
// eg. validates the resources and relation names
func (s *AllowedRelationshipSpec) checkSymbols(relationship *domain.Relationship, lut PolicyLookUpTable) error {
    resName, relName := relationship.Object.Resource, relationship.Relation

    res := lut.GetResource(resName)
    if res == nil {
        return fmt.Errorf("resource %v: %w", resName, ErrResourceNotFound)
    }

    rel := lut.GetRelation(resName, relName)
    if rel == nil {
        return fmt.Errorf("resource %v: relation %v: %w", resName, relName, ErrRelationNotFound)
    }

    switch s := relationship.Subject.Subject.(type) {
    case *domain.Subject_Entity:
        res = lut.GetResource(s.Entity.Resource)
        if res == nil {
            return fmt.Errorf("resource %v: %w", s.Entity.Resource, ErrResourceNotFound)
        }
    case *domain.Subject_EntitySet:
        res, rel := s.EntitySet.Entity.Resource, s.EntitySet.Relation
        resource := lut.GetRelation(res, rel)
        if resource == nil {
            return fmt.Errorf("resource %v: relation %v: %w", res, rel, ErrRelationNotFound)
        }
    case *domain.Subject_ResourceSet:
        resource := lut.GetResource(s.ResourceSet.ResourceName)
        if resource == nil {
            return fmt.Errorf("resource %v: ", s.ResourceSet.ResourceName, ErrResourceNotFound)
        }
    default:
        return fmt.Errorf("subject %v: %w", s, domain.ErrInvalidVariant)
    }

    return nil
}

func (s *AllowedRelationshipSpec) Satisfies(relationship *domain.Relationship, lut PolicyLookUpTable) error {
    err := s.satisfies(relationship, lut)
    if err != nil{
        return fmt.Errorf("relationship %v not allowed: %w", relationship, err)
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
type SubjectRestrictionSpec struct { }

// Satisfies checks whether relationship.Subject is allowed by the Policy's SubjectRestriction rules
func (s *SubjectRestrictionSpec) Satisfies(relationship *domain.Relationship, table PolicyLookUpTable) error {
    relation := table.GetRelation(relationship.Object.Resource, relationship.Relation)
    if relation == nil {
        return fmt.Errorf("relation %v: %w", relationship.Relation, ErrRelationNotFound)
    }

    var err error

    switch restriction := relation.SubjectRestriction.SubjectRestriction.(type) {
    case *domain.SubjectRestriction_UniversalSet:
        err = nil
    case *domain.SubjectRestriction_RestrictionSet:
        err = s.satisfiesRestrictionSet(relationship.Subject, restriction.RestrictionSet)
    default:
        err = fmt.Errorf("SubjectRestriction %v: %w", restriction, domain.ErrInvalidVariant)
    }

    if err != nil{
        return fmt.Errorf("%w: relation %v subject %v: %w", ErrSubjectNotAllowed, relation, relationship.Subject, err)
    }

    return nil
}

func (s *SubjectRestrictionSpec) satisfiesRestrictionSet(subject *domain.Subject, set *domain.SubjectRestrictionSet) error {
    var subjResource, subjRelation string
    var validator func(*domain.SubjectRestrictionSet_Restriction, string, string) error = s.satisfiesRestriction

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
        err := validator(restriction, subjResource, subjRelation)
        if err == nil{
            return nil
        }
    }
    return fmt.Errorf("RestrictionSet does not allow subject")
}

// satisfiesRestrictionForEntitySet verifies whether the given restriction
// allows a Subject of type EntitySet
func (s *SubjectRestrictionSpec) satisfiesRestrictionForEntitySet(restriction *domain.SubjectRestrictionSet_Restriction, resource, relation string) error {
    switch restrictionType := restriction.Entry.(type) {
    case *domain.SubjectRestrictionSet_Restriction_Entity:
        // An Entity Restriction cannot satisfies an EntitySet restriction
        // since the relation is not empty
        return ErrSubjectRestriction
    case *domain.SubjectRestrictionSet_Restriction_EntitySet:
        setRestriction := restrictionType.EntitySet
        if setRestriction.ResourceName == resource && setRestriction.RelationName == relation{
            return nil
        } else {
            return ErrSubjectRestriction
        }
    default:
        return ErrSubjectRestriction
    }
}

// satisfiesRestrictions verifies whether a restriction allows a Subject
// of type Entity or ResourceSet.
func (s *SubjectRestrictionSpec) satisfiesRestriction(restriction *domain.SubjectRestrictionSet_Restriction, resource, _ string) error {
    switch restrictionType := restriction.Entry.(type) {
    case *domain.SubjectRestrictionSet_Restriction_Entity:
        if restrictionType.Entity.ResourceName == resource{
            return nil
        } else {
            return ErrSubjectRestriction
        }
    case *domain.SubjectRestrictionSet_Restriction_EntitySet:
        // Entity or ResourceSet Subjects does not satisfy a EntitySet Restriction
        return ErrSubjectRestriction
    }
    return ErrSubjectRestriction
}
