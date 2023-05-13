package kq

import (
	"MaoerMovie/common/kqueue"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/film-rpc/model"
	"MaoerMovie/service/film-rpc/mq/internal/svc"
	"context"
	jsoniter "github.com/json-iterator/go"
	"strings"
)

type ActorInsertMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActorInsertMq(ctx context.Context, svcCtx *svc.ServiceContext) *ActorInsertMq {
	return &ActorInsertMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 消费函数Consume，集成了Kafka的go-zero框架会自动识别这个消费函数
func (l *ActorInsertMq) Consume(_, val string) error {
	var message kqueue.ActorInsert
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
func (l *ActorInsertMq) execService(message kqueue.ActorInsert) error {
	parts := strings.Split(message.ActorList, "/")
	roleParts := strings.Split(message.RoleList, "/")
	for i, part := range parts {
		actor, err := l.svcCtx.ActorModel.FindByActorName(l.ctx, part)
		if err != nil {
			return err
		}
		filmActor := &model.FilmActor{
			Id:       utils.GenerateNewId(l.svcCtx.RedisClient, "film_actor"),
			FilmId:   message.FilmId,
			ActorId:  actor.Id,
			RoleName: roleParts[i],
		}
		_, err = l.svcCtx.FilmActorModel.InsertWithNewId(l.ctx, filmActor)
		if err != nil {
			return err
		}
	}
	return nil
}
