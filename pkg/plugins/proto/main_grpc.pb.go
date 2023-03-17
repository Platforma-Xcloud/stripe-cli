// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pkg/plugins/proto/main.proto

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

// MainClient is the client API for Main service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainClient interface {
	RunCommand(ctx context.Context, in *RunCommandRequest, opts ...grpc.CallOption) (*RunCommandResponse, error)
}

type mainClient struct {
	cc grpc.ClientConnInterface
}

func NewMainClient(cc grpc.ClientConnInterface) MainClient {
	return &mainClient{cc}
}

func (c *mainClient) RunCommand(ctx context.Context, in *RunCommandRequest, opts ...grpc.CallOption) (*RunCommandResponse, error) {
	out := new(RunCommandResponse)
	err := c.cc.Invoke(ctx, "/proto.Main/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServer is the server API for Main service.
// All implementations must embed UnimplementedMainServer
// for forward compatibility
type MainServer interface {
	RunCommand(context.Context, *RunCommandRequest) (*RunCommandResponse, error)
	mustEmbedUnimplementedMainServer()
}

// UnimplementedMainServer must be embedded to have forward compatible implementations.
type UnimplementedMainServer struct {
}

func (UnimplementedMainServer) RunCommand(context.Context, *RunCommandRequest) (*RunCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}
func (UnimplementedMainServer) mustEmbedUnimplementedMainServer() {}

// UnsafeMainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServer will
// result in compilation errors.
type UnsafeMainServer interface {
	mustEmbedUnimplementedMainServer()
}

func RegisterMainServer(s grpc.ServiceRegistrar, srv MainServer) {
	s.RegisterService(&Main_ServiceDesc, srv)
}

func _Main_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Main/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).RunCommand(ctx, req.(*RunCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Main_ServiceDesc is the grpc.ServiceDesc for Main service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Main_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Main",
	HandlerType: (*MainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _Main_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/plugins/proto/main.proto",
}