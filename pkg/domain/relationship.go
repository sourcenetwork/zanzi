package domain

import (
	"fmt"
	"regexp"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	relationshipRegex = `^(?P<ResourceGroup>\w+):(?P<ResourceID>\w+)#(?P<Relation>\w+)@(?P<SubjectGroup>\w+):(?P<SubjectID>.+)$`
)

func NewEntity(resource, id string) *Entity {
	return &Entity{
		Resource: resource,
		Id:       id,
	}
}

func NewRelationshipRecord(policyId string, relationship *Relationship, data []byte) *RelationshipRecord {
	return &RelationshipRecord{
		PolicyId:     policyId,
		Relationship: relationship,
		AppData:      data,
		CreatedAt:    timestamppb.Now(),
	}
}

type RelationshipBuilder struct{}

func (b *RelationshipBuilder) Relationship(objResource, objId, relation, subjResource, subjId string) Relationship {
	return Relationship{
		Object: &Entity{
			Resource: objResource,
			Id:       objId,
		},
		Relation: relation,
		Subject: &Subject{
			Subject: &Subject_Entity{
				Entity: &Entity{
					Resource: subjResource,
					Id:       subjId,
				},
			},
		},
	}
}

func (b *RelationshipBuilder) RelationshipFromString(relationship string) (Relationship, error) {
	r, err := regexp.Compile(relationshipRegex)
	if err != nil {
		return Relationship{}, err
	}

	if !r.Match([]byte(relationship)) {
		return Relationship{}, fmt.Errorf("relationship validation: %s", relationship)
	}

	results := r.FindStringSubmatch(relationship)
	if len(results) != 6 {
		return Relationship{}, fmt.Errorf("regex submatch size: %d", len(results))
	}

	return b.Relationship(
		results[1], /* resource group */
		results[2], /* resource id */
		results[3], /* relation */
		results[4], /* subject group */
		results[5], /* subject id */
	), nil
}

func (b *RelationshipBuilder) EntitySet(objResource, objId, relation, subjResource, subjId, subjRelation string) Relationship {
	return Relationship{
		Object: &Entity{
			Resource: objResource,
			Id:       objId,
		},
		Relation: relation,
		Subject: &Subject{
			Subject: &Subject_EntitySet{
				EntitySet: &EntitySet{
					Entity: &Entity{
						Resource: subjResource,
						Id:       subjId,
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
			Id:       objId,
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
