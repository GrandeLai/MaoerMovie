package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
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
	CacheRedis  cache.CacheConf
	RedisClient redis.RedisConf
}
