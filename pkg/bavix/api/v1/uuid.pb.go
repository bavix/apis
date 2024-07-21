// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: bavix/api/v1/uuid.proto

package v1

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

// UUID is a message that represents a UUID.
//
// A UUID is composed of two 64-bit values, each representing a part of the UUID.
// The high part is stored in the high field, and the low part is stored in the
// low field.
type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The high part of the UUID.
	//
	// This field contains the high part of the UUID.
	// It is a 64-bit value that is independent of the low part.
	High int64 `protobuf:"varint,1,opt,name=high,proto3" json:"high,omitempty"`
	// The low part of the UUID.
	//
	// This field contains the low part of the UUID.
	// It is a 64-bit value that is independent of the high part.
	Low int64 `protobuf:"varint,2,opt,name=low,proto3" json:"low,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bavix_api_v1_uuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_bavix_api_v1_uuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_bavix_api_v1_uuid_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetHigh() int64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *UUID) GetLow() int64 {
	if x != nil {
		return x.Low
	}
	return 0
}

var File_bavix_api_v1_uuid_proto protoreflect.FileDescriptor

var file_bavix_api_v1_uuid_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x61, 0x76, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x76, 0x69, 0x78,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x2c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x68,
	0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x6c, 0x6f, 0x77, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x76, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x62, 0x61, 0x76, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bavix_api_v1_uuid_proto_rawDescOnce sync.Once
	file_bavix_api_v1_uuid_proto_rawDescData = file_bavix_api_v1_uuid_proto_rawDesc
)

func file_bavix_api_v1_uuid_proto_rawDescGZIP() []byte {
	file_bavix_api_v1_uuid_proto_rawDescOnce.Do(func() {
		file_bavix_api_v1_uuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_bavix_api_v1_uuid_proto_rawDescData)
	})
	return file_bavix_api_v1_uuid_proto_rawDescData
}

var file_bavix_api_v1_uuid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bavix_api_v1_uuid_proto_goTypes = []any{
	(*UUID)(nil), // 0: bavix.api.v1.UUID
}
var file_bavix_api_v1_uuid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bavix_api_v1_uuid_proto_init() }
func file_bavix_api_v1_uuid_proto_init() {
	if File_bavix_api_v1_uuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bavix_api_v1_uuid_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UUID); i {
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
			RawDescriptor: file_bavix_api_v1_uuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bavix_api_v1_uuid_proto_goTypes,
		DependencyIndexes: file_bavix_api_v1_uuid_proto_depIdxs,
		MessageInfos:      file_bavix_api_v1_uuid_proto_msgTypes,
	}.Build()
	File_bavix_api_v1_uuid_proto = out.File
	file_bavix_api_v1_uuid_proto_rawDesc = nil
	file_bavix_api_v1_uuid_proto_goTypes = nil
	file_bavix_api_v1_uuid_proto_depIdxs = nil
}