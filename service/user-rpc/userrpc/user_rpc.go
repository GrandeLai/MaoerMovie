// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userrpc

import (
	"context"

	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoRequest     = user.GetUserInfoRequest
	GetUserInfoResponse    = user.GetUserInfoResponse
	LoginRequest           = user.LoginRequest
	LoginResponse          = user.LoginResponse
	RefreshAuthRequest     = user.RefreshAuthRequest
	RefreshAuthResponse    = user.RefreshAuthResponse
	RegisterRequest        = user.RegisterRequest
	RegisterResponse       = user.RegisterResponse
	SearchUserRequest      = user.SearchUserRequest
	SearchUserResponse     = user.SearchUserResponse
	SendEmailCodeRequest   = user.SendEmailCodeRequest
	SendEmailCodeResponse  = user.SendEmailCodeResponse
	UpdatePasswordRequest  = user.UpdatePasswordRequest
	UpdatePasswordResponse = user.UpdatePasswordResponse
	UpdateUserInfoRequest  = user.UpdateUserInfoRequest
	UpdateUserInfoResponse = user.UpdateUserInfoResponse
	UserPreview            = user.UserPreview

	UserRpc interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
		RefreshAuth(ctx context.Context, in *RefreshAuthRequest, opts ...grpc.CallOption) (*RefreshAuthResponse, error)
		UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
		SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
		GetUserPreview(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserPreview, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

func (m *defaultUserRpc) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserRpc) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserRpc) SendEmailCode(ctx context.Context, in *SendEmailCodeRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.SendEmailCode(ctx, in, opts...)
}

func (m *defaultUserRpc) RefreshAuth(ctx context.Context, in *RefreshAuthRequest, opts ...grpc.CallOption) (*RefreshAuthResponse, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.RefreshAuth(ctx, in, opts...)
}

func (m *defaultUserRpc) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.UpdatePassword(ctx, in, opts...)
}

func (m *defaultUserRpc) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

func (m *defaultUserRpc) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserRpc) SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.SearchUser(ctx, in, opts...)
}

func (m *defaultUserRpc) GetUserPreview(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserPreview, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.GetUserPreview(ctx, in, opts...)
}
