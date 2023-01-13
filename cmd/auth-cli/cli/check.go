package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/sourcenetwork/source-zanzibar/types"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check [namespace] [srcId] [relation] [dstNamespace] [dstId]",
	Short: "Return all entity, relation pairs reachable from specified actor",
	Args:  cobra.ExactArgs(5),
	Run: func(cmd *cobra.Command, args []string) {

		obj := types.Entity{
			Namespace: args[0],
			Id:        args[1],
		}

		actor := types.Entity{
			Namespace: args[3],
			Id:        args[4],
		}

		auth := client.GetAuthorizer()
		ok, err := auth.Check(POLICY_ID, obj, args[2], actor)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", ok)
	},
}
