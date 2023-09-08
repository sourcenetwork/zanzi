// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: zanzi/types/expand_tree.proto

package types

import (
	domain "github.com/sourcenetwork/zanzi/pkg/domain"
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

type Operator int32

const (
	Operator_UNION        Operator = 0
	Operator_DIFFERENCE   Operator = 1
	Operator_INTERSECTION Operator = 2
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "UNION",
		1: "DIFFERENCE",
		2: "INTERSECTION",
	}
	Operator_value = map[string]int32{
		"UNION":        0,
		"DIFFERENCE":   1,
		"INTERSECTION": 2,
	}
)

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}

func (x Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_zanzi_types_expand_tree_proto_enumTypes[0].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_zanzi_types_expand_tree_proto_enumTypes[0]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_zanzi_types_expand_tree_proto_rawDescGZIP(), []int{0}
}

type VerboseExpandTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity   *domain.Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Relation string          `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	Node     *ExpressionNode `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *VerboseExpandTree) Reset() {
	*x = VerboseExpandTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zanzi_types_expand_tree_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerboseExpandTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerboseExpandTree) ProtoMessage() {}

func (x *VerboseExpandTree) ProtoReflect() protoreflect.Message {
	mi := &file_zanzi_types_expand_tree_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerboseExpandTree.ProtoReflect.Descriptor instead.
func (*VerboseExpandTree) Descriptor() ([]byte, []int) {
	return file_zanzi_types_expand_tree_proto_rawDescGZIP(), []int{0}
}

func (x *VerboseExpandTree) GetEntity() *domain.Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *VerboseExpandTree) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *VerboseExpandTree) GetNode() *ExpressionNode {
	if x != nil {
		return x.Node
	}
	return nil
}

type ExpressionNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Node:
	//	*ExpressionNode_FactorNode
	//	*ExpressionNode_OpNode
	Node isExpressionNode_Node `protobuf_oneof:"node"`
}

func (x *ExpressionNode) Reset() {
	*x = ExpressionNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zanzi_types_expand_tree_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionNode) ProtoMessage() {}

