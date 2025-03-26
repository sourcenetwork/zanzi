package integration

import (
	"context"
	"errors"
	"testing"

	rcdb "github.com/sourcenetwork/raccoondb"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/internal/store/kv_store"
	_testing "github.com/sourcenetwork/zanzi/internal/testing"
	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/policy_definition"
)

func setup() (context.Context, api.PolicyServiceServer) {
	kv := rcdb.NewMemKV()
	kvStore, err := kv_store.NewKVStore(kv)
	if err != nil {
		panic(err)
	}

	return context.Background(), policy.NewService(kvStore.GetPolicyRepository())
}

func setupWithPolicy(policies ...*domain.Policy) (context.Context, api.PolicyServiceServer) {
	ctx, service := setup()

	for _, policy := range policies {
		createReq := &api.CreatePolicyRequest{
			PolicyDefinition: &api.PolicyDefinition{
				Definition: &api.PolicyDefinition_Policy{
					Policy: policy,
				},
			},
			AppData: []byte("app data"),
		}
		_, err := service.CreatePolicy(ctx, createReq)
		if err != nil {
			panic(err)
		}
	}
	return ctx, service
}

var relationshipBuilder domain.RelationshipBuilder = domain.RelationshipBuilder{}

var testPolicy *domain.Policy = &domain.Policy{
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

var restrictedPolicy *domain.Policy = &domain.Policy{
	Id:          "11",
	Name:        "test",
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
						SubjectRestriction: &domain.SubjectRestriction_RestrictionSet{
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
						SubjectRestriction: &domain.SubjectRestriction_UniversalSet{
							UniversalSet: &domain.UniversalSet{},
						},
					},
				},
				&domain.Relation{
					Name:        "entity or entity set",
					Description: "userset restriction",
					RelationExpression: &domain.RelationExpression{
						Expression: &domain.RelationExpression_Expr{
							Expr: "_this",
						},
					},
					SubjectRestriction: &domain.SubjectRestriction{
						SubjectRestriction: &domain.SubjectRestriction_RestrictionSet{
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
						SubjectRestriction: &domain.SubjectRestriction_UniversalSet{
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
			Definition: &api.PolicyDefinition_Policy{
				Policy: testPolicy,
			},
		},
		AppData: []byte("app data"),
	}
	t.Log("Create Policy")
	_, errCreate := service.CreatePolicy(ctx, createReq)

	require.Nil(t, errCreate)

	t.Log("Getting Created Policy")
	getReq := &api.GetPolicyRequest{
		Id: testPolicy.Id,
	}
	wantGet := &api.GetPolicyResponse{
		Record: &domain.PolicyRecord{
			Policy:    testPolicy,
			AppData:   createReq.AppData,
			CreatedAt: nil,
		},
	}
	gotGet, errGet := service.GetPolicy(ctx, getReq)
	gotGet.Record.CreatedAt = nil
	require.Nil(t, errGet)
	wantGet.Reset()
	gotGet.Reset()
	require.Equal(t, gotGet, wantGet)
}

func TestCreatePolicyWithIdClashRaisesError(t *testing.T) {
	ctx, service := setup()

	t.Log("Create Policy 10")
	createReq := &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_Policy{
				Policy: testPolicy,
			},
		},
		AppData: []byte("app data"),
	}
	_, errCreate := service.CreatePolicy(ctx, createReq)
	require.Nil(t, errCreate)

	t.Log("Create another Policy 10")
	createReq = &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_Policy{
				Policy: testPolicy,
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
			Definition: &api.PolicyDefinition_Policy{
				Policy: policy,
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
			Definition: &api.PolicyDefinition_Policy{
				Policy: policy,
			},
		},
		AppData:  []byte("more data"),
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
			Policy:    policy,
			AppData:   []byte("more data"),
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

	want := &api.GetPolicyResponse{
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
			Definition: &api.PolicyDefinition_Policy{
				Policy: testPolicy,
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
		Found:                true,
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
		Found:                false,
		RelationshipsRemoved: 0,
	})
}

func TestSetRelationshipAllowedBySubjectRestrictionRulesSavesRelationship(t *testing.T) {
	ctx, service := setupWithPolicy(testPolicy)

	t.Log("Create Relationship")
	relationship := relationshipBuilder.Relationship("file", "readme.txt", "owner", "user", "bob")
	got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     testPolicy.Id,
		Relationship: &relationship,
		AppData:      []byte("app data"),
	})
	require.Nil(t, err)
	_testing.ProtoEq(t, got, &api.SetRelationshipResponse{
		RecordOverwritten: false,
	})

	t.Log("Fetch Relationship")
	getResponse, err := service.GetRelationship(ctx, &api.GetRelationshipRequest{
		PolicyId:     testPolicy.Id,
		Relationship: &relationship,
	})
	require.Nil(t, err)
	_testing.ProtoEq(t, getResponse, &api.GetRelationshipResponse{
		Record: &domain.RelationshipRecord{
			PolicyId:     testPolicy.Id,
			Relationship: &relationship,
			AppData:      []byte("app data"),
			CreatedAt:    getResponse.Record.CreatedAt,
		},
	})
}

