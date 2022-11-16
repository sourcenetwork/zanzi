package simple

import (
	"context"

	"github.com/sourcenetwork/source-zanzibar/authorizer"
	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/tree"
)

var _ authorizer.Checker = (*checker)(nil)

type checker struct {
	expander authorizer.Expander
}

// Simple Check implementation leverages a full expand in order to verify whether
// userset is included in final expand tree
func (c *checker) Check(ctx context.Context, objRel model.AuthNode, user model.AuthNode) (bool, error) {
	root, err := c.expander.Expand(ctx, objRel)
	if err != nil {
		// wrap
		return false, err
	}

	key := user.ToKey()
	usets := tree.Eval(&root)
	return usets.Contains(key), nil
}

// Return an instance of Checker from an Expander
// Checker will perform an expand call and verify whether the final
// expand tree contains the desired user
func CheckerFromExpander(expander authorizer.Expander) authorizer.Checker {
	return &checker{
		expander: expander,
	}
}
