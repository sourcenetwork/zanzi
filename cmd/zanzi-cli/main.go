package main 

import (
    "context"
    "log"
    "os"

    "github.com/spf13/cobra"

    "github.com/sourcenetwork/zanzi/pkg/api"
)

var rootCmd = &cobra.Command{
	Use:   "zanzi-cli",
	Short: "zanzi-cli provides a CLI to zanzi's GRPC services",
	Long: ``,
}


func init() {
    rootCmd.AddCommand(api.PolicyServiceClientCommand())
    rootCmd.AddCommand(api.RelationGraphClientCommand())
}

func main() {
    ctx := context.Background()
    err := rootCmd.ExecuteContext(ctx)
    
    if err != nil {
        log.Printf("%v\n", err)
        os.Exit(1)
    }
}