func TestSetRelationshipNotAllowedByRestrictionGraphErrors(t *testing.T) {
	ctx, service := setupWithPolicy(restrictedPolicy)

	t.Log("Create Relationship")
	relationship := relationshipBuilder.Relationship("a", "blah", "empty_restriction_set", "group", "testers")
	got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     restrictedPolicy.Id,
		Relationship: &relationship,
	})

	require.Nil(t, got)
	require.True(t, errors.Is(err, policy.ErrSubjectNotAllowed))
}

func TestSetRelationshipWithUnknownObjectResourceErrors(t *testing.T) {
	ctx, service := setupWithPolicy(restrictedPolicy)

	relationship := relationshipBuilder.Relationship("unknown", "blah", "empty_restriction_set", "group", "testers")
	got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     restrictedPolicy.Id,
		Relationship: &relationship,
	})

	require.Nil(t, got)
	require.True(t, errors.Is(err, policy.ErrResourceNotFound))
}

func TestSetRelationshipWithUnknownSubjectResourceErrors(t *testing.T) {
	ctx, service := setupWithPolicy(restrictedPolicy)

	relationship := relationshipBuilder.Relationship("a", "blah", "empty_restriction_set", "foo", "testers")
	got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     restrictedPolicy.Id,
		Relationship: &relationship,
	})

	require.Nil(t, got)
	require.True(t, errors.Is(err, policy.ErrResourceNotFound))
}

func TestSetRelationshipWithUnknownRelationErrors(t *testing.T) {
	ctx, service := setupWithPolicy(restrictedPolicy)

	relationship := relationshipBuilder.Relationship("a", "blah", "no", "group", "testers")
	got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     restrictedPolicy.Id,
		Relationship: &relationship,
	})

	require.Nil(t, got)
	require.True(t, errors.Is(err, policy.ErrRelationNotFound))
}

func TestSetRelationshipWithEmptyObjectErrorsOut(t *testing.T) {
	ctx, service := setupWithPolicy(restrictedPolicy)

	relationship := relationshipBuilder.Relationship("a", "", "universal", "group", "testers")
	got, err := service.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     restrictedPolicy.Id,
		Relationship: &relationship,
	})

	require.Nil(t, got)
	t.Logf("error: %v", err)
	require.True(t, errors.Is(err, policy.ErrInvalidRelationship))
}

func TestListPolicyIdsReturnsAllPolicies(t *testing.T) {
	ctx, service := setupWithPolicy(testPolicy, restrictedPolicy)

	got, err := service.ListPolicyIds(ctx, &api.ListPolicyIdsRequest{})

	want := []*api.ListPolicyIdsResponse_Record{
		&api.ListPolicyIdsResponse_Record{Id: testPolicy.Id},
		&api.ListPolicyIdsResponse_Record{Id: restrictedPolicy.Id},
	}
	require.Nil(t, err)
	require.Equal(t, want, got.Records)
}

