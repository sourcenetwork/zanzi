// Userset Rewrite Rules definitions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.21.12
// source: internal/domain/policy/rules.proto

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

// RewriteRule specifies the logical model between relations
type RewriteRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RewriteRule:
	//
	//	*RewriteRule_This
	//	*RewriteRule_ComputedUserset
	//	*RewriteRule_TupleToUserset
	RewriteRule isRewriteRule_RewriteRule `protobuf_oneof:"rewrite_rule"`
}

func (x *RewriteRule) Reset() {
	*x = RewriteRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_policy_rules_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewriteRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewriteRule) ProtoMessage() {}

func (x *RewriteRule) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_rules_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewriteRule.ProtoReflect.Descriptor instead.
func (*RewriteRule) Descriptor() ([]byte, []int) {
	return file_internal_domain_policy_rules_proto_rawDescGZIP(), []int{0}
}

func (m *RewriteRule) GetRewriteRule() isRewriteRule_RewriteRule {
	if m != nil {
		return m.RewriteRule
	}
	return nil
}

func (x *RewriteRule) GetThis() *This {
	if x, ok := x.GetRewriteRule().(*RewriteRule_This); ok {
		return x.This
	}
	return nil
}

func (x *RewriteRule) GetComputedUserset() *ComputedUserset {
	if x, ok := x.GetRewriteRule().(*RewriteRule_ComputedUserset); ok {
		return x.ComputedUserset
	}
	return nil
}

func (x *RewriteRule) GetTupleToUserset() *TupleToUserset {
	if x, ok := x.GetRewriteRule().(*RewriteRule_TupleToUserset); ok {
		return x.TupleToUserset
	}
	return nil
}

type isRewriteRule_RewriteRule interface {
	isRewriteRule_RewriteRule()
}

type RewriteRule_This struct {
	This *This `protobuf:"bytes,1,opt,name=this,proto3,oneof"`
}

type RewriteRule_ComputedUserset struct {
	ComputedUserset *ComputedUserset `protobuf:"bytes,2,opt,name=computed_userset,json=computedUserset,proto3,oneof"`
}

type RewriteRule_TupleToUserset struct {
	TupleToUserset *TupleToUserset `protobuf:"bytes,3,opt,name=tuple_to_userset,json=tupleToUserset,proto3,oneof"`
}

func (*RewriteRule_This) isRewriteRule_RewriteRule() {}

func (*RewriteRule_ComputedUserset) isRewriteRule_RewriteRule() {}

func (*RewriteRule_TupleToUserset) isRewriteRule_RewriteRule() {}

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
		mi := &file_internal_domain_policy_rules_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *This) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*This) ProtoMessage() {}

func (x *This) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_rules_proto_msgTypes[1]
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
	return file_internal_domain_policy_rules_proto_rawDescGZIP(), []int{1}
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
		mi := &file_internal_domain_policy_rules_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedUserset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedUserset) ProtoMessage() {}

func (x *ComputedUserset) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_rules_proto_msgTypes[2]
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
	return file_internal_domain_policy_rules_proto_rawDescGZIP(), []int{2}
}

func (x *ComputedUserset) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

// TupleToUserset is rule which is used to traverse hierarchical relations.
// For a given object `obj` it works as follows:
// 1. Fetch the direct neighbors of the (`obj`, tuplset_relation) pair
// 2. For each fetched tuple `tf`, create a computed userset of (tf.Namespace, tf.ObjId, cu_relation)
//
// Example:
// Let tuplset_relation = "parent"
// Let cu_relation = "owner"
// Let the input object be "doc:readme"
// TupleToUserset would then:
//  1. Lookup all tuples matching (obj="doc:readme", relation="parent").
//     Assume the matching tuples are [(obj="doc:readme", relation="parent", user=(id="dir:root", relation="..."))]
//  2. For each found tuple, it would compute the userset (obj=${result_tuple_userset_obj}, relation=cu_relation)
//
// The result for this example would be the userset: (obj="dir:root", relation="owner")
type TupleToUserset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Relation used to fetch neighbors
	TuplesetRelation string `protobuf:"bytes,1,opt,name=tupleset_relation,json=tuplesetRelation,proto3" json:"tupleset_relation,omitempty"`
	// Namespace to which the Computed Userset relation belongs
	CuRelationNamespace string `protobuf:"bytes,2,opt,name=cu_relation_namespace,json=cuRelationNamespace,proto3" json:"cu_relation_namespace,omitempty"`
	// Relation to use within computed usersets
	CuRelation string `protobuf:"bytes,3,opt,name=cu_relation,json=cuRelation,proto3" json:"cu_relation,omitempty"`
}

