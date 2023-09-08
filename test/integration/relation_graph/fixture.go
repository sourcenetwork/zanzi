package relation_graph

import (
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

var setupPolicy *domain.Policy = &domain.Policy{
	Id:          "10",
	Name:        "test",
	Description: "a test policy",
	Resources: []*domain.Resource{
		&domain.Resource{
			Name:        "file",
			Description: "file resource",
			Relations: []*domain.Relation{
				&domain.Relation{
					Name:        "owner",
					Description: "file owner",
					RelationExpression: &domain.RelationExpression{
						Expression: &domain.RelationExpression_Expr{
							Expr: "_this",
						},
					},
					SubjectRestriction: &domain.SubjectRestriction{
						SubjectRestriction: &domain.SubjectRestriction_UniversalSet{
							UniversalSet: &domain.UniversalSet{},
						},
					},
				},
				&domain.Relation{
					Name:        "read",
					Description: "allowed to read file",
					RelationExpression: &domain.RelationExpression{
						Expression: &domain.RelationExpression_Expr{
							Expr: "owner + directory->owner",
						},
					},
					SubjectRestriction: &domain.SubjectRestriction{
						SubjectRestriction: &domain.SubjectRestriction_RestrictionSet{
							RestrictionSet: &domain.SubjectRestrictionSet{},
						},
					},
				},
				&domain.Relation{
					Name:        "directory",
					Description: "references the directory which contains the file",
					RelationExpression: &domain.RelationExpression{
						Expression: &domain.RelationExpression_Expr{
							Expr: "_this",
						},
					},
					SubjectRestriction: &domain.SubjectRestriction{
						SubjectRestriction: &domain.SubjectRestriction_UniversalSet{
							UniversalSet: &domain.UniversalSet{},
						},
					},
				},
			},
		},
		&domain.Resource{
			Name:        "directory",
			Description: "directories",
			Relations: []*domain.Relation{
				&domain.Relation{
					Name:        "owner",
					Description: "directory owner",
					RelationExpression: &domain.RelationExpression{
						Expression: &domain.RelationExpression_Expr{
							Expr: "_this",
						},
					},
					SubjectRestriction: &domain.SubjectRestriction{
						SubjectRestriction: &domain.SubjectRestriction_UniversalSet{
							UniversalSet: &domain.UniversalSet{},
						},
					},
				},
			},
		},
		&domain.Resource{
			Name: "user",
		},
	},
	Attributes: map[string]string{
		"foo": "bar",
	},
}

var builder domain.RelationshipBuilder

var setupRelationships []domain.Relationship = []domain.Relationship{
	builder.Relationship("file", "readme", "owner", "user", "bob"),
}
