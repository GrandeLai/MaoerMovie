package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/kqueue"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/model"
	"MaoerMovie/service/user-rpc/types/user"
	"context"
	jsoniter "github.com/json-iterator/go"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	//判断输入邮箱验证码是否正确
	verificationCode, err := l.svcCtx.RedisClient.Get(utils.CacheEmailCodeKey + in.Email)
	if err != nil || verificationCode == "" {
		return nil, errorx.NewCodeError(100, "无发送验证码或验证码已到期！")
	}
	if verificationCode != in.EmailCode {
		return nil, errorx.NewCodeError(100, "输入的验证码不一致！")
	}
	//判断输入手机验证码是否正确
	//判断该手机号是否已被注册
	count, err := l.svcCtx.UserModel.CountByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errorx.NewCodeError(100, "该手机号已被注册！")
	}
	//判断该邮箱是否已被注册
	count, err = l.svcCtx.UserModel.CountByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errorx.NewCodeError(100, "该邮箱已被注册！")
	}
	//插入数据库
	res, err := l.svcCtx.UserModel.InsertWithNewId(l.ctx, &model.User{
		Id:     utils.GenerateNewId(l.svcCtx.RedisClient, "user"),
		Name:   in.Name,
		Gender: utils.StringToInt64(in.Gender),
		Phone:  in.Phone,
		//Password: utils.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		Password: in.Password,
		Email:    in.Email,
		Status:   0,
	})
	if err != nil {
		return nil, err
	}
	userId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	err = l.PubKqUserInsertMessage(userId, in.Name)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResponse{Id: utils.Int64ToString(userId)}, nil
}

func (l *RegisterLogic) PubKqUserInsertMessage(id int64, Name string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonStr, err := json.MarshalToString(kqueue.UserInsertMessage{
		Id:   id,
		Name: Name,
	})
	if err != nil {
		return err
	}
	err = l.svcCtx.KqUserInsertClient.Push(jsonStr)
	if err != nil {
		return err
	}
	return nil
}
