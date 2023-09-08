package domain

import (
        "google.golang.org/protobuf/types/known/timestamppb"
)

func NewEntity(resource, id string) *Entity {
    return &Entity{
        Resource: resource,
        Id: id,
    }
}

func NewRelationshipRecord(policyId string, relationship *Relationship, data []byte) *RelationshipRecord {
    return &RelationshipRecord{
        PolicyId: policyId,
        Relationship: relationship,
        AppData: data,
        CreatedAt: timestamppb.Now(),
    }
}

type RelationshipBuilder struct {}

func (b *RelationshipBuilder) Relationship(objResource, objId, relation, subjResource, subjId string) Relationship {
    return Relationship{
        Object: &Entity{
            Resource: objResource,
            Id: objId,
        },
        Relation: relation,
        Subject: &Subject{
            Subject: &Subject_Entity{
                Entity: &Entity{
                    Resource: subjResource,
                    Id: subjId,
                },
            },
        },
    }
}

func (b *RelationshipBuilder) EntitySet(objResource, objId, relation, subjResource, subjId, subjRelation string) Relationship {
    return Relationship{
        Object: &Entity{
            Resource: objResource,
            Id: objId,
        },
        Relation: relation,
        Subject: &Subject{
            Subject: &Subject_EntitySet{
                EntitySet: &EntitySet{
                    Entity: &Entity{
                        Resource: subjResource,
                        Id: subjId,
                    },
                    Relation: subjRelation,
                },
            },
        },
    }
}

func (b *RelationshipBuilder) Wildcard(objResource, objId, relation, subjResource string) Relationship {
    return Relationship{
        Object: &Entity{
            Resource: objResource,
            Id: objId,
        },
        Relation: relation,
        Subject: &Subject{
            Subject: &Subject_ResourceSet{
                ResourceSet: &ResourceSet{
                    ResourceName: subjResource,
                },
            },
        },
    }
}
