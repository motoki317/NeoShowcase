// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: neoshowcase/protobuf/statissite.proto

package pb

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

type ReloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReloadRequest) Reset() {
	*x = ReloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_statissite_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadRequest) ProtoMessage() {}

func (x *ReloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_statissite_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadRequest.ProtoReflect.Descriptor instead.
func (*ReloadRequest) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_statissite_proto_rawDescGZIP(), []int{0}
}

type ReloadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReloadResponse) Reset() {
	*x = ReloadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_protobuf_statissite_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadResponse) ProtoMessage() {}

func (x *ReloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_protobuf_statissite_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadResponse.ProtoReflect.Descriptor instead.
func (*ReloadResponse) Descriptor() ([]byte, []int) {
	return file_neoshowcase_protobuf_statissite_proto_rawDescGZIP(), []int{1}
}

var File_neoshowcase_protobuf_statissite_proto protoreflect.FileDescriptor

var file_neoshowcase_protobuf_statissite_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x73, 0x69, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x0f, 0x0a,
	0x0d, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10,
	0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x68, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53, 0x69, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x06, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x23, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x50, 0x74, 0x69, 0x74,
	0x65, 0x63, 0x68, 0x2f, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_neoshowcase_protobuf_statissite_proto_rawDescOnce sync.Once
	file_neoshowcase_protobuf_statissite_proto_rawDescData = file_neoshowcase_protobuf_statissite_proto_rawDesc
)

func file_neoshowcase_protobuf_statissite_proto_rawDescGZIP() []byte {
	file_neoshowcase_protobuf_statissite_proto_rawDescOnce.Do(func() {
		file_neoshowcase_protobuf_statissite_proto_rawDescData = protoimpl.X.CompressGZIP(file_neoshowcase_protobuf_statissite_proto_rawDescData)
	})
	return file_neoshowcase_protobuf_statissite_proto_rawDescData
}

var file_neoshowcase_protobuf_statissite_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_neoshowcase_protobuf_statissite_proto_goTypes = []interface{}{
	(*ReloadRequest)(nil),  // 0: neoshowcase.protobuf.ReloadRequest
	(*ReloadResponse)(nil), // 1: neoshowcase.protobuf.ReloadResponse
}
var file_neoshowcase_protobuf_statissite_proto_depIdxs = []int32{
	0, // 0: neoshowcase.protobuf.StaticSiteService.Reload:input_type -> neoshowcase.protobuf.ReloadRequest
	1, // 1: neoshowcase.protobuf.StaticSiteService.Reload:output_type -> neoshowcase.protobuf.ReloadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_neoshowcase_protobuf_statissite_proto_init() }
func file_neoshowcase_protobuf_statissite_proto_init() {
	if File_neoshowcase_protobuf_statissite_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_neoshowcase_protobuf_statissite_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadRequest); i {
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
		file_neoshowcase_protobuf_statissite_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadResponse); i {
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
			RawDescriptor: file_neoshowcase_protobuf_statissite_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_neoshowcase_protobuf_statissite_proto_goTypes,
		DependencyIndexes: file_neoshowcase_protobuf_statissite_proto_depIdxs,
		MessageInfos:      file_neoshowcase_protobuf_statissite_proto_msgTypes,
	}.Build()
	File_neoshowcase_protobuf_statissite_proto = out.File
	file_neoshowcase_protobuf_statissite_proto_rawDesc = nil
	file_neoshowcase_protobuf_statissite_proto_goTypes = nil
	file_neoshowcase_protobuf_statissite_proto_depIdxs = nil
}
