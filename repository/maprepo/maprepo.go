// package maperpo implements app's repository interfaces using Go's `map` type
package maprepo

import (
	"fmt"

	"github.com/sourcenetwork/source-zanzibar/model"
	"github.com/sourcenetwork/source-zanzibar/repository"
)

var (
	_ repository.TupleRepository     = (*tupleRepo)(nil)
	_ repository.NamespaceRepository = (*namespaceRepo)(nil)
)

// Return an instance of TupleRepository from a slice of tuples
//
// Note that the repository is not persisted and only lookup operations are implemented
func NewTupleRepo(tuples ...model.Tuple) repository.TupleRepository {
	records := make(map[string]model.Relationship)
	usersets := make(map[model.KeyableUset][]model.Relationship)
	reverse := make(map[model.KeyableUset][]model.Relationship)

	for _, t := range tuples {
		tuple := t
		record := model.Relationship{
			Tuple: &tuple,
		}
		records[tuple.String()] = record

		key := tuple.ObjectRel.ToKey()
		usersets[key] = append(usersets[key], record)

		revKey := tuple.User.Userset.ToKey()
		reverse[revKey] = append(reverse[revKey], record)
	}

	return &tupleRepo{
		tuples:   records,
		usersets: usersets,
		reverse:  reverse,
	}
}

// Return an instance of NamespaceRepository from a slice of namespaces
//
// Note that the repository is not persisted and only lookup operations are implemented
func NewNamespaceRepo(ns ...model.Namespace) repository.NamespaceRepository {
	namespaces := make(map[string]model.Namespace)
	referrers := make(map[string][]model.Relation)

	for _, namespace := range ns {
		namespaces[namespace.Name] = namespace

		for _, rel := range namespace.Relations {

			tree := rel.Rewrite.ExpressionTree
			leaves := tree.GetLeaves()
			for _, leaf := range leaves {
				switch rule := leaf.Rule.Rule.(type) {
				case *model.Rule_This:
					continue
				case *model.Rule_TupleToUserset:
					ttu := rule.TupleToUserset
					key := namespace.Name + ttu.ComputedUsersetRelation
					referrers[key] = append(referrers[key], *rel)

				case *model.Rule_ComputedUserset:
					cu := rule.ComputedUserset
					key := namespace.Name + cu.Relation
					referrers[key] = append(referrers[key], *rel)

				default:
					panic("uknown rule type")
				}
			}
		}
	}

	return &namespaceRepo{
		namespaces: namespaces,
		referrers:  referrers,
	}
}

type tupleRepo struct {
	tuples   map[string]model.Relationship
	usersets map[model.KeyableUset][]model.Relationship
	reverse  map[model.KeyableUset][]model.Relationship
}

func (r *tupleRepo) SetTuple(tuple model.Tuple) (model.Relationship, error) {
	return model.Relationship{}, fmt.Errorf("SetTuple not implemented")
}

func (r *tupleRepo) GetTuple(tuple model.Tuple) (model.Relationship, error) {
	record, ok := r.tuples[tuple.String()]
	if !ok {
		return model.Relationship{}, repository.NewEntityNotFound("Tuple", tuple)
	}
	return record, nil
}

func (r *tupleRepo) GetRelatedUsersets(userset model.AuthNode) ([]model.Relationship, error) {
	usets, ok := r.usersets[userset.ToKey()]
	if !ok {
		return nil, repository.NewEntityNotFound("Userset", userset)
	}
	return usets, nil
}

func (r *tupleRepo) GetIncomingUsersets(userset model.AuthNode) ([]model.Relationship, error) {
	records, _ := r.reverse[userset.ToKey()]
	return records, nil
}

func (r *tupleRepo) RemoveTuple(tuple model.Tuple) error {
	return fmt.Errorf("RemoveTuple not implemented")
}

func (r *tupleRepo) GetTuplesFromRelationAndUserObject(relation string, objNamespace string, objectId string) ([]model.Relationship, error) {
	results := make([]model.Relationship, 0, 10)
	for _, record := range r.tuples {
		tuple := record.Tuple
		userUset := tuple.User.Userset
		if tuple.ObjectRel.Relation == relation && userUset.ObjectId == objectId && userUset.Namespace == objNamespace {
			results = append(results, record)
		}
	}
	return results, nil
}

type namespaceRepo struct {
	namespaces map[string]model.Namespace
	referrers  map[string][]model.Relation
}

func (r *namespaceRepo) GetNamespace(namespace string) (model.Namespace, error) {
	n, ok := r.namespaces[namespace]
	if !ok {
		return model.Namespace{}, repository.NewEntityNotFound("Namespace", namespace)
	}
	return n, nil
}

func (r *namespaceRepo) SetNamespace(namespace model.Namespace) (model.Namespace, error) {
	return model.Namespace{}, fmt.Errorf("SetNamespace not implemented")
}

func (r *namespaceRepo) RemoveNamespace(namespace string) error {
	return fmt.Errorf("RemoveNamespace not implemented")
}

func (r *namespaceRepo) GetRelation(namespace, relation string) (model.Relation, error) {
	n, err := r.GetNamespace(namespace)
	if err != nil {
		return model.Relation{}, err
	}

	for _, r := range n.Relations {
		if r.Name == relation {
			return *r, nil
		}
	}

	return model.Relation{}, repository.NewEntityNotFound("Relation", namespace, relation)
}

func (r *namespaceRepo) GetReferrers(namespace, relation string) ([]model.Relation, error) {
	refs, _ := r.referrers[namespace+relation]
	return refs, nil
}
