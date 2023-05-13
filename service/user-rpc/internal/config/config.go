package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ElasticsearchConf struct {
	Addresses []string
	Username  string
	Password  string
}

type Config struct {
	zrpc.RpcServerConf
	AuthAccess struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
	MinIO struct {
		Endpoint   string
		AccessKey  string
		SecretKey  string
		UseSSL     bool
		BucketName string
	}
	CacheRedis    cache.CacheConf
	RedisClient   redis.RedisConf
	Elasticsearch ElasticsearchConf
	KqUserInsert  KqConfig
	KqUserUpdate  KqConfig
}
