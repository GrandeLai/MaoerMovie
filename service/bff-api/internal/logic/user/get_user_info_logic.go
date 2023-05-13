package user

import (
	"MaoerMovie/common/utils"
	"context"
	"encoding/json"
	"fmt"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResponse, err error) {
	resp = new(types.GetUserInfoResponse)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	data, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &user.GetUserInfoRequest{Id: utils.Int64ToString(userId)})
	if err != nil {
		return
	}
	resp.Id = utils.Int64ToString(userId)
	resp.Name = data.Name
	resp.Gender = data.Gender
	resp.AvatarUrl = data.AvatarUrl
	resp.Email = data.Email
	resp.Phone = data.Phone
	return
}
