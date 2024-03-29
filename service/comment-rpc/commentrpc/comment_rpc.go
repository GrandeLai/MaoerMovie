// Code generated by goctl. DO NOT EDIT!
// Source: comment.proto

package commentrpc

import (
	"context"

	"MaoerMovie/service/comment-rpc/types/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommentPreview         = pb.CommentPreview
	CreateCommentRequest   = pb.CreateCommentRequest
	CreateCommentResponse  = pb.CreateCommentResponse
	DeleteCommentRequest   = pb.DeleteCommentRequest
	DeleteCommentResponse  = pb.DeleteCommentResponse
	GetCommentListRequest  = pb.GetCommentListRequest
	GetCommentListResponse = pb.GetCommentListResponse

	CommentRpc interface {
		CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
		DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
		GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListResponse, error)
	}

	defaultCommentRpc struct {
		cli zrpc.Client
	}
)

func NewCommentRpc(cli zrpc.Client) CommentRpc {
	return &defaultCommentRpc{
		cli: cli,
	}
}

func (m *defaultCommentRpc) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	client := pb.NewCommentRpcClient(m.cli.Conn())
	return client.CreateComment(ctx, in, opts...)
}

func (m *defaultCommentRpc) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	client := pb.NewCommentRpcClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}

func (m *defaultCommentRpc) GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListResponse, error) {
	client := pb.NewCommentRpcClient(m.cli.Conn())
	return client.GetCommentList(ctx, in, opts...)
}
