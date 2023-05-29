// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: proto/demo/demo.proto

package demo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DemoService_SayHello_FullMethodName = "/demo.DemoService/SayHello"
)

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoServiceClient interface {
	SayHello(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error)
}

type demoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoServiceClient(cc grpc.ClientConnInterface) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) SayHello(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error) {
	out := new(DemoResponse)
	err := c.cc.Invoke(ctx, DemoService_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
// All implementations must embed UnimplementedDemoServiceServer
// for forward compatibility
type DemoServiceServer interface {
	SayHello(context.Context, *DemoRequest) (*DemoResponse, error)
	mustEmbedUnimplementedDemoServiceServer()
}

// UnimplementedDemoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (UnimplementedDemoServiceServer) SayHello(context.Context, *DemoRequest) (*DemoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedDemoServiceServer) mustEmbedUnimplementedDemoServiceServer() {}

// UnsafeDemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServiceServer will
// result in compilation errors.
type UnsafeDemoServiceServer interface {
	mustEmbedUnimplementedDemoServiceServer()
}

func RegisterDemoServiceServer(s grpc.ServiceRegistrar, srv DemoServiceServer) {
	s.RegisterService(&DemoService_ServiceDesc, srv)
}

func _DemoService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DemoService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).SayHello(ctx, req.(*DemoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DemoService_ServiceDesc is the grpc.ServiceDesc for DemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _DemoService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/demo/demo.proto",
}
