// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: neoshowcase/services/builder/service.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BuilderStatus int32

const (
	BuilderStatus_UNKNOWN     BuilderStatus = 0
	BuilderStatus_UNAVAILABLE BuilderStatus = 1
	BuilderStatus_WAITING     BuilderStatus = 2
	BuilderStatus_BUILDING    BuilderStatus = 3
)

// Enum value maps for BuilderStatus.
var (
	BuilderStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "UNAVAILABLE",
		2: "WAITING",
		3: "BUILDING",
	}
	BuilderStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"UNAVAILABLE": 1,
		"WAITING":     2,
		"BUILDING":    3,
	}
)

func (x BuilderStatus) Enum() *BuilderStatus {
	p := new(BuilderStatus)
	*p = x
	return p
}

func (x BuilderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuilderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_neoshowcase_services_builder_service_proto_enumTypes[0].Descriptor()
}

func (BuilderStatus) Type() protoreflect.EnumType {
	return &file_neoshowcase_services_builder_service_proto_enumTypes[0]
}

func (x BuilderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuilderStatus.Descriptor instead.
func (BuilderStatus) EnumDescriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{0}
}

type GetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status BuilderStatus `protobuf:"varint,1,opt,name=status,proto3,enum=neoshowcase.proto.services.builder.BuilderStatus" json:"status,omitempty"`
}

func (x *GetStatusResponse) Reset() {
	*x = GetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResponse) ProtoMessage() {}

func (x *GetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatusResponse) GetStatus() BuilderStatus {
	if x != nil {
		return x.Status
	}
	return BuilderStatus_UNKNOWN
}

type BuildImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageName string `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
}

func (x *BuildImageRequest) Reset() {
	*x = BuildImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildImageRequest) ProtoMessage() {}

func (x *BuildImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildImageRequest.ProtoReflect.Descriptor instead.
func (*BuildImageRequest) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{1}
}

func (x *BuildImageRequest) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

type BuildImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BuildImageResponse) Reset() {
	*x = BuildImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neoshowcase_services_builder_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildImageResponse) ProtoMessage() {}

func (x *BuildImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neoshowcase_services_builder_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildImageResponse.ProtoReflect.Descriptor instead.
func (*BuildImageResponse) Descriptor() ([]byte, []int) {
	return file_neoshowcase_services_builder_service_proto_rawDescGZIP(), []int{2}
}

var File_neoshowcase_services_builder_service_proto protoreflect.FileDescriptor

var file_neoshowcase_services_builder_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6e, 0x65,
	0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x31, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x32, 0x0a,
	0x11, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x14, 0x0a, 0x12, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x48, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x32, 0xe9, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x35, 0x2e, 0x6e, 0x65, 0x6f, 0x73,
	0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x7b, 0x0a, 0x0a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x35,
	0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x50,
	0x74, 0x69, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x6e, 0x65, 0x6f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61,
	0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_neoshowcase_services_builder_service_proto_rawDescOnce sync.Once
	file_neoshowcase_services_builder_service_proto_rawDescData = file_neoshowcase_services_builder_service_proto_rawDesc
)

func file_neoshowcase_services_builder_service_proto_rawDescGZIP() []byte {
	file_neoshowcase_services_builder_service_proto_rawDescOnce.Do(func() {
		file_neoshowcase_services_builder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_neoshowcase_services_builder_service_proto_rawDescData)
	})
	return file_neoshowcase_services_builder_service_proto_rawDescData
}

var file_neoshowcase_services_builder_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_neoshowcase_services_builder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_neoshowcase_services_builder_service_proto_goTypes = []interface{}{
	(BuilderStatus)(0),         // 0: neoshowcase.proto.services.builder.BuilderStatus
	(*GetStatusResponse)(nil),  // 1: neoshowcase.proto.services.builder.GetStatusResponse
	(*BuildImageRequest)(nil),  // 2: neoshowcase.proto.services.builder.BuildImageRequest
	(*BuildImageResponse)(nil), // 3: neoshowcase.proto.services.builder.BuildImageResponse
	(*emptypb.Empty)(nil),      // 4: google.protobuf.Empty
}
var file_neoshowcase_services_builder_service_proto_depIdxs = []int32{
	0, // 0: neoshowcase.proto.services.builder.GetStatusResponse.status:type_name -> neoshowcase.proto.services.builder.BuilderStatus
	4, // 1: neoshowcase.proto.services.builder.BuilderService.GetStatus:input_type -> google.protobuf.Empty
	2, // 2: neoshowcase.proto.services.builder.BuilderService.BuildImage:input_type -> neoshowcase.proto.services.builder.BuildImageRequest
	1, // 3: neoshowcase.proto.services.builder.BuilderService.GetStatus:output_type -> neoshowcase.proto.services.builder.GetStatusResponse
	3, // 4: neoshowcase.proto.services.builder.BuilderService.BuildImage:output_type -> neoshowcase.proto.services.builder.BuildImageResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_neoshowcase_services_builder_service_proto_init() }
func file_neoshowcase_services_builder_service_proto_init() {
	if File_neoshowcase_services_builder_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_neoshowcase_services_builder_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResponse); i {
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
		file_neoshowcase_services_builder_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildImageRequest); i {
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
		file_neoshowcase_services_builder_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildImageResponse); i {
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
			RawDescriptor: file_neoshowcase_services_builder_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_neoshowcase_services_builder_service_proto_goTypes,
		DependencyIndexes: file_neoshowcase_services_builder_service_proto_depIdxs,
		EnumInfos:         file_neoshowcase_services_builder_service_proto_enumTypes,
		MessageInfos:      file_neoshowcase_services_builder_service_proto_msgTypes,
	}.Build()
	File_neoshowcase_services_builder_service_proto = out.File
	file_neoshowcase_services_builder_service_proto_rawDesc = nil
	file_neoshowcase_services_builder_service_proto_goTypes = nil
	file_neoshowcase_services_builder_service_proto_depIdxs = nil
}
