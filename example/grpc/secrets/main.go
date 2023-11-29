package main

import (
	"context"
	"encoding/json"
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

	policyServ := api.NewPolicyServiceClient(conn)
	relGraph := api.NewRelationGraphClient(conn)

	ctx := context.Background()

	polReq := &api.CreatePolicyRequest{
		PolicyDefinition: &api.PolicyDefinition{
			PolicyYaml: policyYaml,
		},
	}

	buf, err := json.Marshal(polReq)
	if err != nil {
		panic(err)
	}
	fmt.Println("policy (json):", string(buf))
	// Create Policy
	createResponse, err := policyServ.CreatePolicy(ctx, polReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created policy: %v", createResponse.Record.Policy)

	// Set Relationships
	for _, relationship := range relationships {
		data := &api.SetRelationshipRequest{
			PolicyId:     createResponse.Record.Policy.Id,
			Relationship: &relationship,
		}
		buf, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf))
		_, err = policyServ.SetRelationship(ctx, data)
		if err != nil {
			log.Fatal("set relationship:", err)
		}
	}

	res, err := relGraph.DumpRelationships(ctx, &api.DumpRelationshipsRequest{
		PolicyId: "1",
		Format:   api.DumpRelationshipsRequest_DOT,
	})
	if err != nil {
		log.Fatal("dump relationship:", err)
	}
	dot := res.Dump.(*api.DumpRelationshipResponse_Dot).Dot
	fmt.Println(dot)

	check, err := relGraph.Check(ctx, &api.CheckRequest{
		PolicyId: "1",
		AccessRequest: &domain.AccessRequest{
			Object:   domain.NewEntity("secret", "privateKey"),
			Relation: "read",
			Subject:  domain.NewEntity("user", "alice"),
		},
	})

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
`

var builder domain.RelationshipBuilder

var relationships []domain.Relationship = []domain.Relationship{
	builder.Relationship("secret", "privateKey", "owner", "user", "charlie"),
	builder.Relationship("secret", "privateKey", "collaborator", "user", "alice"),
}
