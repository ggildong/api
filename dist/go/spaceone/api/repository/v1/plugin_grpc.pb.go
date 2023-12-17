// A Plugin is a resource containing data of deployable plugins such as container image and registry URL.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/repository/v1/plugin.proto

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Plugin_Register_FullMethodName    = "/spaceone.api.repository.v1.Plugin/register"
	Plugin_Update_FullMethodName      = "/spaceone.api.repository.v1.Plugin/update"
	Plugin_Deregister_FullMethodName  = "/spaceone.api.repository.v1.Plugin/deregister"
	Plugin_Enable_FullMethodName      = "/spaceone.api.repository.v1.Plugin/enable"
	Plugin_Disable_FullMethodName     = "/spaceone.api.repository.v1.Plugin/disable"
	Plugin_GetVersions_FullMethodName = "/spaceone.api.repository.v1.Plugin/get_versions"
	Plugin_Get_FullMethodName         = "/spaceone.api.repository.v1.Plugin/get"
	Plugin_List_FullMethodName        = "/spaceone.api.repository.v1.Plugin/list"
	Plugin_Stat_FullMethodName        = "/spaceone.api.repository.v1.Plugin/stat"
)

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	// Registers a Plugin. The parameter `registry_type`, meaning container registry type, can be either `DOCKER_HUB` or `AWS_PRIVATE_ECR`. The default value of the `registry_type` is `DOCKER_HUB`. The parameter `registry_url` is required if the `registry_type` is not `DOCKER_HUB`. The parameter `image` is limited to 40 characters.
	Register(ctx context.Context, in *CreatePluginRequest, opts ...grpc.CallOption) (*PluginInfo, error)
	// Updates a specific Plugin registered. A Plugin can be updated only if its Repository's `repository_type` is `local`. You can make changes in Plugin settings, including `template` and its options, `schema`.
	Update(ctx context.Context, in *UpdatePluginRequest, opts ...grpc.CallOption) (*PluginInfo, error)
	// Deregisters and deletes a specific Plugin. You must specify the `plugin_id` of the Plugin to deregister.
	Deregister(ctx context.Context, in *PluginRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Enables a specific Plugin. If the Plugin is enabled, the Plugin can be used as its parameter `state` becomes `ENABLED`.
	Enable(ctx context.Context, in *PluginRequest, opts ...grpc.CallOption) (*PluginInfo, error)
	// Disables a specific Plugin. If the Plugin is disabled, the Plugin cannot be used as its parameter `state` becomes `DISABLED`.
	Disable(ctx context.Context, in *PluginRequest, opts ...grpc.CallOption) (*PluginInfo, error)
	// Gets all version data of a specific Plugin from its Repository. The parameter `plugin_id` is used as an identifier of a Plugin to get version data.
	GetVersions(ctx context.Context, in *RepositoryPluginRequest, opts ...grpc.CallOption) (*VersionsInfo, error)
	// Gets a specific Plugin. Prints detailed information about the Plugin, including  `image`, `registry_url`, and `state`.
	Get(ctx context.Context, in *RepositoryPluginRequest, opts ...grpc.CallOption) (*PluginInfo, error)
	// Gets a list of all Plugins registered in a specific Repository. The parameter `repository_id` is used as an identifier of a Repository to get its list of Plugins. You can use a query to get a filtered list of Plugins.
	List(ctx context.Context, in *PluginQuery, opts ...grpc.CallOption) (*PluginsInfo, error)
	Stat(ctx context.Context, in *PluginStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Register(ctx context.Context, in *CreatePluginRequest, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, Plugin_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Update(ctx context.Context, in *UpdatePluginRequest, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, Plugin_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Deregister(ctx context.Context, in *PluginRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Plugin_Deregister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Enable(ctx context.Context, in *PluginRequest, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, Plugin_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Disable(ctx context.Context, in *PluginRequest, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, Plugin_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) GetVersions(ctx context.Context, in *RepositoryPluginRequest, opts ...grpc.CallOption) (*VersionsInfo, error) {
	out := new(VersionsInfo)
	err := c.cc.Invoke(ctx, Plugin_GetVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Get(ctx context.Context, in *RepositoryPluginRequest, opts ...grpc.CallOption) (*PluginInfo, error) {
	out := new(PluginInfo)
	err := c.cc.Invoke(ctx, Plugin_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) List(ctx context.Context, in *PluginQuery, opts ...grpc.CallOption) (*PluginsInfo, error) {
	out := new(PluginsInfo)
	err := c.cc.Invoke(ctx, Plugin_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Stat(ctx context.Context, in *PluginStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Plugin_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	// Registers a Plugin. The parameter `registry_type`, meaning container registry type, can be either `DOCKER_HUB` or `AWS_PRIVATE_ECR`. The default value of the `registry_type` is `DOCKER_HUB`. The parameter `registry_url` is required if the `registry_type` is not `DOCKER_HUB`. The parameter `image` is limited to 40 characters.
	Register(context.Context, *CreatePluginRequest) (*PluginInfo, error)
	// Updates a specific Plugin registered. A Plugin can be updated only if its Repository's `repository_type` is `local`. You can make changes in Plugin settings, including `template` and its options, `schema`.
	Update(context.Context, *UpdatePluginRequest) (*PluginInfo, error)
	// Deregisters and deletes a specific Plugin. You must specify the `plugin_id` of the Plugin to deregister.
	Deregister(context.Context, *PluginRequest) (*empty.Empty, error)
	// Enables a specific Plugin. If the Plugin is enabled, the Plugin can be used as its parameter `state` becomes `ENABLED`.
	Enable(context.Context, *PluginRequest) (*PluginInfo, error)
	// Disables a specific Plugin. If the Plugin is disabled, the Plugin cannot be used as its parameter `state` becomes `DISABLED`.
	Disable(context.Context, *PluginRequest) (*PluginInfo, error)
	// Gets all version data of a specific Plugin from its Repository. The parameter `plugin_id` is used as an identifier of a Plugin to get version data.
	GetVersions(context.Context, *RepositoryPluginRequest) (*VersionsInfo, error)
	// Gets a specific Plugin. Prints detailed information about the Plugin, including  `image`, `registry_url`, and `state`.
	Get(context.Context, *RepositoryPluginRequest) (*PluginInfo, error)
	// Gets a list of all Plugins registered in a specific Repository. The parameter `repository_id` is used as an identifier of a Repository to get its list of Plugins. You can use a query to get a filtered list of Plugins.
	List(context.Context, *PluginQuery) (*PluginsInfo, error)
	Stat(context.Context, *PluginStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) Register(context.Context, *CreatePluginRequest) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedPluginServer) Update(context.Context, *UpdatePluginRequest) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPluginServer) Deregister(context.Context, *PluginRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (UnimplementedPluginServer) Enable(context.Context, *PluginRequest) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedPluginServer) Disable(context.Context, *PluginRequest) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedPluginServer) GetVersions(context.Context, *RepositoryPluginRequest) (*VersionsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersions not implemented")
}
func (UnimplementedPluginServer) Get(context.Context, *RepositoryPluginRequest) (*PluginInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPluginServer) List(context.Context, *PluginQuery) (*PluginsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPluginServer) Stat(context.Context, *PluginStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Register(ctx, req.(*CreatePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Update(ctx, req.(*UpdatePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Deregister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Deregister(ctx, req.(*PluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Enable(ctx, req.(*PluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Disable(ctx, req.(*PluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_GetVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_GetVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetVersions(ctx, req.(*RepositoryPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Get(ctx, req.(*RepositoryPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).List(ctx, req.(*PluginQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Stat(ctx, req.(*PluginStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.repository.v1.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _Plugin_Register_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Plugin_Update_Handler,
		},
		{
			MethodName: "deregister",
			Handler:    _Plugin_Deregister_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _Plugin_Enable_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _Plugin_Disable_Handler,
		},
		{
			MethodName: "get_versions",
			Handler:    _Plugin_GetVersions_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Plugin_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Plugin_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Plugin_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/repository/v1/plugin.proto",
}
