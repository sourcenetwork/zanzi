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
func NewTupleRepo(ts []model.Tuple) repository.TupleRepository {
	tuples := make(map[string]model.TupleRecord)
	usersets := make(map[string][]model.TupleRecord)

	for _, t := range ts {
		tuple := t
		record := model.TupleRecord{
			Tuple: &tuple,
		}
		tuples[tuple.String()] = record
		usersets[tuple.ObjectRel.String()] = append(usersets[tuple.ObjectRel.String()], record)
	}

	return &tupleRepo{
		tuples:   tuples,
		usersets: usersets,
	}
}

// Return an instance of NamespaceRepository from a slice of namespaces
//
// Note that the repository is not persisted and only lookup operations are implemented
func NewNamespaceRepo(ns []model.Namespace) repository.NamespaceRepository {
	namespaces := make(map[string]model.Namespace)
	for _, namespace := range ns {
		namespaces[namespace.Name] = namespace
	}

	return &namespaceRepo{
		namespaces: namespaces,
	}
}

type tupleRepo struct {
	tuples   map[string]model.TupleRecord
	usersets map[string][]model.TupleRecord
}

func (r *tupleRepo) SetTuple(tuple model.Tuple) (model.TupleRecord, error) {
	return model.TupleRecord{}, fmt.Errorf("SetTuple not implemented")
}

func (r *tupleRepo) GetTuple(tuple model.Tuple) (model.TupleRecord, error) {
	record, ok := r.tuples[tuple.String()]
	if !ok {
		return model.TupleRecord{}, repository.NewEntityNotFound("Tuple", tuple)
	}
	return record, nil
}

func (r *tupleRepo) GetRelatedUsersets(userset model.Userset) ([]model.TupleRecord, error) {
	usets, ok := r.usersets[userset.String()]
	if !ok {
		return nil, repository.NewEntityNotFound("Userset", userset)
	}
	return usets, nil
}

func (r *tupleRepo) GetIncomingUsersets(userset model.Userset) ([]model.TupleRecord, error) {
	return nil, fmt.Errorf("GetIncomingUsersets not implemented")
}

func (r *tupleRepo) RemoveTuple(tuple model.Tuple) error {
	return fmt.Errorf("RemoveTuple not implemented")
}

func (r *tupleRepo) GetTuplesFromRelationAndUserObject(relation string, objNamespace string, objectId string) ([]model.TupleRecord, error) {
	return nil, nil
}

type namespaceRepo struct {
	namespaces map[string]model.Namespace
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
	return nil, fmt.Errorf("GetReferrers not implemented")
}
