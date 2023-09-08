package relation_graph

import (
    "testing"
    "context"
    _ "errors"

    rcdb "github.com/sourcenetwork/raccoondb"
    "github.com/stretchr/testify/require"
        "go.uber.org/zap"

    "github.com/sourcenetwork/zanzi/internal/store/kv_store"
    "github.com/sourcenetwork/zanzi/pkg/domain"
    "github.com/sourcenetwork/zanzi/pkg/api"
    "github.com/sourcenetwork/zanzi/internal/relation_graph"
    "github.com/sourcenetwork/zanzi/internal/policy"
    _testing "github.com/sourcenetwork/zanzi/internal/testing"
)

func setup() (context.Context, api.RelationGraphServer) {
    ctx := context.Background()
	kv := rcdb.NewMemKV()
        kvStore, err := kv_store.NewKVStore(kv)
        if err != nil {
            panic(err)
        }

        zapLogger, err := zap.NewDevelopment()
        if err != nil { panic(err) }
        logger := zapLogger.Sugar()

        polRepo := kvStore.GetPolicyRepository()
        nodeRepo := kvStore.GetRelationNodeRepository()

        polService := policy.NewService(polRepo) 
        service := relation_graph.NewService(nodeRepo, polRepo, logger)

        // setup policy
        createReq := &api.CreatePolicyRequest{
            PolicyDefinition: &api.PolicyDefinition{
                Definition: &api.PolicyDefinition_Policy_{
                    Policy_: setupPolicy,
                },
            },
            AppData: []byte("app data"),
        }
        _, err = polService.CreatePolicy(ctx, createReq)
        if err != nil {
            panic(err)
        }

        // setup relationships
        for _, relationship := range setupRelationships {
            req := api.SetRelationshipRequest{
                PolicyId: setupPolicy.Id,
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
    ctx, service := setup()

    req := &api.CheckRequest{
        PolicyId: setupPolicy.Id,
        AccessRequest: &domain.AccessRequest{
            Object: domain.NewEntity("file", "readme"),
            Relation: "read",
            Subject: domain.NewEntity("user", "bob"),
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
