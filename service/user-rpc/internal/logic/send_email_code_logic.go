package logic

import (
	"MaoerMovie/common/utils"
	"context"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"

	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCodeLogic {
	return &SendEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailCodeLogic) SendEmailCode(in *user.SendEmailCodeRequest) (*user.SearchUserResponse, error) {
	e := email.NewEmail()
	e.From = "VerificationCode by 猫耳电影 <" + utils.ServerEmail + ">"
	e.To = []string{in.Email}
	e.Subject = "This is a VerificationCode!"
	verificationCode := utils.GenerateVerificationCode()
	e.Text = []byte(verificationCode)
	err := e.SendWithTLS(utils.EmailSmtpAddr, smtp.PlainAuth("", utils.ServerEmail, utils.EmailAuthCode, utils.EmailSmtpHost), &tls.Config{InsecureSkipVerify: true, ServerName: utils.EmailSmtpHost})
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.Setex("cache:email:code:"+in.Email, verificationCode, utils.EmailCodeExpireSeconds)
	if err != nil {
		return nil, err
	}
	return &user.SearchUserResponse{}, nil
}
