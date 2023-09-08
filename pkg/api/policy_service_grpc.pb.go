// policy_service.proto defines an RPC service for managing Policies and their Relationships

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: zanzi/api/policy_service.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PolicyService_CreatePolicy_FullMethodName            = "/sourcenetwork.zanzi.api.PolicyService/CreatePolicy"
	PolicyService_UpdatePolicy_FullMethodName            = "/sourcenetwork.zanzi.api.PolicyService/UpdatePolicy"
	PolicyService_DeletePolicy_FullMethodName            = "/sourcenetwork.zanzi.api.PolicyService/DeletePolicy"
	PolicyService_GetPolicy_FullMethodName               = "/sourcenetwork.zanzi.api.PolicyService/GetPolicy"
	PolicyService_ListPolicyIds_FullMethodName           = "/sourcenetwork.zanzi.api.PolicyService/ListPolicyIds"
	PolicyService_SetRelationship_FullMethodName         = "/sourcenetwork.zanzi.api.PolicyService/SetRelationship"
	PolicyService_DeleteRelationship_FullMethodName      = "/sourcenetwork.zanzi.api.PolicyService/DeleteRelationship"
	PolicyService_GetRelationship_FullMethodName         = "/sourcenetwork.zanzi.api.PolicyService/GetRelationship"
	PolicyService_DeleteRelationships_FullMethodName     = "/sourcenetwork.zanzi.api.PolicyService/DeleteRelationships"
	PolicyService_FindRelationshipRecords_FullMethodName = "/sourcenetwork.zanzi.api.PolicyService/FindRelationshipRecords"
)

// PolicyServiceClient is the client API for PolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyServiceClient interface {
	// CreatePolicy creates a new Policy.
	// Supplying a Policy whose ID already exists in the store is an error.
	CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error)
	// UpdatePolicy updates the fields and relations for a Policy.
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyResponse, error)
	// Delete removes the Policy with the given Id from the store
	DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyResponse, error)
	// Get fetches a policy from the Policy store, if it exists
	GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error)
	// List returns all Policies in the Policy store
	ListPolicyIds(ctx context.Context, in *ListPolicyIdsRequest, opts ...grpc.CallOption) (*ListPolicyIdsResponse, error)
	// Set adds a Relationship in a Policy
	SetRelationship(ctx context.Context, in *SetRelationshipRequest, opts ...grpc.CallOption) (*SetRelationshipResponse, error)
	// Remove a Relationship from a Policy
	DeleteRelationship(ctx context.Context, in *DeleteRelationshipRequest, opts ...grpc.CallOption) (*DeleteRelationshipResponse, error)
	// Get fetches a Relationship contained in a Policy, if it exists
	GetRelationship(ctx context.Context, in *GetRelationshipRequest, opts ...grpc.CallOption) (*GetRelationshipResponse, error)
	// DeleteRelationships removes all relationships which matches a SelectorSet
	DeleteRelationships(ctx context.Context, in *DeleteRelationshipsRequest, opts ...grpc.CallOption) (*DeleteRelationshipsResponse, error)
	// FindRelationshipRecords returns all relationships which matches a SelectorSet
	FindRelationshipRecords(ctx context.Context, in *FindRelationshipRecordsRequest, opts ...grpc.CallOption) (*FindRelationshipRecordsResponse, error)
}

type policyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyServiceClient(cc grpc.ClientConnInterface) PolicyServiceClient {
	return &policyServiceClient{cc}
}

