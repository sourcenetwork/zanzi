package builder

import (
	"github.com/sourcenetwork/source-zanzibar/model"
)

type RelationshipBuilder struct {
	objRel        model.AuthNode
	subj          model.AuthNode
	actorNamespace string
}

func WithActorNamespace(ns string) RelationshipBuilder {
	b := RelationshipBuilder{}
	b.actorNamespace = ns
	return b
}

func (b *RelationshipBuilder) ObjRel(namespace, objectId, relation string) *RelationshipBuilder {
        b.objRel = ObjRelation(namespace, objectId, relation)
	return b
}

func (b *RelationshipBuilder) WithActor(userId string) *RelationshipBuilder {
        b.subj = Actor(b.actorNamespace, userId)
	return b
}

func (b *RelationshipBuilder) WithAttribute(namespace, attribute string) *RelationshipBuilder {
        b.subj = Attribute(namespace, attribute)
	return b
}

func (b *RelationshipBuilder) Build() model.Relationship {
	return model.Relationship{
            ObjRelation: b.objRel,
            Subject: b.subj,
	}
}

func ObjRelation(ns, obj, rel string) model.AuthNode {
	return model.AuthNode{
		Namespace: ns,
		ObjectId:  obj,
		Relation:  rel,
                Type: model.NodeType_OBJECT_RELATION,
	}
}

func Attribute(ns, obj string) model.AuthNode {
	return model.AuthNode{
		Namespace: ns,
		ObjectId:  obj,
                Type: model.NodeType_ATTRIBUTE,
	}
}

func Actor(ns, obj string) model.AuthNode {
	return model.AuthNode{
		Namespace: ns,
		ObjectId:  obj,
                Type: model.NodeType_ACTOR,
	}
}
