// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: user.proto

package user

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

// UserRpcClient is the client API for UserRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRpcClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
	RefreshAuth(ctx context.Context, in *RefreshAuthRequest, opts ...grpc.CallOption) (*RefreshAuthResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
}

type userRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRpcClient(cc grpc.ClientConnInterface) UserRpcClient {
	return &userRpcClient{cc}
}

func (c *userRpcClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.userRpc/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user.userRpc/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	out := new(SearchUserResponse)
	err := c.cc.Invoke(ctx, "/user.userRpc/SendEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) RefreshAuth(ctx context.Context, in *RefreshAuthRequest, opts ...grpc.CallOption) (*RefreshAuthResponse, error) {
	out := new(RefreshAuthResponse)
	err := c.cc.Invoke(ctx, "/user.userRpc/RefreshAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, "/user.userRpc/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, "/user.userRpc/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/user.userRpc/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	out := new(SearchUserResponse)
	err := c.cc.Invoke(ctx, "/user.userRpc/SearchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRpcServer is the server API for UserRpc service.
// All implementations must embed UnimplementedUserRpcServer
// for forward compatibility
type UserRpcServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	SendEmailCode(context.Context, *SendEmailCodeRequest) (*SearchUserResponse, error)
	RefreshAuth(context.Context, *RefreshAuthRequest) (*RefreshAuthResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error)
	mustEmbedUnimplementedUserRpcServer()
}

// UnimplementedUserRpcServer must be embedded to have forward compatible implementations.
type UnimplementedUserRpcServer struct {
}

func (UnimplementedUserRpcServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserRpcServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserRpcServer) SendEmailCode(context.Context, *SendEmailCodeRequest) (*SearchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailCode not implemented")
}
func (UnimplementedUserRpcServer) RefreshAuth(context.Context, *RefreshAuthRequest) (*RefreshAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAuth not implemented")
}
func (UnimplementedUserRpcServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserRpcServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserRpcServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserRpcServer) SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedUserRpcServer) mustEmbedUnimplementedUserRpcServer() {}

// UnsafeUserRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRpcServer will
// result in compilation errors.
type UnsafeUserRpcServer interface {
	mustEmbedUnimplementedUserRpcServer()
}

func RegisterUserRpcServer(s grpc.ServiceRegistrar, srv UserRpcServer) {
	s.RegisterService(&UserRpc_ServiceDesc, srv)
}

func _UserRpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_SendEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).SendEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/SendEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).SendEmailCode(ctx, req.(*SendEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_RefreshAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).RefreshAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/RefreshAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).RefreshAuth(ctx, req.(*RefreshAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/SearchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).SearchUser(ctx, req.(*SearchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRpc_ServiceDesc is the grpc.ServiceDesc for UserRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.userRpc",
	HandlerType: (*UserRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserRpc_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserRpc_Register_Handler,
		},
		{
			MethodName: "SendEmailCode",
			Handler:    _UserRpc_SendEmailCode_Handler,
		},
		{
			MethodName: "RefreshAuth",
			Handler:    _UserRpc_RefreshAuth_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserRpc_UpdatePassword_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserRpc_UpdateUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserRpc_GetUserInfo_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _UserRpc_SearchUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}