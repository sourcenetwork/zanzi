package policy

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/zanzi/pkg/domain"
)

// operationResult models the output of applying a
// PolicyMutationOperation
type operationResult struct {
	RelationshipsRemoved uint64
}

// PolicyMutationOperation models an individual operation
// which when applied, transforms the state of a corresponding
// policy.
type PolicyMutationOperation interface {
	// Execute applies the mutation to the Policy scope.
	// The mutation changes the set of relationships
	// defined within the given policy
	Execute(ctx context.Context, repo Repository, policyId string) (operationResult, error)
}

// AddRelationOp models the actions that
// must be applied after adding a relation to a policy
type AddRelationOp struct {
	Resource string
	Relation string
}

func (o *AddRelationOp) Execute(ctx context.Context, repo Repository, polId string) (operationResult, error) {
	return operationResult{}, nil
}

// RemoveRelationOp models the actions that
// must be applied after adding a relation to a policy
type RemoveRelationOp struct {
	Resource string
	Relation string
}

// Execute removes all relationships for the given resource
// which match the removed relation
func (o *RemoveRelationOp) Execute(ctx context.Context, repo Repository, polId string) (operationResult, error) {
	builder := domain.SelectorBuilder{}

	// remove relationships where the relationship's object is from the given resource
	// and the relation is the given relation
	selector := builder.WithResource(o.Resource).WithRelation(o.Relation).AnySubject().Build()
	count1, err := repo.DeleteRelationships(ctx, polId, &selector)
	if err != nil {
		return operationResult{}, err
	}

	// now remove relationships whose subject are subject sets of the removed relation
	selector = builder.AnyObject().AnyRelation().WithSubjectGroup(o.Resource, o.Relation).Build()
	count2, err := repo.DeleteRelationships(ctx, polId, &selector)
	if err != nil {
		return operationResult{}, err
	}

	return operationResult{
		RelationshipsRemoved: count1 + count2,
	}, nil
}

// AddResourceOperation models the actions that
// must be applied after adding a relation to a policy
type AddResourceOperation struct {
	Resource string
}

func (o *AddResourceOperation) Execute(ctx context.Context, repo Repository, polId string) (operationResult, error) {
	return operationResult{}, nil
}

// RemoveResourceOp models the actions that
// must be applied after adding a relation to a policy
type RemoveResourceOp struct {
	Resource string
}

// Execute removes all relationships whose object
// belonged to the removed resource
func (o *RemoveResourceOp) Execute(ctx context.Context, repo Repository, polId string) (operationResult, error) {
	builder := domain.SelectorBuilder{}

	// remove all relationships with objects from the resource being removed
	builder.WithResource(o.Resource).AnyRelation().AnySubject()
	selector := builder.Build()
	count1, err := repo.DeleteRelationships(ctx, polId, &selector)
	if err != nil {
		return operationResult{}, err
	}

	// removes all relationships whose subjects are any object in the
	// removed resource
	builder.AnyObject().AnyRelation().WithSubjectResource(o.Resource)
	selector = builder.Build()
	count2, err := repo.DeleteRelationships(ctx, polId, &selector)
	if err != nil {
		return operationResult{}, err
	}

	return operationResult{
		RelationshipsRemoved: count1 + count2,
	}, nil
}

// MutationRelationOp models the actions that
// must be applied after adding a relation to a policy
type MutateRelationOp struct {
	Resource string
	Relation string
	Record   domain.Relation
}

func (o *MutateRelationOp) Execute(ctx context.Context, repo Repository, polId string) (operationResult, error) {
	return operationResult{}, nil
}

// mapOperationToComparable maps an operation to an integer
// which models the order in which the mutation operation
// needs to be applied.
//
// Lower numbers means that the operation has higher precedence.
// ie. sorting the operations produces a slice of operations
// in order which they need to be applied
func mapOperationToComparable(op PolicyMutationOperation) int {
	switch op.(type) {
	case *RemoveResourceOp:
		// remove resource should happen before any other ops
		return 0
	case *RemoveRelationOp:
		// followed by remove relation
		return 1
	case *AddRelationOp:
		return 2
	case *AddResourceOperation:
		return 2
	case *MutateRelationOp:
		return 3
	default:
		panic(fmt.Sprintf("invalid op: %v", op))
	}
}
