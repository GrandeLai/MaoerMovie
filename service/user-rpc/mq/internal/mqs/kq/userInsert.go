package kq

import (
	"MaoerMovie/common/kqueue"
	"MaoerMovie/service/user-rpc/mq/internal/svc"
	"context"
	jsoniter "github.com/json-iterator/go"
)

type UserInsertMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInsertMq(ctx context.Context, svcCtx *svc.ServiceContext) *UserInsertMq {
	return &UserInsertMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消费函数Consume，集成了Kafka的go-zero框架会自动识别这个消费函数
func (l *UserInsertMq) Consume(_, val string) error {
	var message kqueue.UserInsertMessage
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
func (l *UserInsertMq) execService(message kqueue.UserInsertMessage) error {
	return l.svcCtx.EsModel.InsertUser(message)
}
