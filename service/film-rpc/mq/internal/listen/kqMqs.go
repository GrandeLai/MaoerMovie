package listen

import (
	"MaoerMovie/service/film-rpc/mq/internal/config"
	kq2 "MaoerMovie/service/film-rpc/mq/internal/mqs/kq"
	"MaoerMovie/service/film-rpc/mq/internal/svc"
	"context"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// 创建kafka消息队列
func KqMqs(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqFilmInsert, kq2.NewFilmInsertMq(ctx, svcCtx)),
		kq.MustNewQueue(c.KqFilmUpdate, kq2.NewFilmUpdateMq(ctx, svcCtx)),
		kq.MustNewQueue(c.KqActorInsert, kq2.NewActorInsertMq(ctx, svcCtx)),
	}
}
