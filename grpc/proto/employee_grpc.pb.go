// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: employee.proto

package routeGuide

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

// EmployeeDBClient is the client API for EmployeeDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeDBClient interface {
	AddEmployee(ctx context.Context, in *EmployeeToBeAdded, opts ...grpc.CallOption) (*Employee, error)
	GetEmployees(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllEmployees, error)
	UpdateEmailOfEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error)
	DeleteEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Empty, error)
}

type employeeDBClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeDBClient(cc grpc.ClientConnInterface) EmployeeDBClient {
	return &employeeDBClient{cc}
}

func (c *employeeDBClient) AddEmployee(ctx context.Context, in *EmployeeToBeAdded, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/routeGuide.EmployeeDB/AddEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeDBClient) GetEmployees(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AllEmployees, error) {
	out := new(AllEmployees)
	err := c.cc.Invoke(ctx, "/routeGuide.EmployeeDB/GetEmployees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeDBClient) UpdateEmailOfEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Employee, error) {
	out := new(Employee)
	err := c.cc.Invoke(ctx, "/routeGuide.EmployeeDB/UpdateEmailOfEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeDBClient) DeleteEmployee(ctx context.Context, in *Employee, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/routeGuide.EmployeeDB/DeleteEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeDBServer is the server API for EmployeeDB service.
// All implementations must embed UnimplementedEmployeeDBServer
// for forward compatibility
type EmployeeDBServer interface {
	AddEmployee(context.Context, *EmployeeToBeAdded) (*Employee, error)
	GetEmployees(context.Context, *Empty) (*AllEmployees, error)
	UpdateEmailOfEmployee(context.Context, *Employee) (*Employee, error)
	DeleteEmployee(context.Context, *Employee) (*Empty, error)
	mustEmbedUnimplementedEmployeeDBServer()
}

// UnimplementedEmployeeDBServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeDBServer struct {
}

func (UnimplementedEmployeeDBServer) AddEmployee(context.Context, *EmployeeToBeAdded) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEmployee not implemented")
}
func (UnimplementedEmployeeDBServer) GetEmployees(context.Context, *Empty) (*AllEmployees, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployees not implemented")
}
func (UnimplementedEmployeeDBServer) UpdateEmailOfEmployee(context.Context, *Employee) (*Employee, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailOfEmployee not implemented")
}
func (UnimplementedEmployeeDBServer) DeleteEmployee(context.Context, *Employee) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmployee not implemented")
}
func (UnimplementedEmployeeDBServer) mustEmbedUnimplementedEmployeeDBServer() {}

// UnsafeEmployeeDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeDBServer will
// result in compilation errors.
type UnsafeEmployeeDBServer interface {
	mustEmbedUnimplementedEmployeeDBServer()
}

func RegisterEmployeeDBServer(s grpc.ServiceRegistrar, srv EmployeeDBServer) {
	s.RegisterService(&EmployeeDB_ServiceDesc, srv)
}

func _EmployeeDB_AddEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeToBeAdded)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeDBServer).AddEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeGuide.EmployeeDB/AddEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeDBServer).AddEmployee(ctx, req.(*EmployeeToBeAdded))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeDB_GetEmployees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeDBServer).GetEmployees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeGuide.EmployeeDB/GetEmployees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeDBServer).GetEmployees(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeDB_UpdateEmailOfEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeDBServer).UpdateEmailOfEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeGuide.EmployeeDB/UpdateEmailOfEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeDBServer).UpdateEmailOfEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeDB_DeleteEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Employee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeDBServer).DeleteEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeGuide.EmployeeDB/DeleteEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeDBServer).DeleteEmployee(ctx, req.(*Employee))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeDB_ServiceDesc is the grpc.ServiceDesc for EmployeeDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routeGuide.EmployeeDB",
	HandlerType: (*EmployeeDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEmployee",
			Handler:    _EmployeeDB_AddEmployee_Handler,
		},
		{
			MethodName: "GetEmployees",
			Handler:    _EmployeeDB_GetEmployees_Handler,
		},
		{
			MethodName: "UpdateEmailOfEmployee",
			Handler:    _EmployeeDB_UpdateEmailOfEmployee_Handler,
		},
		{
			MethodName: "DeleteEmployee",
			Handler:    _EmployeeDB_DeleteEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "employee.proto",
}
