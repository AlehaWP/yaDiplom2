// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: gophepass.proto

package proto

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

// GophePassClient is the client API for GophePass service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophePassClient interface {
	AddFile(ctx context.Context, in *AddFileRequest, opts ...grpc.CallOption) (*AddResponse, error)
	AddAcc(ctx context.Context, in *AddAccRequest, opts ...grpc.CallOption) (*AddResponse, error)
	AddCard(ctx context.Context, in *AddCardRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type gophePassClient struct {
	cc grpc.ClientConnInterface
}

func NewGophePassClient(cc grpc.ClientConnInterface) GophePassClient {
	return &gophePassClient{cc}
}

func (c *gophePassClient) AddFile(ctx context.Context, in *AddFileRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.GophePass/AddFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophePassClient) AddAcc(ctx context.Context, in *AddAccRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.GophePass/AddAcc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophePassClient) AddCard(ctx context.Context, in *AddCardRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.GophePass/AddCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophePassServer is the server API for GophePass service.
// All implementations must embed UnimplementedGophePassServer
// for forward compatibility
type GophePassServer interface {
	AddFile(context.Context, *AddFileRequest) (*AddResponse, error)
	AddAcc(context.Context, *AddAccRequest) (*AddResponse, error)
	AddCard(context.Context, *AddCardRequest) (*AddResponse, error)
	mustEmbedUnimplementedGophePassServer()
}

// UnimplementedGophePassServer must be embedded to have forward compatible implementations.
type UnimplementedGophePassServer struct {
}

func (UnimplementedGophePassServer) AddFile(context.Context, *AddFileRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFile not implemented")
}
func (UnimplementedGophePassServer) AddAcc(context.Context, *AddAccRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAcc not implemented")
}
func (UnimplementedGophePassServer) AddCard(context.Context, *AddCardRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCard not implemented")
}
func (UnimplementedGophePassServer) mustEmbedUnimplementedGophePassServer() {}

// UnsafeGophePassServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophePassServer will
// result in compilation errors.
type UnsafeGophePassServer interface {
	mustEmbedUnimplementedGophePassServer()
}

func RegisterGophePassServer(s grpc.ServiceRegistrar, srv GophePassServer) {
	s.RegisterService(&GophePass_ServiceDesc, srv)
}

func _GophePass_AddFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophePassServer).AddFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GophePass/AddFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophePassServer).AddFile(ctx, req.(*AddFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophePass_AddAcc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAccRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophePassServer).AddAcc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GophePass/AddAcc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophePassServer).AddAcc(ctx, req.(*AddAccRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophePass_AddCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophePassServer).AddCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GophePass/AddCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophePassServer).AddCard(ctx, req.(*AddCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GophePass_ServiceDesc is the grpc.ServiceDesc for GophePass service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophePass_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GophePass",
	HandlerType: (*GophePassServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFile",
			Handler:    _GophePass_AddFile_Handler,
		},
		{
			MethodName: "AddAcc",
			Handler:    _GophePass_AddAcc_Handler,
		},
		{
			MethodName: "AddCard",
			Handler:    _GophePass_AddCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gophepass.proto",
}
