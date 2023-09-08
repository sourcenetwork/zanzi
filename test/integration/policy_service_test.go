package integration

import (
    "context"
    "testing"
    "errors"

    rcdb "github.com/sourcenetwork/raccoondb"
    "github.com/stretchr/testify/require"
    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/zanzi/pkg/domain"
    "github.com/sourcenetwork/zanzi/pkg/api"
    "github.com/sourcenetwork/zanzi/internal/store/kv_store"
    "github.com/sourcenetwork/zanzi/internal/policy"
    _testing "github.com/sourcenetwork/zanzi/internal/testing"
)


func setup() (context.Context, api.PolicyServiceServer) {
	kv := rcdb.NewMemKV()
        kvStore, err := kv_store.NewKVStore(kv)
        if err != nil {
            panic(err)
        }

        return context.Background(), policy.NewService(kvStore.GetPolicyRepository())
}

func setupWithPolicy(policy *domain.Policy)  (context.Context, api.PolicyServiceServer) {
    ctx, service := setup()

    createReq := &api.CreatePolicyRequest{
        PolicyDefinition: &api.PolicyDefinition{
            Definition: &api.PolicyDefinition_Policy_{
                Policy_: policy,
            },
        },
        AppData: []byte("app data"),
    }
    _, err := service.CreatePolicy(ctx, createReq)
    if err != nil {
        panic(err)
    }
    return ctx, service
}

var relationshipBuilder domain.RelationshipBuilder = domain.RelationshipBuilder{}

