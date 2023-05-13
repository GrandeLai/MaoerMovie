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

type UpdatePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePasswordLogic {
	return &UpdatePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePasswordLogic) UpdatePassword(req *types.UpdatePasswordRequest) (resp *types.UpdatePasswordResponse, err error) {
	resp = new(types.UpdatePasswordResponse)
	password := utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	newPassword := utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.NewPassword)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserRPC.UpdatePassword(l.ctx, &user.UpdatePasswordRequest{
		Id:          utils.Int64ToString(userId),
		Password:    password,
		NewPassword: newPassword,
	})
	if err != nil {
		return
	}
	return
}
