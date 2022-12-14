// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: demo/demo.proto

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

// DemoGatewayClient is the client API for DemoGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoGatewayClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type demoGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoGatewayClient(cc grpc.ClientConnInterface) DemoGatewayClient {
	return &demoGatewayClient{cc}
}

func (c *demoGatewayClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := c.cc.Invoke(ctx, "/demo.DemoGateway/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGatewayClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/demo.DemoGateway/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoGatewayClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/demo.DemoGateway/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoGatewayServer is the server API for DemoGateway service.
// All implementations must embed UnimplementedDemoGatewayServer
// for forward compatibility
type DemoGatewayServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedDemoGatewayServer()
}

// UnimplementedDemoGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedDemoGatewayServer struct {
}

func (UnimplementedDemoGatewayServer) Echo(context.Context, *StringMessage) (*StringMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedDemoGatewayServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedDemoGatewayServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDemoGatewayServer) mustEmbedUnimplementedDemoGatewayServer() {}

// UnsafeDemoGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoGatewayServer will
// result in compilation errors.
type UnsafeDemoGatewayServer interface {
	mustEmbedUnimplementedDemoGatewayServer()
}

func RegisterDemoGatewayServer(s grpc.ServiceRegistrar, srv DemoGatewayServer) {
	s.RegisterService(&DemoGateway_ServiceDesc, srv)
}

func _DemoGateway_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoGatewayServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoGateway/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoGatewayServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoGateway_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoGatewayServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoGateway/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoGatewayServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoGateway_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoGatewayServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoGateway/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoGatewayServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DemoGateway_ServiceDesc is the grpc.ServiceDesc for DemoGateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoGateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.DemoGateway",
	HandlerType: (*DemoGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _DemoGateway_Echo_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _DemoGateway_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _DemoGateway_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo/demo.proto",
}