func TestFindRelationshipRecords_ObjectSelectorReferencingUnknownRelationReturnsErr(t *testing.T) {
	builder := domain.SelectorBuilder{}
	builder.WithObject(domain.NewEntity("unknown-resource", "abc"))
	builder.AnyRelation()
	builder.AnySubject()
	selector := builder.Build()

	ctx, service := setupWithPolicy(testPolicy)

	resp, err := service.FindRelationshipRecords(ctx, &api.FindRelationshipRecordsRequest{
		PolicyId: testPolicy.Id,
		Selector: &selector,
	})

	require.Nil(t, resp)
	require.ErrorIs(t, err, policy.ErrResourceNotFound)
}

func TestFindRelationshipRecords_RelationSpecToNonExistingRelationReturnsError(t *testing.T) {
	builder := domain.SelectorBuilder{}
	builder.WithObject(domain.NewEntity("file", "abc"))
	builder.WithRelation("a-relation")
	builder.AnySubject()
	selector := builder.Build()

	ctx, service := setupWithPolicy(testPolicy)

	resp, err := service.FindRelationshipRecords(ctx, &api.FindRelationshipRecordsRequest{
		PolicyId: testPolicy.Id,
		Selector: &selector,
	})

	require.Nil(t, resp)
	require.ErrorIs(t, err, policy.ErrRelationNotFound)
}

func TestFindRelationshipRecords_SubjectSpecReferencingNonExistingResourceReturnsError(t *testing.T) {
	builder := domain.SelectorBuilder{}
	builder.WithObject(domain.NewEntity("file", "abc"))
	builder.WithRelation("owner")
	builder.WithSubject(&domain.Subject{
		Subject: &domain.Subject_Entity{
			Entity: domain.NewEntity("missing-resource", "abc"),
		},
	})
	selector := builder.Build()

	ctx, service := setupWithPolicy(testPolicy)

	resp, err := service.FindRelationshipRecords(ctx, &api.FindRelationshipRecordsRequest{
		PolicyId: testPolicy.Id,
		Selector: &selector,
	})

	require.Nil(t, resp)
	require.ErrorIs(t, err, policy.ErrResourceNotFound)
}

func Test_EditPolicy_RemovingPolicyRelation_RemovesForwardRelationshipsForThatRelation(t *testing.T) {
	ctx, serv := setup()

	old :=
		`
id: test
name: test
resources:
  file:
    relations:
      owner:
        expr: _this
        types: 
          - "*"
  user:
    relations:
`
	new :=
		`
id: test
name: test
resources:
  file:
    relations:
  user:
    relations:
`

	// Given policy with relation owner for resource file
	// and relationships for owner
	_, err := serv.CreatePolicy(ctx, &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: old,
			},
		},
	})
	require.NoError(t, err)

	rel := relationshipBuilder.Relationship("file", "foo", "owner", "user", "bob")
	_, err = serv.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     "test",
		Relationship: &rel,
	})
	require.NoError(t, err)
	rel = relationshipBuilder.Relationship("file", "bar", "owner", "user", "alice")
	_, err = serv.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     "test",
		Relationship: &rel,
	})
	require.NoError(t, err)

	// When I edit Policy
	resp, err := serv.EditPolicy(ctx, &api.EditPolicyRequest{
		PolicyId: "test",
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: new,
			},
		},
	})

	// Then I get no errors and 2 relationships are removed
	require.NoError(t, err)
	require.Equal(t, uint64(2), resp.RemovedRelationshipsCount)
}

