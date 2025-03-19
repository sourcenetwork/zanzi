package domain

// SelectorBuilder builds RelationshipSelector objects
type SelectorBuilder struct {
	objSelector  *ObjectSelector
	relSelector  *RelationSelector
	subjSelector *SubjectSelector
}

// AnyObject sets the Selector to match any system Entity as the
// Relationship's Object
func (s *SelectorBuilder) AnyObject() *SelectorBuilder {
	s.objSelector = &ObjectSelector{
		Selector: &ObjectSelector_Wildcard{
			Wildcard: &WildcardSelector{},
		},
	}
	return s
}

// WithObject sets the Selector to match Relationships with the given entity as the Object
func (s *SelectorBuilder) WithObject(obj *Entity) *SelectorBuilder {
	s.objSelector = &ObjectSelector{
		Selector: &ObjectSelector_ObjectSpec{
			ObjectSpec: obj,
		},
	}
	return s
}

// WithResource sets the Selector to match Relationships with the given Resource for the object
func (s *SelectorBuilder) WithResource(resource string) *SelectorBuilder {
	s.objSelector = &ObjectSelector{
		Selector: &ObjectSelector_ResourceSpec{
			ResourceSpec: resource,
		},
	}
	return s
}

// WithRelation sets the Selector to match the given relation
func (s *SelectorBuilder) WithRelation(relation string) *SelectorBuilder {
	s.relSelector = &RelationSelector{
		Selector: &RelationSelector_RelationName{
			RelationName: relation,
		},
	}
	return s
}

// AnyRelation sets the Selector to include all relations
func (s *SelectorBuilder) AnyRelation() *SelectorBuilder {
	s.relSelector = &RelationSelector{
		Selector: &RelationSelector_Wildcard{
			Wildcard: &WildcardSelector{},
		},
	}
	return s
}

// AnySubject sets the Selector to match any subject
func (s *SelectorBuilder) AnySubject() *SelectorBuilder {
	s.subjSelector = &SubjectSelector{
		Selector: &SubjectSelector_Wildcard{
			Wildcard: &WildcardSelector{},
		},
	}
	return s
}

// WithSubject sets the Selector to match the given subject
func (s *SelectorBuilder) WithSubject(subject *Subject) *SelectorBuilder {
	s.subjSelector = &SubjectSelector{
		Selector: &SubjectSelector_SubjectSpec{
			SubjectSpec: subject,
		},
	}
	return s
}

// WithSubjectResource sets the Selector to subjects from the given resource
func (s *SelectorBuilder) WithSubjectResource(resource string) *SelectorBuilder {
	s.subjSelector = &SubjectSelector{
		Selector: &SubjectSelector_ResourceSpec{
			ResourceSpec: resource,
		},
	}
	return s
}

// WithSubjectGroup sets the Selector to match relationships subjects
// which are subject groups for the given resource, relation pair
func (s *SelectorBuilder) WithSubjectGroup(resourceName, relationName string) *SelectorBuilder {
	s.subjSelector = &SubjectSelector{
		Selector: &SubjectSelector_SubjectGroup{
			SubjectGroup: &SubjectGroupSelector{
				ResourceName: resourceName,
				RelationName: relationName,
			},
		},
	}
	return s
}

// Build a RelationshipSelector with the given parameters
func (s *SelectorBuilder) Build() RelationshipSelector {
	return RelationshipSelector{
		ObjectSelector:   s.objSelector,
		RelationSelector: s.relSelector,
		SubjectSelector:  s.subjSelector,
	}
}

// Reset the builder state to its zero value
func (s *SelectorBuilder) Reset() {
	s.objSelector = nil
	s.relSelector = nil
	s.subjSelector = nil
}
