package cli
/*

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/sourcenetwork/source-zanzibar/authorizer/simple"
	"github.com/sourcenetwork/source-zanzibar/model"
)

func init() {
	rootCmd.AddCommand(reverseCmd)
}

var reverseCmd = &cobra.Command{
	Use:   "reverse [namespace] [userId]",
	Short: "Return all nodes reachable from specified user",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		user := model.User{
			Userset: &model.AuthNode{
				Namespace: args[0],
				ObjectId:  args[1],
				Relation:  model.EMPTY_REL,
			},
			Type: model.UserType_USER,
		}

		ctx := context.Background()
		tr := buildSampleTupleRepo()
		nr := buildSampleNamespaceRepo()
		rev := simple.NewReverser(nr, tr)

		nodes, err := rev.ReverseLookup(ctx, user)
		if err != nil {
			log.Fatal(err)
		}

		for _, node := range nodes {
			fmt.Printf("%v\n", node.ToKey())
		}
	},
}
*/