func Test_EditPolicy_EditingPolicy_NewPolicyIsReturned(t *testing.T) {
	ctx, serv := setup()

	old := `
id: test
name: test
resources:
  file:
    relations:
      owner:
        expr: _this

  user:
    relations:
      foo:
        expr: _this
      mutate_expr:
        expr: foo + _this
`
	new := `
id: test
name: test
resources:
  user:
    relations:
      mutate_expr:
        expr: _this
  new_file:
    relations:
      new_owner:
        expr: _this
`

	_, err := serv.CreatePolicy(ctx, &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: old,
			},
		},
	})
	require.NoError(t, err)

	resp, err := serv.EditPolicy(ctx, &api.EditPolicyRequest{
		PolicyId: "test",
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: new,
			},
		},
	})
	require.NoError(t, err)
	wantPol, err := policy_definition.PolicyFromYaml(new)
	require.NoError(t, err)

	require.Equal(t, uint64(0), resp.RemovedRelationshipsCount)
	wantPol.Reset()
	resp.Record.Policy.Reset()
	require.Equal(t, wantPol, resp.Record.Policy)
}

func Test_EditPolicy_EditingPolicyWithInvalidId_ReturnsErr(t *testing.T) {
	ctx, serv := setup()

	new := `
id: test
name: test
resources:
  user:
    relations:
`

	resp, err := serv.EditPolicy(ctx, &api.EditPolicyRequest{
		PolicyId: "not-defined",
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: new,
			},
		},
	})
	require.ErrorIs(t, err, policy.ErrPolicyNotFound)
	require.Nil(t, resp)
}

func Test_EditPolicy_RemovingResource_RemovesRelationships(t *testing.T) {
	ctx, serv := setup()

	old :=
		`
id: test
name: test
resources:
  file:
    relations:
      reader:
        expr: _this
        types: 
          - "*"
      owner:
        expr: _this
        types: 
          - "*"
  user:
    relations:
`
	new :=
		`
id: test
name: test
resources:
  user:
    relations:
`

	// Given policy with relation owner for resource file
	// and relationships for owner
	_, err := serv.CreatePolicy(ctx, &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: old,
			},
		},
	})
	require.NoError(t, err)

	rel := relationshipBuilder.Relationship("file", "foo", "owner", "user", "bob")
	_, err = serv.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     "test",
		Relationship: &rel,
	})
	require.NoError(t, err)
	rel = relationshipBuilder.Relationship("file", "bar", "reader", "user", "alice")
	_, err = serv.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     "test",
		Relationship: &rel,
	})
	require.NoError(t, err)

	// When I edit Policy
	resp, err := serv.EditPolicy(ctx, &api.EditPolicyRequest{
		PolicyId: "test",
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: new,
			},
		},
	})

	// Then I get no errors and 2 relationships are removed
	require.NoError(t, err)
	require.Equal(t, uint64(2), resp.RemovedRelationshipsCount)
}

func Test_EditPolicy_RemovingPolicyRelation_RemovesBackwardsRelationshipsForThatRelation(t *testing.T) {
	ctx, serv := setup()

	old :=
		`
id: test
name: test
resources:
  file:
    relations:
      owner:
        expr: _this
        types: 
          - "*"
  group:
    relations:
      member:
        expr: _this
        types:
          - "*"
`
	new :=
		`
id: test
name: test
resources:
  file:
    relations:
      owner:
        expr: _this
        types: 
          - "*"
  group:
    relations:
`

	_, err := serv.CreatePolicy(ctx, &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: old,
			},
		},
	})
	require.NoError(t, err)

	rel := relationshipBuilder.EntitySet("file", "foo", "owner", "group", "admin", "member")
	_, err = serv.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     "test",
		Relationship: &rel,
	})
	require.NoError(t, err)
	rel = relationshipBuilder.Relationship("file", "bar", "owner", "group", "test")
	_, err = serv.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     "test",
		Relationship: &rel,
	})
	require.NoError(t, err)

	// When I edit Policy
	resp, err := serv.EditPolicy(ctx, &api.EditPolicyRequest{
		PolicyId: "test",
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: new,
			},
		},
	})

	// Then I get no errors and the 1 relationship which
	// has group#member as subjects are removed
	require.NoError(t, err)
	require.Equal(t, uint64(1), resp.RemovedRelationshipsCount)
}

