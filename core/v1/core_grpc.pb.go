// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: core/v1/core.proto

package core

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

// PasteServiceClient is the client API for PasteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PasteServiceClient interface {
	List(ctx context.Context, in *ListPasteRequest, opts ...grpc.CallOption) (*ListPasteReply, error)
	Create(ctx context.Context, in *CreatePasteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *ByIDPasteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Put(ctx context.Context, in *PutPasteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *ByIDPasteRequest, opts ...grpc.CallOption) (*Paste, error)
}

type pasteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPasteServiceClient(cc grpc.ClientConnInterface) PasteServiceClient {
	return &pasteServiceClient{cc}
}

func (c *pasteServiceClient) List(ctx context.Context, in *ListPasteRequest, opts ...grpc.CallOption) (*ListPasteReply, error) {
	out := new(ListPasteReply)
	err := c.cc.Invoke(ctx, "/PasteService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pasteServiceClient) Create(ctx context.Context, in *CreatePasteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/PasteService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pasteServiceClient) Delete(ctx context.Context, in *ByIDPasteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/PasteService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pasteServiceClient) Put(ctx context.Context, in *PutPasteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/PasteService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pasteServiceClient) Get(ctx context.Context, in *ByIDPasteRequest, opts ...grpc.CallOption) (*Paste, error) {
	out := new(Paste)
	err := c.cc.Invoke(ctx, "/PasteService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PasteServiceServer is the server API for PasteService service.
// All implementations must embed UnimplementedPasteServiceServer
// for forward compatibility
type PasteServiceServer interface {
	List(context.Context, *ListPasteRequest) (*ListPasteReply, error)
	Create(context.Context, *CreatePasteRequest) (*emptypb.Empty, error)
	Delete(context.Context, *ByIDPasteRequest) (*emptypb.Empty, error)
	Put(context.Context, *PutPasteRequest) (*emptypb.Empty, error)
	Get(context.Context, *ByIDPasteRequest) (*Paste, error)
	mustEmbedUnimplementedPasteServiceServer()
}

// UnimplementedPasteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPasteServiceServer struct {
}

func (UnimplementedPasteServiceServer) List(context.Context, *ListPasteRequest) (*ListPasteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPasteServiceServer) Create(context.Context, *CreatePasteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPasteServiceServer) Delete(context.Context, *ByIDPasteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPasteServiceServer) Put(context.Context, *PutPasteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedPasteServiceServer) Get(context.Context, *ByIDPasteRequest) (*Paste, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPasteServiceServer) mustEmbedUnimplementedPasteServiceServer() {}

// UnsafePasteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PasteServiceServer will
// result in compilation errors.
type UnsafePasteServiceServer interface {
	mustEmbedUnimplementedPasteServiceServer()
}

func RegisterPasteServiceServer(s grpc.ServiceRegistrar, srv PasteServiceServer) {
	s.RegisterService(&PasteService_ServiceDesc, srv)
}

func _PasteService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPasteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasteServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PasteService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasteServiceServer).List(ctx, req.(*ListPasteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PasteService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePasteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasteServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PasteService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasteServiceServer).Create(ctx, req.(*CreatePasteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PasteService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIDPasteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasteServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PasteService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasteServiceServer).Delete(ctx, req.(*ByIDPasteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PasteService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPasteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasteServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PasteService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasteServiceServer).Put(ctx, req.(*PutPasteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PasteService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIDPasteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PasteServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PasteService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PasteServiceServer).Get(ctx, req.(*ByIDPasteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PasteService_ServiceDesc is the grpc.ServiceDesc for PasteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PasteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PasteService",
	HandlerType: (*PasteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PasteService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PasteService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PasteService_Delete_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _PasteService_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PasteService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/v1/core.proto",
}
