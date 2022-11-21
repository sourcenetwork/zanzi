// Policy definition
//
// Policy defines a client's application relation model.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: internal/domain/policy/policy.proto

package policy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ActorIdType specify an identifier type for an actor object.
// Actor ids are enforced as system constants.
type ActorIdType int32

const (
	ActorIdType_STRING ActorIdType = 0
	ActorIdType_NUMBER ActorIdType = 1
	ActorIdType_DID    ActorIdType = 2
)

// Enum value maps for ActorIdType.
var (
	ActorIdType_name = map[int32]string{
		0: "STRING",
		1: "NUMBER",
		2: "DID",
	}
	ActorIdType_value = map[string]int32{
		"STRING": 0,
		"NUMBER": 1,
		"DID":    2,
	}
)

func (x ActorIdType) Enum() *ActorIdType {
	p := new(ActorIdType)
	*p = x
	return p
}

func (x ActorIdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActorIdType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_domain_policy_policy_proto_enumTypes[0].Descriptor()
}

func (ActorIdType) Type() protoreflect.EnumType {
	return &file_internal_domain_policy_policy_proto_enumTypes[0]
}

func (x ActorIdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActorIdType.Descriptor instead.
func (ActorIdType) EnumDescriptor() ([]byte, []int) {
	return file_internal_domain_policy_policy_proto_rawDescGZIP(), []int{0}
}

// Specifies the type of a rule.
// Rule Types have no runtime difference, they are a public interface aid.
type RuleType int32

const (
	RuleType_RELATION   RuleType = 0
	RuleType_PERMISSION RuleType = 1
)

// Enum value maps for RuleType.
var (
	RuleType_name = map[int32]string{
		0: "RELATION",
		1: "PERMISSION",
	}
	RuleType_value = map[string]int32{
		"RELATION":   0,
		"PERMISSION": 1,
	}
)

func (x RuleType) Enum() *RuleType {
	p := new(RuleType)
	*p = x
	return p
}

func (x RuleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RuleType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_domain_policy_policy_proto_enumTypes[1].Descriptor()
}

func (RuleType) Type() protoreflect.EnumType {
	return &file_internal_domain_policy_policy_proto_enumTypes[1]
}

func (x RuleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RuleType.Descriptor instead.
func (RuleType) EnumDescriptor() ([]byte, []int) {
	return file_internal_domain_policy_policy_proto_rawDescGZIP(), []int{1}
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// creation date
	// last modified
	// creator
	Resources  []*Resource       `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty"`
	Actors     []*Actor          `protobuf:"bytes,5,rep,name=actors,proto3" json:"actors,omitempty"`
	Attributes map[string]string `protobuf:"bytes,10,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_policy_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_internal_domain_policy_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Policy) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Policy) GetActors() []*Actor {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *Policy) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type TypeConstraint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule string `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"` // TODO this could also include objectId / actor contraints
}

func (x *TypeConstraint) Reset() {
	*x = TypeConstraint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_policy_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeConstraint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeConstraint) ProtoMessage() {}

func (x *TypeConstraint) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeConstraint.ProtoReflect.Descriptor instead.
func (*TypeConstraint) Descriptor() ([]byte, []int) {
	return file_internal_domain_policy_policy_proto_rawDescGZIP(), []int{1}
}

func (x *TypeConstraint) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

// Actor represents a class of system users
type Actor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Constraints []ActorIdType `protobuf:"varint,2,rep,packed,name=constraints,proto3,enum=internal.domain.policy.ActorIdType" json:"constraints,omitempty"`
}

func (x *Actor) Reset() {
	*x = Actor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_policy_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Actor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Actor) ProtoMessage() {}

func (x *Actor) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Actor.ProtoReflect.Descriptor instead.
func (*Actor) Descriptor() ([]byte, []int) {
	return file_internal_domain_policy_policy_proto_rawDescGZIP(), []int{2}
}

func (x *Actor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Actor) GetConstraints() []ActorIdType {
	if x != nil {
		return x.Constraints
	}
	return nil
}

// Resources represents a class of objects in the relation graph
// Example resources: files, directories, repositories, groups
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rules []*Rule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_policy_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_internal_domain_policy_policy_proto_rawDescGZIP(), []int{3}
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// A Rule defines a Relation Type in the relation graph.
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           RuleType          `protobuf:"varint,1,opt,name=type,proto3,enum=internal.domain.policy.RuleType" json:"type,omitempty"`
	Name           string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ExpressionTree *Tree             `protobuf:"bytes,3,opt,name=expression_tree,json=expressionTree,proto3" json:"expression_tree,omitempty"`
	RewriteExpr    string            `protobuf:"bytes,4,opt,name=rewrite_expr,json=rewriteExpr,proto3" json:"rewrite_expr,omitempty"`
	Constraints    []*TypeConstraint `protobuf:"bytes,5,rep,name=constraints,proto3" json:"constraints,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_policy_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_internal_domain_policy_policy_proto_rawDescGZIP(), []int{4}
}

func (x *Rule) GetType() RuleType {
	if x != nil {
		return x.Type
	}
	return RuleType_RELATION
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetExpressionTree() *Tree {
	if x != nil {
		return x.ExpressionTree
	}
	return nil
}

func (x *Rule) GetRewriteExpr() string {
	if x != nil {
		return x.RewriteExpr
	}
	return ""
}

func (x *Rule) GetConstraints() []*TypeConstraint {
	if x != nil {
		return x.Constraints
	}
	return nil
}

var File_internal_domain_policy_policy_proto protoreflect.FileDescriptor

var file_internal_domain_policy_policy_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x21, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd4, 0x02, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3e, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x35, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x24, 0x0a, 0x0e, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x62, 0x0a,
	0x05, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x22, 0x52, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x32, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x34,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x52,
	0x0e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x78,
	0x70, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x2a, 0x2e, 0x0a, 0x0b,
	0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53,
	0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x49, 0x44, 0x10, 0x02, 0x2a, 0x28, 0x0a, 0x08,
	0x52, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x62,
	0x61, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_domain_policy_policy_proto_rawDescOnce sync.Once
	file_internal_domain_policy_policy_proto_rawDescData = file_internal_domain_policy_policy_proto_rawDesc
)

func file_internal_domain_policy_policy_proto_rawDescGZIP() []byte {
	file_internal_domain_policy_policy_proto_rawDescOnce.Do(func() {
		file_internal_domain_policy_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_domain_policy_policy_proto_rawDescData)
	})
	return file_internal_domain_policy_policy_proto_rawDescData
}

var file_internal_domain_policy_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_domain_policy_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_domain_policy_policy_proto_goTypes = []interface{}{
	(ActorIdType)(0),       // 0: internal.domain.policy.ActorIdType
	(RuleType)(0),          // 1: internal.domain.policy.RuleType
	(*Policy)(nil),         // 2: internal.domain.policy.Policy
	(*TypeConstraint)(nil), // 3: internal.domain.policy.TypeConstraint
	(*Actor)(nil),          // 4: internal.domain.policy.Actor
	(*Resource)(nil),       // 5: internal.domain.policy.Resource
	(*Rule)(nil),           // 6: internal.domain.policy.Rule
	nil,                    // 7: internal.domain.policy.Policy.AttributesEntry
	(*Tree)(nil),           // 8: internal.domain.policy.Tree
}
var file_internal_domain_policy_policy_proto_depIdxs = []int32{
	5, // 0: internal.domain.policy.Policy.resources:type_name -> internal.domain.policy.Resource
	4, // 1: internal.domain.policy.Policy.actors:type_name -> internal.domain.policy.Actor
	7, // 2: internal.domain.policy.Policy.attributes:type_name -> internal.domain.policy.Policy.AttributesEntry
	0, // 3: internal.domain.policy.Actor.constraints:type_name -> internal.domain.policy.ActorIdType
	6, // 4: internal.domain.policy.Resource.rules:type_name -> internal.domain.policy.Rule
	1, // 5: internal.domain.policy.Rule.type:type_name -> internal.domain.policy.RuleType
	8, // 6: internal.domain.policy.Rule.expression_tree:type_name -> internal.domain.policy.Tree
	3, // 7: internal.domain.policy.Rule.constraints:type_name -> internal.domain.policy.TypeConstraint
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_internal_domain_policy_policy_proto_init() }
func file_internal_domain_policy_policy_proto_init() {
	if File_internal_domain_policy_policy_proto != nil {
		return
	}
	file_internal_domain_policy_tree_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_domain_policy_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_domain_policy_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeConstraint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_domain_policy_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Actor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_domain_policy_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_domain_policy_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_domain_policy_policy_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_domain_policy_policy_proto_goTypes,
		DependencyIndexes: file_internal_domain_policy_policy_proto_depIdxs,
		EnumInfos:         file_internal_domain_policy_policy_proto_enumTypes,
		MessageInfos:      file_internal_domain_policy_policy_proto_msgTypes,
	}.Build()
	File_internal_domain_policy_policy_proto = out.File
	file_internal_domain_policy_policy_proto_rawDesc = nil
	file_internal_domain_policy_policy_proto_goTypes = nil
	file_internal_domain_policy_policy_proto_depIdxs = nil
}
