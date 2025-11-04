package relation_graph

import (
	"context"
	_ "errors"
	"testing"

	rcdb "github.com/sourcenetwork/raccoondb"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/relation_graph"
	"github.com/sourcenetwork/zanzi/internal/store/kv_store"
	_testing "github.com/sourcenetwork/zanzi/internal/testing"
	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

func setup(pol domain.Policy, rels []domain.Relationship) (context.Context, api.RelationGraphServer) {
	ctx := context.Background()
	kv := rcdb.NewMemKV()
	kvStore, err := kv_store.NewKVStore(kv)
	if err != nil {
		panic(err)
	}

	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger := zapLogger.Sugar()

	polRepo := kvStore.GetPolicyRepository()
	nodeRepo := kvStore.GetRelationNodeRepository()

	polService := policy.NewService(polRepo)
	service := relation_graph.NewService(nodeRepo, polRepo, logger)

	// setup policy
	createReq := &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_Policy{
				Policy: &pol,
			},
		},
		AppData: []byte("app data"),
	}
	_, err = polService.CreatePolicy(ctx, createReq)
	if err != nil {
		panic(err)
	}

	// setup relationships
	for _, relationship := range rels {
		req := api.SetRelationshipRequest{
			PolicyId:     pol.Id,
			Relationship: &relationship,
		}
		_, err = polService.SetRelationship(ctx, &req)
		if err != nil {
			panic(err)
		}
	}

	return ctx, service
}

func TestCheckForSimpleObject(t *testing.T) {
	ctx, service := setup(*setupPolicy, setupRelationships)

	req := &api.CheckRequest{
		PolicyId: setupPolicy.Id,
		AccessRequest: &domain.AccessRequest{
			Object:   domain.NewEntity("file", "readme"),
			Relation: "read",
			Subject:  domain.NewEntity("user", "bob"),
		},
	}
	res, err := service.Check(ctx, req)

	require.Nil(t, err)
	_testing.ProtoEq(t, res, &api.CheckResponse{
		Result: &api.CheckResponse_Result{
			Authorized: true,
		},
	})
}

func TestCheckExpr(t *testing.T) {
	pol := domain.Policy{
		Id:   "1",
		Name: "pol",
		Resources: []*domain.Resource{
			&domain.Resource{
				Name: "file",
				Relations: []*domain.Relation{
					{
						Name: "owner",
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
					{
						Name: "reader",
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
	}
	rels := []domain.Relationship{
		builder.Relationship("file", "readme", "owner", "user", "owner"),
		builder.Relationship("file", "readme", "reader", "user", "reader"),
	}

	ctx, srv := setup(pol, rels)
	result, err := srv.CheckExpression(ctx, &api.CheckExpressionRequest{
		PolicyId:           pol.Id,
		Object:             domain.NewEntity("file", "readme"),
		Subject:            domain.NewEntity("user", "reader"),
		RelationExpression: "owner + reader",
	})
	require.NoError(t, err)
	require.True(t, result.Authorized)

	result, err = srv.CheckExpression(ctx, &api.CheckExpressionRequest{
		PolicyId:           pol.Id,
		Object:             domain.NewEntity("file", "readme"),
		Subject:            domain.NewEntity("user", "owner"),
		RelationExpression: "owner + reader",
	})
	require.NoError(t, err)
	require.True(t, result.Authorized)
}
