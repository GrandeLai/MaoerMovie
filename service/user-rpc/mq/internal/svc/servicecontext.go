package svc

import (
	"MaoerMovie/service/user-rpc/model"
	"MaoerMovie/service/user-rpc/mq/internal/config"
	"MaoerMovie/service/user-rpc/mq/internal/mqs/esClient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	UserModel   model.UserModel
	EsModel     *esClient.EsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.RedisClient.Host, func(r *redis.Redis) {
			r.Type = c.RedisClient.Type
			r.Pass = c.RedisClient.Pass
		}),
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		EsModel:   esClient.NewEsModel(c.Elasticsearch.Addresses, c.Elasticsearch.Username, c.Elasticsearch.Password),
	}
}
