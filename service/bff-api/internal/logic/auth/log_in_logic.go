package auth

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/user-rpc/types/user"
	"context"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogInLogic {
	return &LogInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogInLogic) LogIn(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	resp = new(types.LoginResponse)
	password := utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	rpcResp, err := l.svcCtx.UserRPC.Login(l.ctx, &user.LoginRequest{
		Email:    req.Email,
		Password: password,
	})
	if err != nil {
		return
	}
	userId := utils.StringToInt64(rpcResp.UserId)
	auth := l.svcCtx.Config.Auth
	accessExpire := int64(0)
	resp.AccessToken, accessExpire, err = utils.GenerateJwtToken(auth.AccessSecret, auth.AccessExpire, userId)
	resp.AccessExpire = utils.Int64ToString(accessExpire)
	resp.UserId = rpcResp.UserId
	return
}
