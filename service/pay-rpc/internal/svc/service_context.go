package svc

import (
	"MaoerMovie/service/pay-rpc/internal/config"
	"MaoerMovie/service/pay-rpc/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	PayModel    model.PayModel
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
		PayModel:    model.NewPayModel(conn, c.CacheRedis),
	}
}
