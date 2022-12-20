package types

func NewEntity(namespace, id string) Entity {
	return Entity{
		Namespace: namespace,
		Id:        id,
	}
}

type RelationshipBuilder struct {
	policyId string
}

func NewBuilder(policyId string) RelationshipBuilder {
	return RelationshipBuilder{
		policyId: policyId,
	}
}

func (b *RelationshipBuilder) Grant(obj Entity, relation string, actor Entity) Relationship {
	return b.buildRel(obj, relation, actor, "", RelationshipType_GRANT)
}

func (b *RelationshipBuilder) Delegate(obj Entity, relation string, subject Entity, subjectRelation string) Relationship {
	return b.buildRel(obj, relation, subject, subjectRelation, RelationshipType_DELEGATE)
}

func (b *RelationshipBuilder) Attribute(obj Entity, relation string, subject Entity) Relationship {
	return b.buildRel(obj, relation, subject, "", RelationshipType_ATTRIBUTE)
}

func (b *RelationshipBuilder) buildRel(source Entity, relation string, dest Entity, destRel string, t RelationshipType) Relationship {
	return Relationship{
		PolicyId:        b.policyId,
		Type:            t,
		Object:          source,
		Relation:        relation,
		Subject:         dest,
		SubjectRelation: destRel,
	}
}

type PolicyBuilder struct {
}
