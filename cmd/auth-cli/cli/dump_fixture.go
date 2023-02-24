package cli

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var dumpFixtureCmd = &cobra.Command{
	Use:   "dump-fixture",
	Short: "Dump the default JSON Fixture to stdout. Can be edited and loaded with the fixture flag.",
	Run: func(cmd *cobra.Command, args []string) {
		fixture := buildDefaultFixture()
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "  ")
		err := encoder.Encode(fixture)
		if err != nil {
			log.Fatalf("could not unmarshal default fixture: %v", err)
		}
	},
}
