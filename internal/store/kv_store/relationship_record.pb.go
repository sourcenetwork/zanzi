// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: kv_store/relationship_record.proto

package kv_store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RelationType int32

const (
	RelationType_UNKNOWN      RelationType = 0
	RelationType_OBJECT       RelationType = 1
	RelationType_OBJECT_SET   RelationType = 2
	RelationType_RESOURCE_SET RelationType = 3
)

// Enum value maps for RelationType.
var (
	RelationType_name = map[int32]string{
		0: "UNKNOWN",
		1: "OBJECT",
		2: "OBJECT_SET",
		3: "RESOURCE_SET",
	}
	RelationType_value = map[string]int32{
		"UNKNOWN":      0,
		"OBJECT":       1,
		"OBJECT_SET":   2,
		"RESOURCE_SET": 3,
	}
)

func (x RelationType) Enum() *RelationType {
	p := new(RelationType)
	*p = x
	return p
}

func (x RelationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationType) Descriptor() protoreflect.EnumDescriptor {
	return file_kv_store_relationship_record_proto_enumTypes[0].Descriptor()
}

func (RelationType) Type() protoreflect.EnumType {
	return &file_kv_store_relationship_record_proto_enumTypes[0]
}

func (x RelationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationType.Descriptor instead.
func (RelationType) EnumDescriptor() ([]byte, []int) {
	return file_kv_store_relationship_record_proto_rawDescGZIP(), []int{0}
}

type RelationshipData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelationshipId []byte `protobuf:"bytes,1,opt,name=relationship_id,json=relationshipId,proto3" json:"relationship_id,omitempty"`
	// timestamp for creation time of relationship (based on the local system clock)
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// app_data opaque application data sent by the client
	AppData []byte `protobuf:"bytes,3,opt,name=app_data,json=appData,proto3" json:"app_data,omitempty"`
}

func (x *RelationshipData) Reset() {
	*x = RelationshipData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_store_relationship_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationshipData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationshipData) ProtoMessage() {}

func (x *RelationshipData) ProtoReflect() protoreflect.Message {
	mi := &file_kv_store_relationship_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationshipData.ProtoReflect.Descriptor instead.
func (*RelationshipData) Descriptor() ([]byte, []int) {
	return file_kv_store_relationship_record_proto_rawDescGZIP(), []int{0}
}

func (x *RelationshipData) GetRelationshipId() []byte {
	if x != nil {
		return x.RelationshipId
	}
	return nil
}

func (x *RelationshipData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RelationshipData) GetAppData() []byte {
	if x != nil {
		return x.AppData
	}
	return nil
}

type Relationship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *RelationNode `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Dest   *RelationNode `protobuf:"bytes,2,opt,name=dest,proto3" json:"dest,omitempty"`
}

func (x *Relationship) Reset() {
	*x = Relationship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_store_relationship_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relationship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relationship) ProtoMessage() {}

func (x *Relationship) ProtoReflect() protoreflect.Message {
	mi := &file_kv_store_relationship_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relationship.ProtoReflect.Descriptor instead.
func (*Relationship) Descriptor() ([]byte, []int) {
	return file_kv_store_relationship_record_proto_rawDescGZIP(), []int{1}
}

func (x *Relationship) GetSource() *RelationNode {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Relationship) GetDest() *RelationNode {
	if x != nil {
		return x.Dest
	}
	return nil
}

type RelationNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource string       `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Id       string       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Relation string       `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	Type     RelationType `protobuf:"varint,4,opt,name=type,proto3,enum=sourcenetwork.zanzi.internal.kv_store.RelationType" json:"type,omitempty"`
}

func (x *RelationNode) Reset() {
	*x = RelationNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_store_relationship_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationNode) ProtoMessage() {}

func (x *RelationNode) ProtoReflect() protoreflect.Message {
	mi := &file_kv_store_relationship_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationNode.ProtoReflect.Descriptor instead.
func (*RelationNode) Descriptor() ([]byte, []int) {
	return file_kv_store_relationship_record_proto_rawDescGZIP(), []int{2}
}

func (x *RelationNode) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *RelationNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RelationNode) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *RelationNode) GetType() RelationType {
	if x != nil {
		return x.Type
	}
	return RelationType_UNKNOWN
}

var File_kv_store_relationship_record_proto protoreflect.FileDescriptor

var file_kv_store_relationship_record_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a,
	0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61,
	0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69,
	0x70, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x33, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x04, 0x64, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a, 0x61, 0x6e,
	0x7a, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6b, 0x76, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x64, 0x65, 0x73, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x47, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33,
	0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x7a,
	0x61, 0x6e, 0x7a, 0x69, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x6b, 0x76,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x49, 0x0a, 0x0c, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x45, 0x54,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53,
	0x45, 0x54, 0x10, 0x03, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x7a, 0x61, 0x6e, 0x7a, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6b, 0x76, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kv_store_relationship_record_proto_rawDescOnce sync.Once
	file_kv_store_relationship_record_proto_rawDescData = file_kv_store_relationship_record_proto_rawDesc
)

func file_kv_store_relationship_record_proto_rawDescGZIP() []byte {
	file_kv_store_relationship_record_proto_rawDescOnce.Do(func() {
		file_kv_store_relationship_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_store_relationship_record_proto_rawDescData)
	})
	return file_kv_store_relationship_record_proto_rawDescData
}

var file_kv_store_relationship_record_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kv_store_relationship_record_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kv_store_relationship_record_proto_goTypes = []interface{}{
	(RelationType)(0),             // 0: sourcenetwork.zanzi.internal.kv_store.RelationType
	(*RelationshipData)(nil),      // 1: sourcenetwork.zanzi.internal.kv_store.RelationshipData
	(*Relationship)(nil),          // 2: sourcenetwork.zanzi.internal.kv_store.Relationship
	(*RelationNode)(nil),          // 3: sourcenetwork.zanzi.internal.kv_store.RelationNode
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_kv_store_relationship_record_proto_depIdxs = []int32{
	4, // 0: sourcenetwork.zanzi.internal.kv_store.RelationshipData.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: sourcenetwork.zanzi.internal.kv_store.Relationship.source:type_name -> sourcenetwork.zanzi.internal.kv_store.RelationNode
	3, // 2: sourcenetwork.zanzi.internal.kv_store.Relationship.dest:type_name -> sourcenetwork.zanzi.internal.kv_store.RelationNode
	0, // 3: sourcenetwork.zanzi.internal.kv_store.RelationNode.type:type_name -> sourcenetwork.zanzi.internal.kv_store.RelationType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kv_store_relationship_record_proto_init() }
func file_kv_store_relationship_record_proto_init() {
	if File_kv_store_relationship_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_store_relationship_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationshipData); i {
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
		file_kv_store_relationship_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relationship); i {
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
		file_kv_store_relationship_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationNode); i {
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
			RawDescriptor: file_kv_store_relationship_record_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kv_store_relationship_record_proto_goTypes,
		DependencyIndexes: file_kv_store_relationship_record_proto_depIdxs,
		EnumInfos:         file_kv_store_relationship_record_proto_enumTypes,
		MessageInfos:      file_kv_store_relationship_record_proto_msgTypes,
	}.Build()
	File_kv_store_relationship_record_proto = out.File
	file_kv_store_relationship_record_proto_rawDesc = nil
	file_kv_store_relationship_record_proto_goTypes = nil
	file_kv_store_relationship_record_proto_depIdxs = nil
}
