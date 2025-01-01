// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: transport.proto

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
	CompanyService_CreateCompany_FullMethodName  = "/transport.CompanyService/CreateCompany"
	CompanyService_UpdateCompany_FullMethodName  = "/transport.CompanyService/UpdateCompany"
	CompanyService_DeleteCompany_FullMethodName  = "/transport.CompanyService/DeleteCompany"
	CompanyService_GetByIDCompany_FullMethodName = "/transport.CompanyService/GetByIDCompany"
)

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyServiceClient interface {
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error)
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*DeleteCompanyResponse, error)
	GetByIDCompany(ctx context.Context, in *GetByIDCompanyRequest, opts ...grpc.CallOption) (*GetByIDCompanyResponse, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_CreateCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_UpdateCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*DeleteCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_DeleteCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetByIDCompany(ctx context.Context, in *GetByIDCompanyRequest, opts ...grpc.CallOption) (*GetByIDCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetByIDCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_GetByIDCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
// All implementations must embed UnimplementedCompanyServiceServer
// for forward compatibility.
type CompanyServiceServer interface {
	CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error)
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*DeleteCompanyResponse, error)
	GetByIDCompany(context.Context, *GetByIDCompanyRequest) (*GetByIDCompanyResponse, error)
	mustEmbedUnimplementedCompanyServiceServer()
}

// UnimplementedCompanyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCompanyServiceServer struct{}

func (UnimplementedCompanyServiceServer) CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedCompanyServiceServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCompanyServiceServer) DeleteCompany(context.Context, *DeleteCompanyRequest) (*DeleteCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (UnimplementedCompanyServiceServer) GetByIDCompany(context.Context, *GetByIDCompanyRequest) (*GetByIDCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIDCompany not implemented")
}
func (UnimplementedCompanyServiceServer) mustEmbedUnimplementedCompanyServiceServer() {}
func (UnimplementedCompanyServiceServer) testEmbeddedByValue()                        {}

// UnsafeCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServiceServer will
// result in compilation errors.
type UnsafeCompanyServiceServer interface {
	mustEmbedUnimplementedCompanyServiceServer()
}

func RegisterCompanyServiceServer(s grpc.ServiceRegistrar, srv CompanyServiceServer) {
	// If the following call pancis, it indicates UnimplementedCompanyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CompanyService_ServiceDesc, srv)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_CreateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_UpdateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_DeleteCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetByIDCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetByIDCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_GetByIDCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetByIDCompany(ctx, req.(*GetByIDCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyService_ServiceDesc is the grpc.ServiceDesc for CompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transport.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyService_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _CompanyService_DeleteCompany_Handler,
		},
		{
			MethodName: "GetByIDCompany",
			Handler:    _CompanyService_GetByIDCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}

const (
	TourService_CreateTour_FullMethodName  = "/transport.TourService/CreateTour"
	TourService_UpdateTour_FullMethodName  = "/transport.TourService/UpdateTour"
	TourService_DeleteTour_FullMethodName  = "/transport.TourService/DeleteTour"
	TourService_GetByIDTour_FullMethodName = "/transport.TourService/GetByIDTour"
)

// TourServiceClient is the client API for TourService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TourServiceClient interface {
	CreateTour(ctx context.Context, in *CreateTourRequest, opts ...grpc.CallOption) (*CreateTourResponse, error)
	UpdateTour(ctx context.Context, in *UpdateTourRequest, opts ...grpc.CallOption) (*UpdateTourResponse, error)
	DeleteTour(ctx context.Context, in *DeleteTourRequest, opts ...grpc.CallOption) (*DeleteTourResponse, error)
	GetByIDTour(ctx context.Context, in *GetByIDTourRequest, opts ...grpc.CallOption) (*GetByIDTourResponse, error)
}

type tourServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTourServiceClient(cc grpc.ClientConnInterface) TourServiceClient {
	return &tourServiceClient{cc}
}

