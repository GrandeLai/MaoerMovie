package svc

import (
	"MaoerMovie/service/cinema-rpc/internal/config"
	"MaoerMovie/service/cinema-rpc/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	RedisClient     *redis.Redis
	BrandModel      model.BrandModel
	CinemaModel     model.CinemaModel
	DistrictModel   model.DistrictModel
	HallModel       model.HallModel
	CinemaFilmModel model.CinemaFilmModel
	ShowModel       model.ShowModel
	MinioClient     *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	redisClient := redis.New(c.RedisClient.Host, func(r *redis.Redis) {
		r.Type = c.RedisClient.Type
		r.Pass = c.RedisClient.Pass
	})
	minioClient, err := minio.New(c.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.MinIO.AccessKey, c.MinIO.SecretKey, ""),
		Secure: c.MinIO.UseSSL})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		RedisClient:     redisClient,
		Config:          c,
		BrandModel:      model.NewBrandModel(conn, c.CacheRedis),
		CinemaModel:     model.NewCinemaModel(conn, c.CacheRedis),
		DistrictModel:   model.NewDistrictModel(conn, c.CacheRedis),
		HallModel:       model.NewHallModel(conn, c.CacheRedis),
		CinemaFilmModel: model.NewCinemaFilmModel(conn, c.CacheRedis),
		ShowModel:       model.NewShowModel(conn, c.CacheRedis),
		MinioClient:     minioClient,
	}
}
