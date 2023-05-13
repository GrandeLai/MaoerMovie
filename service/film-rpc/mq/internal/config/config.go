package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ElasticsearchConf struct {
	Addresses []string
	Username  string
	Password  string
}

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis    cache.CacheConf
	RedisClient   redis.RedisConf
	Elasticsearch ElasticsearchConf
	KqFilmInsert  kq.KqConf
	KqFilmUpdate  kq.KqConf
	KqActorInsert kq.KqConf
}
