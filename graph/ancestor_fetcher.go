package graph

import (
	"context"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/utils"
)

// Ancestor Fetcher is used to fetch a node's stored and logical ancestors
type AncestorFetcher struct {
	logicalAncestors []model.Userset
}

// Return all ancestors nodes of uset
// Includes direct ancestors and logical ones (obtained by reverting rewrite rules)
func (f *AncestorFetcher) FetchAll(ctx context.Context, uset model.Userset) ([]model.Userset, error) {
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

// fetchLogicalAncestors return TupleRecords which are "logical ancestors" of uset.
// Logical Ancestors are edges reachable through userset rewrite rules.
func (f *AncestorFetcher) FetchLogicalAncestors(ctx context.Context, uset model.Userset) ([]model.Userset, error) {
	f.logicalAncestors = nil

	nsRepo := utils.GetNamespaceRepo(ctx)

	referringRels, err := nsRepo.GetReferrers(uset.Namespace, uset.Relation)
	if err != nil {
		return nil, err
	}

	for _, relation := range referringRels {
		err := f.buildAncestorsFromRel(ctx, uset, relation)
		if err != nil {
			return nil, err
		}
	}

	return f.logicalAncestors, nil
}

// Fetch all directly accessible Ancestors of uset
func (f *AncestorFetcher) FetchAncestors(ctx context.Context, uset model.Userset) ([]model.Userset, error) {
	repo := utils.GetTupleRepo(ctx)
	records, err := repo.GetIncomingUsersets(uset)
	if err != nil {
		return nil, err
	}
	return utils.MapSlice(records, recordToObjRel), nil
}

// Extracts all rules from relation which reference uset.Relation
// for each matching rule, build possible ancestors and
// append results to f.logicalAncestors
func (f *AncestorFetcher) buildAncestorsFromRel(ctx context.Context, uset model.Userset, relation model.Relation) error {
	rules := f.getReferrers(relation, uset)

	for _, rule := range rules {
		err := f.buildAncestorsFromRule(ctx, uset, relation, rule)
		if err != nil {
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
func (f *AncestorFetcher) getReferrers(relation model.Relation, uset model.Userset) []model.Rule {
	leaves := relation.Rewrite.ExpressionTree.GetLeaves()

	referencedRel := uset.Relation

	var rules = make([]model.Rule, 0, len(leaves))
	for _, leaf := range leaves {
		switch rule := leaf.Rule.Rule.(type) {
		case *model.Rule_This:
			// "This" rules never references any relation
			continue
		case *model.Rule_TupleToUserset:
			// Tuple To Userset references a relation through the
			// ComputedUserset property
			ttu := rule.TupleToUserset
			if ttu.ComputedUsersetRelation == referencedRel {
				rules = append(rules, *leaf.Rule)
			}
		case *model.Rule_ComputedUserset:
			// Computed Userset directly references a relation
			cu := rule.ComputedUserset
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
func (f *AncestorFetcher) buildAncestorsFromRule(ctx context.Context, uset model.Userset, relation model.Relation, rule model.Rule) error {
	switch rule := rule.Rule.(type) {
	case *model.Rule_TupleToUserset:
		ttu := rule.TupleToUserset

		usets, err := revertTTU(ctx, uset, relation.Name, *ttu)
		if err != nil {
			return err
		}

		f.logicalAncestors = append(f.logicalAncestors, usets...)

	case *model.Rule_ComputedUserset:
		uset := revertCU(uset, relation.Name)
		f.logicalAncestors = append(f.logicalAncestors, uset)

	default:
		panic("invalid rule type")
	}
	return nil
}

// Apply a CU rule backwards and return the original node
// Returned node is not guaranteed to exist
func revertCU(uset model.Userset, referer string) model.Userset {
	return model.Userset{
		Namespace: uset.Namespace,
		ObjectId:  uset.ObjectId,
		Relation:  referer,
	}
}

// Apply a TTU rule backwards and return nodes that may reach uset
// Returned nodes are not guaranteed to exist
func revertTTU(ctx context.Context, uset model.Userset, referer string, ttu model.TupleToUserset) ([]model.Userset, error) {
	repo := utils.GetTupleRepo(ctx)

	tuples, err := repo.GetTuplesFromRelationAndUserObject(ttu.TuplesetRelation, uset.Namespace, uset.ObjectId)
	if err != nil {
		return nil, err
	}

	usets := utils.MapSlice(tuples, recordToObjRel)

	for _, u := range usets {
		u.Relation = referer
	}

	return usets, nil
}

// Map the object, rel userset in a tuple record to a userset
func recordToObjRel(record model.TupleRecord) model.Userset {
	return model.Userset{
		Namespace: record.Tuple.ObjectRel.Namespace,
		ObjectId:  record.Tuple.ObjectRel.ObjectId,
		Relation:  record.Tuple.ObjectRel.Relation,
	}
}
