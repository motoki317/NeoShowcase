// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: neoshowcase/protobuf/builder.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BuilderServiceClient is the client API for BuilderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuilderServiceClient interface {
	GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStatusResponse, error)
	ConnectEventStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (BuilderService_ConnectEventStreamClient, error)
	StartBuildImage(ctx context.Context, in *StartBuildImageRequest, opts ...grpc.CallOption) (*StartBuildImageResponse, error)
	StartBuildStatic(ctx context.Context, in *StartBuildStaticRequest, opts ...grpc.CallOption) (*StartBuildStaticResponse, error)
	CancelTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CancelTaskResponse, error)
}

type builderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBuilderServiceClient(cc grpc.ClientConnInterface) BuilderServiceClient {
	return &builderServiceClient{cc}
}

func (c *builderServiceClient) GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/neoshowcase.protobuf.BuilderService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderServiceClient) ConnectEventStream(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (BuilderService_ConnectEventStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &BuilderService_ServiceDesc.Streams[0], "/neoshowcase.protobuf.BuilderService/ConnectEventStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &builderServiceConnectEventStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BuilderService_ConnectEventStreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type builderServiceConnectEventStreamClient struct {
	grpc.ClientStream
}

func (x *builderServiceConnectEventStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *builderServiceClient) StartBuildImage(ctx context.Context, in *StartBuildImageRequest, opts ...grpc.CallOption) (*StartBuildImageResponse, error) {
	out := new(StartBuildImageResponse)
	err := c.cc.Invoke(ctx, "/neoshowcase.protobuf.BuilderService/StartBuildImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderServiceClient) StartBuildStatic(ctx context.Context, in *StartBuildStaticRequest, opts ...grpc.CallOption) (*StartBuildStaticResponse, error) {
	out := new(StartBuildStaticResponse)
	err := c.cc.Invoke(ctx, "/neoshowcase.protobuf.BuilderService/StartBuildStatic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderServiceClient) CancelTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CancelTaskResponse, error) {
	out := new(CancelTaskResponse)
	err := c.cc.Invoke(ctx, "/neoshowcase.protobuf.BuilderService/CancelTask", in, out, opts...)
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
	ConnectEventStream(*emptypb.Empty, BuilderService_ConnectEventStreamServer) error
	StartBuildImage(context.Context, *StartBuildImageRequest) (*StartBuildImageResponse, error)
	StartBuildStatic(context.Context, *StartBuildStaticRequest) (*StartBuildStaticResponse, error)
	CancelTask(context.Context, *emptypb.Empty) (*CancelTaskResponse, error)
	mustEmbedUnimplementedBuilderServiceServer()
}

// UnimplementedBuilderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBuilderServiceServer struct {
}

func (UnimplementedBuilderServiceServer) GetStatus(context.Context, *emptypb.Empty) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedBuilderServiceServer) ConnectEventStream(*emptypb.Empty, BuilderService_ConnectEventStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectEventStream not implemented")
}
func (UnimplementedBuilderServiceServer) StartBuildImage(context.Context, *StartBuildImageRequest) (*StartBuildImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBuildImage not implemented")
}
func (UnimplementedBuilderServiceServer) StartBuildStatic(context.Context, *StartBuildStaticRequest) (*StartBuildStaticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBuildStatic not implemented")
}
func (UnimplementedBuilderServiceServer) CancelTask(context.Context, *emptypb.Empty) (*CancelTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTask not implemented")
}
func (UnimplementedBuilderServiceServer) mustEmbedUnimplementedBuilderServiceServer() {}

// UnsafeBuilderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuilderServiceServer will
// result in compilation errors.
type UnsafeBuilderServiceServer interface {
	mustEmbedUnimplementedBuilderServiceServer()
}

func RegisterBuilderServiceServer(s grpc.ServiceRegistrar, srv BuilderServiceServer) {
	s.RegisterService(&BuilderService_ServiceDesc, srv)
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
		FullMethod: "/neoshowcase.protobuf.BuilderService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).GetStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderService_ConnectEventStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuilderServiceServer).ConnectEventStream(m, &builderServiceConnectEventStreamServer{stream})
}

type BuilderService_ConnectEventStreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type builderServiceConnectEventStreamServer struct {
	grpc.ServerStream
}

func (x *builderServiceConnectEventStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _BuilderService_StartBuildImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBuildImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).StartBuildImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neoshowcase.protobuf.BuilderService/StartBuildImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).StartBuildImage(ctx, req.(*StartBuildImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderService_StartBuildStatic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBuildStaticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).StartBuildStatic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neoshowcase.protobuf.BuilderService/StartBuildStatic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).StartBuildStatic(ctx, req.(*StartBuildStaticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BuilderService_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServiceServer).CancelTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neoshowcase.protobuf.BuilderService/CancelTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServiceServer).CancelTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BuilderService_ServiceDesc is the grpc.ServiceDesc for BuilderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuilderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "neoshowcase.protobuf.BuilderService",
	HandlerType: (*BuilderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _BuilderService_GetStatus_Handler,
		},
		{
			MethodName: "StartBuildImage",
			Handler:    _BuilderService_StartBuildImage_Handler,
		},
		{
			MethodName: "StartBuildStatic",
			Handler:    _BuilderService_StartBuildStatic_Handler,
		},
		{
			MethodName: "CancelTask",
			Handler:    _BuilderService_CancelTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectEventStream",
			Handler:       _BuilderService_ConnectEventStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "neoshowcase/protobuf/builder.proto",
}
