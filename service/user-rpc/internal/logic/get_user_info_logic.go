package logic

import (
	"MaoerMovie/common/utils"
	"context"
	"strconv"

	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	userId, err := strconv.ParseInt(in.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	return &user.GetUserInfoResponse{
		Id:        in.Id,
		Name:      userInfo.Name,
		Gender:    utils.Int64ToString(userInfo.Gender),
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		AvatarUrl: userInfo.AvatarUrl.String,
	}, nil
}
