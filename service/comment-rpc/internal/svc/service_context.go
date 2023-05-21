package svc

import (
	"MaoerMovie/service/comment-rpc/internal/config"
	"MaoerMovie/service/comment-rpc/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	RedisClient  *redis.Redis
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.RedisClient.Host, func(r *redis.Redis) {
		r.Type = c.RedisClient.Type
		r.Pass = c.RedisClient.Pass
	})
	return &ServiceContext{
		Config:       c,
		RedisClient:  redisClient,
		CommentModel: model.NewCommentModel(conn, c.CacheRedis),
	}
}