var testPolicy *domain.Policy = &domain.Policy{
    Id: "10",
    Name: "test",
    Description: "a test policy",
    Resources: []*domain.Resource{
        &domain.Resource{
            Name: "file",
            Description: "file resource",
            Relations: []*domain.Relation{
                &domain.Relation{
                    Name: "owner",
                    Description: "file owner",
                    RelationExpression: &domain.RelationExpression{
                        Expression: &domain.RelationExpression_Expr{
                            Expr: "_this",
                        },
                    },
                    SubjectRestriction: &domain.SubjectRestriction{
                        SubjectRestriction: &domain.SubjectRestriction_UniversalSet {
                            UniversalSet: &domain.UniversalSet{},
                        },
                    },
                },
                &domain.Relation{
                    Name: "read",
                    Description: "allowed to read file",
                    RelationExpression: &domain.RelationExpression{
                        Expression: &domain.RelationExpression_Expr{
                            Expr: "owner + directory->owner",
                        },
                    },
                    SubjectRestriction: &domain.SubjectRestriction{
                        SubjectRestriction: &domain.SubjectRestriction_RestrictionSet {
                            RestrictionSet: &domain.SubjectRestrictionSet{},
                        },
                    },
                },
                &domain.Relation{
                    Name: "directory",
                    Description: "references the directory which contains the file",
                    RelationExpression: &domain.RelationExpression{
                        Expression: &domain.RelationExpression_Expr{
                            Expr: "_this",
                        },
                    },
                    SubjectRestriction: &domain.SubjectRestriction{
                        SubjectRestriction: &domain.SubjectRestriction_UniversalSet {
                            UniversalSet: &domain.UniversalSet{},
                        },
                    },
                },
            },
        },
        &domain.Resource{
            Name: "directory",
            Description: "directories",
            Relations: []*domain.Relation{
                &domain.Relation{
                    Name: "owner",
                    Description: "directory owner",
                    RelationExpression: &domain.RelationExpression{
                        Expression: &domain.RelationExpression_Expr{
                            Expr: "_this",
                        },
                    },
                    SubjectRestriction: &domain.SubjectRestriction{
                        SubjectRestriction: &domain.SubjectRestriction_UniversalSet {
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
    Attributes: map[string]string {
        "foo": "bar",
    },
}


var restrictedPolicy *domain.Policy = &domain.Policy{
    Id: "11",
    Name: "test",
    Description: "policy with restrictions",
    Resources: []*domain.Resource{
        &domain.Resource{
            Name: "a",
            Relations: []*domain.Relation{
                &domain.Relation{
                    Name: "empty_restriction_set",
                    RelationExpression: &domain.RelationExpression{
                        Expression: &domain.RelationExpression_Expr{
                            Expr: "_this",
                        },
                    },
                    SubjectRestriction: &domain.SubjectRestriction{
                        SubjectRestriction: &domain.SubjectRestriction_RestrictionSet {
                            RestrictionSet: &domain.SubjectRestrictionSet{},
                        },
                    },
                },
                &domain.Relation{
                    Name: "universal",
                    RelationExpression: &domain.RelationExpression{
                        Expression: &domain.RelationExpression_Expr{
                            Expr: "_this",
                        },
                    },
                    SubjectRestriction: &domain.SubjectRestriction{
                        SubjectRestriction: &domain.SubjectRestriction_UniversalSet {
                            UniversalSet: &domain.UniversalSet{},
                        },
                    },
                },
                &domain.Relation{
                    Name: "entity or entity set",
                    Description: "userset restriction",
                    RelationExpression: &domain.RelationExpression{
                        Expression: &domain.RelationExpression_Expr{
                            Expr: "_this",
                        },
                    },
                    SubjectRestriction: &domain.SubjectRestriction{
                        SubjectRestriction: &domain.SubjectRestriction_RestrictionSet {
                            RestrictionSet: &domain.SubjectRestrictionSet{
                                Restrictions: []*domain.SubjectRestrictionSet_Restriction{
                                    &domain.SubjectRestrictionSet_Restriction{
                                        Entry: &domain.SubjectRestrictionSet_Restriction_Entity{
                                            Entity: &domain.EntityRestriction{
                                                ResourceName: "group",
                                            },
                                        },
                                    },
                                    &domain.SubjectRestrictionSet_Restriction{
                                        Entry: &domain.SubjectRestrictionSet_Restriction_EntitySet{
                                            EntitySet: &domain.EntitySetRestriction{
                                                ResourceName: "group",
                                                RelationName: "member",
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
        &domain.Resource{
            Name: "group",
            Relations: []*domain.Relation{
                &domain.Relation{
                    Name: "member",
                    RelationExpression: &domain.RelationExpression{
                        Expression: &domain.RelationExpression_Expr{
                            Expr: "_this",
                        },
                    },
                    SubjectRestriction: &domain.SubjectRestriction{
                        SubjectRestriction: &domain.SubjectRestriction_UniversalSet {
                            UniversalSet: &domain.UniversalSet{},
                        },
                    },
                },
            },
        },
    },
}

// FIXME should find a way to test the timestamps

// TODO Test Invalid Policy scenarios:
// - inconsistent subject restrictions
// - duplicated relations
// - missing fields
// - broken relation expr

func TestCreatePolicy(t *testing.T) {
    ctx, service := setup()

    createReq := &api.CreatePolicyRequest{
        PolicyDefinition: &api.PolicyDefinition{
            Definition: &api.PolicyDefinition_Policy_{
                Policy_: testPolicy,
            },
        },
        AppData: []byte("app data"),
    }
    t.Log("Create Policy")
    gotCreate, errCreate := service.CreatePolicy(ctx, createReq)

    wantCreate := &api.CreatePolicyResponse{}
    require.Nil(t, errCreate)
    _testing.ProtoEq(t, gotCreate, wantCreate)

    t.Log("Getting Created Policy")
    getReq := &api.GetPolicyRequest{
        Id: testPolicy.Id,
    }
    wantGet := &api.GetPolicyResponse{
        Record: &domain.PolicyRecord{
            Policy: testPolicy,
            AppData: createReq.AppData,
            CreatedAt: nil,
        },
    }
    gotGet, errGet := service.GetPolicy(ctx, getReq)
    gotGet.Record.CreatedAt = nil
    require.Nil(t, errGet)
    _testing.ProtoEq(t, gotGet, wantGet)
}

func TestCreatePolicyWithIdClashRaisesError(t *testing.T) {
    ctx, service := setup()

    t.Log("Create Policy 10")
    createReq := &api.CreatePolicyRequest{
        PolicyDefinition: &api.PolicyDefinition{
            Definition: &api.PolicyDefinition_Policy_{
                Policy_: testPolicy,
            },
        },
        AppData: []byte("app data"),
    }
    _, errCreate := service.CreatePolicy(ctx, createReq)
    require.Nil(t, errCreate)

    t.Log("Create another Policy 10")
    createReq = &api.CreatePolicyRequest{
        PolicyDefinition: &api.PolicyDefinition{
            Definition: &api.PolicyDefinition_Policy_{
                Policy_: testPolicy,
            },
        },
        AppData: []byte("more data"),
    }
    got, err := service.CreatePolicy(ctx, createReq)
    require.Nil(t, got)
    require.NotNil(t, err)
    require.True(t, errors.Is(err, policy.ErrPolicyExists))
}

func TestUpdatePolicyUpdatesPolicy(t *testing.T) {
    ctx, service := setup()

    policy := proto.Clone(testPolicy).(*domain.Policy)

    t.Log("Create Policy 10")
    createReq := &api.CreatePolicyRequest{
        PolicyDefinition: &api.PolicyDefinition{
            Definition: &api.PolicyDefinition_Policy_{
                Policy_: policy,
            },
        },
        AppData: []byte("app data"),
    }
    _, errCreate := service.CreatePolicy(ctx, createReq)
    require.Nil(t, errCreate)

    t.Log("Update Policy 10")
    policy.Name = "some name"
    updateReq := &api.UpdatePolicyRequest{
        PolicyDefinition: &api.PolicyDefinition{
            Definition: &api.PolicyDefinition_Policy_{
                Policy_: policy,
            },
        },
        AppData: []byte("more data"),
        Strategy: api.UpdatePolicyRequest_IGNORE_ORPHANS,
    }
    updateGot, err := service.UpdatePolicy(ctx, updateReq)
    require.Nil(t, err)
    _testing.ProtoEq(t, updateGot, new(api.UpdatePolicyResponse))

    t.Log("Retrieve Updated Policy 10")
    getReq := &api.GetPolicyRequest{
        Id: policy.Id,
    }
    wantGet := &api.GetPolicyResponse{
        Record: &domain.PolicyRecord{
            Policy: policy,
            AppData: []byte("more data"),
            CreatedAt: nil,
        },
    }
    gotGet, errGet := service.GetPolicy(ctx, getReq)
    gotGet.Record.CreatedAt = nil
    require.Nil(t, errGet)
    _testing.ProtoEq(t, gotGet, wantGet)
}

func TestGetNonExistingPolicyReturnsNil(t *testing.T) {
    ctx, service := setup()

    req := &api.GetPolicyRequest{
        Id: "10",
    }
    got, err := service.GetPolicy(ctx, req)

    want := &api.GetPolicyResponse {
        Record: nil,
    }
    require.Nil(t, err)
    _testing.ProtoEq(t, want, got)
}


func TestDeletingPolicyRemovesFromStore(t *testing.T) {
    ctx, service := setup()

    t.Log("Create Policy")
    createReq := &api.CreatePolicyRequest{
        PolicyDefinition: &api.PolicyDefinition{
            Definition: &api.PolicyDefinition_Policy_{
                Policy_: testPolicy,
            },
        },
        AppData: []byte("app data"),
    }
    _, errCreate := service.CreatePolicy(ctx, createReq)
    require.Nil(t, errCreate)

    t.Log("Delete Policy 10")
    deleteRequest := &api.DeletePolicyRequest{
        Id: testPolicy.Id,
    }
    got, err := service.DeletePolicy(ctx, deleteRequest)

    // Then Policy no longer exists in Store
    require.Nil(t, err)
    require.NotNil(t, got)
    _testing.ProtoEq(t, got, &api.DeletePolicyResponse{ 
        Found: true,
        RelationshipsRemoved: 0,
    })
}

func TestDeletingNonExistingPolicyReturnsNotFound(t *testing.T) {
    ctx, service := setup()

    t.Log("Delete Policy 10")
    deleteRequest := &api.DeletePolicyRequest{
        Id: "10",
    }
    got, err := service.DeletePolicy(ctx, deleteRequest)

    // Then Policy no longer exists in Store
    require.Nil(t, err)
    require.NotNil(t, got)
    _testing.ProtoEq(t, got, &api.DeletePolicyResponse{ 
        Found: false,
        RelationshipsRemoved: 0,
    })
}


func TestSetRelationshipAllowedBySubjectRestrictionRulesSavesRelationship(t *testing.T) {
    ctx, service := setupWithPolicy(testPolicy)

    t.Log("Create Relationship")
    relationship := relationshipBuilder.Relationship("file", "readme.txt", "owner", "user", "bob")
    got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
        PolicyId: testPolicy.Id,
        Relationship: &relationship,
        AppData: []byte("app data"),
    })
    require.Nil(t, err)
    _testing.ProtoEq(t, got, &api.SetRelationshipResponse{
        RecordOverwritten: false,
    })

    t.Log("Fetch Relationship")
    getResponse, err := service.GetRelationship(ctx, &api.GetRelationshipRequest{
        PolicyId: testPolicy.Id,
        Relationship: &relationship,
    })
    require.Nil(t, err)
    _testing.ProtoEq(t, getResponse, &api.GetRelationshipResponse{
        Record: &domain.RelationshipRecord{
            PolicyId: testPolicy.Id,
            Relationship: &relationship,
            AppData: []byte("app data"),
            CreatedAt: getResponse.Record.CreatedAt,
        },
    })
}

func TestSetRelationshipNotAllowedByRestrictionGraphErrors(t *testing.T) { 
    ctx, service := setupWithPolicy(restrictedPolicy)

    t.Log("Create Relationship")
    relationship := relationshipBuilder.Relationship("a", "blah", "empty_restriction_set", "group", "testers")
    got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
        PolicyId: restrictedPolicy.Id,
        Relationship: &relationship,
    })

    require.Nil(t, got)
    require.True(t, errors.Is(err, policy.ErrSubjectNotAllowed))
}

func TestSetRelationshipWithUnknownObjectResourceErrors(t *testing.T) {
    ctx, service := setupWithPolicy(restrictedPolicy)

    relationship := relationshipBuilder.Relationship("unknown", "blah", "empty_restriction_set", "group", "testers")
    got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
        PolicyId: restrictedPolicy.Id,
        Relationship: &relationship,
    })

    require.Nil(t, got)
    require.True(t, errors.Is(err, policy.ErrResourceNotFound))
}

func TestSetRelationshipWithUnknownSubjectResourceErrors(t *testing.T) {
    ctx, service := setupWithPolicy(restrictedPolicy)

    relationship := relationshipBuilder.Relationship("a", "blah", "empty_restriction_set", "foo", "testers")
    got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
        PolicyId: restrictedPolicy.Id,
        Relationship: &relationship,
    })

    require.Nil(t, got)
    require.True(t, errors.Is(err, policy.ErrResourceNotFound))
}

func TestSetRelationshipWithUnknownRelationErrors(t *testing.T) {
    ctx, service := setupWithPolicy(restrictedPolicy)

    relationship := relationshipBuilder.Relationship("a", "blah", "no", "group", "testers")
    got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
        PolicyId: restrictedPolicy.Id,
        Relationship: &relationship,
    })

    require.Nil(t, got)
    require.True(t, errors.Is(err, policy.ErrRelationNotFound))
}

func TestSetRelationshipWithEmptyObjectErrorsOut(t *testing.T) {
    ctx, service := setupWithPolicy(restrictedPolicy)

    relationship := relationshipBuilder.Relationship("a", "", "universal", "group", "testers")
    got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
        PolicyId: restrictedPolicy.Id,
        Relationship: &relationship,
    })

    require.Nil(t, got)
    t.Logf("error: %v", err)
    require.True(t, errors.Is(err, policy.ErrInvalidRelationship))
}

// list policy returns all policies

/*


func (s *RelationshipServiceTestSuite) TestDeleteRelationship() {
}

func (s *RelationshipServiceTestSuite) TestGetNonExistingRelationship() {
}


func (s *RelationshipServiceTestSuite) TestDeleteNonExistingRelationship() {
}

func (s *RelationshipServiceTestSuite) TestUpdatingRelationship() {
}
*/
