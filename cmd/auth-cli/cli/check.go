package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/sourcenetwork/source-zanzibar/types"
)

var checkCmd = &cobra.Command{
	Use:   "check [policyId] [namespace] [srcId] [relation] [dstNamespace] [dstId]",
	Short: "Return all entity, relation pairs reachable from specified actor",
	Args:  cobra.ExactArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		policyId := args[0]

		obj := types.Entity{
			Namespace: args[1],
			Id:        args[2],
		}

		actor := types.Entity{
			Namespace: args[4],
			Id:        args[5],
		}

		auth := client.GetAuthorizer()
		ok, err := auth.Check(policyId, obj, args[3], actor)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", ok)
	},
}
