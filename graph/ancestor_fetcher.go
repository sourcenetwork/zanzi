package graph

import (
	"context"
	"fmt"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/repository"
	"github.com/sourcenetwork/source-zanzibar/utils"
)

// Ancestor Fetcher is used to fetch a node's stored and logical ancestors
type AncestorFetcher[T model.Relationship] struct {
	logicalAncestors []model.AuthNode
	nsRepo           repository.NamespaceRepository
	tupleRepo        repository.TupleRepository[T]
}

func NewFetcher[T model.Relationship](nsRepo repository.NamespaceRepository, tupleRepo repository.TupleRepository[T]) AncestorFetcher[T] {
	return AncestorFetcher[T]{
		nsRepo:    nsRepo,
		tupleRepo: tupleRepo,
	}
}

// Return all ancestors nodes of uset
// Includes direct ancestors and logical ones (obtained by reverting rewrite rules)
func (f *AncestorFetcher[T]) FetchAll(ctx context.Context, uset model.AuthNode) ([]model.AuthNode, error) {
	directAncestors, err := f.FetchAncestors(ctx, uset)
	if err != nil {
		return nil, err
	}

	logicalAncestors, err := f.FetchLogicalAncestors(ctx, uset)
	if err != nil {
		return nil, err
	}

	return append(directAncestors, logicalAncestors...), nil
}

// fetchLogicalAncestors return Relationships which are "logical ancestors" of uset.
// Logical Ancestors are edges reachable through userset rewrite rules.
func (f *AncestorFetcher[T]) FetchLogicalAncestors(ctx context.Context, uset model.AuthNode) ([]model.AuthNode, error) {
	f.logicalAncestors = nil

	referringRels, err := f.nsRepo.GetReferrers(uset.Namespace, uset.Relation)
	if err != nil {
		err = fmt.Errorf("failed FetchLogicalAncestors for %v: %v", uset, err)
		return nil, err
	}

	for _, relation := range referringRels {
		err := f.buildAncestorsFromRel(ctx, uset, relation)
		if err != nil {
			err = fmt.Errorf("failed building ancestors for %v: %v", uset, err)
			return nil, err
		}
	}

	return f.logicalAncestors, nil
}

// Fetch all directly accessible Ancestors of uset
func (f *AncestorFetcher[T]) FetchAncestors(ctx context.Context, uset model.AuthNode) ([]model.AuthNode, error) {
	records, err := f.tupleRepo.GetIncomingAuthNodes(uset)
	if err != nil {
		err = fmt.Errorf("fetch ancestors failed for %v: %v", uset, err)
		return nil, err
	}
	return utils.MapSlice(records, recordToObjRel), nil
}

// Extracts all rules from relation which reference uset.Relation
// for each matching rule, build possible ancestors and
// append results to f.logicalAncestors
func (f *AncestorFetcher[T]) buildAncestorsFromRel(ctx context.Context, uset model.AuthNode, relation model.Relation) error {
	rules := f.getReferrers(relation, uset)

	for _, rule := range rules {
		err := f.buildAncestorsFromRule(ctx, uset, relation, rule)
		if err != nil {
			fmt.Errorf("failed building ancestors for relation %v: %v", relation.Name, err)
			return err
		}
	}
	return nil
}

// Extract all `Rule`s from a `Relation` userset expression tree
// which references the relation specificied in uset
// eg:
// let uset.Relation = "Owner"
// let relation exp tree contain a rule of type CU(Owner)
// getReferrers would return only the Owner CU rule
func (f *AncestorFetcher[T]) getReferrers(relation model.Relation, uset model.AuthNode) []model.Rule {
	leaves := relation.Rewrite.ExpressionTree.GetLeaves()

	referencedRel := uset.Relation

	var rules = make([]model.Rule, 0, len(leaves))
	for _, leaf := range leaves {
		switch rule := leaf.Rule.Rule.(type) {
		case *model.Rule_This:
			// "This" rules never references any relation
			continue
		case *model.Rule_TupleToUserset:
			// Tuple To AuthNode references a relation through the
			// ComputedAuthNode property
			ttu := rule.TupleToUserset
			if ttu.ComputedAuthNodeRelation == referencedRel {
				rules = append(rules, *leaf.Rule)
			}
		case *model.Rule_ComputedAuthNode:
			// Computed AuthNode directly references a relation
			cu := rule.ComputedAuthNode
			if cu.Relation == referencedRel {
				rules = append(rules, *leaf.Rule)
			}
		default:
			panic("uknown rule type")
		}
	}
	return rules
}

// Append potential logical ancestors of uset for a relation, rule pair
// Assume rules were previously filtered to only contain TTU and CU rules
// which applies to the given uset
func (f *AncestorFetcher[T]) buildAncestorsFromRule(ctx context.Context, uset model.AuthNode, relation model.Relation, rule model.Rule) error {
	switch rule := rule.Rule.(type) {
	case *model.Rule_TupleToUserset:
		ttu := rule.TupleToUserset

		usets, err := f.revertTTU(ctx, uset, relation.Name, *ttu)
		if err != nil {
			return err
		}

		f.logicalAncestors = append(f.logicalAncestors, usets...)

	case *model.Rule_ComputedAuthNode:
		uset := f.revertCU(uset, relation.Name)
		f.logicalAncestors = append(f.logicalAncestors, uset)

	default:
		panic("invalid rule type")
	}
	return nil
}

// Apply a CU rule backwards and return the original node
// Returned node is not guaranteed to exist
func (f *AncestorFetcher) revertCU(uset model.AuthNode, referer string) model.AuthNode {
	return model.AuthNode{
		Namespace: uset.Namespace,
		ObjectId:  uset.ObjectId,
		Relation:  referer,
                Type: uset.Type,
	}
}

// Apply a TTU rule backwards and return nodes that may reach uset
// Returned nodes are not guaranteed to exist
func (f *AncestorFetcher) revertTTU(ctx context.Context, uset model.AuthNode, referer string, ttu model.TupleToUserset) ([]model.AuthNode, error) {
	tuples, err := f.tupleRepo.GetTuplesFromRelationAndUserObject(ttu.TuplesetRelation, uset.Namespace, uset.ObjectId)
	if err != nil {
		err = fmt.Errorf("failed reverting TTU rule %v: %v", ttu, err)
		return nil, err
	}

	usets := utils.MapSlice(tuples, recordToObjRel)

	for i, _ := range usets {
		usets[i].Relation = referer
	}

	return usets, nil
}

// Map the object, rel userset in a tuple record to a userset
func recordToObjRel[T model.Relationship](record T) model.AuthNode {
	return record.GetObject()
}
