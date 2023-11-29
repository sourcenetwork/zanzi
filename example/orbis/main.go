package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

func main() {
	target := "127.0.0.1:8080"
	cred := insecure.NewCredentials()
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal("dial:", err)
	}

	// policyServ := api.NewPolicyServiceClient(conn)
	relGraph := api.NewRelationGraphClient(conn)

	ctx := context.Background()

	// Create Policy
	// createResponse, err := policyServ.CreatePolicy(ctx, &api.CreatePolicyRequest{
	// 	PolicyDefinition: &api.PolicyDefinition{
	// 		Definition: &api.PolicyDefinition_PolicyYaml{
	// 			PolicyYaml: policyYaml,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Created policy: %v", createResponse.Record.Policy)

	// Set Relationships
	// for _, relationship := range relationships {
	// 	_, err = policyServ.SetRelationship(ctx, &api.SetRelationshipRequest{
	// 		PolicyId:     createResponse.Record.Policy.Id,
	// 		Relationship: &relationship,
	// 	})
	// 	if err != nil {
	// 		log.Fatal("set relationship:", err)
	// 	}
	// }

	// res, err := relGraph.DumpRelationships(ctx, &api.DumpRelationshipsRequest{
	// 	PolicyId: "10",
	// 	Format:   api.DumpRelationshipsRequest_DOT,
	// })
	// if err != nil {
	// 	log.Fatal("dump relationship:", err)
	// }
	// dot := res.Dump.(*api.DumpRelationshipResponse_Dot).Dot
	// fmt.Println(dot)

	check, err := relGraph.Check(ctx, &api.CheckRequest{
		PolicyId: "1",
		AccessRequest: &domain.AccessRequest{
			Object:   domain.NewEntity("secret", "mysupersecret"),
			Relation: "read",
			Subject:  domain.NewEntity("user", "did:key:z6MkpFhGHvoRuTRRqKMXaPiZAqp2ypnhV9jnWt4uijkFYpfy"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Check Result:", check.Result.Authorized)

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

var policyYaml = `
id: 1
name: test
doc: test policy

resources:
  secret:
    doc: sensitive secret
    relations:
      owner:
        expr: _this
        types:
          - '*'
      collaborator:
        expr: _this
        types:
          - user
      read:
        expr: owner + collaborator
        types: []
  user:

attributes:
  foo: bar
`

// var builder domain.RelationshipBuilder

// var relationships []domain.Relationship = []domain.Relationship{
// 	builder.Relationship("file", "readme", "owner", "user", "charlie"),
// 	builder.Relationship("file", "readme", "parent", "directory", "proj"),
// 	builder.EntitySet("directory", "proj", "owner", "group", "eng", "owner"),
// 	builder.EntitySet("directory", "proj", "reader", "group", "eng", "member"),
// 	builder.Relationship("group", "eng", "owner", "user", "alice"),
// 	builder.Relationship("group", "eng", "member", "user", "bob"),
// }
