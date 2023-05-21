package comment

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"MaoerMovie/service/comment-rpc/types/pb"
	"MaoerMovie/service/user-rpc/types/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.GetCommentListRequest) (resp *types.GetCommentListResponse, err error) {
	resp = new(types.GetCommentListResponse)
	rpcResp, err := l.svcCtx.CommentRPC.GetCommentList(l.ctx, &pb.GetCommentListRequest{
		FilmId: req.FilmId,
		Page:   req.Page,
		Size:   req.Size,
	})
	var commentList []types.UserComment
	for _, comment := range rpcResp.List {
		user, err := l.svcCtx.UserRPC.GetUserPreview(l.ctx, &user.GetUserInfoRequest{Id: comment.UserId})
		if err != nil {
			return nil, err
		}
		commentList = append(commentList, types.UserComment{
			UserPreview: types.UserPreview{
				Id:        user.Id,
				Name:      user.Name,
				AvatarUrl: user.AvatarUrl,
			},
			CommentId: comment.CommentId,
			Content:   comment.Content,
			Score:     comment.Score,
		})
	}
	resp.List = commentList
	resp.Total = utils.IntToString(len(commentList))
	return
}
