package user

import (
	"MaoerMovie/service/user-rpc/types/user"
	"context"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUserLogic) SearchUser(req *types.SearchUserRequest) (resp *types.SearchUserResponse, err error) {
	resp = new(types.SearchUserResponse)
	data, err := l.svcCtx.UserRPC.SearchUser(l.ctx, &user.SearchUserRequest{
		Keyword: req.Keyword,
		Page:    req.Limit.Page,
		Size:    req.Limit.Size,
	})
	if err != nil {
		return nil, err
	}
	resp.Total = data.Total
	resp.Users = make([]types.UserPreview, 0, len(data.Users))
	for _, user := range data.Users {
		resp.Users = append(resp.Users, types.UserPreview{
			Id:        user.Id,
			Name:      user.Name,
			AvatarUrl: user.AvatarUrl,
		})
	}
	return
}
