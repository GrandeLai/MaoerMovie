package kq

import (
	"MaoerMovie/common/kqueue"
	"MaoerMovie/service/user-rpc/mq/internal/svc"
	"context"
	jsoniter "github.com/json-iterator/go"
)

type UserUpdateMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateMq {
	return &UserUpdateMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消费函数Consume，集成了Kafka的go-zero框架会自动识别这个消费函数
func (l *UserUpdateMq) Consume(_, val string) error {
	var message kqueue.UserUpdateMessage
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.UnmarshalFromString(val, &message)
	if err != nil {
		return err
	}
	err = l.execService(message)
	if err != nil {
		return err
	}
	return nil
}

// 执行业务
func (l *UserUpdateMq) execService(message kqueue.UserUpdateMessage) error {
	return l.svcCtx.EsModel.UpdateUser(message)
}
