package svc

import (
	"MaoerMovie/service/film-rpc/model"
	"MaoerMovie/service/film-rpc/mq/internal/config"
	"MaoerMovie/service/film-rpc/mq/internal/mqs/esClient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Redis
	EsModel        *esClient.EsModel
	FilmActorModel model.FilmActorModel
	ActorModel     model.ActorModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.RedisClient.Host, func(r *redis.Redis) {
			r.Type = c.RedisClient.Type
			r.Pass = c.RedisClient.Pass
		}),
		ActorModel:     model.NewActorModel(conn, c.CacheRedis),
		FilmActorModel: model.NewFilmActorModel(conn, c.CacheRedis),
		EsModel:        esClient.NewEsModel(c.Elasticsearch.Addresses, c.Elasticsearch.Username, c.Elasticsearch.Password),
	}
}
