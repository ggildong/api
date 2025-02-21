// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/alert_manager/v1/webhook.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Webhook_Create_FullMethodName       = "/spaceone.api.alert_manager.v1.Webhook/create"
	Webhook_Update_FullMethodName       = "/spaceone.api.alert_manager.v1.Webhook/update"
	Webhook_UpdatePlugin_FullMethodName = "/spaceone.api.alert_manager.v1.Webhook/update_plugin"
	Webhook_Enable_FullMethodName       = "/spaceone.api.alert_manager.v1.Webhook/enable"
	Webhook_Disable_FullMethodName      = "/spaceone.api.alert_manager.v1.Webhook/disable"
	Webhook_Delete_FullMethodName       = "/spaceone.api.alert_manager.v1.Webhook/delete"
	Webhook_Get_FullMethodName          = "/spaceone.api.alert_manager.v1.Webhook/get"
	Webhook_List_FullMethodName         = "/spaceone.api.alert_manager.v1.Webhook/list"
	Webhook_ListErrors_FullMethodName   = "/spaceone.api.alert_manager.v1.Webhook/list_errors"
	Webhook_Stat_FullMethodName         = "/spaceone.api.alert_manager.v1.Webhook/stat"
)

// WebhookClient is the client API for Webhook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhookClient interface {
	Create(ctx context.Context, in *WebhookCreateRequest, opts ...grpc.CallOption) (*WebhookInfo, error)
	Update(ctx context.Context, in *WebhookUpdateRequest, opts ...grpc.CallOption) (*WebhookInfo, error)
	UpdatePlugin(ctx context.Context, in *WebhookUpdatePluginRequest, opts ...grpc.CallOption) (*WebhookInfo, error)
	Enable(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookInfo, error)
	Disable(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookInfo, error)
	Delete(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookInfo, error)
	List(ctx context.Context, in *WebhookSearchQuery, opts ...grpc.CallOption) (*WebhooksInfo, error)
	ListErrors(ctx context.Context, in *WebhookErrorSearchQuery, opts ...grpc.CallOption) (*WebhookErrorsInfo, error)
	Stat(ctx context.Context, in *WebhookStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type webhookClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhookClient(cc grpc.ClientConnInterface) WebhookClient {
	return &webhookClient{cc}
}

func (c *webhookClient) Create(ctx context.Context, in *WebhookCreateRequest, opts ...grpc.CallOption) (*WebhookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookInfo)
	err := c.cc.Invoke(ctx, Webhook_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Update(ctx context.Context, in *WebhookUpdateRequest, opts ...grpc.CallOption) (*WebhookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookInfo)
	err := c.cc.Invoke(ctx, Webhook_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) UpdatePlugin(ctx context.Context, in *WebhookUpdatePluginRequest, opts ...grpc.CallOption) (*WebhookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookInfo)
	err := c.cc.Invoke(ctx, Webhook_UpdatePlugin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Enable(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookInfo)
	err := c.cc.Invoke(ctx, Webhook_Enable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Disable(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookInfo)
	err := c.cc.Invoke(ctx, Webhook_Disable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Delete(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Webhook_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Get(ctx context.Context, in *WebhookRequest, opts ...grpc.CallOption) (*WebhookInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookInfo)
	err := c.cc.Invoke(ctx, Webhook_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) List(ctx context.Context, in *WebhookSearchQuery, opts ...grpc.CallOption) (*WebhooksInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhooksInfo)
	err := c.cc.Invoke(ctx, Webhook_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) ListErrors(ctx context.Context, in *WebhookErrorSearchQuery, opts ...grpc.CallOption) (*WebhookErrorsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookErrorsInfo)
	err := c.cc.Invoke(ctx, Webhook_ListErrors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Stat(ctx context.Context, in *WebhookStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Webhook_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookServer is the server API for Webhook service.
// All implementations must embed UnimplementedWebhookServer
// for forward compatibility.
type WebhookServer interface {
	Create(context.Context, *WebhookCreateRequest) (*WebhookInfo, error)
	Update(context.Context, *WebhookUpdateRequest) (*WebhookInfo, error)
	UpdatePlugin(context.Context, *WebhookUpdatePluginRequest) (*WebhookInfo, error)
	Enable(context.Context, *WebhookRequest) (*WebhookInfo, error)
	Disable(context.Context, *WebhookRequest) (*WebhookInfo, error)
	Delete(context.Context, *WebhookRequest) (*empty.Empty, error)
	Get(context.Context, *WebhookRequest) (*WebhookInfo, error)
	List(context.Context, *WebhookSearchQuery) (*WebhooksInfo, error)
	ListErrors(context.Context, *WebhookErrorSearchQuery) (*WebhookErrorsInfo, error)
	Stat(context.Context, *WebhookStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedWebhookServer()
}

// UnimplementedWebhookServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWebhookServer struct{}

func (UnimplementedWebhookServer) Create(context.Context, *WebhookCreateRequest) (*WebhookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWebhookServer) Update(context.Context, *WebhookUpdateRequest) (*WebhookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWebhookServer) UpdatePlugin(context.Context, *WebhookUpdatePluginRequest) (*WebhookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlugin not implemented")
}
func (UnimplementedWebhookServer) Enable(context.Context, *WebhookRequest) (*WebhookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedWebhookServer) Disable(context.Context, *WebhookRequest) (*WebhookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedWebhookServer) Delete(context.Context, *WebhookRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWebhookServer) Get(context.Context, *WebhookRequest) (*WebhookInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedWebhookServer) List(context.Context, *WebhookSearchQuery) (*WebhooksInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedWebhookServer) ListErrors(context.Context, *WebhookErrorSearchQuery) (*WebhookErrorsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListErrors not implemented")
}
func (UnimplementedWebhookServer) Stat(context.Context, *WebhookStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedWebhookServer) mustEmbedUnimplementedWebhookServer() {}
func (UnimplementedWebhookServer) testEmbeddedByValue()                 {}

// UnsafeWebhookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhookServer will
// result in compilation errors.
type UnsafeWebhookServer interface {
	mustEmbedUnimplementedWebhookServer()
}

func RegisterWebhookServer(s grpc.ServiceRegistrar, srv WebhookServer) {
	// If the following call pancis, it indicates UnimplementedWebhookServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Webhook_ServiceDesc, srv)
}

func _Webhook_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Create(ctx, req.(*WebhookCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Update(ctx, req.(*WebhookUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_UpdatePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookUpdatePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).UpdatePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_UpdatePlugin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).UpdatePlugin(ctx, req.(*WebhookUpdatePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Enable(ctx, req.(*WebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Disable(ctx, req.(*WebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Delete(ctx, req.(*WebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Get(ctx, req.(*WebhookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).List(ctx, req.(*WebhookSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_ListErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookErrorSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).ListErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_ListErrors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).ListErrors(ctx, req.(*WebhookErrorSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Webhook_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Stat(ctx, req.(*WebhookStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Webhook_ServiceDesc is the grpc.ServiceDesc for Webhook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Webhook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.alert_manager.v1.Webhook",
	HandlerType: (*WebhookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Webhook_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Webhook_Update_Handler,
		},
		{
			MethodName: "update_plugin",
			Handler:    _Webhook_UpdatePlugin_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _Webhook_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _Webhook_Disable_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Webhook_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Webhook_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Webhook_List_Handler,
		},
		{
			MethodName: "list_errors",
			Handler:    _Webhook_ListErrors_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Webhook_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/alert_manager/v1/webhook.proto",
}
