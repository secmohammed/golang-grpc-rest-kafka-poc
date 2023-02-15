// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: handlers/grpc/proto/companies.proto

package companies

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

// CompaniesClient is the client API for Companies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompaniesClient interface {
	GetCompanyList(ctx context.Context, in *GetCompaniesListRequest, opts ...grpc.CallOption) (*GetCompaniesListResponse, error)
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*DeleteCompanyResponse, error)
}

type companiesClient struct {
	cc grpc.ClientConnInterface
}

func NewCompaniesClient(cc grpc.ClientConnInterface) CompaniesClient {
	return &companiesClient{cc}
}

func (c *companiesClient) GetCompanyList(ctx context.Context, in *GetCompaniesListRequest, opts ...grpc.CallOption) (*GetCompaniesListResponse, error) {
	out := new(GetCompaniesListResponse)
	err := c.cc.Invoke(ctx, "/Companies/GetCompanyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companiesClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, "/Companies/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companiesClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, "/Companies/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companiesClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, "/Companies/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companiesClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*DeleteCompanyResponse, error) {
	out := new(DeleteCompanyResponse)
	err := c.cc.Invoke(ctx, "/Companies/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompaniesServer is the server API for Companies service.
// All implementations must embed UnimplementedCompaniesServer
// for forward compatibility
type CompaniesServer interface {
	GetCompanyList(context.Context, *GetCompaniesListRequest) (*GetCompaniesListResponse, error)
	GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)
	CreateCompany(context.Context, *CreateCompanyRequest) (*GetCompanyResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*GetCompanyResponse, error)
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*DeleteCompanyResponse, error)
	mustEmbedUnimplementedCompaniesServer()
}

// UnimplementedCompaniesServer must be embedded to have forward compatible implementations.
type UnimplementedCompaniesServer struct {
}

func (UnimplementedCompaniesServer) GetCompanyList(context.Context, *GetCompaniesListRequest) (*GetCompaniesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyList not implemented")
}
func (UnimplementedCompaniesServer) GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedCompaniesServer) CreateCompany(context.Context, *CreateCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedCompaniesServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCompaniesServer) DeleteCompany(context.Context, *DeleteCompanyRequest) (*DeleteCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (UnimplementedCompaniesServer) mustEmbedUnimplementedCompaniesServer() {}

// UnsafeCompaniesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompaniesServer will
// result in compilation errors.
type UnsafeCompaniesServer interface {
	mustEmbedUnimplementedCompaniesServer()
}

func RegisterCompaniesServer(s grpc.ServiceRegistrar, srv CompaniesServer) {
	s.RegisterService(&Companies_ServiceDesc, srv)
}

func _Companies_GetCompanyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompaniesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).GetCompanyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Companies/GetCompanyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).GetCompanyList(ctx, req.(*GetCompaniesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Companies_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Companies/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Companies_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Companies/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Companies_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Companies/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Companies_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Companies/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Companies_ServiceDesc is the grpc.ServiceDesc for Companies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Companies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Companies",
	HandlerType: (*CompaniesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyList",
			Handler:    _Companies_GetCompanyList_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _Companies_GetCompany_Handler,
		},
		{
			MethodName: "CreateCompany",
			Handler:    _Companies_CreateCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _Companies_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _Companies_DeleteCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "handlers/grpc/proto/companies.proto",
}
