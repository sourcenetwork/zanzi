// Package simple provides a simplified implementation of authorizers methods
//
// The primary implementation is done through the Expand call.
// Expand call is used by the Check call.
// Check is used by ReverseLookup
package simple

import (
	"github.com/sourcenetwork/source-zanzibar/authorizer"
)

func NewChecker() authorizer.Checker {
	expander := NewExpander()
	return CheckerFromExpander(expander)
}

func NewExpander() authorizer.Expander {
	return &expander{}
}

func NewReverser() authorizer.Reverser {
	checker := NewChecker()
	return ReverserFromChecker(checker)
}
