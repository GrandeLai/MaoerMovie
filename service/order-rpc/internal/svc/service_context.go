package svc

import (
	"MaoerMovie/service/order-rpc/internal/config"
	"MaoerMovie/service/order-rpc/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	OrderModel  model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.RedisClient.Host, func(r *redis.Redis) {
		r.Type = c.RedisClient.Type
		r.Pass = c.RedisClient.Pass
	})
	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,
		OrderModel:  model.NewOrderModel(conn, c.CacheRedis),
	}
}