func (x *ExpressionNode) ProtoReflect() protoreflect.Message {
	mi := &file_zanzi_types_expand_tree_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionNode.ProtoReflect.Descriptor instead.
func (*ExpressionNode) Descriptor() ([]byte, []int) {
	return file_zanzi_types_expand_tree_proto_rawDescGZIP(), []int{1}
}

func (m *ExpressionNode) GetNode() isExpressionNode_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (x *ExpressionNode) GetFactorNode() *FactorNode {
	if x, ok := x.GetNode().(*ExpressionNode_FactorNode); ok {
		return x.FactorNode
	}
	return nil
}

func (x *ExpressionNode) GetOpNode() *OpNode {
	if x, ok := x.GetNode().(*ExpressionNode_OpNode); ok {
		return x.OpNode
	}
	return nil
}

type isExpressionNode_Node interface {
	isExpressionNode_Node()
}

type ExpressionNode_FactorNode struct {
	FactorNode *FactorNode `protobuf:"bytes,1,opt,name=factor_node,json=factorNode,proto3,oneof"`
}

type ExpressionNode_OpNode struct {
	OpNode *OpNode `protobuf:"bytes,2,opt,name=op_node,json=opNode,proto3,oneof"`
}

func (*ExpressionNode_FactorNode) isExpressionNode_Node() {}

func (*ExpressionNode_OpNode) isExpressionNode_Node() {}

type FactorNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewriteRule string               `protobuf:"bytes,1,opt,name=rewrite_rule,json=rewriteRule,proto3" json:"rewrite_rule,omitempty"`
	Children    []*VerboseExpandTree `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *FactorNode) Reset() {
	*x = FactorNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zanzi_types_expand_tree_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FactorNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FactorNode) ProtoMessage() {}

func (x *FactorNode) ProtoReflect() protoreflect.Message {
	mi := &file_zanzi_types_expand_tree_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FactorNode.ProtoReflect.Descriptor instead.
func (*FactorNode) Descriptor() ([]byte, []int) {
	return file_zanzi_types_expand_tree_proto_rawDescGZIP(), []int{2}
}

func (x *FactorNode) GetRewriteRule() string {
	if x != nil {
		return x.RewriteRule
	}
	return ""
}

func (x *FactorNode) GetChildren() []*VerboseExpandTree {
	if x != nil {
		return x.Children
	}
	return nil
}

type OpNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left     *ExpressionNode `protobuf:"bytes,1,opt,name=left,proto3" json:"left,omitempty"`
	Operator Operator        `protobuf:"varint,2,opt,name=operator,proto3,enum=sourcenetwork.zanzi.types.Operator" json:"operator,omitempty"`
	Right    *ExpressionNode `protobuf:"bytes,3,opt,name=right,proto3" json:"right,omitempty"`
}

func (x *OpNode) Reset() {
	*x = OpNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zanzi_types_expand_tree_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpNode) ProtoMessage() {}

func (x *OpNode) ProtoReflect() protoreflect.Message {
	mi := &file_zanzi_types_expand_tree_proto_msgTypes[3]
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
	return file_zanzi_types_expand_tree_proto_rawDescGZIP(), []int{3}
}

func (x *OpNode) GetLeft() *ExpressionNode {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *OpNode) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_UNION
}

func (x *OpNode) GetRight() *ExpressionNode {
	if x != nil {
		return x.Right
	}
	return nil
}

var File_zanzi_types_expand_tree_proto protoreflect.FileDescriptor

var file_zanzi_types_expand_tree_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x78,
	0x70, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a,
	0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x7a, 0x61, 0x6e, 0x7a,
	0x69, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x11,
	0x56, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x54, 0x72, 0x65,
	0x65, 0x12, 0x3a, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4f, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x70, 0x4e,
	0x6f, 0x64, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x79, 0x0a, 0x0a, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a,
	0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x65, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x54, 0x72, 0x65, 0x65, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xc9, 0x01, 0x0a, 0x06, 0x4f, 0x70, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x3d, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74,
	0x12, 0x3f, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x3f, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2e, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x72, 0x69, 0x67,
	0x68, 0x74, 0x2a, 0x37, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x09,
	0x0a, 0x05, 0x55, 0x4e, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x46,
	0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x54,
	0x45, 0x52, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zanzi_types_expand_tree_proto_rawDescOnce sync.Once
	file_zanzi_types_expand_tree_proto_rawDescData = file_zanzi_types_expand_tree_proto_rawDesc
)

func file_zanzi_types_expand_tree_proto_rawDescGZIP() []byte {
	file_zanzi_types_expand_tree_proto_rawDescOnce.Do(func() {
		file_zanzi_types_expand_tree_proto_rawDescData = protoimpl.X.CompressGZIP(file_zanzi_types_expand_tree_proto_rawDescData)
	})
	return file_zanzi_types_expand_tree_proto_rawDescData
}

var file_zanzi_types_expand_tree_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zanzi_types_expand_tree_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_zanzi_types_expand_tree_proto_goTypes = []interface{}{
	(Operator)(0),             // 0: sourcenetwork.zanzi.types.Operator
	(*VerboseExpandTree)(nil), // 1: sourcenetwork.zanzi.types.VerboseExpandTree
	(*ExpressionNode)(nil),    // 2: sourcenetwork.zanzi.types.ExpressionNode
	(*FactorNode)(nil),        // 3: sourcenetwork.zanzi.types.FactorNode
	(*OpNode)(nil),            // 4: sourcenetwork.zanzi.types.OpNode
	(*domain.Entity)(nil),     // 5: sourcenetwork.zanzi.domain.Entity
}
var file_zanzi_types_expand_tree_proto_depIdxs = []int32{
	5, // 0: sourcenetwork.zanzi.types.VerboseExpandTree.entity:type_name -> sourcenetwork.zanzi.domain.Entity
	2, // 1: sourcenetwork.zanzi.types.VerboseExpandTree.node:type_name -> sourcenetwork.zanzi.types.ExpressionNode
	3, // 2: sourcenetwork.zanzi.types.ExpressionNode.factor_node:type_name -> sourcenetwork.zanzi.types.FactorNode
	4, // 3: sourcenetwork.zanzi.types.ExpressionNode.op_node:type_name -> sourcenetwork.zanzi.types.OpNode
	1, // 4: sourcenetwork.zanzi.types.FactorNode.children:type_name -> sourcenetwork.zanzi.types.VerboseExpandTree
	2, // 5: sourcenetwork.zanzi.types.OpNode.left:type_name -> sourcenetwork.zanzi.types.ExpressionNode
	0, // 6: sourcenetwork.zanzi.types.OpNode.operator:type_name -> sourcenetwork.zanzi.types.Operator
	2, // 7: sourcenetwork.zanzi.types.OpNode.right:type_name -> sourcenetwork.zanzi.types.ExpressionNode
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_zanzi_types_expand_tree_proto_init() }
func file_zanzi_types_expand_tree_proto_init() {
	if File_zanzi_types_expand_tree_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zanzi_types_expand_tree_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerboseExpandTree); i {
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
		file_zanzi_types_expand_tree_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionNode); i {
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
		file_zanzi_types_expand_tree_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FactorNode); i {
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
		file_zanzi_types_expand_tree_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_zanzi_types_expand_tree_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExpressionNode_FactorNode)(nil),
		(*ExpressionNode_OpNode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zanzi_types_expand_tree_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zanzi_types_expand_tree_proto_goTypes,
		DependencyIndexes: file_zanzi_types_expand_tree_proto_depIdxs,
		EnumInfos:         file_zanzi_types_expand_tree_proto_enumTypes,
		MessageInfos:      file_zanzi_types_expand_tree_proto_msgTypes,
	}.Build()
	File_zanzi_types_expand_tree_proto = out.File
	file_zanzi_types_expand_tree_proto_rawDesc = nil
	file_zanzi_types_expand_tree_proto_goTypes = nil
	file_zanzi_types_expand_tree_proto_depIdxs = nil
}
