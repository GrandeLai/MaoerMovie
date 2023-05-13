package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/utils"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(in *user.UpdatePasswordRequest) (*user.UpdatePasswordResponse, error) {
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, utils.StringToInt64(in.Id))
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的用户！")
	default:
		return nil, err
	}
	if in.Password != userInfo.Password {
		return nil, errorx.NewCodeError(100, "原密码错误！")
	}
	userInfo.Password = in.NewPassword
	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &user.UpdatePasswordResponse{}, nil
}
