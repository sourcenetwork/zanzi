package mappers

import (
	"github.com/sourcenetwork/zanzi/internal/domain/tuple"
	"github.com/sourcenetwork/zanzi/types"
)

type RelationshipMapper struct{}

func (m *RelationshipMapper) FromRelationship(rel types.Relationship) tuple.Tuple {
	src := tuple.TupleNode{
		Namespace: rel.Object.Namespace,
		Id:        rel.Object.Id,
		Relation:  rel.Relation,
		Type:      tuple.NodeType_RELATION_SOURCE,
	}

	dst := tuple.TupleNode{
		Namespace: rel.Actor.Namespace,
		Id:        rel.Actor.Id,
		Relation:  rel.ActorRelation,
	}

	switch rel.Type {
	case types.RelationshipType_ATTRIBUTE:
		dst.Type = tuple.NodeType_OBJECT
	case types.RelationshipType_GRANT:
		dst.Type = tuple.NodeType_ACTOR
	case types.RelationshipType_DELEGATE:
		dst.Type = tuple.NodeType_RELATION_SOURCE
	}

	return tuple.Tuple{
		//CreatedAt: rel.CreatedAt, FIXME
		Partition: rel.PolicyId,
		Source:    src,
		Dest:      dst,
	}
}

func (m *RelationshipMapper) ToRelationship(t tuple.Tuple) types.Relationship {
	var relType types.RelationshipType
	switch t.Dest.Type {
	case tuple.NodeType_ACTOR:
		relType = types.RelationshipType_GRANT
	case tuple.NodeType_OBJECT:
		relType = types.RelationshipType_ATTRIBUTE
	case tuple.NodeType_RELATION_SOURCE:
		relType = types.RelationshipType_DELEGATE
	}

	return types.Relationship{
		PolicyId: t.Partition,
		Type:     relType,
		Object: &types.Entity{
			Namespace: t.Source.Namespace,
			Id:        t.Source.Id,
		},
		Relation: t.Source.Relation,
		Actor: &types.Entity{
			Namespace: t.Dest.Namespace,
			Id:        t.Dest.Id,
		},
		ActorRelation: t.Dest.Relation,
		//CreatedAt:       t.CreatedAt, FIXME
	}
}
