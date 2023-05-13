package listen

import (
	"MaoerMovie/service/film-rpc/mq/internal/config"
	"MaoerMovie/service/film-rpc/mq/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/service"
)

// 返回所有消息队列
func Mqs(c config.Config) []service.Service {
	ctx := context.Background()
	svcCtx := svc.NewServiceContext(c)
	var services []service.Service
	//kafka消息队列
	services = append(services, KqMqs(c, ctx, svcCtx)...)

	return services
}
