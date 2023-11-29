package api

import "context"

type PolicyClient interface {
	// CreatePolicy creates a new Policy.
	// Supplying a Policy whose ID already exists in the store is an error.
	CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error)
	// UpdatePolicy updates the fields and relations for a Policy.
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyResponse, error)
	// Delete removes the Policy with the given Id from the store
	DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error)
	// Get fetches a policy from the Policy store, if it exists
	GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error)
	// List returns all Policies in the Policy store
	ListPolicyIds(context.Context, *ListPolicyIdsRequest) (*ListPolicyIdsResponse, error)
	// Set adds a Relationship in a Policy
	SetRelationship(context.Context, *SetRelationshipRequest) (*SetRelationshipResponse, error)
	// Remove a Relationship from a Policy
	DeleteRelationship(context.Context, *DeleteRelationshipRequest) (*DeleteRelationshipResponse, error)
	// Get fetches a Relationship contained in a Policy, if it exists
	GetRelationship(context.Context, *GetRelationshipRequest) (*GetRelationshipResponse, error)
	// DeleteRelationships removes all relationships which matches a SelectorSet
	DeleteRelationships(context.Context, *DeleteRelationshipsRequest) (*DeleteRelationshipsResponse, error)
	// FindRelationshipRecords returns all relationships which matches a SelectorSet
	FindRelationshipRecords(context.Context, *FindRelationshipRecordsRequest) (*FindRelationshipRecordsResponse, error)
}

type policyService struct {
	client policyServiceClient
}

// CreatePolicy creates a new Policy.
// Supplying a Policy whose ID already exists in the store is an error.
func (p *policyService) CreatePolicy(ctx context.Context, req *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	return p.client.CreatePolicy(ctx, req)
}

// UpdatePolicy updates the fields and relations for a Policy.
func (p *policyService) UpdatePolicy(ctx context.Context, req *UpdatePolicyRequest) (*UpdatePolicyResponse, error) {
	return p.client.UpdatePolicy(ctx, req)
}

// Delete removes the Policy with the given Id from the store
func (p *policyService) DeletePolicy(ctx context.Context, req *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	return p.client.DeletePolicy(ctx, req)
}

// Get fetches a policy from the Policy store, if it exists
func (p *policyService) GetPolicy(ctx context.Context, req *GetPolicyRequest) (*GetPolicyResponse, error) {
	return p.client.GetPolicy(ctx, req)
}

// List returns all Policies in the Policy store
func (p *policyService) ListPolicyIds(ctx context.Context, req *ListPolicyIdsRequest) (*ListPolicyIdsResponse, error) {
	return p.client.ListPolicyIds(ctx, req)
}

// Set adds a Relationship in a Policy
func (p *policyService) SetRelationship(ctx context.Context, req *SetRelationshipRequest) (*SetRelationshipResponse, error) {
	return p.client.SetRelationship(ctx, req)
}

// Remove a Relationship from a Policy
func (p *policyService) DeleteRelationship(ctx context.Context, req *DeleteRelationshipRequest) (*DeleteRelationshipResponse, error) {
	return p.client.DeleteRelationship(ctx, req)
}

// Get fetches a Relationship contained in a Policy, if it exists
func (p *policyService) GetRelationship(ctx context.Context, req *GetRelationshipRequest) (*GetRelationshipResponse, error) {
	return p.client.GetRelationship(ctx, req)
}

// DeleteRelationships removes all relationships which matches a SelectorSet
func (p *policyService) DeleteRelationships(ctx context.Context, req *DeleteRelationshipsRequest) (*DeleteRelationshipsResponse, error) {
	return p.client.DeleteRelationships(ctx, req)
}

// FindRelationshipRecords returns all relationships which matches a SelectorSet
func (p *policyService) FindRelationshipRecords(ctx context.Context, req *FindRelationshipRecordsRequest) (*FindRelationshipRecordsResponse, error) {
	return p.client.FindRelationshipRecords(ctx, req)
}

func PolicyServiceFromClient(client policyServiceClient) PolicyClient {
	return &policyService{client}
}
