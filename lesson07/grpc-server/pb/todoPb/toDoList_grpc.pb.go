// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todoPb

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

// ToDoListClient is the client API for ToDoList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ToDoListClient interface {
	CreateToDo(ctx context.Context, in *ToDoContent, opts ...grpc.CallOption) (*Result, error)
}

type toDoListClient struct {
	cc grpc.ClientConnInterface
}

func NewToDoListClient(cc grpc.ClientConnInterface) ToDoListClient {
	return &toDoListClient{cc}
}

func (c *toDoListClient) CreateToDo(ctx context.Context, in *ToDoContent, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ToDoList/CreateToDo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToDoListServer is the server API for ToDoList service.
// All implementations must embed UnimplementedToDoListServer
// for forward compatibility
type ToDoListServer interface {
	CreateToDo(context.Context, *ToDoContent) (*Result, error)
	mustEmbedUnimplementedToDoListServer()
}

// UnimplementedToDoListServer must be embedded to have forward compatible implementations.
type UnimplementedToDoListServer struct {
}

func (UnimplementedToDoListServer) CreateToDo(context.Context, *ToDoContent) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToDo not implemented")
}
func (UnimplementedToDoListServer) mustEmbedUnimplementedToDoListServer() {}

// UnsafeToDoListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ToDoListServer will
// result in compilation errors.
type UnsafeToDoListServer interface {
	mustEmbedUnimplementedToDoListServer()
}

func RegisterToDoListServer(s grpc.ServiceRegistrar, srv ToDoListServer) {
	s.RegisterService(&ToDoList_ServiceDesc, srv)
}

func _ToDoList_CreateToDo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToDoContent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoListServer).CreateToDo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ToDoList/CreateToDo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoListServer).CreateToDo(ctx, req.(*ToDoContent))
	}
	return interceptor(ctx, in, info, handler)
}

// ToDoList_ServiceDesc is the grpc.ServiceDesc for ToDoList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ToDoList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ToDoList",
	HandlerType: (*ToDoListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToDo",
			Handler:    _ToDoList_CreateToDo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/toDoList.proto",
}
