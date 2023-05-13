package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLogic) SearchUser(in *user.SearchUserRequest) (*user.SearchUserResponse, error) {
	data, total, err := l.svcCtx.EsModel.SearchUser(l.ctx, in.Keyword, utils.StringToInt64(in.Page), utils.StringToInt64(in.Size))
	if err != nil {
		return nil, err
	}
	size := utils.StringToInt64(in.Size)
	res := make([]*user.UserPreview, 0, size)
	for _, d := range data {
		m := &user.UserPreview{
			Id:        utils.Int64ToString(d.Id),
			AvatarUrl: d.AvatarUrl.String,
			Name:      d.Name,
		}
		res = append(res, m)
	}
	return &user.SearchUserResponse{
		Users: res,
		Total: utils.Int64ToString(total),
	}, nil
}
