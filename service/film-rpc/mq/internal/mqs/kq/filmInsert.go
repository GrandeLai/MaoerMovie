package kq

import (
	"MaoerMovie/common/kqueue"
	"MaoerMovie/service/film-rpc/mq/internal/svc"
	"context"
	jsoniter "github.com/json-iterator/go"
)

type FilmInsertMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmInsertMq(ctx context.Context, svcCtx *svc.ServiceContext) *FilmInsertMq {
	return &FilmInsertMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消费函数Consume，集成了Kafka的go-zero框架会自动识别这个消费函数
func (l *FilmInsertMq) Consume(_, val string) error {
	var message kqueue.FilmInsertMessage
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
func (l *FilmInsertMq) execService(message kqueue.FilmInsertMessage) error {
	return l.svcCtx.EsModel.InsertFilmToES(message)
}