func (c *tourServiceClient) CreateTour(ctx context.Context, in *CreateTourRequest, opts ...grpc.CallOption) (*CreateTourResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTourResponse)
	err := c.cc.Invoke(ctx, TourService_CreateTour_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourServiceClient) UpdateTour(ctx context.Context, in *UpdateTourRequest, opts ...grpc.CallOption) (*UpdateTourResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTourResponse)
	err := c.cc.Invoke(ctx, TourService_UpdateTour_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourServiceClient) DeleteTour(ctx context.Context, in *DeleteTourRequest, opts ...grpc.CallOption) (*DeleteTourResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTourResponse)
	err := c.cc.Invoke(ctx, TourService_DeleteTour_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tourServiceClient) GetByIDTour(ctx context.Context, in *GetByIDTourRequest, opts ...grpc.CallOption) (*GetByIDTourResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetByIDTourResponse)
	err := c.cc.Invoke(ctx, TourService_GetByIDTour_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TourServiceServer is the server API for TourService service.
// All implementations must embed UnimplementedTourServiceServer
// for forward compatibility.
type TourServiceServer interface {
	CreateTour(context.Context, *CreateTourRequest) (*CreateTourResponse, error)
	UpdateTour(context.Context, *UpdateTourRequest) (*UpdateTourResponse, error)
	DeleteTour(context.Context, *DeleteTourRequest) (*DeleteTourResponse, error)
	GetByIDTour(context.Context, *GetByIDTourRequest) (*GetByIDTourResponse, error)
	mustEmbedUnimplementedTourServiceServer()
}

// UnimplementedTourServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTourServiceServer struct{}

func (UnimplementedTourServiceServer) CreateTour(context.Context, *CreateTourRequest) (*CreateTourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTour not implemented")
}
func (UnimplementedTourServiceServer) UpdateTour(context.Context, *UpdateTourRequest) (*UpdateTourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTour not implemented")
}
func (UnimplementedTourServiceServer) DeleteTour(context.Context, *DeleteTourRequest) (*DeleteTourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTour not implemented")
}
func (UnimplementedTourServiceServer) GetByIDTour(context.Context, *GetByIDTourRequest) (*GetByIDTourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIDTour not implemented")
}
func (UnimplementedTourServiceServer) mustEmbedUnimplementedTourServiceServer() {}
func (UnimplementedTourServiceServer) testEmbeddedByValue()                     {}

// UnsafeTourServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TourServiceServer will
// result in compilation errors.
type UnsafeTourServiceServer interface {
	mustEmbedUnimplementedTourServiceServer()
}

func RegisterTourServiceServer(s grpc.ServiceRegistrar, srv TourServiceServer) {
	// If the following call pancis, it indicates UnimplementedTourServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TourService_ServiceDesc, srv)
}

func _TourService_CreateTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).CreateTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TourService_CreateTour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).CreateTour(ctx, req.(*CreateTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TourService_UpdateTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).UpdateTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TourService_UpdateTour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).UpdateTour(ctx, req.(*UpdateTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TourService_DeleteTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).DeleteTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TourService_DeleteTour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).DeleteTour(ctx, req.(*DeleteTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TourService_GetByIDTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TourServiceServer).GetByIDTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TourService_GetByIDTour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TourServiceServer).GetByIDTour(ctx, req.(*GetByIDTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TourService_ServiceDesc is the grpc.ServiceDesc for TourService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TourService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transport.TourService",
	HandlerType: (*TourServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTour",
			Handler:    _TourService_CreateTour_Handler,
		},
		{
			MethodName: "UpdateTour",
			Handler:    _TourService_UpdateTour_Handler,
		},
		{
			MethodName: "DeleteTour",
			Handler:    _TourService_DeleteTour_Handler,
		},
		{
			MethodName: "GetByIDTour",
			Handler:    _TourService_GetByIDTour_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}
