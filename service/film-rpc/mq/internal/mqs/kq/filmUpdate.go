package kq

import (
	"MaoerMovie/common/kqueue"
	"MaoerMovie/service/film-rpc/mq/internal/svc"
	"context"
	jsoniter "github.com/json-iterator/go"
)

type FilmUpdateMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmUpdateMq(ctx context.Context, svcCtx *svc.ServiceContext) *FilmUpdateMq {
	return &FilmUpdateMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消费函数Consume，集成了Kafka的go-zero框架会自动识别这个消费函数
func (l *FilmUpdateMq) Consume(_, val string) error {
	var message kqueue.FilmUpdateMessage
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
func (l *FilmUpdateMq) execService(message kqueue.FilmUpdateMessage) error {
	return l.svcCtx.EsModel.UpdateFilmToES(message)
}
