package listen

import (
	"MaoerMovie/service/user-rpc/mq/internal/config"
	kq2 "MaoerMovie/service/user-rpc/mq/internal/mqs/kq"
	"MaoerMovie/service/user-rpc/mq/internal/svc"
	"context"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// 创建kafka消息队列
func KqMqs(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.KqUserInsert, kq2.NewUserInsertMq(ctx, svcCtx)),
		kq.MustNewQueue(c.KqUserUpdate, kq2.NewUserUpdateMq(ctx, svcCtx)),
	}
}
