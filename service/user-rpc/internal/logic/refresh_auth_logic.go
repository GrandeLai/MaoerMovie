package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/utils"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"

	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthLogic {
	return &RefreshAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefreshAuthLogic) RefreshAuth(in *user.RefreshAuthRequest) (*user.RefreshAuthResponse, error) {
	//获得原token的剩余信息
	restClaims := make(jwt.MapClaims)
	judgeValid, err := jwt.ParseWithClaims(in.Authorization, restClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.svcCtx.Config.AuthAccess.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	//判断是否token有效
	if !judgeValid.Valid {
		return nil, errorx.NewCodeError(100, "Token已失效！")
	}
	//利用过期token的其他值，生成新token等信息
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.AuthAccess.AccessExpire
	userId, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	jwtToken, toExpire, err := utils.GenerateJwtToken(l.svcCtx.Config.AuthAccess.AccessSecret, accessExpire, userId)
	if err != nil {
		return nil, err
	}

	return &user.RefreshAuthResponse{
		AccessToken:  jwtToken,
		AccessExpire: utils.Int64ToString(toExpire),
		RefreshAfter: utils.Int64ToString(now + accessExpire/2),
	}, nil
}
