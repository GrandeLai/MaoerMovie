package logic

import (
	"MaoerMovie/common/kqueue"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/model"
	"MaoerMovie/service/film-rpc/types/pb"
	"bytes"
	"context"
	"database/sql"
	jsoniter "github.com/json-iterator/go"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"path"
	"strconv"
	"time"
)

type CreateFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFilmLogic {
	return &CreateFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFilmLogic) CreateFilm(in *pb.FilmCreateRequest) (*pb.FilmCreateResponse, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	filmTime, err := time.ParseInLocation("2006-01-02", in.FilmTime, time.Local)
	if err != nil {
		return nil, err
	}
	filmId := utils.GenerateNewId(l.svcCtx.RedisClient, "film")
	film := &model.Film{
		Id:         filmId,
		FilmName:   in.FilmName,
		FilmEnName: in.FilmEnglishName,
		FilmType:   utils.StringToInt64(in.FilmType),
		FilmLength: utils.StringToInt64(in.FilmLength),
		CategoryId: utils.StringToInt64(in.CategoryId),
		FilmArea:   in.FilmArea,
		FilmTime:   sql.NullTime{Time: filmTime},
		DirectorId: utils.StringToInt64(in.DirectorId),
		Biography:  in.Biography,
	}
	cacheFilmKey := utils.CacheFilmKey + strconv.FormatInt(film.Id, 10)
	cacheActorListKey := utils.CacheActorListKey + strconv.FormatInt(film.Id, 10)
	if len(in.FilmCover) != 0 {
		objectName := utils.GenerateUUID() + path.Ext(in.CoverName)
		_, err := l.svcCtx.MinioClient.PutObject(context.Background(), l.svcCtx.Config.MinIO.BucketName, objectName, bytes.NewReader(in.FilmCover), int64(len(in.FilmCover)),
			minio.PutObjectOptions{ContentType: "binary/octet-stream"})
		if err != nil {
			//return nil, errorx.NewDefaultError("服务器处理上传失败！")
			return nil, err
		}
		cover := l.svcCtx.Config.MinIO.BucketName + "/" + objectName
		film.FilmCover = cover
	}
	cate, err := l.svcCtx.CategoryModel.FindOne(l.ctx, film.CategoryId)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.FilmModel.InsertWithNewId(l.ctx, film)
	if err != nil {
		return nil, err
	}
	//发生消息插入演员表
	err = l.PubKqActorInsertMessage(film.Id, in.ActorList, in.RoleList)
	if err != nil {
		return nil, err
	}

	filmJson, err := json.MarshalToString(film)
	if err != nil {
		return nil, err
	}
	err = l.PubKqFilmInsertMessage(film, in.ActorList, cate.CateName)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.Set(cacheFilmKey, filmJson)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.Set(cacheActorListKey, in.ActorList)
	if err != nil {
		return nil, err
	}
	return &pb.FilmCreateResponse{Id: utils.Int64ToString(film.Id)}, nil
}

func (l *CreateFilmLogic) PubKqFilmInsertMessage(film *model.Film, actorList string, category string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonStr, err := json.MarshalToString(kqueue.FilmInsertMessage{
		FilmId:        film.Id,
		FilmName:      film.FilmName,
		FilmTime:      utils.TimeToString(film.FilmTime.Time),
		FilmCategory:  category,
		FilmArea:      film.FilmArea,
		FilmCoverUrl:  film.FilmCover,
		ActorNameList: actorList,
	})
	if err != nil {
		return err
	}
	err = l.svcCtx.KqFilmInsertClient.Push(jsonStr)
	if err != nil {
		return err
	}
	return nil
}

func (l *CreateFilmLogic) PubKqActorInsertMessage(filmId int64, actorList string, roleList string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonStr, err := json.MarshalToString(kqueue.ActorInsert{
		ActorList: actorList,
		FilmId:    filmId,
		RoleList:  roleList,
	})
	if err != nil {
		return err
	}
	err = l.svcCtx.KqActorInsertClient.Push(jsonStr)
	if err != nil {
		return err
	}
	return nil
}
