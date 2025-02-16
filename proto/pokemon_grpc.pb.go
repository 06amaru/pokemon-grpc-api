// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: proto/pokemon.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PokemonService_CreatePokemon_FullMethodName = "/pokemon.PokemonService/CreatePokemon"
	PokemonService_GetPokemon_FullMethodName    = "/pokemon.PokemonService/GetPokemon"
	PokemonService_UpdatePokemon_FullMethodName = "/pokemon.PokemonService/UpdatePokemon"
	PokemonService_DeletePokemon_FullMethodName = "/pokemon.PokemonService/DeletePokemon"
	PokemonService_ListPokemon_FullMethodName   = "/pokemon.PokemonService/ListPokemon"
)

// PokemonServiceClient is the client API for PokemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokemonServiceClient interface {
	CreatePokemon(ctx context.Context, in *CreatePokemonRequest, opts ...grpc.CallOption) (*Pokemon, error)
	GetPokemon(ctx context.Context, in *GetPokemonRequest, opts ...grpc.CallOption) (*Pokemon, error)
	UpdatePokemon(ctx context.Context, in *UpdatePokemonRequest, opts ...grpc.CallOption) (*Pokemon, error)
	DeletePokemon(ctx context.Context, in *DeletePokemonRequest, opts ...grpc.CallOption) (*DeletePokemonResponse, error)
	ListPokemon(ctx context.Context, in *ListPokemonRequest, opts ...grpc.CallOption) (*ListPokemonResponse, error)
}

type pokemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonServiceClient(cc grpc.ClientConnInterface) PokemonServiceClient {
	return &pokemonServiceClient{cc}
}

func (c *pokemonServiceClient) CreatePokemon(ctx context.Context, in *CreatePokemonRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, PokemonService_CreatePokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) GetPokemon(ctx context.Context, in *GetPokemonRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, PokemonService_GetPokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) UpdatePokemon(ctx context.Context, in *UpdatePokemonRequest, opts ...grpc.CallOption) (*Pokemon, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Pokemon)
	err := c.cc.Invoke(ctx, PokemonService_UpdatePokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) DeletePokemon(ctx context.Context, in *DeletePokemonRequest, opts ...grpc.CallOption) (*DeletePokemonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePokemonResponse)
	err := c.cc.Invoke(ctx, PokemonService_DeletePokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) ListPokemon(ctx context.Context, in *ListPokemonRequest, opts ...grpc.CallOption) (*ListPokemonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPokemonResponse)
	err := c.cc.Invoke(ctx, PokemonService_ListPokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonServiceServer is the server API for PokemonService service.
// All implementations must embed UnimplementedPokemonServiceServer
// for forward compatibility
type PokemonServiceServer interface {
	CreatePokemon(context.Context, *CreatePokemonRequest) (*Pokemon, error)
	GetPokemon(context.Context, *GetPokemonRequest) (*Pokemon, error)
	UpdatePokemon(context.Context, *UpdatePokemonRequest) (*Pokemon, error)
	DeletePokemon(context.Context, *DeletePokemonRequest) (*DeletePokemonResponse, error)
	ListPokemon(context.Context, *ListPokemonRequest) (*ListPokemonResponse, error)
	mustEmbedUnimplementedPokemonServiceServer()
}

// UnimplementedPokemonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPokemonServiceServer struct {
}

func (UnimplementedPokemonServiceServer) CreatePokemon(context.Context, *CreatePokemonRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePokemon not implemented")
}
func (UnimplementedPokemonServiceServer) GetPokemon(context.Context, *GetPokemonRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}
func (UnimplementedPokemonServiceServer) UpdatePokemon(context.Context, *UpdatePokemonRequest) (*Pokemon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePokemon not implemented")
}
func (UnimplementedPokemonServiceServer) DeletePokemon(context.Context, *DeletePokemonRequest) (*DeletePokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePokemon not implemented")
}
func (UnimplementedPokemonServiceServer) ListPokemon(context.Context, *ListPokemonRequest) (*ListPokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPokemon not implemented")
}
func (UnimplementedPokemonServiceServer) mustEmbedUnimplementedPokemonServiceServer() {}

// UnsafePokemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokemonServiceServer will
// result in compilation errors.
type UnsafePokemonServiceServer interface {
	mustEmbedUnimplementedPokemonServiceServer()
}

func RegisterPokemonServiceServer(s grpc.ServiceRegistrar, srv PokemonServiceServer) {
	s.RegisterService(&PokemonService_ServiceDesc, srv)
}

func _PokemonService_CreatePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).CreatePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_CreatePokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).CreatePokemon(ctx, req.(*CreatePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_GetPokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).GetPokemon(ctx, req.(*GetPokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_UpdatePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).UpdatePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_UpdatePokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).UpdatePokemon(ctx, req.(*UpdatePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_DeletePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).DeletePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_DeletePokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).DeletePokemon(ctx, req.(*DeletePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_ListPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).ListPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_ListPokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).ListPokemon(ctx, req.(*ListPokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PokemonService_ServiceDesc is the grpc.ServiceDesc for PokemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PokemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokemon.PokemonService",
	HandlerType: (*PokemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePokemon",
			Handler:    _PokemonService_CreatePokemon_Handler,
		},
		{
			MethodName: "GetPokemon",
			Handler:    _PokemonService_GetPokemon_Handler,
		},
		{
			MethodName: "UpdatePokemon",
			Handler:    _PokemonService_UpdatePokemon_Handler,
		},
		{
			MethodName: "DeletePokemon",
			Handler:    _PokemonService_DeletePokemon_Handler,
		},
		{
			MethodName: "ListPokemon",
			Handler:    _PokemonService_ListPokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pokemon.proto",
}
