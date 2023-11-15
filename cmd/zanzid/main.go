package main

import (
	"log"

	"github.com/spf13/cobra"

	zanzi "github.com/sourcenetwork/zanzi"
	"github.com/sourcenetwork/zanzi/pkg/server"
)

func main() {
	rootCmd.Execute()
}

var dataDir string
var address string

var rootCmd = &cobra.Command{
	Use:   "zanzid [flags]",
	Short: "zanzid starts a grpc server for zanzi",
	Run: func(cmd *cobra.Command, args []string) {
		z, err := zanzi.New(
			zanzi.WithDefaultLogger(),
			zanzi.WithDefaultKVStore(dataDir),
		)
		if err != nil {
			log.Fatalf("error initializing zanzi: %v", err)
		}

		server := server.NewServer(address)
		log.Printf("Initializing Zanzi Service on %v", address)

		if err := server.Init(&z); err != nil {
			log.Fatalf("error starting grpc server: %v", err)
		}

		server.Run()
	},
}

func init() {
	rootCmd.Flags().StringVar(&address, "address", "0.0.0.0:8080", "sets the address zanzi will listen on")
	rootCmd.Flags().StringVar(&dataDir, "data_dir", "~/.zanzi", "sets the directory zanzi will store application data")
}
