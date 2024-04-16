//go:build !wasm

package zanzi

import (
	"fmt"
	"os/user"
	"strings"

	rcdb "github.com/sourcenetwork/raccoondb"
)

// WithDefaultKVStore configures Zanzi to use its default KV Store (LevelDB). Data will be stored in dataDir
func WithDefaultKVStore(path string) option {
	return func(z *Zanzi) error {
		if strings.HasPrefix(path, "~") {
			usr, err := user.Current()
			if err != nil {
				return fmt.Errorf("error identifying user: %w", err)
			}
			home := usr.HomeDir

			path = strings.Replace(path, "~", home, 1)
		}

		kv, err := rcdb.NewLevelDB(path, dataFile)
		if err != nil {
			return fmt.Errorf("error initializing kv store: %v", err)
		}
		return WithKVStore(kv)(z)
	}
}
