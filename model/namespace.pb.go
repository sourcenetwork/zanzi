// File contains types representing Namespaces.
//
// Namespaces declare client application specific contexts and relations.
// A Namespace definition cointains a name and a set of relations.
// Namespaces are identified by their name, therefore names should be unique.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: proto/model/namespace.proto

package model

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

type Operation int32

const (
	Operation_UNION        Operation = 0
	Operation_INTERSECTION Operation = 1
	Operation_DIFFERENCE   Operation = 2
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "UNION",
		1: "INTERSECTION",
		2: "DIFFERENCE",
	}
	Operation_value = map[string]int32{
		"UNION":        0,
		"INTERSECTION": 1,
		"DIFFERENCE":   2,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_model_namespace_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_proto_model_namespace_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{0}
}

type Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Relations []*Relation `protobuf:"bytes,2,rep,name=relations,proto3" json:"relations,omitempty"`
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *Namespace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Namespace) GetRelations() []*Relation {
	if x != nil {
		return x.Relations
	}
	return nil
}

type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rewrite *UsersetRewrite `protobuf:"bytes,2,opt,name=rewrite,proto3" json:"rewrite,omitempty"`
}

func (x *Relation) Reset() {
	*x = Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *Relation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Relation) GetRewrite() *UsersetRewrite {
	if x != nil {
		return x.Rewrite
	}
	return nil
}

type UsersetRewrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpressionTree *RewriteNode `protobuf:"bytes,1,opt,name=expressionTree,proto3" json:"expressionTree,omitempty"`
}

func (x *UsersetRewrite) Reset() {
	*x = UsersetRewrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersetRewrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersetRewrite) ProtoMessage() {}

func (x *UsersetRewrite) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersetRewrite.ProtoReflect.Descriptor instead.
func (*UsersetRewrite) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{2}
}

func (x *UsersetRewrite) GetExpressionTree() *RewriteNode {
	if x != nil {
		return x.ExpressionTree
	}
	return nil
}

type RewriteNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Node:
	//	*RewriteNode_Opnode
	//	*RewriteNode_Leaf
	Node isRewriteNode_Node `protobuf_oneof:"node"`
}

func (x *RewriteNode) Reset() {
	*x = RewriteNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewriteNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewriteNode) ProtoMessage() {}

func (x *RewriteNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewriteNode.ProtoReflect.Descriptor instead.
func (*RewriteNode) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{3}
}

func (m *RewriteNode) GetNode() isRewriteNode_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (x *RewriteNode) GetOpnode() *OpNode {
	if x, ok := x.GetNode().(*RewriteNode_Opnode); ok {
		return x.Opnode
	}
	return nil
}

func (x *RewriteNode) GetLeaf() *Leaf {
	if x, ok := x.GetNode().(*RewriteNode_Leaf); ok {
		return x.Leaf
	}
	return nil
}

type isRewriteNode_Node interface {
	isRewriteNode_Node()
}

type RewriteNode_Opnode struct {
	Opnode *OpNode `protobuf:"bytes,1,opt,name=opnode,proto3,oneof"`
}

type RewriteNode_Leaf struct {
	Leaf *Leaf `protobuf:"bytes,2,opt,name=leaf,proto3,oneof"`
}

func (*RewriteNode_Opnode) isRewriteNode_Node() {}

func (*RewriteNode_Leaf) isRewriteNode_Node() {}

type OpNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left  *RewriteNode `protobuf:"bytes,1,opt,name=left,proto3" json:"left,omitempty"`
	Right *RewriteNode `protobuf:"bytes,2,opt,name=right,proto3" json:"right,omitempty"`
	Op    Operation    `protobuf:"varint,3,opt,name=op,proto3,enum=model.Operation" json:"op,omitempty"`
}

func (x *OpNode) Reset() {
	*x = OpNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpNode) ProtoMessage() {}

func (x *OpNode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpNode.ProtoReflect.Descriptor instead.
func (*OpNode) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{4}
}

func (x *OpNode) GetLeft() *RewriteNode {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *OpNode) GetRight() *RewriteNode {
	if x != nil {
		return x.Right
	}
	return nil
}

func (x *OpNode) GetOp() Operation {
	if x != nil {
		return x.Op
	}
	return Operation_UNION
}

type Leaf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule *Rule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *Leaf) Reset() {
	*x = Leaf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Leaf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leaf) ProtoMessage() {}

func (x *Leaf) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leaf.ProtoReflect.Descriptor instead.
func (*Leaf) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{5}
}

