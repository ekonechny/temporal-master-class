// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	temporal "temporal-master-class/internal/generated/temporal"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Customer_NewCustomer_FullMethodName         = "/server.Customer/NewCustomer"
	Customer_GetProfile_FullMethodName          = "/server.Customer/GetProfile"
	Customer_UpdateProfile_FullMethodName       = "/server.Customer/UpdateProfile"
	Customer_DeleteProfile_FullMethodName       = "/server.Customer/DeleteProfile"
	Customer_SetAddress_FullMethodName          = "/server.Customer/SetAddress"
	Customer_GetCart_FullMethodName             = "/server.Customer/GetCart"
	Customer_UpdateCart_FullMethodName          = "/server.Customer/UpdateCart"
	Customer_DeleteCart_FullMethodName          = "/server.Customer/DeleteCart"
	Customer_GetOrder_FullMethodName            = "/server.Customer/GetOrder"
	Customer_GetOrders_FullMethodName           = "/server.Customer/GetOrders"
	Customer_Checkout_FullMethodName            = "/server.Customer/Checkout"
	Customer_PaymentCallback_FullMethodName     = "/server.Customer/PaymentCallback"
	Customer_VendorOrderCallback_FullMethodName = "/server.Customer/VendorOrderCallback"
)

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerClient interface {
	// Здесь мы начинаем жизненный цикл пользователя
	NewCustomer(ctx context.Context, in *NewCustomerRequest, opts ...grpc.CallOption) (*temporal.Profile, error)
	// Получить профиль пользователя
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*temporal.Profile, error)
	// Обновить профиль пользователя
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*temporal.Profile, error)
	// Удалить профиль пользователя
	DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Задать юзеру адрес
	SetAddress(ctx context.Context, in *SetAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Получить корзину пользователя
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*temporal.Cart, error)
	// Обновить корзина пользователя целиком
	UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*temporal.Cart, error)
	// Очистить корзину пользователя
	DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Получить заказ пользователя
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*temporal.Order, error)
	// Получить все заказы пользователя
	GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
	// Создать заказ
	Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*temporal.Order, error)
	// Платежный колбек
	PaymentCallback(ctx context.Context, in *PaymentCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Колбек для апдейта от вендора
	VendorOrderCallback(ctx context.Context, in *VendorOrderCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) NewCustomer(ctx context.Context, in *NewCustomerRequest, opts ...grpc.CallOption) (*temporal.Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(temporal.Profile)
	err := c.cc.Invoke(ctx, Customer_NewCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*temporal.Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(temporal.Profile)
	err := c.cc.Invoke(ctx, Customer_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*temporal.Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(temporal.Profile)
	err := c.cc.Invoke(ctx, Customer_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) DeleteProfile(ctx context.Context, in *DeleteProfileRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Customer_DeleteProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) SetAddress(ctx context.Context, in *SetAddressRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Customer_SetAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*temporal.Cart, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(temporal.Cart)
	err := c.cc.Invoke(ctx, Customer_GetCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*temporal.Cart, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(temporal.Cart)
	err := c.cc.Invoke(ctx, Customer_UpdateCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Customer_DeleteCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*temporal.Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(temporal.Order)
	err := c.cc.Invoke(ctx, Customer_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) GetOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, Customer_GetOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*temporal.Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(temporal.Order)
	err := c.cc.Invoke(ctx, Customer_Checkout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) PaymentCallback(ctx context.Context, in *PaymentCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Customer_PaymentCallback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) VendorOrderCallback(ctx context.Context, in *VendorOrderCallbackRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Customer_VendorOrderCallback_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
// All implementations must embed UnimplementedCustomerServer
// for forward compatibility.
type CustomerServer interface {
	// Здесь мы начинаем жизненный цикл пользователя
	NewCustomer(context.Context, *NewCustomerRequest) (*temporal.Profile, error)
	// Получить профиль пользователя
	GetProfile(context.Context, *GetProfileRequest) (*temporal.Profile, error)
	// Обновить профиль пользователя
	UpdateProfile(context.Context, *UpdateProfileRequest) (*temporal.Profile, error)
	// Удалить профиль пользователя
	DeleteProfile(context.Context, *DeleteProfileRequest) (*emptypb.Empty, error)
	// Задать юзеру адрес
	SetAddress(context.Context, *SetAddressRequest) (*emptypb.Empty, error)
	// Получить корзину пользователя
	GetCart(context.Context, *GetCartRequest) (*temporal.Cart, error)
	// Обновить корзина пользователя целиком
	UpdateCart(context.Context, *UpdateCartRequest) (*temporal.Cart, error)
	// Очистить корзину пользователя
	DeleteCart(context.Context, *DeleteCartRequest) (*emptypb.Empty, error)
	// Получить заказ пользователя
	GetOrder(context.Context, *GetOrderRequest) (*temporal.Order, error)
	// Получить все заказы пользователя
	GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	// Создать заказ
	Checkout(context.Context, *CheckoutRequest) (*temporal.Order, error)
	// Платежный колбек
	PaymentCallback(context.Context, *PaymentCallbackRequest) (*emptypb.Empty, error)
	// Колбек для апдейта от вендора
	VendorOrderCallback(context.Context, *VendorOrderCallbackRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCustomerServer()
}

// UnimplementedCustomerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCustomerServer struct{}

func (UnimplementedCustomerServer) NewCustomer(context.Context, *NewCustomerRequest) (*temporal.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewCustomer not implemented")
}
func (UnimplementedCustomerServer) GetProfile(context.Context, *GetProfileRequest) (*temporal.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedCustomerServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*temporal.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedCustomerServer) DeleteProfile(context.Context, *DeleteProfileRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (UnimplementedCustomerServer) SetAddress(context.Context, *SetAddressRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAddress not implemented")
}
func (UnimplementedCustomerServer) GetCart(context.Context, *GetCartRequest) (*temporal.Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCustomerServer) UpdateCart(context.Context, *UpdateCartRequest) (*temporal.Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCart not implemented")
}
func (UnimplementedCustomerServer) DeleteCart(context.Context, *DeleteCartRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (UnimplementedCustomerServer) GetOrder(context.Context, *GetOrderRequest) (*temporal.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedCustomerServer) GetOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedCustomerServer) Checkout(context.Context, *CheckoutRequest) (*temporal.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (UnimplementedCustomerServer) PaymentCallback(context.Context, *PaymentCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentCallback not implemented")
}
func (UnimplementedCustomerServer) VendorOrderCallback(context.Context, *VendorOrderCallbackRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VendorOrderCallback not implemented")
}
func (UnimplementedCustomerServer) mustEmbedUnimplementedCustomerServer() {}
func (UnimplementedCustomerServer) testEmbeddedByValue()                  {}

// UnsafeCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServer will
// result in compilation errors.
type UnsafeCustomerServer interface {
	mustEmbedUnimplementedCustomerServer()
}

func RegisterCustomerServer(s grpc.ServiceRegistrar, srv CustomerServer) {
	// If the following call pancis, it indicates UnimplementedCustomerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Customer_ServiceDesc, srv)
}

func _Customer_NewCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).NewCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_NewCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).NewCustomer(ctx, req.(*NewCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_DeleteProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).DeleteProfile(ctx, req.(*DeleteProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_SetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).SetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_SetAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).SetAddress(ctx, req.(*SetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_UpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).UpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_UpdateCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).UpdateCart(ctx, req.(*UpdateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_DeleteCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).DeleteCart(ctx, req.(*DeleteCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_Checkout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).Checkout(ctx, req.(*CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_PaymentCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).PaymentCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_PaymentCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).PaymentCallback(ctx, req.(*PaymentCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_VendorOrderCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VendorOrderCallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).VendorOrderCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_VendorOrderCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).VendorOrderCallback(ctx, req.(*VendorOrderCallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_ServiceDesc is the grpc.ServiceDesc for Customer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewCustomer",
			Handler:    _Customer_NewCustomer_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Customer_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _Customer_UpdateProfile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _Customer_DeleteProfile_Handler,
		},
		{
			MethodName: "SetAddress",
			Handler:    _Customer_SetAddress_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _Customer_GetCart_Handler,
		},
		{
			MethodName: "UpdateCart",
			Handler:    _Customer_UpdateCart_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _Customer_DeleteCart_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Customer_GetOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _Customer_GetOrders_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _Customer_Checkout_Handler,
		},
		{
			MethodName: "PaymentCallback",
			Handler:    _Customer_PaymentCallback_Handler,
		},
		{
			MethodName: "VendorOrderCallback",
			Handler:    _Customer_VendorOrderCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
