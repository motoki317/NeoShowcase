// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BuilderServiceClient is the client API for BuilderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuilderServiceClient interface {
	GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStatusResponse, error)
	BuildImage(ctx context.Context, in *BuildImageRequest, opts ...grpc.CallOption) (*BuildImageResponse, error)
}

type builderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuilderServiceClient(cc grpc.ClientConnInterface) BuilderServiceClient {
	return &builderServiceClient{cc}
}

func (c *builderServiceClient) GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/neoshowcase.proto.services.builder.BuilderService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderServiceClient) BuildImage(ctx context.Context, in *BuildImageRequest, opts ...grpc.CallOption) (*BuildImageResponse, error) {
	out := new(BuildImageResponse)
	err := c.cc.Invoke(ctx, "/neoshowcase.proto.services.builder.BuilderService/BuildImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuilderServiceServer is the server API for BuilderService service.
// All implementations must embed UnimplementedBuilderServiceServer
// for forward compatibility
type BuilderServiceServer interface {
	GetStatus(context.Context, *emptypb.Empty) (*GetStatusResponse, error)
	BuildImage(context.Context, *BuildImageRequest) (*BuildImageResponse, error)
	mustEmbedUnimplementedBuilderServiceServer()
}

// UnimplementedBuilderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBuilderServiceServer struct {
}

func (UnimplementedBuilderServiceServer) GetStatus(context.Context, *emptypb.Empty) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedBuilderServiceServer) BuildImage(context.Context, *BuildImageRequest) (*BuildImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildImage not implemented")
}
func (UnimplementedBuilderServiceServer) mustEmbedUnimplementedBuilderServiceServer() {}

// UnsafeBuilderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuilderServiceServer will
// result in compilation errors.
type UnsafeBuilderServiceServer interface {
	mustEmbedUnimplementedBuilderServiceServer()
}

func RegisterBuilderServiceServer(s *grpc.Server, srv BuilderServiceServer) {
	s.RegisterService(&_BuilderService_serviceDesc, srv)
}

func _BuilderService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neoshowcase.proto.services.builder.BuilderService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).GetStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderService_BuildImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).BuildImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neoshowcase.proto.services.builder.BuilderService/BuildImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).BuildImage(ctx, req.(*BuildImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuilderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neoshowcase.proto.services.builder.BuilderService",
	HandlerType: (*BuilderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _BuilderService_GetStatus_Handler,
		},
		{
			MethodName: "BuildImage",
			Handler:    _BuilderService_BuildImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "neoshowcase/services/builder/service.proto",
}
