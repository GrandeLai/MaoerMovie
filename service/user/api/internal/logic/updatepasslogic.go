package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"MaoerMovie/service/user/api/internal/svc"
	"MaoerMovie/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePassLogic {
	return &UpdatePassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePassLogic) UpdatePass(req *types.UpdatePassRequest) (resp *types.UpdatePassResponse, err error) {
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的用户！")
	default:
		return nil, err
	}
	password := utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	if password != userInfo.Password {
		return nil, errorx.NewCodeError(100, "原密码错误！")
	}
	userInfo.Password = utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.NewPassword)
	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &types.UpdatePassResponse{}, nil
}
