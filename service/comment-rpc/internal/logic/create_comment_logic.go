package logic

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/comment-rpc/internal/svc"
	"MaoerMovie/service/comment-rpc/model"
	"MaoerMovie/service/comment-rpc/types/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCommentLogic) CreateComment(in *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	commentId := utils.GenerateNewId(l.svcCtx.RedisClient, "comment")
	userId := utils.StringToInt64(in.UserId)
	filmId := utils.StringToInt64(in.FilmId)
	score := utils.StringToFloat64(in.Score)
	_, err := l.svcCtx.CommentModel.InsertWithNewId(l.ctx, &model.Comment{
		Id:      commentId,
		UserId:  userId,
		FilmId:  filmId,
		Content: in.Content,
		Score:   score,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateCommentResponse{CommentId: utils.Int64ToString(commentId)}, nil
}
