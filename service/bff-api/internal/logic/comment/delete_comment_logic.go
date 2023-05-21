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

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentRequest) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	_, err = l.svcCtx.CommentRPC.DeleteComment(l.ctx, &pb.DeleteCommentRequest{
		CommentId: req.CommentId,
		UserId:    utils.Int64ToString(userId),
	})
	if err != nil {
		return nil, err
	}
	return
}
