package types

func NewEntity(namespace, id string) Entity {
	return Entity{
		Namespace: namespace,
		Id:        id,
	}
}

type relationshipBuilder struct {
	policyId string
}

func RelationshipBuilder(policyId string) relationshipBuilder {
	return relationshipBuilder{
		policyId: policyId,
	}
}

// Grant builds a Relationship which relates an object to an actor.
// Effectively grants actor access to obj.
func (b relationshipBuilder) Grant(objNamespace, objId, relation, actorNamespace, actorId string) Relationship {
	obj := NewEntity(objNamespace, objId)
	actor := NewEntity(actorNamespace, actorId)
	return b.buildRel(obj, relation, actor, "", RelationshipType_GRANT)
}

// Delegate builds a Relationship which extends a relationship
// between a node and its actor set to another node.
// Ie. all actors related to the node subject, subjectRelation
// will also be related to obj, relation.
func (b relationshipBuilder) Delegate(objNamespace, objId, relation, subjectNamespace, subjectId, subjectRelation string) Relationship {
	obj := NewEntity(objNamespace, objId)
	subject := NewEntity(subjectNamespace, subjectId)
	return b.buildRel(obj, relation, subject, subjectRelation, RelationshipType_DELEGATE)
}

// Atribute builds a Relationship between two objects.
// Attribute is often used to store metadata, such as the parent
// directory of a file.
func (b relationshipBuilder) Attribute(objNamespace, objId, relation, subjectNamespace, subjectId string) Relationship {
	obj := NewEntity(objNamespace, objId)
	subject := NewEntity(subjectNamespace, subjectId)
	return b.buildRel(obj, relation, subject, "", RelationshipType_ATTRIBUTE)
}

func (b relationshipBuilder) buildRel(source Entity, relation string, dest Entity, destRel string, t RelationshipType) Relationship {
	return Relationship{
		PolicyId:        b.policyId,
		Type:            t,
		Object:          source,
		Relation:        relation,
		Subject:         dest,
		SubjectRelation: destRel,
	}
}

func (b relationshipBuilder) Entity(namespace string, id string) Entity {
	return Entity{
		Namespace: namespace,
		Id:        id,
	}
}

type ResourceBuilder struct {
	perms []*Permission
	rels  []*Relation
	name  string
}

func (b *ResourceBuilder) Name(name string) {
	b.name = name
}

func (b *ResourceBuilder) Perm(name string, expression string) {
	perm := &Permission{
		Name:       name,
		PermissionExpr: expression,
	}
	b.perms = append(b.perms, perm)
}

func (b *ResourceBuilder) Relations(relations ...string) {
	for _, name := range relations {
		rel := &Relation{
			Name: name,
		}
		b.rels = append(b.rels, rel)
	}
}

func (b *ResourceBuilder) Build() Resource {
	resource := Resource{
		Name:        b.name,
		Relations:   b.rels,
		Permissions: b.perms,
	}
	b.name = ""
	b.rels = nil
	b.perms = nil
	return resource
}

func NewActor(actor string) Actor {
	return Actor{
		Name:       actor,
	}
}

type policyBuilder struct {
	name        string
	id          string
	resources   []*Resource
	actors      []*Actor
	attrs       map[string]string
	description string
}

func PolicyBuilder() policyBuilder {
	return policyBuilder{
		attrs: make(map[string]string),
	}
}

func (b *policyBuilder) IdNameDescription(id string, name string, description string) {
	b.name = name
	b.id = id
	b.description = description
}

func (b *policyBuilder) Actors(actors ...Actor) {
	for _, actor := range actors {
		b.actors = append(b.actors, &actor)
	}
}

func (b *policyBuilder) Resources(resources ...Resource) {
	for _, resource := range resources {
		b.resources = append(b.resources, &resource)
	}
}

func (b *policyBuilder) Resource(resource Resource) {
	b.Resources(resource)
}

func (b *policyBuilder) Attr(key, value string) {
	b.attrs[key] = value
}

func (b *policyBuilder) Build() Policy {
	policy := Policy{
		Name:        b.name,
		Id:          b.id,
		Resources:   b.resources,
		Actors:      b.actors,
		Attributes:  b.attrs,
		Description: b.description,
	}
	b.name = ""
	b.id = ""
	b.resources = nil
	b.description = ""
	b.actors = nil
	b.attrs = make(map[string]string)
	return policy
}
