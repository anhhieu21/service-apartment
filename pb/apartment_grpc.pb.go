// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: api/proto/apartment.proto

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

// ApartmentServiceClient is the client API for ApartmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApartmentServiceClient interface {
	GetApartment(ctx context.Context, in *GetApartmentRequest, opts ...grpc.CallOption) (*GetApartmentResponse, error)
	CreateApartment(ctx context.Context, in *CreateApartmentRequest, opts ...grpc.CallOption) (*CreateApartmentResponse, error)
	UpdateApartment(ctx context.Context, in *UpdateApartmentRequest, opts ...grpc.CallOption) (*UpdateApartmentResponse, error)
	DeleteApartment(ctx context.Context, in *DeleteApartmentRequest, opts ...grpc.CallOption) (*DeleteApartmentResponse, error)
}

type apartmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApartmentServiceClient(cc grpc.ClientConnInterface) ApartmentServiceClient {
	return &apartmentServiceClient{cc}
}

func (c *apartmentServiceClient) GetApartment(ctx context.Context, in *GetApartmentRequest, opts ...grpc.CallOption) (*GetApartmentResponse, error) {
	out := new(GetApartmentResponse)
	err := c.cc.Invoke(ctx, "/proto.ApartmentService/GetApartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apartmentServiceClient) CreateApartment(ctx context.Context, in *CreateApartmentRequest, opts ...grpc.CallOption) (*CreateApartmentResponse, error) {
	out := new(CreateApartmentResponse)
	err := c.cc.Invoke(ctx, "/proto.ApartmentService/CreateApartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apartmentServiceClient) UpdateApartment(ctx context.Context, in *UpdateApartmentRequest, opts ...grpc.CallOption) (*UpdateApartmentResponse, error) {
	out := new(UpdateApartmentResponse)
	err := c.cc.Invoke(ctx, "/proto.ApartmentService/UpdateApartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apartmentServiceClient) DeleteApartment(ctx context.Context, in *DeleteApartmentRequest, opts ...grpc.CallOption) (*DeleteApartmentResponse, error) {
	out := new(DeleteApartmentResponse)
	err := c.cc.Invoke(ctx, "/proto.ApartmentService/DeleteApartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApartmentServiceServer is the server API for ApartmentService service.
// All implementations must embed UnimplementedApartmentServiceServer
// for forward compatibility
type ApartmentServiceServer interface {
	GetApartment(context.Context, *GetApartmentRequest) (*GetApartmentResponse, error)
	CreateApartment(context.Context, *CreateApartmentRequest) (*CreateApartmentResponse, error)
	UpdateApartment(context.Context, *UpdateApartmentRequest) (*UpdateApartmentResponse, error)
	DeleteApartment(context.Context, *DeleteApartmentRequest) (*DeleteApartmentResponse, error)
	mustEmbedUnimplementedApartmentServiceServer()
}

// UnimplementedApartmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedApartmentServiceServer struct {
}

func (UnimplementedApartmentServiceServer) GetApartment(context.Context, *GetApartmentRequest) (*GetApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApartment not implemented")
}
func (UnimplementedApartmentServiceServer) CreateApartment(context.Context, *CreateApartmentRequest) (*CreateApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApartment not implemented")
}
func (UnimplementedApartmentServiceServer) UpdateApartment(context.Context, *UpdateApartmentRequest) (*UpdateApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApartment not implemented")
}
func (UnimplementedApartmentServiceServer) DeleteApartment(context.Context, *DeleteApartmentRequest) (*DeleteApartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApartment not implemented")
}
func (UnimplementedApartmentServiceServer) mustEmbedUnimplementedApartmentServiceServer() {}

// UnsafeApartmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApartmentServiceServer will
// result in compilation errors.
type UnsafeApartmentServiceServer interface {
	mustEmbedUnimplementedApartmentServiceServer()
}

func RegisterApartmentServiceServer(s grpc.ServiceRegistrar, srv ApartmentServiceServer) {
	s.RegisterService(&ApartmentService_ServiceDesc, srv)
}

func _ApartmentService_GetApartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApartmentServiceServer).GetApartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ApartmentService/GetApartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApartmentServiceServer).GetApartment(ctx, req.(*GetApartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApartmentService_CreateApartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApartmentServiceServer).CreateApartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ApartmentService/CreateApartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApartmentServiceServer).CreateApartment(ctx, req.(*CreateApartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApartmentService_UpdateApartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApartmentServiceServer).UpdateApartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ApartmentService/UpdateApartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApartmentServiceServer).UpdateApartment(ctx, req.(*UpdateApartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApartmentService_DeleteApartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApartmentServiceServer).DeleteApartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ApartmentService/DeleteApartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApartmentServiceServer).DeleteApartment(ctx, req.(*DeleteApartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApartmentService_ServiceDesc is the grpc.ServiceDesc for ApartmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApartmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ApartmentService",
	HandlerType: (*ApartmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApartment",
			Handler:    _ApartmentService_GetApartment_Handler,
		},
		{
			MethodName: "CreateApartment",
			Handler:    _ApartmentService_CreateApartment_Handler,
		},
		{
			MethodName: "UpdateApartment",
			Handler:    _ApartmentService_UpdateApartment_Handler,
		},
		{
			MethodName: "DeleteApartment",
			Handler:    _ApartmentService_DeleteApartment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/apartment.proto",
}

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerResponse, error)
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerResponse, error)
	DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*DeleteCustomerResponse, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetCustomer(ctx context.Context, in *GetCustomerRequest, opts ...grpc.CallOption) (*GetCustomerResponse, error) {
	out := new(GetCustomerResponse)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/GetCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error) {
	out := new(CreateCustomerResponse)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerRequest, opts ...grpc.CallOption) (*UpdateCustomerResponse, error) {
	out := new(UpdateCustomerResponse)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/UpdateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerRequest, opts ...grpc.CallOption) (*DeleteCustomerResponse, error) {
	out := new(DeleteCustomerResponse)
	err := c.cc.Invoke(ctx, "/proto.CustomerService/DeleteCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerResponse, error)
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerResponse, error)
	UpdateCustomer(context.Context, *UpdateCustomerRequest) (*UpdateCustomerResponse, error)
	DeleteCustomer(context.Context, *DeleteCustomerRequest) (*DeleteCustomerResponse, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) GetCustomer(context.Context, *GetCustomerRequest) (*GetCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateCustomer(context.Context, *UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) DeleteCustomer(context.Context, *DeleteCustomerRequest) (*DeleteCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/GetCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetCustomer(ctx, req.(*GetCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/UpdateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomer(ctx, req.(*UpdateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CustomerService/DeleteCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).DeleteCustomer(ctx, req.(*DeleteCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _CustomerService_GetCustomer_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerService_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _CustomerService_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _CustomerService_DeleteCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/apartment.proto",
}
