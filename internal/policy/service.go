package policy

import (
    "fmt"
	"context"

	"github.com/sourcenetwork/zanzi/pkg/api"
	"github.com/sourcenetwork/zanzi/pkg/domain"
	"github.com/sourcenetwork/zanzi/internal/utils"
)

var _ api.PolicyServiceServer = (*Service)(nil)


// Service implements Zanzi's PolicyServiceServer
type Service struct {
	api.UnimplementedPolicyServiceServer
        repository Repository
}

func NewService(repository Repository) api.PolicyServiceServer {
    return &Service{
        repository: repository,
    }
}

func (s *Service) getPolicyRepository() Repository {
    return s.repository
}

func (s *Service) CreatePolicy(
	ctx context.Context,
        req *api.CreatePolicyRequest) (*api.CreatePolicyResponse, error) {

        policy, err := GetPolicyFromDefinition(req.PolicyDefinition)
        if err != nil {
            return nil, fmt.Errorf("create policy: %w", err)
        }

        repo := s.getPolicyRepository()

        fetched, err := repo.GetPolicy(ctx, policy.Id)
        if fetched != nil {
            return nil, fmt.Errorf("create policy: policy %v: %w", policy.Id, ErrPolicyExists)
        }
        if err != nil {
            return nil, fmt.Errorf("create policy: %w", err)
        }

        spec := ValidPolicySpec{}
        err = spec.Verify(policy)
        if err != nil {
            return nil, fmt.Errorf("create policy: %w", err)
        }

        rec := domain.NewPolicyRecord(policy, req.AppData)
        _, err = repo.SetPolicy(ctx, rec)
        if err != nil {
            return nil, fmt.Errorf("create policy: %w", err)
        }
            
	return &api.CreatePolicyResponse{
            Record: rec,
        }, nil
}

func (s *Service) UpdatePolicy(
	ctx context.Context,
        req *api.UpdatePolicyRequest) (*api.UpdatePolicyResponse, error) {

        policy, err := GetPolicyFromDefinition(req.PolicyDefinition)
        if err != nil {
            return nil, fmt.Errorf("update policy: %w", err)
        }

        repo := s.getPolicyRepository()

        record, err := repo.GetPolicy(ctx, policy.Id)
        if err != nil {
            return nil, fmt.Errorf("update policy: %w", err)
        }
        if record == nil  {
            return nil, fmt.Errorf("update policy: policy %v: %w", policy.Id, ErrPolicyNotFound)
        }

        spec := ValidPolicySpec{}
        err = spec.Verify(policy)
        if err != nil {
            return nil, err
        }

        record.Policy = policy
        record.AppData = req.AppData
        _, err = repo.SetPolicy(ctx, record)
        if err != nil {
            return nil, err
        }
            
	return &api.UpdatePolicyResponse{
        }, nil
}


func (s *Service) DeletePolicy(
	ctx context.Context,
	req *api.DeletePolicyRequest) (*api.DeletePolicyResponse, error) {
        repo := s.getPolicyRepository()

        // TODO wrap in Tx
	found, err := repo.DeletePolicy(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete policy: %w", err)
	} 
        if !found {
            return &api.DeletePolicyResponse{
                Found: false,
                RelationshipsRemoved: 0,
            }, nil
        }

        selector := allRelationshipsSelector()
        count, err := repo.DeleteRelationships(ctx, req.Id, &selector)
	if err != nil {
		return nil, fmt.Errorf("failed to delete policy: %w", err)
	}

	return &api.DeletePolicyResponse{
            Found: true,
            RelationshipsRemoved: count,
        }, nil
}

func (s *Service) GetPolicy(
	ctx context.Context,
	req *api.GetPolicyRequest) (*api.GetPolicyResponse, error) {
        repo := s.getPolicyRepository()

        record, err := repo.GetPolicy(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get policy: %w", err)
	} 

        return &api.GetPolicyResponse{
            Record: record,
        }, nil
}

func (s *Service) ListPolicyIds(
	ctx context.Context,
	req *api.ListPolicyIdsRequest) (*api.ListPolicyIdsResponse, error) {
        repo := s.getPolicyRepository()

        ids, err := repo.ListPolicyIds(ctx)
        if err != nil {
            return nil, fmt.Errorf("list policy ids: %w", err)
        }

        records := utils.MapSlice(ids, func(id string) *api.ListPolicyIdsResponse_Record{
            return &api.ListPolicyIdsResponse_Record{
                Id: id,
            }
        })

	return &api.ListPolicyIdsResponse{
            Records: records,
	}, nil
}

func (s *Service) SetRelationship(
	ctx context.Context,
	req *api.SetRelationshipRequest) (*api.SetRelationshipResponse, error) {
        repo := s.getPolicyRepository()

        record, err := repo.GetPolicy(ctx, req.PolicyId)
        if err != nil  {
            return nil, fmt.Errorf("set relationship: %w", err)
        } else if record == nil {
            return nil, fmt.Errorf("set relationship: policy %v: %w", req.PolicyId, ErrPolicyNotFound)
        }

        lut := NewPolicyLookUpTable(record.Policy)
        spec := AllowedRelationshipSpec{}

        err = spec.Satisfies(req.Relationship, lut)
        if err != nil {
            return nil, fmt.Errorf("set relationship: %w", err)
        }

        rec := domain.NewRelationshipRecord(req.PolicyId, req.Relationship, req.AppData)
        updated, err := repo.SetRelationship(ctx, rec)
        if err != nil {
            return nil, fmt.Errorf("set relationship: store: %w", err)
        }

        return &api.SetRelationshipResponse{
            RecordOverwritten: updated,
        }, nil
}

func (s *Service) DeleteRelationship(
	ctx context.Context,
	req *api.DeleteRelationshipRequest) (*api.DeleteRelationshipResponse, error) {
        repo := s.getPolicyRepository()

        found, err := repo.DeleteRelationship(ctx, req.PolicyId, req.Relationship)
        if err != nil {
            return nil, fmt.Errorf("delete Relationship: %w", err)
        }

	return &api.DeleteRelationshipResponse{
            Found: bool(found),
        }, nil
}

func (s *Service) GetRelationship(
	ctx context.Context,
	req *api.GetRelationshipRequest) (*api.GetRelationshipResponse, error) {
        repo := s.getPolicyRepository()

        record, err := repo.GetRelationship(ctx, req.PolicyId, req.Relationship)
        if err != nil {
            return nil, fmt.Errorf("get relationship: %w", err)
        }

	return &api.GetRelationshipResponse{
            Record: record,
	}, nil
}

func (s *Service) FindRelationshipRecords(
	ctx context.Context,
	req *api.FindRelationshipRecordsRequest) (*api.FindRelationshipRecordsResponse, error) {
        repo := s.getPolicyRepository()

        records, err := repo.FindRelationshipRecords(ctx, req.PolicyId, req.Selector)
        if err != nil {
            return nil, fmt.Errorf("find relationship records: %w", err)
        }
        

        return &api.FindRelationshipRecordsResponse{
            Result: &api.RelationshipRecordSet{
                Records: records,
            },
        }, nil
}

func (s *Service) DeleteRelationships(
	ctx context.Context,
	req *api.DeleteRelationshipsRequest) (*api.DeleteRelationshipsResponse, error) {
        repo := s.getPolicyRepository()

        count, err := repo.DeleteRelationships(ctx, req.PolicyId, req.Selector)
        if err != nil {
            return nil, fmt.Errorf("delete relationships: %w", err)
        }

        return &api.DeleteRelationshipsResponse{
            RecordsAffected: count,
        }, nil
}
