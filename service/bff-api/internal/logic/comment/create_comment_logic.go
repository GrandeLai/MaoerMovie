package comment

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/comment-rpc/types/pb"
	"context"
	"encoding/json"
	"fmt"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentRequest) (resp *types.CreateCommentResponse, err error) {
	resp = new(types.CreateCommentResponse)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	rpcResp, err := l.svcCtx.CommentRPC.CreateComment(l.ctx, &pb.CreateCommentRequest{
		FilmId:  req.FilmId,
		UserId:  utils.Int64ToString(userId),
		Content: req.Content,
		Score:   req.Score,
	})
	resp.CommentId = rpcResp.CommentId
	return
}
