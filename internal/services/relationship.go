package services

import (
	_ "fmt"

	"github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
	o "github.com/sourcenetwork/source-zanzibar/pkg/option"
	"github.com/sourcenetwork/source-zanzibar/types"
)

var _ types.RelationshipService = (*relationshipService)(nil)

// Return a RelationshipService implementation by wrapping a TupleStore
func RelationshipServiceFromTupleStore(tStore tuple.TupleStore) types.RelationshipService {
	return &relationshipService{
		tStore: tStore,
	}
}

// relationshipService implements the RelationService interface by wrapping a TupleStore
type relationshipService struct {
	tStore tuple.TupleStore
	mapper RelationshipMapper
}

func (s *relationshipService) Set(rel types.Relationship) error {
	t := s.mapper.FromRelationship(rel)
	return s.tStore.SetTuple(t)
}

func (s *relationshipService) Delete(rel types.Relationship) error {
	t := s.mapper.FromRelationship(rel)
	return s.tStore.DeleteTuple(t.Partition, t.Source, t.Dest)
}

// thsi makes no sense
func (s *relationshipService) Get(rel types.Relationship) (o.Option[types.Relationship], error) {
	t := s.mapper.FromRelationship(rel)
	opt, err := s.tStore.GetTuple(t.Partition, t.Source, t.Dest)
	if err != nil || opt.IsEmpty() {
		return o.None[types.Relationship](), err
	}

	rec := s.mapper.ToRelationship(t)
	return o.Some[types.Relationship](rec), nil
}

func (s *relationshipService) Has(rel types.Relationship) (bool, error) {
	// TODO
	return false, nil
}
