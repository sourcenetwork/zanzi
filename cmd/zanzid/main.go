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
var grpcAddress string
var restAddress string

var rootCmd = &cobra.Command{
	Use:   "zanzid [flags]",
	Short: "zanzid starts a grpc server for zanzi",
	Run: func(cmd *cobra.Command, args []string) {
		z, err := zanzi.New(
			zanzi.WithDefaultLogger(),
			// zanzi.WithDefaultKVStore(dataDir),
		)
		if err != nil {
			log.Fatalf("error initializing zanzi: %v", err)
		}

		svr := server.NewServer(grpcAddress)
		log.Printf("Initializing Zanzi Service on %v", grpcAddress)

		if err := svr.Init(&z); err != nil {
			log.Fatalf("error starting grpc server: %v", err)
		}
		go func() {
			log.Printf("Serving gRPC-Service on http://%s", grpcAddress)
			svr.Run()
		}()

		// Serve REST with gRPC-Gateway.
		gwServer, err := server.NewGRPCGatewayServer(grpcAddress, restAddress)
		if err != nil {
			log.Fatalf("create gRPC gateway server: %v", err)
		}
		log.Printf("Serving gRPC-Gateway on http://%s", restAddress)
		gwServer.ListenAndServe()
	},
}

func init() {
	rootCmd.Flags().StringVar(&grpcAddress, "address", "0.0.0.0:8080", "sets the GRPC address zanzi will listen on")
	rootCmd.Flags().StringVar(&restAddress, "rest", "0.0.0.0:8090", "sets the REST address zanzi will listen on")
	rootCmd.Flags().StringVar(&dataDir, "data_dir", "~/.zanzi", "sets the directory zanzi will store application data")
}
