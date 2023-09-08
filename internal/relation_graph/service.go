package relation_graph

import (
    "fmt"
	"context"

	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

var _ api.RelationGraphServer = (*Service)(nil)

func NewService(repository NodeRepository, policyRepository policy.Repository, logger types.Logger) api.RelationGraphServer{
    return &Service{
        logger: logger,
        repository: repository,
        policyRepository: policyRepository,
    }
}

type Service struct {
    api.UnimplementedRelationGraphServer
    logger types.Logger
    repository NodeRepository
    policyRepository policy.Repository
}

func (s *Service) Check(
    ctx context.Context,
    req *api.CheckRequest) (*api.CheckResponse, error) {
        pol, err := s.policyRepository.GetPolicy(ctx, req.PolicyId)
        if err != nil {
            return nil, fmt.Errorf("check: %w: %w", api.ErrInternal, err)
        } 
        if pol == nil {
            return nil, fmt.Errorf("check: %w: policy %v: %w", api.ErrInvalidRequest, req.PolicyId, policy.ErrPolicyNotFound)
        }

        evaluator := newEvaluator(s.repository, s.logger)
        builder := newGoalTreeBuilder(evaluator, s.logger)
        searcher := NewSearcher(builder, s.logger)

        origin := domain.RelationNode {
            Node: &domain.RelationNode_EntitySet{
                EntitySet: &domain.EntitySetNode {
                    Object: req.AccessRequest.Object,
                    Relation: req.AccessRequest.Relation,
                },
            },
        }
        goal := Goal{
            Target: &domain.RelationNode{
                Node: &domain.RelationNode_Entity{
                    Entity: &domain.EntityNode{
                        Object: req.AccessRequest.Subject,
                    },
                },
            },
        }
        tree, err := searcher.Search(ctx, pol.Policy, &origin, &goal)

        if err != nil{
            return nil, err
        }

        found := false
        if tree.GetResult() == SearchResult_SUCCESS {
            found = true
        }

	return &api.CheckResponse{
            Result: &api.CheckResponse_Result{
                Authorized: found,
            },
        }, nil
}

func (s *Service) ExplainCheck(
    ctx context.Context,
    req *api.ExplainCheckRequest) (*api.ExplainCheckResponse, error) {
        serializer, err := SerializerFactory(req.Format)
        if err != nil {
            return nil, fmt.Errorf("check: %w: %w", api.ErrInvalidRequest, err)
        }

        pol, err := s.policyRepository.GetPolicy(ctx, req.PolicyId)
        if err != nil {
            return nil, fmt.Errorf("check: %w: %w", api.ErrInternal, err)
        } 
        if pol == nil {
            return nil, fmt.Errorf("check: %w: policy %v: %w", api.ErrInvalidRequest, req.PolicyId, policy.ErrPolicyNotFound)
        }

        evaluator := newEvaluator(s.repository, s.logger)
        builder := newGoalTreeBuilder(evaluator, s.logger)
        searcher := NewSearcher(builder, s.logger)

        origin := domain.RelationNode {
            Node: &domain.RelationNode_EntitySet{
                EntitySet: &domain.EntitySetNode {
                    Object: req.AccessRequest.Object,
                    Relation: req.AccessRequest.Relation,
                },
            },
        }
        goal := Goal{
            Target: &domain.RelationNode{
                Node: &domain.RelationNode_Entity{
                    Entity: &domain.EntityNode{
                        Object: req.AccessRequest.Subject,
                    },
                },
            },
        }
        tree, err := searcher.Search(ctx, pol.Policy, &origin, &goal)

        if err != nil{
            return nil, err
        }

        found := false
        if tree.GetResult() == SearchResult_SUCCESS {
            found = true
        }

        serialized, err := serializer.Serialize(tree)
        if err != nil {
            return nil, fmt.Errorf("explain check: %w: %w", api.ErrInternal, err)
        }

	return &api.ExplainCheckResponse{
            GoalTree: serialized,
            Authorized: found,
        }, nil
}

func (s *Service) Expand(
    ctx context.Context,
    req *api.ExpandRequest) (*api.ExpandResponse, error) {
        serializer, err := SerializerFactory(req.Format)
        if err != nil {
            return nil, fmt.Errorf("expand: %w: %w", api.ErrInvalidRequest, err)
        }

        pol, err := s.policyRepository.GetPolicy(ctx, req.PolicyId)
        if err != nil {
            return nil, fmt.Errorf("expand: %w: %w", api.ErrInternal, err)
        } 
        if pol == nil {
            return nil, fmt.Errorf("expand: %w: policy %v: %w", api.ErrInvalidRequest, req.PolicyId, policy.ErrPolicyNotFound)
        }

        evaluator := newEvaluator(s.repository, s.logger)
        builder := newGoalTreeBuilder(evaluator, s.logger)
        searcher := NewSearcher(builder, s.logger)

        goal := Goal{
            Target: nil,
        }
        tree, err := searcher.Search(ctx, pol.Policy, req.Root, &goal)
        if err != nil{
            return nil, fmt.Errorf("expand: %w", err)
        }

        serialized, err := serializer.Serialize(tree)
        if err != nil {
            return nil, fmt.Errorf("expand: %w: %w", api.ErrInternal, err)
        }

	return &api.ExpandResponse{
            GoalTree: serialized,
            Format: req.Format,
        }, nil
}

func (s *Service) DumpRelationships(
    ctx context.Context,
    req *api.DumpRelationshipsRequest) (*api.DumpRelationshipResponse, error) {
        rec, err := s.policyRepository.GetPolicy(ctx, req.PolicyId)
        if err != nil {
            return nil, fmt.Errorf("DumpRelationships: %w: %w", api.ErrInternal, err)
        } 
        if rec == nil {
            return nil, fmt.Errorf("DumpRelationships: %w: policy %v: %w", api.ErrInvalidRequest, req.PolicyId, policy.ErrPolicyNotFound)
        }

        walker := newWalker(s.repository, s.logger)
        tree, err := walker.Walk(ctx, rec.Policy)
        if err != nil {
            return nil, fmt.Errorf("DumpRelationships: %w", err)
        }

        response := &api.DumpRelationshipResponse{}
        switch req.Format {
        case api.DumpRelationshipsRequest_DOT:
            serializer := RelationTreeDOTSerializer{}
            treeStr, err := serializer.Serialize(tree)
            if err != nil {
                return nil, fmt.Errorf("DumpRelationships: %w", err)
            }
            response.Dump = &api.DumpRelationshipResponse_Dot {
                Dot: treeStr,
            }
        default:
            return nil, fmt.Errorf("DumpRelationships: format %v: %w", req.Format, domain.ErrInvalidVariant)
        }

        return response, nil
}
