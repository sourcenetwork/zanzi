package source_zanzibar

import (
    "ctx"

    "google.golang.org/protobuf/proto"
    cosmos "github.com/cosmos/cosmos-sdk/store/types"

    "github.com/sourcenetwork/source-zanzibar/types"
    rg "github.com/sourcenetwork/source-zanzibar/internal/domain/relation_graph"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/tuple"
    "github.com/sourcenetwork/source-zanzibar/internal/domain/policy"
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

// Build a Zanzibar service from a cosmos-sdk kv store
func InitFromCosmosKV[T proto.Message](
    policyStore cosmos.KVStore,
    policyPrefix string,
    relationStore cosmos.KVStore,
    relationPrefix string) Service[T] {
    return Service {
    }
}

// authorizer implements the Authorizer interface by wrapping a relation graph
type authorizer struct {
    rg rg.RelationGraph
}

func (a *authorizer) Check(policyId string, obj types.Id, relation string, actor types.Id) (bool, error) {
    ctx := context.Background()
    source := tuple.TupleNode {
        Namespace: obj.Namespace,
        Id: obj.Id,
        Relation: relation,
    }
    reachable, err := rg.IsReachable(ctx, policyId, source)
    if err != nil {
        return false, fmt.Errorf("check failed: %w", err)
    }

    return reachable, nil
}


func (a* authorizer) Reverse(policyId string, actor types.Id, relation string) ([]types.EntityRelPair, error) {
    ctx := context.Background()
    source := tuple.TupleNode {
        Namespace: obj.Namespace,
        Id: obj.Id,
        Relation: relation,
    }
    ancestors, err := a.rg.GetAncestors(ctx, policyId, source)
    if err != nil {
        return nil, fmt.Errorf("reverse failed: %w", err)
    }

    // TODO define output data type
    return ancestors, nil
}


func (a* authorizer) Expand(obj types.Id, relation string) (types.ExpandTree, error) {
    /*
    ctx := context.Background()
    source := tuple.TupleNode {
        Namespace: obj.Namespace,
        Id: obj.Id,
        Relation: relation,
    }
    tree, err := a.rg.Walk(ctx, policyId, source)
    if err != nil {
        return nil, fmt.Errorf("expand failed: %w", err)
    }
    */
    return nil, nil
}

// relService implements the RelationService interface by wrapping a TupleStore
type relService[T proto.Message] struct {
    tStore tuple.TupleStore[T]
}

func (s *relService[T]) Set(rel Relationship, data T) error {
    tuple := toTuple[T](rel, data)
    return s.tStore.SetTuple(tuple)
}

func (s *relService[T]) Delete(rel Relationship) error {
    t := toTuple[T](rel, nil)
    return s.tStore.DeleteTuple(rel.PolicyId, t.Source, t.Dest)
}

func (s *relService[T]) Get(rel Relationship) (o.Option[Record[T]], error) {
    t := toTuple[T](rel)
    opt, err := s.tStore.GetTuple(t.Partition, t.Source, t.Dest)
    if err != nil {
        return o.None[Record[T]](), err
    }
    if opt.IsEmpty() {
        return o.None[Record[T]](), nil
    }
    
    // build record
}

func (s *relService[T]) GetRelationships(entity Entity, relation string) ([]Relationship, error) {

    // TODO
}

func toTuple[T proto.Message](rel types.Relationship) tuple.Tuple[T] {
    // TODO
}

// TODO implement policy service

type policyService struct {
    pStore policy.PolicyStore
}

func (s *policyService) Set(policy Policy) error {}

func (s *policyService) Get(policyId string) (o.Option[Policy], error) {}

func (s *policyService) Delete(policy Policy) error {}
