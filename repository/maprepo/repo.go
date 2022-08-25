package maprepo

import (
    "fmt"

    "github.com/sourcenetwork/source-zanzibar/model"
    "github.com/sourcenetwork/source-zanzibar/repository"
)

var (
    _ repository.TupleRepository = (*Tuples)(nil)
    _ repository.NamespaceRepository = (*Namespaces)(nil)
)


type Tuples struct {
    tuples map[string]model.TupleRecord
    usersets map[string][]model.TupleRecord
}

type Namespaces struct {
    namespaces map[string]model.Namespace
}

func NewTupleRepo(ts []model.Tuple) repository.TupleRepository {
    tuples := make(map[string]model.TupleRecord)
    usersets := make(map[string][]model.TupleRecord)
    
    for _, t := range ts {
        tuple := t
        record := model.TupleRecord {
            Tuple: &tuple,
        }
        tuples[tuple.String()] = record
        usersets[tuple.ObjectRel.String()] = append(usersets[tuple.ObjectRel.String()], record)
    }

    return &Tuples {
        tuples: tuples,
        usersets: usersets,
    }
}

func NewNamespaceRepo(ns []model.Namespace) repository.NamespaceRepository {
    namespaces := make(map[string]model.Namespace)
    for _, namespace := range ns {
        namespaces[namespace.Name] = namespace
    }

    return &Namespaces {
        namespaces: namespaces,
    }
}


func (r *Tuples) SetTuple(tuple model.Tuple) error {
    return fmt.Errorf("SetTuple not implemented")
}

func (r *Tuples) GetTuple(tuple model.Tuple) (model.TupleRecord, error){
    record, ok := r.tuples[tuple.String()]
    if !ok {
        return model.TupleRecord{}, fmt.Errorf("tuple %#v not found", tuple)
    }
    return record, nil
}

func (r *Tuples) GetRelatedUsersets(userset model.Userset) ([]model.TupleRecord, error){
    usets, ok := r.usersets[userset.String()]
    if !ok {
        return nil, fmt.Errorf("Userset %#v not found", userset)

    }
    return usets, nil
}

func (r *Tuples) GetParentTuples(userset model.Userset) ([]model.TupleRecord, error){ 
    return nil, fmt.Errorf("GetParentTuples not implemented")
}

func (r *Tuples) RemoveTuple(tuple model.Tuple) error{ 
    return fmt.Errorf("RemoveTuple not implemented")
}

func (r *Namespaces) GetNamespace(namespace string) (model.Namespace, error) { 
    n, ok := r.namespaces[namespace]
    if !ok {
        return model.Namespace{}, fmt.Errorf("Namespace %#v not found", namespace)
    }
    return n, nil
}

func (r *Namespaces) SetNamespace(namespace model.Namespace) error { 
    return fmt.Errorf("SetNamespace not implemented")
}

func (r *Namespaces) RemoveNamespace(namespace string) error { 
    return fmt.Errorf("RemoveNamespace not implemented")
}

func (r *Namespaces) GetRelation(namespace, relation string) (model.Relation, error) {
    n, err := r.GetNamespace(namespace)
    if err != nil {
        return model.Relation{}, err
    }

    for _, r := range n.Relations {
        if r.Name == relation {
            return *r, nil
        }
    }

    return model.Relation{}, fmt.Errorf("Relation %#v not found", relation)
}