func (x *Leaf) GetRule() *Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Rule:
	//	*Rule_This
	//	*Rule_ComputedUserset
	//	*Rule_TupleToUserset
	Rule isRule_Rule `protobuf_oneof:"rule"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[6]
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
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{6}
}

func (m *Rule) GetRule() isRule_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (x *Rule) GetThis() *This {
	if x, ok := x.GetRule().(*Rule_This); ok {
		return x.This
	}
	return nil
}

func (x *Rule) GetComputedUserset() *ComputedUserset {
	if x, ok := x.GetRule().(*Rule_ComputedUserset); ok {
		return x.ComputedUserset
	}
	return nil
}

func (x *Rule) GetTupleToUserset() *TupleToUserset {
	if x, ok := x.GetRule().(*Rule_TupleToUserset); ok {
		return x.TupleToUserset
	}
	return nil
}

type isRule_Rule interface {
	isRule_Rule()
}

type Rule_This struct {
	This *This `protobuf:"bytes,1,opt,name=this,proto3,oneof"`
}

type Rule_ComputedUserset struct {
	ComputedUserset *ComputedUserset `protobuf:"bytes,2,opt,name=computedUserset,proto3,oneof"`
}

type Rule_TupleToUserset struct {
	TupleToUserset *TupleToUserset `protobuf:"bytes,3,opt,name=tupleToUserset,proto3,oneof"`
}

func (*Rule_This) isRule_Rule() {}

func (*Rule_ComputedUserset) isRule_Rule() {}

func (*Rule_TupleToUserset) isRule_Rule() {}

// This specifies a Rule which returns all users for a given (object, relation) pair.
// The rule performs userset chasing.
type This struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *This) Reset() {
	*x = This{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *This) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*This) ProtoMessage() {}

func (x *This) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use This.ProtoReflect.Descriptor instead.
func (*This) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{7}
}

// ComputerUserset specifies a rule to dynamically create a userset for a given object.
// The created Userset may then be used to perform additional lookups.
// It's functionally similar to "This", except it does checks using a volatile relation tuple.
type ComputedUserset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation string `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *ComputedUserset) Reset() {
	*x = ComputedUserset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedUserset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedUserset) ProtoMessage() {}

func (x *ComputedUserset) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputedUserset.ProtoReflect.Descriptor instead.
func (*ComputedUserset) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{8}
}

func (x *ComputedUserset) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

// The TupleToUserset is a multistep rule that:
// - Fetches a Tupleset for the given object using TuplesetRelation
// - For each fetched tuple compute an userset using UsersetRelation
//
// For Check and Expand calls, the computed usersets are used to perform further lookups.
//
// A practical example:
// Let TuplesetRelation = "parent"
// Let UsersetRelation = "owner"
// Let the input object be "doc:readme"
// TupleToUserset would then:
// - Lookup all tuples matching (obj="doc:readme", relation="parent").
//   assume the matching tuples are [(obj="doc:readme", relation="parent", user=(id="dir:root", relation="..."))]
// - For each found tuple, it would compute the userset (obj=${result_tuple_userset_obj}, relation=UsersetRelation);
//   eg. it would compute the userset (obj="dir:root", relation="owner")
type TupleToUserset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TuplesetRelation        string `protobuf:"bytes,1,opt,name=tuplesetRelation,proto3" json:"tuplesetRelation,omitempty"`
	ComputerUsersetRelation string `protobuf:"bytes,2,opt,name=computerUsersetRelation,proto3" json:"computerUsersetRelation,omitempty"`
}

func (x *TupleToUserset) Reset() {
	*x = TupleToUserset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_namespace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TupleToUserset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TupleToUserset) ProtoMessage() {}

func (x *TupleToUserset) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_namespace_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TupleToUserset.ProtoReflect.Descriptor instead.
func (*TupleToUserset) Descriptor() ([]byte, []int) {
	return file_proto_model_namespace_proto_rawDescGZIP(), []int{9}
}

func (x *TupleToUserset) GetTuplesetRelation() string {
	if x != nil {
		return x.TuplesetRelation
	}
	return ""
}

func (x *TupleToUserset) GetComputerUsersetRelation() string {
	if x != nil {
		return x.ComputerUsersetRelation
	}
	return ""
}

var File_proto_model_namespace_proto protoreflect.FileDescriptor