func (c *policyServiceClient) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...grpc.CallOption) (*CreatePolicyResponse, error) {
	out := new(CreatePolicyResponse)
	err := c.cc.Invoke(ctx, PolicyService_CreatePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (*UpdatePolicyResponse, error) {
	out := new(UpdatePolicyResponse)
	err := c.cc.Invoke(ctx, PolicyService_UpdatePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...grpc.CallOption) (*DeletePolicyResponse, error) {
	out := new(DeletePolicyResponse)
	err := c.cc.Invoke(ctx, PolicyService_DeletePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...grpc.CallOption) (*GetPolicyResponse, error) {
	out := new(GetPolicyResponse)
	err := c.cc.Invoke(ctx, PolicyService_GetPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) ListPolicyIds(ctx context.Context, in *ListPolicyIdsRequest, opts ...grpc.CallOption) (*ListPolicyIdsResponse, error) {
	out := new(ListPolicyIdsResponse)
	err := c.cc.Invoke(ctx, PolicyService_ListPolicyIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) SetRelationship(ctx context.Context, in *SetRelationshipRequest, opts ...grpc.CallOption) (*SetRelationshipResponse, error) {
	out := new(SetRelationshipResponse)
	err := c.cc.Invoke(ctx, PolicyService_SetRelationship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) DeleteRelationship(ctx context.Context, in *DeleteRelationshipRequest, opts ...grpc.CallOption) (*DeleteRelationshipResponse, error) {
	out := new(DeleteRelationshipResponse)
	err := c.cc.Invoke(ctx, PolicyService_DeleteRelationship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) GetRelationship(ctx context.Context, in *GetRelationshipRequest, opts ...grpc.CallOption) (*GetRelationshipResponse, error) {
	out := new(GetRelationshipResponse)
	err := c.cc.Invoke(ctx, PolicyService_GetRelationship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) DeleteRelationships(ctx context.Context, in *DeleteRelationshipsRequest, opts ...grpc.CallOption) (*DeleteRelationshipsResponse, error) {
	out := new(DeleteRelationshipsResponse)
	err := c.cc.Invoke(ctx, PolicyService_DeleteRelationships_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyServiceClient) FindRelationshipRecords(ctx context.Context, in *FindRelationshipRecordsRequest, opts ...grpc.CallOption) (*FindRelationshipRecordsResponse, error) {
	out := new(FindRelationshipRecordsResponse)
	err := c.cc.Invoke(ctx, PolicyService_FindRelationshipRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServiceServer is the server API for PolicyService service.
// All implementations must embed UnimplementedPolicyServiceServer
// for forward compatibility
type PolicyServiceServer interface {
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
	mustEmbedUnimplementedPolicyServiceServer()
}

// UnimplementedPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyServiceServer struct {
}

func (UnimplementedPolicyServiceServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedPolicyServiceServer) UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePolicy not implemented")
}
func (UnimplementedPolicyServiceServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedPolicyServiceServer) GetPolicy(context.Context, *GetPolicyRequest) (*GetPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedPolicyServiceServer) ListPolicyIds(context.Context, *ListPolicyIdsRequest) (*ListPolicyIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicyIds not implemented")
}
func (UnimplementedPolicyServiceServer) SetRelationship(context.Context, *SetRelationshipRequest) (*SetRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRelationship not implemented")
}
func (UnimplementedPolicyServiceServer) DeleteRelationship(context.Context, *DeleteRelationshipRequest) (*DeleteRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelationship not implemented")
}
func (UnimplementedPolicyServiceServer) GetRelationship(context.Context, *GetRelationshipRequest) (*GetRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelationship not implemented")
}
func (UnimplementedPolicyServiceServer) DeleteRelationships(context.Context, *DeleteRelationshipsRequest) (*DeleteRelationshipsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelationships not implemented")
}
func (UnimplementedPolicyServiceServer) FindRelationshipRecords(context.Context, *FindRelationshipRecordsRequest) (*FindRelationshipRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRelationshipRecords not implemented")
}
func (UnimplementedPolicyServiceServer) mustEmbedUnimplementedPolicyServiceServer() {}

// UnsafePolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyServiceServer will
// result in compilation errors.
type UnsafePolicyServiceServer interface {
	mustEmbedUnimplementedPolicyServiceServer()
}

func RegisterPolicyServiceServer(s grpc.ServiceRegistrar, srv PolicyServiceServer) {
	s.RegisterService(&PolicyService_ServiceDesc, srv)
}

func _PolicyService_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_CreatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).CreatePolicy(ctx, req.(*CreatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_UpdatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).UpdatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_UpdatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_DeletePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).DeletePolicy(ctx, req.(*DeletePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_GetPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).GetPolicy(ctx, req.(*GetPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_ListPolicyIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPolicyIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).ListPolicyIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_ListPolicyIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).ListPolicyIds(ctx, req.(*ListPolicyIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_SetRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).SetRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_SetRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).SetRelationship(ctx, req.(*SetRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_DeleteRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).DeleteRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_DeleteRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).DeleteRelationship(ctx, req.(*DeleteRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_GetRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).GetRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_GetRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).GetRelationship(ctx, req.(*GetRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_DeleteRelationships_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelationshipsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).DeleteRelationships(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_DeleteRelationships_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).DeleteRelationships(ctx, req.(*DeleteRelationshipsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyService_FindRelationshipRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRelationshipRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServiceServer).FindRelationshipRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyService_FindRelationshipRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServiceServer).FindRelationshipRecords(ctx, req.(*FindRelationshipRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyService_ServiceDesc is the grpc.ServiceDesc for PolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sourcenetwork.zanzi.api.PolicyService",
	HandlerType: (*PolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePolicy",
			Handler:    _PolicyService_CreatePolicy_Handler,
		},
		{
			MethodName: "UpdatePolicy",
			Handler:    _PolicyService_UpdatePolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _PolicyService_DeletePolicy_Handler,
		},
		{
			MethodName: "GetPolicy",
			Handler:    _PolicyService_GetPolicy_Handler,
		},
		{
			MethodName: "ListPolicyIds",
			Handler:    _PolicyService_ListPolicyIds_Handler,
		},
		{
			MethodName: "SetRelationship",
			Handler:    _PolicyService_SetRelationship_Handler,
		},
		{
			MethodName: "DeleteRelationship",
			Handler:    _PolicyService_DeleteRelationship_Handler,
		},
		{
			MethodName: "GetRelationship",
			Handler:    _PolicyService_GetRelationship_Handler,
		},
		{
			MethodName: "DeleteRelationships",
			Handler:    _PolicyService_DeleteRelationships_Handler,
		},
		{
			MethodName: "FindRelationshipRecords",
			Handler:    _PolicyService_FindRelationshipRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zanzi/api/policy_service.proto",
}
