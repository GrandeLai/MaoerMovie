package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/comment-rpc/internal/svc"
	"MaoerMovie/service/comment-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *pb.GetCommentListRequest) (*pb.GetCommentListResponse, error) {
	page := utils.StringToInt64(in.Page)
	size := utils.StringToInt64(in.Size)
	filmId := utils.StringToInt64(in.FilmId)
	commentList, err := l.svcCtx.CommentModel.FindAllInPageByFilmId(l.ctx, filmId, page, size)
	if err != nil {
		return nil, err
	}
	var preList []*pb.CommentPreview
	for _, comment := range commentList {
		preList = append(preList, &pb.CommentPreview{
			CommentId: utils.Int64ToString(comment.Id),
			Content:   comment.Content,
			Score:     utils.Float64ToString(comment.Score),
			UserId:    utils.Int64ToString(comment.UserId),
		})
	}
	return &pb.GetCommentListResponse{List: preList}, nil
}
