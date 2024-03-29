package store

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sourcenetwork/zanzi/internal/policy"
	_testing "github.com/sourcenetwork/zanzi/internal/testing"
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

type PolicyRepositoryTestSuite struct {
	factory func() policy.Repository
}

func NewPolicyRepositoryTestSuite(repoFactory func() policy.Repository) PolicyRepositoryTestSuite {
	return PolicyRepositoryTestSuite{
		factory: repoFactory,
	}
}

var policyDef *domain.Policy = &domain.Policy{
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

var policyRecord *domain.PolicyRecord = &domain.PolicyRecord{
	CreatedAt: timestamppb.New(time.Date(2023, time.November, 10, 0, 0, 0, 0, time.UTC)),
	AppData:   []byte("app_data"),
	Policy:    policyDef,
}

func (s *PolicyRepositoryTestSuite) TestPolicySetGet(t *testing.T) {
	ctx, repository := s.setup()

	found, setErr := repository.SetPolicy(ctx, policyRecord)

	got, getErr := repository.GetPolicy(ctx, policyDef.Id)

	require.Nil(t, setErr)
	require.Nil(t, getErr)
	require.False(t, bool(found))
	_testing.ProtoEq(t, policyRecord, got)
}

func (s *PolicyRepositoryTestSuite) TestDeletingASetPolicy(t *testing.T) {
	ctx, repository := s.setup()

	_, setErr := repository.SetPolicy(ctx, policyRecord)

	found, getErr := repository.DeletePolicy(ctx, policyDef.Id)

	require.Nil(t, setErr)
	require.Nil(t, getErr)
	require.True(t, bool(found))
}

func (s *PolicyRepositoryTestSuite) TestDeletingAnUnknownPolicyDoesNotErrorOut(t *testing.T) {
	ctx, repository := s.setup()

	found, getErr := repository.DeletePolicy(ctx, "")

	require.Nil(t, getErr)
	require.False(t, bool(found))
}

func (s *PolicyRepositoryTestSuite) setup() (context.Context, policy.Repository) {
	return context.Background(), s.factory()
}

func (s *PolicyRepositoryTestSuite) TestListingPolicyIdsReturnAllIds(t *testing.T) {
	ctx, repository := s.setup()

        rec1 := &domain.PolicyRecord{
            Policy:    &domain.Policy{
                Id: "1",
            },
        }
        rec2 := &domain.PolicyRecord{
            Policy:    &domain.Policy{
                Id: "2",
            },
        }
        _, err := repository.SetPolicy(ctx, rec1)
        require.Nil(t, err)
        _, err = repository.SetPolicy(ctx, rec2)
        require.Nil(t, err)

        ids, err := repository.ListPolicyIds(ctx)

        require.Nil(t, err)
        require.Equal(t, []string{"1", "2"}, ids)
}
