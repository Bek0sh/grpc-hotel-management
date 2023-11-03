// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: room_type.proto

package pb

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

// RoomTypeServiceClient is the client API for RoomTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomTypeServiceClient interface {
	CreateRoomType(ctx context.Context, in *CreateRoomTypeReq, opts ...grpc.CallOption) (*CreateRoomTypeRes, error)
	DeleteRoomType(ctx context.Context, in *DeleteRoomTypeReq, opts ...grpc.CallOption) (*Empty, error)
	UpdateRoomType(ctx context.Context, in *UpdateRoomTypeReq, opts ...grpc.CallOption) (*Empty, error)
	GetAllRoomTypes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllRoomTypesRes, error)
	GetByRoomType(ctx context.Context, in *GetByRoomTypeReq, opts ...grpc.CallOption) (*GetByRoomTypeRes, error)
}

type roomTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomTypeServiceClient(cc grpc.ClientConnInterface) RoomTypeServiceClient {
	return &roomTypeServiceClient{cc}
}

func (c *roomTypeServiceClient) CreateRoomType(ctx context.Context, in *CreateRoomTypeReq, opts ...grpc.CallOption) (*CreateRoomTypeRes, error) {
	out := new(CreateRoomTypeRes)
	err := c.cc.Invoke(ctx, "/proto.RoomTypeService/CreateRoomType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeServiceClient) DeleteRoomType(ctx context.Context, in *DeleteRoomTypeReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.RoomTypeService/DeleteRoomType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeServiceClient) UpdateRoomType(ctx context.Context, in *UpdateRoomTypeReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.RoomTypeService/UpdateRoomType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeServiceClient) GetAllRoomTypes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetAllRoomTypesRes, error) {
	out := new(GetAllRoomTypesRes)
	err := c.cc.Invoke(ctx, "/proto.RoomTypeService/GetAllRoomTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomTypeServiceClient) GetByRoomType(ctx context.Context, in *GetByRoomTypeReq, opts ...grpc.CallOption) (*GetByRoomTypeRes, error) {
	out := new(GetByRoomTypeRes)
	err := c.cc.Invoke(ctx, "/proto.RoomTypeService/GetByRoomType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomTypeServiceServer is the server API for RoomTypeService service.
// All implementations must embed UnimplementedRoomTypeServiceServer
// for forward compatibility
type RoomTypeServiceServer interface {
	CreateRoomType(context.Context, *CreateRoomTypeReq) (*CreateRoomTypeRes, error)
	DeleteRoomType(context.Context, *DeleteRoomTypeReq) (*Empty, error)
	UpdateRoomType(context.Context, *UpdateRoomTypeReq) (*Empty, error)
	GetAllRoomTypes(context.Context, *Empty) (*GetAllRoomTypesRes, error)
	GetByRoomType(context.Context, *GetByRoomTypeReq) (*GetByRoomTypeRes, error)
	mustEmbedUnimplementedRoomTypeServiceServer()
}

// UnimplementedRoomTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoomTypeServiceServer struct {
}

func (UnimplementedRoomTypeServiceServer) CreateRoomType(context.Context, *CreateRoomTypeReq) (*CreateRoomTypeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoomType not implemented")
}
func (UnimplementedRoomTypeServiceServer) DeleteRoomType(context.Context, *DeleteRoomTypeReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoomType not implemented")
}
func (UnimplementedRoomTypeServiceServer) UpdateRoomType(context.Context, *UpdateRoomTypeReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoomType not implemented")
}
func (UnimplementedRoomTypeServiceServer) GetAllRoomTypes(context.Context, *Empty) (*GetAllRoomTypesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoomTypes not implemented")
}
func (UnimplementedRoomTypeServiceServer) GetByRoomType(context.Context, *GetByRoomTypeReq) (*GetByRoomTypeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByRoomType not implemented")
}
func (UnimplementedRoomTypeServiceServer) mustEmbedUnimplementedRoomTypeServiceServer() {}

// UnsafeRoomTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomTypeServiceServer will
// result in compilation errors.
type UnsafeRoomTypeServiceServer interface {
	mustEmbedUnimplementedRoomTypeServiceServer()
}

func RegisterRoomTypeServiceServer(s grpc.ServiceRegistrar, srv RoomTypeServiceServer) {
	s.RegisterService(&RoomTypeService_ServiceDesc, srv)
}

func _RoomTypeService_CreateRoomType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTypeServiceServer).CreateRoomType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RoomTypeService/CreateRoomType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTypeServiceServer).CreateRoomType(ctx, req.(*CreateRoomTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTypeService_DeleteRoomType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTypeServiceServer).DeleteRoomType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RoomTypeService/DeleteRoomType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTypeServiceServer).DeleteRoomType(ctx, req.(*DeleteRoomTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTypeService_UpdateRoomType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTypeServiceServer).UpdateRoomType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RoomTypeService/UpdateRoomType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTypeServiceServer).UpdateRoomType(ctx, req.(*UpdateRoomTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTypeService_GetAllRoomTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTypeServiceServer).GetAllRoomTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RoomTypeService/GetAllRoomTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTypeServiceServer).GetAllRoomTypes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomTypeService_GetByRoomType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByRoomTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomTypeServiceServer).GetByRoomType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RoomTypeService/GetByRoomType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomTypeServiceServer).GetByRoomType(ctx, req.(*GetByRoomTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomTypeService_ServiceDesc is the grpc.ServiceDesc for RoomTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RoomTypeService",
	HandlerType: (*RoomTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoomType",
			Handler:    _RoomTypeService_CreateRoomType_Handler,
		},
		{
			MethodName: "DeleteRoomType",
			Handler:    _RoomTypeService_DeleteRoomType_Handler,
		},
		{
			MethodName: "UpdateRoomType",
			Handler:    _RoomTypeService_UpdateRoomType_Handler,
		},
		{
			MethodName: "GetAllRoomTypes",
			Handler:    _RoomTypeService_GetAllRoomTypes_Handler,
		},
		{
			MethodName: "GetByRoomType",
			Handler:    _RoomTypeService_GetByRoomType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room_type.proto",
}
