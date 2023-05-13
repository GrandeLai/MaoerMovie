package auth

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/user-rpc/types/user"
	"context"
	"encoding/json"
	"fmt"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthLogic {
	return &RefreshAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthLogic) RefreshAuth(Authorization string) (resp *types.RefreshAuthResponse, err error) {
	resp = new(types.RefreshAuthResponse)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	rpcResp, err := l.svcCtx.UserRPC.RefreshAuth(l.ctx, &user.RefreshAuthRequest{
		Authorization: Authorization,
		Id:            utils.Int64ToString(userId),
	})
	if err != nil {
		return
	}
	resp.AccessToken = rpcResp.AccessToken
	resp.RefreshAfter = rpcResp.RefreshAfter
	resp.AccessExpire = rpcResp.AccessExpire
	return
}
