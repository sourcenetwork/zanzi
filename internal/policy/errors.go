package policy

import "github.com/sourcenetwork/zanzi/pkg/errors"

var (
	ErrInvalidRelationship = errors.New("invalid relationship", errors.BadInput)
	ErrDuplicateDefinition = errors.New("duplicate definition", errors.BadInput)
	ErrRelExpTree          = errors.New("relation expression tree", errors.BadInput)

	ErrPolicyNotFound = errors.New("policy not found", errors.NotFound)
)
