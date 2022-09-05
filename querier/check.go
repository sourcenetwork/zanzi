package querier

import (
    "context"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/tree"
)

// Simple Check implementation leverages a full expand in order to verify whether 
// userset is included in final expand tree
func Check(ctx context.Context, objRel model.Userset, user model.Userset) (bool, error) {
	root, err := Expand(ctx, objRel)
	if err != nil {
		// wrap
		return false, err
	}

	key := user.ToKey()
	usets := tree.Eval(root)
	return usets.Contains(key), nil
}