func Test_EditPolicy_RemovingPolicyResource_RemovesAllRelationshipsWithSubjectsFromThatResource(t *testing.T) {
	ctx, serv := setup()

	old :=
		`
id: test
name: test
resources:
  file:
    relations:
      owner:
        expr: _this
        types: 
          - "*"
  group:
    relations:
      member:
        expr: _this
        types:
          - "*"
`
	new :=
		`
id: test
name: test
resources:
  file:
    relations:
      owner:
        expr: _this
        types: 
          - "*"
`

	_, err := serv.CreatePolicy(ctx, &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: old,
			},
		},
	})
	require.NoError(t, err)

	rel := relationshipBuilder.EntitySet("file", "foo", "owner", "group", "admin", "member")
	_, err = serv.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     "test",
		Relationship: &rel,
	})
	require.NoError(t, err)
	rel = relationshipBuilder.Relationship("file", "bar", "owner", "group", "test")
	_, err = serv.SetRelationship(ctx, &api.SetRelationshipRequest{
		PolicyId:     "test",
		Relationship: &rel,
	})
	require.NoError(t, err)

	// When I edit Policy
	resp, err := serv.EditPolicy(ctx, &api.EditPolicyRequest{
		PolicyId: "test",
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: new,
			},
		},
	})

	// Then I get no errors and the 1 relationship which
	// has group#member as subjects are removed
	require.NoError(t, err)
	require.Equal(t, uint64(2), resp.RemovedRelationshipsCount)
}

func Test_EditPolicy_PreservesAppData(t *testing.T) {
	ctx, serv := setup()

	appData := []byte{0, 0, 1}
	policy :=
		`
id: test
name: test
`
	_, err := serv.CreatePolicy(ctx, &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: policy,
			},
		},
		AppData: appData,
	})
	require.NoError(t, err)

	// When I edit the Policy
	response, err := serv.EditPolicy(ctx, &api.EditPolicyRequest{
		PolicyId: "test",
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: policy,
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, appData, response.Record.AppData)

	getResp, err := serv.GetPolicy(ctx, &api.GetPolicyRequest{
		Id: "test",
	})
	require.NoError(t, err)
	require.Equal(t, appData, getResp.Record.AppData)
}

func Test_EditPolicyAppData_UpdatesAppData(t *testing.T) {
	ctx, serv := setup()

	policy :=
		`
id: test
name: test
`
	_, err := serv.CreatePolicy(ctx, &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: policy,
			},
		},
		AppData: []byte{0, 0, 1},
	})
	require.NoError(t, err)

	// When I edit the Policy AppData
	newData := []byte{0, 0, 2}
	resp, err := serv.EditPolicyAppData(ctx, &api.EditPolicyAppDataRequest{
		PolicyId: "test",
		AppData:  newData,
	})

	// Then response contains the new app data and record is updated
	require.NoError(t, err)
	require.Equal(t, newData, resp.Record.AppData)

	getResp, err := serv.GetPolicy(ctx, &api.GetPolicyRequest{
		Id: "test",
	})
	require.NoError(t, err)
	require.Equal(t, newData, getResp.Record.AppData)
}

func Test_EditPolicyAppData_SendingNilErrasesAppData(t *testing.T) {
	ctx, serv := setup()

	policy :=
		`
id: test
name: test
`
	_, err := serv.CreatePolicy(ctx, &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			Definition: &api.PolicyDefinition_PolicyYaml{
				PolicyYaml: policy,
			},
		},
		AppData: []byte{0, 0, 1},
	})
	require.NoError(t, err)

	// When I edit the Policy AppData
	var newData []byte
	resp, err := serv.EditPolicyAppData(ctx, &api.EditPolicyAppDataRequest{
		PolicyId: "test",
		AppData:  newData,
	})

	// Then response contains the new app data and record is updated
	require.NoError(t, err)
	require.Equal(t, newData, resp.Record.AppData)

	getResp, err := serv.GetPolicy(ctx, &api.GetPolicyRequest{
		Id: "test",
	})
	require.NoError(t, err)
	require.Equal(t, newData, getResp.Record.AppData)
}
