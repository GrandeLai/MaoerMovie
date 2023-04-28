package logic

import (
	"MaoerMovie/common/utils"
	"context"
	"crypto/tls"
	"net/smtp"

	"MaoerMovie/service/user/api/internal/svc"
	"MaoerMovie/service/user/api/internal/types"
	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCodeSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCodeSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeSendLogic {
	return &EmailCodeSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCodeSendLogic) EmailCodeSend(req *types.EmailCodeSendRequest) (resp *types.EmailCodeSendResponse, err error) {
	e := email.NewEmail()
	e.From = "VerificationCode by 猫耳电影 <" + utils.ServerEmail + ">"
	e.To = []string{req.Email}
	e.Subject = "This is a VerificationCode!"
	verificationCode := utils.GenerateVerificationCode()
	e.Text = []byte(verificationCode)
	err = e.SendWithTLS(utils.EmailSmtpAddr, smtp.PlainAuth("", utils.ServerEmail, utils.EmailAuthCode, utils.EmailSmtpHost), &tls.Config{InsecureSkipVerify: true, ServerName: utils.EmailSmtpHost})
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.Setex("cache:email:code:"+req.Email, verificationCode, utils.EmailCodeExpireSeconds)
	if err != nil {
		return nil, err
	}
	return &types.EmailCodeSendResponse{}, nil
}
