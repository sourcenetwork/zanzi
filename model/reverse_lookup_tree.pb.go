// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: proto/reverse_lookup_tree.proto

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

// ReasonReason represents the origin of an edge, explaining why it was included.
// Reasons are used to explain whether an edge was found via direct lookup or
// was added due to an userset rewrite rule
type SourceReason int32

const (
	SourceReason_DIRECT           SourceReason = 0 // Direct relation stored in the database
	SourceReason_COMPUTED_USERSET SourceReason = 1
	SourceReason_TUPLE_TO_USERSET SourceReason = 2
)

// Enum value maps for SourceReason.
var (
	SourceReason_name = map[int32]string{
		0: "DIRECT",
		1: "COMPUTED_USERSET",
		2: "TUPLE_TO_USERSET",
	}
	SourceReason_value = map[string]int32{
		"DIRECT":           0,
		"COMPUTED_USERSET": 1,
		"TUPLE_TO_USERSET": 2,
	}
)

func (x SourceReason) Enum() *SourceReason {
	p := new(SourceReason)
	*p = x
	return p
}

func (x SourceReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SourceReason) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_reverse_lookup_tree_proto_enumTypes[0].Descriptor()
}

func (SourceReason) Type() protoreflect.EnumType {
	return &file_proto_reverse_lookup_tree_proto_enumTypes[0]
}

func (x SourceReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SourceReason.Descriptor instead.
func (SourceReason) EnumDescriptor() ([]byte, []int) {
	return file_proto_reverse_lookup_tree_proto_rawDescGZIP(), []int{0}
}

// Node represents the result from a reverse lookup call.
// Reverse Lookup performs a BFS search through the graph, which yields a BFS tree.
// Node contains the (object, relation) pair as the node data and a list of edges.
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *Userset `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Edges []*Edge  `protobuf:"bytes,2,rep,name=edges,proto3" json:"edges,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reverse_lookup_tree_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reverse_lookup_tree_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_reverse_lookup_tree_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetData() *Userset {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Node) GetEdges() []*Edge {
	if x != nil {
		return x.Edges
	}
	return nil
}

// Edge represents a connection from a a node to another.
// It contains the reason the edge was included in the Tree,
// as well as an explanatory string with aditional details.
type Edge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node    *Node        `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Reason  SourceReason `protobuf:"varint,2,opt,name=reason,proto3,enum=tree.SourceReason" json:"reason,omitempty"`
	Details string       `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Edge) Reset() {
	*x = Edge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reverse_lookup_tree_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Edge) ProtoMessage() {}

func (x *Edge) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reverse_lookup_tree_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Edge.ProtoReflect.Descriptor instead.
func (*Edge) Descriptor() ([]byte, []int) {
	return file_proto_reverse_lookup_tree_proto_rawDescGZIP(), []int{1}
}

func (x *Edge) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *Edge) GetReason() SourceReason {
	if x != nil {
		return x.Reason
	}
	return SourceReason_DIRECT
}

func (x *Edge) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

var File_proto_reverse_lookup_tree_proto protoreflect.FileDescriptor

var file_proto_reverse_lookup_tree_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f,
	0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x72, 0x65, 0x65, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x04, 0x4e,
	0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x65, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x05, 0x65, 0x64, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x45,
	0x64, 0x67, 0x65, 0x52, 0x05, 0x65, 0x64, 0x67, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x04, 0x45, 0x64,
	0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2a, 0x46, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x49, 0x52, 0x45,
	0x43, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x44,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x55,
	0x50, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x45, 0x54, 0x10, 0x02,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2d, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x62, 0x61, 0x72, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reverse_lookup_tree_proto_rawDescOnce sync.Once
	file_proto_reverse_lookup_tree_proto_rawDescData = file_proto_reverse_lookup_tree_proto_rawDesc
)

func file_proto_reverse_lookup_tree_proto_rawDescGZIP() []byte {
	file_proto_reverse_lookup_tree_proto_rawDescOnce.Do(func() {
		file_proto_reverse_lookup_tree_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reverse_lookup_tree_proto_rawDescData)
	})
	return file_proto_reverse_lookup_tree_proto_rawDescData
}

var file_proto_reverse_lookup_tree_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_reverse_lookup_tree_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_reverse_lookup_tree_proto_goTypes = []interface{}{
	(SourceReason)(0), // 0: tree.SourceReason
	(*Node)(nil),      // 1: tree.Node
	(*Edge)(nil),      // 2: tree.Edge
	(*Userset)(nil),   // 3: tuples.Userset
}
var file_proto_reverse_lookup_tree_proto_depIdxs = []int32{
	3, // 0: tree.Node.data:type_name -> tuples.Userset
	2, // 1: tree.Node.edges:type_name -> tree.Edge
	1, // 2: tree.Edge.node:type_name -> tree.Node
	0, // 3: tree.Edge.reason:type_name -> tree.SourceReason
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_reverse_lookup_tree_proto_init() }
func file_proto_reverse_lookup_tree_proto_init() {
	if File_proto_reverse_lookup_tree_proto != nil {
		return
	}
	file_proto_tuples_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_reverse_lookup_tree_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_proto_reverse_lookup_tree_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Edge); i {
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
			RawDescriptor: file_proto_reverse_lookup_tree_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_reverse_lookup_tree_proto_goTypes,
		DependencyIndexes: file_proto_reverse_lookup_tree_proto_depIdxs,
		EnumInfos:         file_proto_reverse_lookup_tree_proto_enumTypes,
		MessageInfos:      file_proto_reverse_lookup_tree_proto_msgTypes,
	}.Build()
	File_proto_reverse_lookup_tree_proto = out.File
	file_proto_reverse_lookup_tree_proto_rawDesc = nil
	file_proto_reverse_lookup_tree_proto_goTypes = nil
	file_proto_reverse_lookup_tree_proto_depIdxs = nil
}
