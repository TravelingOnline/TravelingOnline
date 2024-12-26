// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: vehicle.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VehicleService_CreateVehicle_FullMethodName = "/vehicle.VehicleService/CreateVehicle"
)

// VehicleServiceClient is the client API for VehicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleServiceClient interface {
	CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...grpc.CallOption) (*CreateVehicleResponse, error)
}

type vehicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleServiceClient(cc grpc.ClientConnInterface) VehicleServiceClient {
	return &vehicleServiceClient{cc}
}

func (c *vehicleServiceClient) CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...grpc.CallOption) (*CreateVehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVehicleResponse)
	err := c.cc.Invoke(ctx, VehicleService_CreateVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServiceServer is the server API for VehicleService service.
// All implementations must embed UnimplementedVehicleServiceServer
// for forward compatibility.
type VehicleServiceServer interface {
	CreateVehicle(context.Context, *CreateVehicleRequest) (*CreateVehicleResponse, error)
	mustEmbedUnimplementedVehicleServiceServer()
}

// UnimplementedVehicleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVehicleServiceServer struct{}

func (UnimplementedVehicleServiceServer) CreateVehicle(context.Context, *CreateVehicleRequest) (*CreateVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicle not implemented")
}
func (UnimplementedVehicleServiceServer) mustEmbedUnimplementedVehicleServiceServer() {}
func (UnimplementedVehicleServiceServer) testEmbeddedByValue()                        {}

// UnsafeVehicleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleServiceServer will
// result in compilation errors.
type UnsafeVehicleServiceServer interface {
	mustEmbedUnimplementedVehicleServiceServer()
}

func RegisterVehicleServiceServer(s grpc.ServiceRegistrar, srv VehicleServiceServer) {
	// If the following call pancis, it indicates UnimplementedVehicleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VehicleService_ServiceDesc, srv)
}

func _VehicleService_CreateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).CreateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleService_CreateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).CreateVehicle(ctx, req.(*CreateVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VehicleService_ServiceDesc is the grpc.ServiceDesc for VehicleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vehicle.VehicleService",
	HandlerType: (*VehicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVehicle",
			Handler:    _VehicleService_CreateVehicle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vehicle.proto",
}
