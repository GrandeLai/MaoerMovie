package auth

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/user-rpc/types/user"
	"context"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	resp = new(types.RegisterResponse)
	password := utils.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	rpcResp, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Name:      req.Name,
		Gender:    req.Gender,
		Phone:     req.Phone,
		Email:     req.Email,
		Password:  password,
		EmailCode: req.EmailCode,
	})
	if err != nil {
		return nil, err
	}
	resp.Id = string(rpcResp.Id)
	return
}
