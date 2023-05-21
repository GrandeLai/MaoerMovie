package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPreviewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPreviewLogic {
	return &GetUserPreviewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPreviewLogic) GetUserPreview(in *user.GetUserInfoRequest) (*user.UserPreview, error) {
	userId := utils.StringToInt64(in.Id)
	userModel, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	return &user.UserPreview{
		Id:        in.Id,
		AvatarUrl: userModel.AvatarUrl.String,
		Name:      userModel.Name,
	}, nil
}
