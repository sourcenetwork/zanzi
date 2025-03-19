package policy

import (
	"context"

	"github.com/sourcenetwork/raccoondb/v2/utils"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/sourcenetwork/zanzi/pkg/domain"
)

// policyComparator models a component which given two policies,
// computes the difference between them as Operations to be executed
// to transform the old policy into the new policy
type policyComparator struct{}

// Compare computes the difference in resources and relations between
// old and new policies.
// Returns slice of mutation operation which will transform old into new
func (c *policyComparator) Compare(ctx context.Context, old, new *domain.Policy) (ops []PolicyMutationOperation) {
	oldResources := sets.New(old.GetResourcesNames()...)
	newResources := sets.New(new.GetResourcesNames()...)
	removed := oldResources.Difference(newResources)

	for _, name := range removed.UnsortedList() {
		op := &RemoveResourceOp{
			Resource: name,
		}
		ops = append(ops, op)
	}

	for _, resource := range new.Resources {
		oldResource := old.GetResourceByName(resource.Name)
		if oldResource == nil {
			op := &AddResourceOperation{
				Resource: resource.Name,
			}
			ops = append(ops, op)
			continue
		}

		currentRelations := sets.New(resource.GetRelationsNames()...)
		oldRelations := sets.New(oldResource.GetRelationsNames()...)
		removed := oldRelations.Difference(currentRelations)
		added := currentRelations.Difference(oldRelations)
		maybeModified := oldRelations.Intersection(currentRelations)

		for _, relation := range removed.UnsortedList() {
			op := &RemoveRelationOp{
				Resource: resource.Name,
				Relation: relation,
			}
			ops = append(ops, op)
		}

		for _, relation := range added.UnsortedList() {
			op := &AddRelationOp{
				Resource: resource.Name,
				Relation: relation,
			}
			ops = append(ops, op)
		}

		for _, relName := range maybeModified.UnsortedList() {
			op := &MutateRelationOp{
				Resource: resource.Name,
				Relation: relName,
			}
			ops = append(ops, op)
		}
	}

	sortable := utils.FromExtractor(ops, mapOperationToComparable)
	sortable.SortInPlace()
	return ops
}
