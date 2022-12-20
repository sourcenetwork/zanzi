package services

import (
    _ "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/types"
)

type RelationshipMapper struct { }

func (m *RelationshipMapper) FromRelationship(rel types.Relationship) tuple.Tuple {
    src := tuple.TupleNode {
        Namespace: rel.Object.Namespace,
        Id: rel.Object.Id,
        Relation: rel.Relation,
        Type: tuple.NodeType_RELATION_SOURCE,
    }

    dst := tuple.TupleNode {
        Namespace: rel.Subject.Namespace,
        Id: rel.Subject.Id,
        Relation: rel.SubjectRelation,
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
        Partition: rel.PolicyId,
        Source: src,
        Dest: dst,
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
        Type: relType,
        Object: types.Entity{
            Namespace: t.Source.Namespace,
            Id: t.Source.Id,
        },
        Relation: t.Source.Relation,
        Subject: types.Entity{
            Namespace: t.Dest.Namespace,
            Id: t.Dest.Id,
        },
        SubjectRelation: t.Dest.Relation,
    }
}
