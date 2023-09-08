package kv_store

import (
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

type relationshipMapper struct { }

func (m *relationshipMapper) ToInternal(relationship *domain.Relationship) Relationship {
    var subjType RelationType
    var subjResource, subjId, subjRelation string

    switch subject := relationship.Subject.Subject.(type) {
    case *domain.Subject_Entity:
        subjResource = subject.Entity.Resource
        subjId = subject.Entity.Id
        subjType = RelationType_OBJECT
    case *domain.Subject_EntitySet:
        subjResource = subject.EntitySet.Entity.Resource
        subjId = subject.EntitySet.Entity.Id
        subjRelation = subject.EntitySet.Relation
        subjType = RelationType_OBJECT_SET
    case *domain.Subject_ResourceSet:
        subjResource = subject.ResourceSet.ResourceName
        subjType = RelationType_RESOURCE_SET
    default:
        subjType = RelationType_UNKNOWN
    }

    return Relationship{
        Source: &RelationNode {
            Resource: relationship.Object.Resource,
            Id: relationship.Object.Id,
            Relation: relationship.Relation,
            Type: RelationType_OBJECT_SET,
        },
        Dest: &RelationNode{
            Resource: subjResource,
            Id: subjId,
            Relation: subjRelation,
            Type: subjType,
        },
    }
}

func (m *relationshipMapper) FromInternal(relationship *Relationship) domain.Relationship {
    subject := &domain.Subject{}

    dest := relationship.Dest
    switch relationship.Dest.Type {
    case RelationType_OBJECT:
        subject.Subject = &domain.Subject_Entity{
            Entity: &domain.Entity{
                Resource: dest.Resource,
                Id: dest.Id,
            },
        }
    case RelationType_OBJECT_SET:
        subject.Subject = &domain.Subject_EntitySet{
            EntitySet: &domain.EntitySet{
                Entity: &domain.Entity{
                    Resource: dest.Resource,
                    Id: dest.Id,
                },
                Relation: dest.Relation,
            },
        }
    case RelationType_RESOURCE_SET:
        subject.Subject = &domain.Subject_ResourceSet{
            ResourceSet: &domain.ResourceSet{
                ResourceName: dest.Resource,
            },
        }
    }
    return domain.Relationship{
        Object: &domain.Entity {
            Resource: relationship.Source.Resource,
            Id: relationship.Source.Id,
        },
        Relation: relationship.Source.Relation,
        Subject: subject,
    }
}


func (m *relationshipMapper) FromPublicRelationNode(node *domain.RelationNode) RelationNode {
    var resource, id, relation string
    var relType RelationType

    switch nodeType := node.Node.(type) {
    case *domain.RelationNode_Entity:
        resource = nodeType.Entity.Object.Resource
        id = nodeType.Entity.Object.Id
        relType = RelationType_OBJECT
    case *domain.RelationNode_EntitySet:
        resource = nodeType.EntitySet.Object.Resource
        id = nodeType.EntitySet.Object.Id
        relation = nodeType.EntitySet.Relation
        relType = RelationType_OBJECT_SET
    case *domain.RelationNode_Wildcard:
        resource = nodeType.Wildcard.Resource
        relType = RelationType_RESOURCE_SET
    default:
        relType = RelationType_UNKNOWN
    }

    return RelationNode{
        Resource: resource,
        Id: id,
        Relation: relation,
        Type: relType,
    }
}

func (m *relationshipMapper) ToPublicRelationNode(node *RelationNode) domain.RelationNode {
    var relationNode domain.RelationNode

    switch node.Type{
    case RelationType_OBJECT:
        relationNode.Node = &domain.RelationNode_Entity{
            Entity: &domain.EntityNode{
                Object: &domain.Entity{
                    Resource: node.Resource,
                    Id: node.Id,
                },
            },
        }
    case RelationType_OBJECT_SET:
        relationNode.Node = &domain.RelationNode_EntitySet{
            EntitySet: &domain.EntitySetNode{
                Object: &domain.Entity{
                    Resource: node.Resource,
                    Id: node.Id,
                },
                Relation: node.Relation,
            },
        }
    case RelationType_RESOURCE_SET:
        relationNode.Node = &domain.RelationNode_Wildcard{
            Wildcard: &domain.WildcardNode{
                Resource: node.Resource,
            },
        }
    }

    return relationNode
}

func (m *relationshipMapper) FromInternalRecord(policyId string, relationship *Relationship, data *RelationshipData) domain.RelationshipRecord {
    rel := m.FromInternal(relationship)
    return domain.RelationshipRecord {
        PolicyId: policyId,
        Relationship: &rel,
        CreatedAt: data.CreatedAt,
        AppData: data.AppData,
    }
}
