package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/sourcenetwork/source-zanzibar/authorizer/simple"
	"github.com/sourcenetwork/source-zanzibar/model"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check [namespace] [objId] [relation] [userNamespace] [userId]",
	Short: "Return all nodes reachable from specified user",
	Args:  cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {

		objUset := model.Userset{
			Namespace: args[0],
			ObjectId:  args[1],
			Relation:  args[2],
		}

		user := model.Userset{
			Namespace: args[3],
			ObjectId:  args[4],
			Relation:  model.EMPTY_REL,
		}

		ctx := context.Background()
		tr := buildSampleTupleRepo()
		nr := buildSampleNamespaceRepo()
		checker := simple.NewChecker(nr, tr)

		result, err := checker.Check(ctx, objUset, user)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", result)
	},
}
