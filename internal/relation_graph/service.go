package relation_graph

import (
	"context"

	"github.com/google/uuid"

	"github.com/sourcenetwork/zanzi/internal/policy"
	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/pkg/errors"
	"github.com/sourcenetwork/zanzi/pkg/types"
)

var _ api.RelationGraphServer = (*Service)(nil)

func NewService(repository NodeRepository, policyRepository policy.Repository, logger types.Logger) api.RelationGraphServer {
	return &Service{
		logger:           logger,
		repository:       repository,
		policyRepository: policyRepository,
	}
}

type Service struct {
	api.UnimplementedRelationGraphServer
	logger           types.Logger
	repository       NodeRepository
	policyRepository policy.Repository
}

func (s *Service) Check(
	ctx context.Context,
	req *api.CheckRequest) (*api.CheckResponse, error) {
	pol, err := s.policyRepository.GetPolicy(ctx, req.PolicyId)
	if err != nil {
		return nil, err
	}
	if pol == nil {
		return nil, errors.ErrPolicyNotFound(req.PolicyId)
	}

	evaluator := newEvaluator(s.repository, s.logger)
	builder := newGoalTreeBuilder(evaluator, s.logger)
	searcher := NewSearcher(builder, s.logger)

	origin := domain.RelationNode{
		Node: &domain.RelationNode_EntitySet{
			EntitySet: &domain.EntitySetNode{
				Object:   req.AccessRequest.Object,
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

	if err != nil {
		return nil, err
	}

	return &api.CheckResponse{
		Result: &api.CheckResponse_Result{
			Authorized: tree.GetResult().Authorized,
		},
	}, nil
}

func (s *Service) ExplainCheck(
	ctx context.Context,
	req *api.ExplainCheckRequest) (*api.ExplainCheckResponse, error) {
	pol, err := s.policyRepository.GetPolicy(ctx, req.PolicyId)
	if err != nil {
		return nil, err
	}
	if pol == nil {
		return nil, errors.ErrPolicyNotFound(req.PolicyId)
	}

	evaluator := newEvaluator(s.repository, s.logger)
	builder := newGoalTreeBuilder(evaluator, s.logger)
	searcher := NewSearcher(builder, s.logger)

	origin := domain.RelationNode{
		Node: &domain.RelationNode_EntitySet{
			EntitySet: &domain.EntitySetNode{
				Object:   req.AccessRequest.Object,
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

	if err != nil {
		return nil, err
	}

	mapper := ExplainCheckTreeMapper{}

	return &api.ExplainCheckResponse{
		Graph:      mapper.Map(tree),
		Authorized: tree.GetResult().Authorized,
	}, nil
}

func (s *Service) DOTExplainCheck(
	ctx context.Context,
	req *api.DOTExplainCheckRequest) (*api.DOTExplainCheckResponse, error) {
	tree, err := s.check(ctx, req.PolicyId, req.AccessRequest)
	if err != nil {
		return nil, err
	}

	mapper := ExplainCheckTreeMapper{}
	explainTree := mapper.Map(tree)
	serializer := DOTMapper{}
	out, err := serializer.Map(explainTree, req.OmitUnknown)
	if err != nil {
		return nil, err
	}
	return &api.DOTExplainCheckResponse{
		Tree:       out,
		Authorized: tree.GetResult().Authorized,
	}, nil
}

func (s *Service) check(ctx context.Context, polId string, req *domain.AccessRequest) (GoalTree, error) {
	pol, err := s.policyRepository.GetPolicy(ctx, polId)
	if err != nil {
		return nil, err
	}
	if pol == nil {
		return nil, errors.ErrPolicyNotFound(polId)
	}

	evaluator := newEvaluator(s.repository, s.logger)
	builder := newGoalTreeBuilder(evaluator, s.logger)
	searcher := NewSearcher(builder, s.logger)

	origin := domain.RelationNode{
		Node: &domain.RelationNode_EntitySet{
			EntitySet: &domain.EntitySetNode{
				Object:   req.Object,
				Relation: req.Relation,
			},
		},
	}
	goal := Goal{
		Target: &domain.RelationNode{
			Node: &domain.RelationNode_Entity{
				Entity: &domain.EntityNode{
					Object: req.Subject,
				},
			},
		},
	}
	return searcher.Search(ctx, pol.Policy, &origin, &goal)
}

func (s *Service) DumpRelationships(
	ctx context.Context,
	req *api.DumpRelationshipsRequest) (*api.DumpRelationshipResponse, error) {
	rec, err := s.policyRepository.GetPolicy(ctx, req.PolicyId)
	if err != nil {
		return nil, err
	}
	if rec == nil {
		return nil, errors.ErrPolicyNotFound(req.PolicyId)
	}

	walker := newWalker(s.repository, s.logger)
	tree, err := walker.Walk(ctx, rec.Policy)
	if err != nil {
		return nil, err
	}

	response := &api.DumpRelationshipResponse{}
	switch req.Format {
	case api.DumpRelationshipsRequest_DOT:
		serializer := RelationTreeDOTSerializer{}
		treeStr, err := serializer.Serialize(tree)
		if err != nil {
			return nil, err
		}
		response.Dump = &api.DumpRelationshipResponse_Dot{
			Dot: treeStr,
		}
	default:
		return nil, errors.Wrap("invalid format", errors.ErrInvalidVariant)
	}

	return response, nil
}

func (s *Service) CheckExpression(
	ctx context.Context,
	req *api.CheckExpressionRequest) (*api.CheckExpressionResponse, error) {
	record, err := s.policyRepository.GetPolicy(ctx, req.PolicyId)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, errors.ErrPolicyNotFound(req.PolicyId)
	}

	// augment policy with the provided expression
	id := uuid.NewString()
	res := record.Policy.GetResourceByName(req.Object.Resource)
	if res == nil {
		return nil, errors.New("resource not found", errors.BadInput,
			errors.Pair("policy", req.PolicyId),
			errors.Pair("resource", req.Object.Resource))
	}
	tmpRel := domain.Relation{
		Name:        id,
		Description: "temporary relation for check expression",
		SubjectRestriction: &domain.SubjectRestriction{
			SubjectRestriction: &domain.SubjectRestriction_UniversalSet{},
		},
		RelationExpression: &domain.RelationExpression{
			Expression: &domain.RelationExpression_Expr{
				Expr: req.RelationExpression,
			},
		},
	}
	res.Relations = append(res.Relations, &tmpRel)

	evaluator := newEvaluator(s.repository, s.logger)
	builder := newGoalTreeBuilder(evaluator, s.logger)
	searcher := NewSearcher(builder, s.logger)

	origin := domain.RelationNode{
		Node: &domain.RelationNode_EntitySet{
			EntitySet: &domain.EntitySetNode{
				Object:   req.Object,
				Relation: id,
			},
		},
	}
	goal := Goal{
		Target: &domain.RelationNode{
			Node: &domain.RelationNode_Entity{
				Entity: &domain.EntityNode{
					Object: req.Subject,
				},
			},
		},
	}
	tree, err := searcher.Search(ctx, record.Policy, &origin, &goal)

	if err != nil {
		return nil, err
	}

	return &api.CheckExpressionResponse{
		Authorized: tree.GetResult().Authorized,
	}, nil
}
