package auth

import (
	"MaoerMovie/service/user-rpc/types/user"
	"context"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	return &SendEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(req *types.SendEmailCodeRequest) (resp *types.SendEmailCodeResponse, err error) {
	resp = new(types.SendEmailCodeResponse)
	_, err = l.svcCtx.UserRPC.SendEmailCode(l.ctx, &user.SendEmailCodeRequest{Email: req.Email})
	if err != nil {
		return
	}
	return
}
