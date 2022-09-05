package querier

import (
    "context"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/tree"
)

// Simple Check implementation leverages a full expand in order to verify whether 
// userset is included in final expand tree
func Check(ctx context.Context, userset model.Userset) (bool, error) {
	usetNode, err := Expand(ctx, userset)
	if err != nil {
		// wrap
		return false, err
	}

	usets := tree.Eval(usetNode)
	key := userset.ToKey()
	return usets.Contains(key), nil
}