func (x *TupleToUserset) Reset() {
	*x = TupleToUserset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_domain_policy_rules_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TupleToUserset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TupleToUserset) ProtoMessage() {}

func (x *TupleToUserset) ProtoReflect() protoreflect.Message {
	mi := &file_internal_domain_policy_rules_proto_msgTypes[3]
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
	return file_internal_domain_policy_rules_proto_rawDescGZIP(), []int{3}
}

func (x *TupleToUserset) GetTuplesetRelation() string {
	if x != nil {
		return x.TuplesetRelation
	}
	return ""
}

func (x *TupleToUserset) GetCuRelationNamespace() string {
	if x != nil {
		return x.CuRelationNamespace
	}
	return ""
}

func (x *TupleToUserset) GetCuRelation() string {
	if x != nil {
		return x.CuRelation
	}
	return ""
}

var File_internal_domain_policy_rules_proto protoreflect.FileDescriptor

var file_internal_domain_policy_rules_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0xfb, 0x01, 0x0a,
	0x0b, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x74, 0x68, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x54, 0x68, 0x69, 0x73, 0x48, 0x00, 0x52, 0x04, 0x74, 0x68, 0x69, 0x73,
	0x12, 0x54, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x12, 0x52, 0x0a, 0x10, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x5f,
	0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x75, 0x70, 0x6c,
	0x65, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x65, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x72, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x54, 0x68,
	0x69, 0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x92, 0x01, 0x0a, 0x0e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x54, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x74,
	0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x75, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x63, 0x75, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x5f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x62,
	0x61, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_domain_policy_rules_proto_rawDescOnce sync.Once
	file_internal_domain_policy_rules_proto_rawDescData = file_internal_domain_policy_rules_proto_rawDesc
)

func file_internal_domain_policy_rules_proto_rawDescGZIP() []byte {
	file_internal_domain_policy_rules_proto_rawDescOnce.Do(func() {
		file_internal_domain_policy_rules_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_domain_policy_rules_proto_rawDescData)
	})
	return file_internal_domain_policy_rules_proto_rawDescData
}

var file_internal_domain_policy_rules_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_domain_policy_rules_proto_goTypes = []interface{}{
	(*RewriteRule)(nil),     // 0: internal.domain.policy.RewriteRule
	(*This)(nil),            // 1: internal.domain.policy.This
	(*ComputedUserset)(nil), // 2: internal.domain.policy.ComputedUserset
	(*TupleToUserset)(nil),  // 3: internal.domain.policy.TupleToUserset
}
var file_internal_domain_policy_rules_proto_depIdxs = []int32{
	1, // 0: internal.domain.policy.RewriteRule.this:type_name -> internal.domain.policy.This
	2, // 1: internal.domain.policy.RewriteRule.computed_userset:type_name -> internal.domain.policy.ComputedUserset
	3, // 2: internal.domain.policy.RewriteRule.tuple_to_userset:type_name -> internal.domain.policy.TupleToUserset
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_domain_policy_rules_proto_init() }
func file_internal_domain_policy_rules_proto_init() {
	if File_internal_domain_policy_rules_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_domain_policy_rules_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewriteRule); i {
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
		file_internal_domain_policy_rules_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_domain_policy_rules_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_domain_policy_rules_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	file_internal_domain_policy_rules_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RewriteRule_This)(nil),
		(*RewriteRule_ComputedUserset)(nil),
		(*RewriteRule_TupleToUserset)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_domain_policy_rules_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_domain_policy_rules_proto_goTypes,
		DependencyIndexes: file_internal_domain_policy_rules_proto_depIdxs,
		MessageInfos:      file_internal_domain_policy_rules_proto_msgTypes,
	}.Build()
	File_internal_domain_policy_rules_proto = out.File
	file_internal_domain_policy_rules_proto_rawDesc = nil
	file_internal_domain_policy_rules_proto_goTypes = nil
	file_internal_domain_policy_rules_proto_depIdxs = nil
}
