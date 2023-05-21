package logic

import (
	"MaoerMovie/common/utils"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"MaoerMovie/service/comment-rpc/internal/svc"
	"MaoerMovie/service/comment-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	commonId := utils.StringToInt64(in.CommentId)
	userId := utils.StringToInt64(in.UserId)
	_, err := l.svcCtx.CommentModel.FindOneByUserIdAndCommentId(l.ctx, userId, commonId)
	switch err {
	case nil:
		break
	case sqlc.ErrNotFound:
		return nil, errors.New("用户无此评论")
	default:
		return nil, err
	}
	err = l.svcCtx.CommentModel.Delete(l.ctx, commonId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCommentResponse{}, nil
}
