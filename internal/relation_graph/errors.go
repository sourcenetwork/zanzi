package relation_graph

import "github.com/sourcenetwork/zanzi/pkg/errors"

var ErrWildcardGoal = errors.Wrap("invalid goal: goal target cannot be wildcard, use an userset instead", errors.BadInput)
