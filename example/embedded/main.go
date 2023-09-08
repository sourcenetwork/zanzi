package main

import (
    "fmt"
    "log"
    "context"

    "github.com/sourcenetwork/zanzi/pkg/domain"
    "github.com/sourcenetwork/zanzi/pkg/api"
    _zanzi "github.com/sourcenetwork/zanzi"
)

func main() {
    // Init Zanzi
    zanzi, err := _zanzi.New(
        _zanzi.WithDevelopmentLogger(),
        _zanzi.WithMemKVStore(),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()
    policyServ := zanzi.GetPolicyService()
    relGraph := zanzi.GetRelationGraphService()

    setupData(ctx, policyServ)

    res, err := relGraph.DumpRelationships(ctx, &api.DumpRelationshipsRequest{
        PolicyId: "10",
        Format: api.DumpRelationshipsRequest_DOT,
    })
    dot := res.Dump.(*api.DumpRelationshipResponse_Dot).Dot
    fmt.Println(dot)

    /*
    res, err := relGraph.Check(ctx, &api.CheckRequest{
        PolicyId: "10",
        AccessRequest: &domain.AccessRequest{
            Object: domain.NewEntity("file", "readme"),
            Relation: "read",
            Subject: domain.NewEntity("user", "bob"),
        },
    })

    fmt.Println(res.Result.Authorized)
    */
        
    /*
    res, err := relGraph.Expand(ctx, &api.ExpandRequest{
        PolicyId: "10",
        Root: &domain.RelationNode{
            Node: &domain.RelationNode_EntitySet{
                EntitySet: &domain.EntitySetNode{
                    Object: domain.NewEntity("file", "readme"),
                    Relation: "read",
                },
            },
        },
        Format: api.ExplainFormat_DOT,
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res.GoalTree)
    */
}

func setupData(ctx context.Context, service api.PolicyServiceServer) {
    // Create Policy
    createResponse, err := service.CreatePolicy(ctx, &api.CreatePolicyRequest{
            PolicyDefinition: &api.PolicyDefinition{
                Definition: &api.PolicyDefinition_PolicyYaml{
                    PolicyYaml: policyYaml,
                },
            },
    })
    if err != nil {
        log.Fatal(err)
    } 
    log.Printf("Created policy: %v", createResponse.Record.Policy)

    // Set Relationships
    for _, relationship := range relationships {
        _, err = service.SetRelationship(ctx, &api.SetRelationshipRequest{
            PolicyId: createResponse.Record.Policy.Id,
            Relationship: &relationship,
        })
        if err != nil {
            log.Fatal(err)
        }
    }
}

var policyYaml = `
id: 10
name: test
doc: test policy

resources:
  file:
    doc: file resource
    relations:
      owner:
        expr: _this
        types:
          - '*'
      parent:
        expr: _this
        types:
          - directory
      read:
        expr: owner + parent->read
        types: []
  directory:
    relations:
      owner:
        types:
          - user
          - group:member
          - group:owner
        expr: _this
      reader:
        expr: _this
        types:
          - user
          - group:member
      read:
        expr: owner + reader

  group:
    relations:
      member:
        expr: _this + owner
        types:
          - user
      owner:
        expr: _this
        types:
          - user
  user:

attributes:
  foo: bar
`

var builder domain.RelationshipBuilder

var relationships []domain.Relationship = []domain.Relationship{
    builder.Relationship("file", "readme", "owner", "user", "charlie"),
    builder.Relationship("file", "readme", "parent", "directory", "proj"),
    builder.EntitySet("directory", "proj", "owner", "group", "eng", "owner"),
    builder.EntitySet("directory", "proj", "reader", "group", "eng", "member"),
    builder.Relationship("group", "eng", "owner", "user", "alice"),
    builder.Relationship("group", "eng", "member", "user", "bob"),
}
