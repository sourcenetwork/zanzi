package main

import (
	"context"
	"fmt"
	"log"
	"os"

	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
)

var rootCmd = &cobra.Command{
	Use:   "zanzi-cli",
	Short: "zanzi-cli provides a CLI to zanzi's GRPC services",
	Long:  ``,
}

var relationshipCmd = &cobra.Command{
	Use:   "relationship",
	Short: "relationship builder",
}

func createRelationshipCommand() *cobra.Command {
	var (
		policyid string
	)
	cfg := client.NewConfig()
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a new relationship",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			if policyid == "" {
				return fmt.Errorf("missing policyID")
			}

			var builder domain.RelationshipBuilder
			relationships := make([]domain.Relationship, len(args))
			var err error
			for i, arg := range args {
				relationships[i], err = builder.RelationshipFromString(arg)
				if err != nil {
					return err
				}
			}

			cred := insecure.NewCredentials()
			cc, err := grpc.DialContext(cmd.Context(), cfg.ServerAddr, grpc.WithTransportCredentials(cred))
			if err != nil {
				return fmt.Errorf("grpc dial: %w", err)
			}

			cli := api.NewPolicyServiceClient(cc)
			for _, rel := range relationships {
				_, err = cli.SetRelationship(cmd.Context(), &api.SetRelationshipRequest{
					PolicyId:     policyid,
					Relationship: &rel,
				})
				if err != nil {
					return fmt.Errorf("set relationship: %w", err)
				}
			}

			return nil
		},
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.PersistentFlags().StringVar(&policyid, "policy-id", "", "policy identifer")
	cmd.MarkFlagRequired("policy-id")
	return cmd
}

func checkCommand() *cobra.Command {
	var (
		policyid string
	)
	cfg := client.NewConfig()
	cmd := &cobra.Command{
		Use:   "check",
		Short: "check a relationship permission",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			if policyid == "" {
				return fmt.Errorf("missing policyID")
			}

			var builder domain.RelationshipBuilder
			relationships := make([]domain.Relationship, len(args))
			var err error
			for i, arg := range args {
				relationships[i], err = builder.RelationshipFromString(arg)
				if err != nil {
					return err
				}
			}

			cred := insecure.NewCredentials()
			cc, err := grpc.DialContext(cmd.Context(), cfg.ServerAddr, grpc.WithTransportCredentials(cred))
			if err != nil {
				return fmt.Errorf("grpc dial: %w", err)
			}

			cli := api.NewRelationGraphClient(cc)
			for i, rel := range relationships {
				relstring := args[i]
				check, err := cli.Check(cmd.Context(), &api.CheckRequest{
					PolicyId: policyid,
					AccessRequest: &domain.AccessRequest{
						Object:   domain.NewEntity(rel.Object.Resource, rel.Object.Id),
						Relation: rel.Relation,
						Subject:  domain.NewEntity(rel.Subject.GetEntity().Resource, rel.Subject.GetEntity().Id),
					},
				})
				if err != nil {
					log.Fatalf("check(%v): %v", relstring, err)
				}
				fmt.Printf("check(%v): %v\n", relstring, check.Result.Authorized)
			}

			return nil
		},
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.PersistentFlags().StringVar(&policyid, "policy-id", "", "policy identifer")
	cmd.MarkFlagRequired("policy-id")
	return cmd
}

func init() {
	relationshipCmd.AddCommand(createRelationshipCommand())
	relationshipCmd.AddCommand(checkCommand())
	rootCmd.AddCommand(api.PolicyServiceClientCommand())
	rootCmd.AddCommand(api.RelationGraphClientCommand())
	rootCmd.AddCommand(relationshipCmd)
}

func main() {
	ctx := context.Background()
	err := rootCmd.ExecuteContext(ctx)

	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
}
