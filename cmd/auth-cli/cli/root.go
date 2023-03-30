package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var fixtureFile string = ""

var rootCmd = &cobra.Command{
	Use:   "auth-cli [command]",
	Short: "Perform Zanzibar auth operations",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fixture := buildDefaultFixture()
		if fixtureFile != "" {
			loadFixtureFile(fixtureFile)
		}
		initService(fixture)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&fixtureFile, "fixture", "f", "", "JSON file with fixture data, use [dump-fixture] command to generate sample. Uses default fixture if not specified.")

	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(dumpFixtureCmd)

}

func loadFixtureFile(file string) Fixture {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("could not open fixture: %v", err)
	}

	fixtureBytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("could not read fixture: %v", err)
	}

	fixture := Fixture{}

	err = json.Unmarshal(fixtureBytes, &fixture)
	if err != nil {
		log.Fatalf("could not unmarshal fixture: %v", err)
	}

	log.Printf("loaded fixture from %v", file)

	return fixture
}
