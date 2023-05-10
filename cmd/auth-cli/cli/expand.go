package cli

/*
import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/sourcenetwork/zanzi/authorizer/simple"
	"github.com/sourcenetwork/zanzi/model"
	"github.com/sourcenetwork/zanzi/tree"
)

func init() {
	rootCmd.AddCommand(expandCmd)
}

var expandCmd = &cobra.Command{
	Use:   "expand [namespace] [objId] [relation]",
	Short: "Expand the Relation graph for a root userset and target user",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		uset := model.AuthNode{
			Namespace: args[0],
			ObjectId:  args[1],
			Relation:  args[2],
		}

		ctx := context.Background()
		tr := buildSampleTupleRepo()
		nr := buildSampleNamespaceRepo()
		expander := simple.NewExpander(nr, tr)

		usetNode, err := expander.Expand(ctx, uset)
		if err != nil {
			log.Fatal(err)
		}

		printNode(1, &usetNode)
	},
}

func printNode(lvl int, node tree.Node) {
	header := strings.Repeat(" ", lvl)

	switch n := node.(type) {

	case *tree.RuleNode:
		fmt.Printf("%v RuleNode rule=%v\n", header, n.Rule)
		for _, child := range n.Children {
			printNode(lvl+1, child)
		}

	case *tree.UsersetNode:
		uset := fmt.Sprintf("{%v, %v, %v}", n.Userset.Namespace, n.Userset.ObjectId, n.Userset.Relation)
		fmt.Printf("%v UsersetNode userset=%v\n", header, uset)
		printNode(lvl+1, n.Child)

	case *tree.OpNode:
		fmt.Printf("%v OpNode OP=%v\n", header, n.JoinOp)
		printNode(lvl+1, n.Left)
		printNode(lvl+1, n.Right)
	}
}
*/
