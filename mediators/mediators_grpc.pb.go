// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: pb/mediators.proto

package mediators

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

// MediateClient is the client API for Mediate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediateClient interface {
	Mediate(ctx context.Context, in *MediationInput, opts ...grpc.CallOption) (*MediationOutput, error)
}

type mediateClient struct {
	cc grpc.ClientConnInterface
}

func NewMediateClient(cc grpc.ClientConnInterface) MediateClient {
	return &mediateClient{cc}
}

func (c *mediateClient) Mediate(ctx context.Context, in *MediationInput, opts ...grpc.CallOption) (*MediationOutput, error) {
	out := new(MediationOutput)
	err := c.cc.Invoke(ctx, "/Mediate/Mediate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediateServer is the server API for Mediate service.
// All implementations should embed UnimplementedMediateServer
// for forward compatibility
type MediateServer interface {
	Mediate(context.Context, *MediationInput) (*MediationOutput, error)
}

// UnimplementedMediateServer should be embedded to have forward compatible implementations.
type UnimplementedMediateServer struct {
}

func (UnimplementedMediateServer) Mediate(context.Context, *MediationInput) (*MediationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mediate not implemented")
}

// UnsafeMediateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediateServer will
// result in compilation errors.
type UnsafeMediateServer interface {
	mustEmbedUnimplementedMediateServer()
}

func RegisterMediateServer(s grpc.ServiceRegistrar, srv MediateServer) {
	s.RegisterService(&Mediate_ServiceDesc, srv)
}

func _Mediate_Mediate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediateServer).Mediate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mediate/Mediate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediateServer).Mediate(ctx, req.(*MediationInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Mediate_ServiceDesc is the grpc.ServiceDesc for Mediate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mediate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Mediate",
	HandlerType: (*MediateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mediate",
			Handler:    _Mediate_Mediate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/mediators.proto",
}
