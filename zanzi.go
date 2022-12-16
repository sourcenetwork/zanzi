package source_zanzibar

import (
    "context"
    "fmt"

    "google.golang.org/protobuf/proto"

    "github.com/sourcenetwork/source-zanzibar/types"
    rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
    o "github.com/sourcenetwork/source-zanzibar/pkg/option"
)

// Exposes Zanzibar-like functionality to client applications
type Service[T proto.Message] struct {
    relationshipService types.RelationshipService[T]
    policyService types.PolicyService
    authorizer types.Authorizer
}

func (s *Service[T]) GetAuthorizer() types.Authorizer {
    return s.authorizer
}

func (s *Service[T]) GetPolicyService() types.PolicyService {
    return s.policyService
}

func (s *Service[T]) GetRelationshipService() types.RelationshipService[T] {
    return s.relationshipService
}


// authorizer implements the Authorizer interface by wrapping a relation graph
type authorizer[T proto.Message] struct {
    rg rg.RelationGraph
    builder tuple.TupleBuilder[T]
    mapper treeMapper
}

func (a *authorizer[T]) Check(policyId string, obj types.Entity, relation string, actor types.Entity) (bool, error) {
    ctx := context.Background()
    src := a.builder.RelSource(obj.Namespace, obj.Id, relation)
    dst := a.builder.ActorWithNamespace(actor.Namespace, actor.Id)
    reachable, err := a.rg.IsReachable(ctx, policyId, src, dst)
    if err != nil {
        return false, fmt.Errorf("check failed: %w", err)
    }

    return reachable, nil
}

func (a *authorizer[T]) Reverse(policyId string, actor types.Entity) ([]types.EntityRelPair, error) {
    return nil, nil
}

func (a *authorizer[T]) Expand(policyId string, obj types.Entity, relation string) (types.ExpandTree, error) {
    ctx := context.Background()
    src := a.builder.RelSource(obj.Namespace, obj.Id, relation)

    tree, err := a.rg.Walk(ctx, policyId, src)
    if err != nil {
        return types.ExpandTree{}, fmt.Errorf("expand failed: %w", err)
    }

    return a.mapper.ToExpandTree(&tree), nil
}


// relService implements the RelationService interface by wrapping a TupleStore
type relService[T proto.Message] struct {
    tStore tuple.TupleStore[T]
    mapper tupleMapper[T]
}

func (s *relService[T]) Set(rel types.Relationship, data T) error {
    t := s.mapper.FromRelationship(rel)
    t.SetData(data)
    return s.tStore.SetTuple(t)
}

func (s *relService[T]) Delete(rel types.Relationship) error {
    t := s.mapper.FromRelationship(rel)
    return s.tStore.DeleteTuple(t.Partition, t.Source, t.Dest)
}

func (s *relService[T]) Get(rel types.Relationship) (o.Option[types.Record[T]], error) {
    t := s.mapper.FromRelationship(rel)
    opt, err := s.tStore.GetTuple(t.Partition, t.Source, t.Dest)
    if err != nil || opt.IsEmpty() {
        return o.None[types.Record[T]](), err
    }

    rec := s.mapper.ToRecord(t)
    return o.Some[types.Record[T]](rec), nil
}


// policyService wraps a PolicyStore in order to implement PolicyService
type policyService struct {
    pStore policy.PolicyStore
    mapper policyMapper
}

func (s *policyService) Set(p types.Policy) error {
    mapped := s.mapper.ToInternal(p)
    return s.pStore.SetPolicy(&mapped)
}

func (s *policyService) Get(id string) (o.Option[types.Policy], error) {
    polOpt, err := s.pStore.GetPolicy(id)
    if err != nil || polOpt.IsEmpty() {
        return o.None[types.Policy](), err
    }

    pol := s.mapper.FromInternal(polOpt.Value())

    return o.Some[types.Policy](pol), nil
}

func (s *policyService) Delete(id string) error {
    return s.pStore.DeletePolicy(id)
}


var _ types.PolicyService = (*policyService)(nil)
var _ types.Authorizer = (*authorizer[proto.Message])(nil)
var _ types.RelationshipService[proto.Message] = (*relService[proto.Message])(nil)
