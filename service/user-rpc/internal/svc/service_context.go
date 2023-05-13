package svc

import (
	"MaoerMovie/service/user-rpc/internal/config"
	"MaoerMovie/service/user-rpc/model"
	"MaoerMovie/service/user-rpc/model/esClient"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	RedisClient        *redis.Redis
	UserModel          model.UserModel
	MinioClient        *minio.Client
	KqUserInsertClient *kq.Pusher
	KqUserUpdateClient *kq.Pusher
	EsModel            *esClient.EsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	//esClient, err := elasticsearch.NewClient(elasticsearch.Config{
	//	Addresses: c.Elasticsearch.Addresses,
	//	Username:  c.Elasticsearch.Username,
	//	Password:  c.Elasticsearch.Password,
	//	Transport: &http.Transport{
	//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}
	minioClient, err := minio.New(c.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.MinIO.AccessKey, c.MinIO.SecretKey, ""),
		Secure: c.MinIO.UseSSL})

	redisClient := redis.New(c.RedisClient.Host, func(r *redis.Redis) {
		r.Type = c.RedisClient.Type
		r.Pass = c.RedisClient.Pass
	})
	if err != nil {
		return nil
	}
	return &ServiceContext{
		Config:             c,
		RedisClient:        redisClient,
		UserModel:          model.NewUserModel(conn, c.CacheRedis),
		MinioClient:        minioClient,
		EsModel:            esClient.NewEsModel(c.Elasticsearch.Addresses, c.Elasticsearch.Username, c.Elasticsearch.Password),
		KqUserInsertClient: kq.NewPusher(c.KqUserInsert.Brokers, c.KqUserInsert.Topic),
		KqUserUpdateClient: kq.NewPusher(c.KqUserUpdate.Brokers, c.KqUserUpdate.Topic),
	}
}
