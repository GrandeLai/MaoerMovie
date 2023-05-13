package svc

import (
	"MaoerMovie/service/film-rpc/internal/config"
	"MaoerMovie/service/film-rpc/model"
	"MaoerMovie/service/film-rpc/model/esClient"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	ActorModel     model.ActorModel
	CategoryModel  model.CategoryModel
	FilmModel      model.FilmModel
	FilmActorModel model.FilmActorModel
	FilmInfoModel  model.FilmInfoModel
	FilmScoreModel model.FilmScoreModel

	KqFilmInsertClient  *kq.Pusher
	KqFilmUpdateClient  *kq.Pusher
	KqActorInsertClient *kq.Pusher
	MinioClient         *minio.Client
	EsModel             *esClient.EsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
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
		Config:              c,
		RedisClient:         redisClient,
		ActorModel:          model.NewActorModel(conn, c.CacheRedis),
		CategoryModel:       model.NewCategoryModel(conn, c.CacheRedis),
		FilmModel:           model.NewFilmModel(conn, c.CacheRedis),
		FilmActorModel:      model.NewFilmActorModel(conn, c.CacheRedis),
		FilmInfoModel:       model.NewFilmInfoModel(conn, c.CacheRedis),
		FilmScoreModel:      model.NewFilmScoreModel(conn, c.CacheRedis),
		KqFilmInsertClient:  kq.NewPusher(c.KqFilmInsert.Brokers, c.KqFilmInsert.Topic),
		KqFilmUpdateClient:  kq.NewPusher(c.KqFilmUpdate.Brokers, c.KqFilmUpdate.Topic),
		KqActorInsertClient: kq.NewPusher(c.KqActorInsert.Brokers, c.KqActorInsert.Topic),
		EsModel:             esClient.NewEsModel(c.Elasticsearch.Addresses, c.Elasticsearch.Username, c.Elasticsearch.Password),
		MinioClient:         minioClient,
	}
}
