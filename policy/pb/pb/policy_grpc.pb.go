// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PolicyClient is the client API for Policy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyClient interface {
	Add(ctx context.Context, in *Input, opts ...grpc.CallOption) (*AddRep, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Placeholder, error)
	Check(ctx context.Context, in *RuleOptions, opts ...grpc.CallOption) (*CheckRep, error)
}

type policyClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyClient(cc grpc.ClientConnInterface) PolicyClient {
	return &policyClient{cc}
}

func (c *policyClient) Add(ctx context.Context, in *Input, opts ...grpc.CallOption) (*AddRep, error) {
	out := new(AddRep)
	err := c.cc.Invoke(ctx, "/policy.Policy/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*Placeholder, error) {
	out := new(Placeholder)
	err := c.cc.Invoke(ctx, "/policy.Policy/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyClient) Check(ctx context.Context, in *RuleOptions, opts ...grpc.CallOption) (*CheckRep, error) {
	out := new(CheckRep)
	err := c.cc.Invoke(ctx, "/policy.Policy/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyServer is the server API for Policy service.
// All implementations must embed UnimplementedPolicyServer
// for forward compatibility
type PolicyServer interface {
	Add(context.Context, *Input) (*AddRep, error)
	Delete(context.Context, *DeleteReq) (*Placeholder, error)
	Check(context.Context, *RuleOptions) (*CheckRep, error)
	mustEmbedUnimplementedPolicyServer()
}

// UnimplementedPolicyServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyServer struct {
}

func (UnimplementedPolicyServer) Add(context.Context, *Input) (*AddRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedPolicyServer) Delete(context.Context, *DeleteReq) (*Placeholder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPolicyServer) Check(context.Context, *RuleOptions) (*CheckRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedPolicyServer) mustEmbedUnimplementedPolicyServer() {}

// UnsafePolicyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyServer will
// result in compilation errors.
type UnsafePolicyServer interface {
	mustEmbedUnimplementedPolicyServer()
}

func RegisterPolicyServer(s grpc.ServiceRegistrar, srv PolicyServer) {
	s.RegisterService(&Policy_ServiceDesc, srv)
}

func _Policy_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Input)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy.Policy/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).Add(ctx, req.(*Input))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy.Policy/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Policy_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/policy.Policy/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyServer).Check(ctx, req.(*RuleOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// Policy_ServiceDesc is the grpc.ServiceDesc for Policy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Policy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "policy.Policy",
	HandlerType: (*PolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Policy_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Policy_Delete_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Policy_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policy.proto",
}
