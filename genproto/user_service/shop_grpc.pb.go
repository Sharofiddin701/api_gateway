// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: shop.proto

package user_service

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
	ShopService_Create_FullMethodName  = "/user_service.ShopService/Create"
	ShopService_GetByID_FullMethodName = "/user_service.ShopService/GetByID"
	ShopService_GetList_FullMethodName = "/user_service.ShopService/GetList"
	ShopService_Update_FullMethodName  = "/user_service.ShopService/Update"
	ShopService_Delete_FullMethodName  = "/user_service.ShopService/Delete"
)

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopServiceClient interface {
	Create(ctx context.Context, in *CreateShop, opts ...grpc.CallOption) (*Shop, error)
	GetByID(ctx context.Context, in *ShopPrimaryKey, opts ...grpc.CallOption) (*Shop, error)
	GetList(ctx context.Context, in *GetListShopRequest, opts ...grpc.CallOption) (*GetListShopResponse, error)
	Update(ctx context.Context, in *UpdateShop, opts ...grpc.CallOption) (*Shop, error)
	Delete(ctx context.Context, in *ShopPrimaryKey, opts ...grpc.CallOption) (*Empty4, error)
}

type shopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShopServiceClient(cc grpc.ClientConnInterface) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) Create(ctx context.Context, in *CreateShop, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, ShopService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetByID(ctx context.Context, in *ShopPrimaryKey, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, ShopService_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetList(ctx context.Context, in *GetListShopRequest, opts ...grpc.CallOption) (*GetListShopResponse, error) {
	out := new(GetListShopResponse)
	err := c.cc.Invoke(ctx, ShopService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) Update(ctx context.Context, in *UpdateShop, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, ShopService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) Delete(ctx context.Context, in *ShopPrimaryKey, opts ...grpc.CallOption) (*Empty4, error) {
	out := new(Empty4)
	err := c.cc.Invoke(ctx, ShopService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
// All implementations should embed UnimplementedShopServiceServer
// for forward compatibility
type ShopServiceServer interface {
	Create(context.Context, *CreateShop) (*Shop, error)
	GetByID(context.Context, *ShopPrimaryKey) (*Shop, error)
	GetList(context.Context, *GetListShopRequest) (*GetListShopResponse, error)
	Update(context.Context, *UpdateShop) (*Shop, error)
	Delete(context.Context, *ShopPrimaryKey) (*Empty4, error)
}

// UnimplementedShopServiceServer should be embedded to have forward compatible implementations.
type UnimplementedShopServiceServer struct {
}

func (UnimplementedShopServiceServer) Create(context.Context, *CreateShop) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedShopServiceServer) GetByID(context.Context, *ShopPrimaryKey) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedShopServiceServer) GetList(context.Context, *GetListShopRequest) (*GetListShopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedShopServiceServer) Update(context.Context, *UpdateShop) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedShopServiceServer) Delete(context.Context, *ShopPrimaryKey) (*Empty4, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServiceServer will
// result in compilation errors.
type UnsafeShopServiceServer interface {
	mustEmbedUnimplementedShopServiceServer()
}

func RegisterShopServiceServer(s grpc.ServiceRegistrar, srv ShopServiceServer) {
	s.RegisterService(&ShopService_ServiceDesc, srv)
}

func _ShopService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).Create(ctx, req.(*CreateShop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetByID(ctx, req.(*ShopPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetList(ctx, req.(*GetListShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).Update(ctx, req.(*UpdateShop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShopService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).Delete(ctx, req.(*ShopPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopService_ServiceDesc is the grpc.ServiceDesc for ShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ShopService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _ShopService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ShopService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ShopService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ShopService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}
