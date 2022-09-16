// Package simple provides a simplified implementation of authorizers methods
//
// The primary implementation is done through the Expand call.
// Expand call is used by the Check call.
// Check is used by ReverseLookup
package simple

import (
	"github.com/sourcenetwork/source-zanzibar/authorizer"
	"github.com/sourcenetwork/source-zanzibar/repository"
)


func NewChecker(nsRepo repository.NamespaceRepository, tupleRepo repository.TupleRepository) authorizer.Checker {
	expander := NewExpander(nsRepo, tupleRepo)
	return CheckerFromExpander(expander)
}

func NewExpander(nsRepo repository.NamespaceRepository, tupleRepo repository.TupleRepository) authorizer.Expander {
	return &expander{
            tupleRepo: tupleRepo,
            nsRepo: nsRepo,
        }
}

func NewReverser(nsRepo repository.NamespaceRepository, tupleRepo repository.TupleRepository) authorizer.Reverser {
	checker := NewChecker(nsRepo, tupleRepo)
	return ReverserFromChecker(nsRepo, tupleRepo, checker)
}
