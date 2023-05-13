package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/utils"
	"context"
	"time"

	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	//根据mobile判断用户是否存在
	count, err := l.svcCtx.UserModel.CountByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errorx.NewCodeError(100, "该手机号码尚未注册过！")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, err
	}
	//用户密码是否正确
	if in.Password != userInfo.Password {
		return nil, errorx.NewCodeError(100, "输入密码不正确！")
	}
	//生成token并返回
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.AuthAccess.AccessExpire
	jwtToken, expireTime, err := utils.GenerateJwtToken(l.svcCtx.Config.AuthAccess.AccessSecret, accessExpire, userInfo.Id)
	if err != nil {
		return nil, err
	}
	return &user.LoginResponse{
		UserId:       utils.Int64ToString(userInfo.Id),
		AccessToken:  jwtToken,
		AccessExpire: utils.Int64ToString(expireTime),
		RefreshAfter: utils.Int64ToString(now + accessExpire/2),
	}, nil
}