var file_proto_model_namespace_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x4e, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4f, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x74, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x07, 0x72, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0x4c, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x72, 0x65, 0x65, 0x22, 0x61, 0x0a, 0x0b, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x70, 0x4e, 0x6f, 0x64,
	0x65, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6c,
	0x65, 0x61, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4c, 0x65, 0x61, 0x66, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x42, 0x06,
	0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x7c, 0x0a, 0x06, 0x4f, 0x70, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x72, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x20, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x02, 0x6f, 0x70, 0x22, 0x27, 0x0a, 0x04, 0x4c, 0x65, 0x61, 0x66, 0x12, 0x1f, 0x0a, 0x04,
	0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0xb6, 0x01,
	0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x68, 0x69, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x68, 0x69,
	0x73, 0x48, 0x00, 0x52, 0x04, 0x74, 0x68, 0x69, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x12, 0x3f, 0x0a,
	0x0e, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0e,
	0x74, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x42, 0x06,
	0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x54, 0x68, 0x69, 0x73, 0x22, 0x2d,
	0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a,
	0x0e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x12,
	0x2a, 0x0a, 0x10, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x75, 0x70, 0x6c, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x17, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x38, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x46, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2d, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x62, 0x61, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_model_namespace_proto_rawDescOnce sync.Once
	file_proto_model_namespace_proto_rawDescData = file_proto_model_namespace_proto_rawDesc
)

func file_proto_model_namespace_proto_rawDescGZIP() []byte {
	file_proto_model_namespace_proto_rawDescOnce.Do(func() {
		file_proto_model_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_model_namespace_proto_rawDescData)
	})
	return file_proto_model_namespace_proto_rawDescData
}

var file_proto_model_namespace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_model_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_model_namespace_proto_goTypes = []interface{}{
	(Operation)(0),          // 0: model.Operation
	(*Namespace)(nil),       // 1: model.Namespace
	(*Relation)(nil),        // 2: model.Relation
	(*UsersetRewrite)(nil),  // 3: model.UsersetRewrite
	(*RewriteNode)(nil),     // 4: model.RewriteNode
	(*OpNode)(nil),          // 5: model.OpNode
	(*Leaf)(nil),            // 6: model.Leaf
	(*Rule)(nil),            // 7: model.Rule
	(*This)(nil),            // 8: model.This
	(*ComputedUserset)(nil), // 9: model.ComputedUserset
	(*TupleToUserset)(nil),  // 10: model.TupleToUserset
}
var file_proto_model_namespace_proto_depIdxs = []int32{
	2,  // 0: model.Namespace.relations:type_name -> model.Relation
	3,  // 1: model.Relation.rewrite:type_name -> model.UsersetRewrite
	4,  // 2: model.UsersetRewrite.expressionTree:type_name -> model.RewriteNode
	5,  // 3: model.RewriteNode.opnode:type_name -> model.OpNode
	6,  // 4: model.RewriteNode.leaf:type_name -> model.Leaf
	4,  // 5: model.OpNode.left:type_name -> model.RewriteNode
	4,  // 6: model.OpNode.right:type_name -> model.RewriteNode
	0,  // 7: model.OpNode.op:type_name -> model.Operation
	7,  // 8: model.Leaf.rule:type_name -> model.Rule
	8,  // 9: model.Rule.this:type_name -> model.This
	9,  // 10: model.Rule.computedUserset:type_name -> model.ComputedUserset
	10, // 11: model.Rule.tupleToUserset:type_name -> model.TupleToUserset
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_proto_model_namespace_proto_init() }
func file_proto_model_namespace_proto_init() {
	if File_proto_model_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_model_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Namespace); i {
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
		file_proto_model_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relation); i {
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
		file_proto_model_namespace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersetRewrite); i {
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
		file_proto_model_namespace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewriteNode); i {
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
		file_proto_model_namespace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpNode); i {
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
		file_proto_model_namespace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Leaf); i {
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
		file_proto_model_namespace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_model_namespace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*This); i {
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
		file_proto_model_namespace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputedUserset); i {
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
		file_proto_model_namespace_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TupleToUserset); i {
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
	file_proto_model_namespace_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*RewriteNode_Opnode)(nil),
		(*RewriteNode_Leaf)(nil),
	}
	file_proto_model_namespace_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Rule_This)(nil),
		(*Rule_ComputedUserset)(nil),
		(*Rule_TupleToUserset)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_model_namespace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_model_namespace_proto_goTypes,
		DependencyIndexes: file_proto_model_namespace_proto_depIdxs,
		EnumInfos:         file_proto_model_namespace_proto_enumTypes,
		MessageInfos:      file_proto_model_namespace_proto_msgTypes,
	}.Build()
	File_proto_model_namespace_proto = out.File
	file_proto_model_namespace_proto_rawDesc = nil
	file_proto_model_namespace_proto_goTypes = nil
	file_proto_model_namespace_proto_depIdxs = nil
}
