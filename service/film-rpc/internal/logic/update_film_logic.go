package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/kqueue"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/film-rpc/model"
	"bytes"
	"context"
	"database/sql"
	jsoniter "github.com/json-iterator/go"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"path"
	"strings"
	"time"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFilmLogic {
	return &UpdateFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFilmLogic) UpdateFilm(in *pb.FilmUpdateRequest) (*pb.FilmResponse, error) {
	film, err := l.svcCtx.FilmModel.FindOne(l.ctx, utils.StringToInt64(in.FilmId))
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的电影！")
	default:
		return nil, err
	}

	if in.FilmName != "" {
		film.FilmName = in.FilmName
	}
	if in.FilmEnglishName != "" {
		film.FilmEnName = in.FilmEnglishName
	}
	if in.FilmType != "" {
		film.FilmType = utils.StringToInt64(in.FilmType)
	}
	cateName := ""
	if in.FilmCategory != "" {
		film.CategoryId = utils.StringToInt64(in.FilmCategory)
		category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, film.CategoryId)
		if err != nil {
			return nil, err
		}
		cateName = category.CateName
	}
	if in.FilmArea != "" {
		film.FilmArea = in.FilmArea
	}
	if in.FilmArea != "" {
		film.FilmLength = utils.StringToInt64(in.FilmLength)
	}
	if in.FilmTime != "" {
		filmTime, err := time.ParseInLocation("2006-01-02 15:04:05", in.FilmTime, time.Local)
		if err != nil {
			return nil, err
		}
		film.FilmTime = sql.NullTime{Time: filmTime}
	}
	if in.Director != "" {
		film.DirectorId = utils.StringToInt64(in.Director)
	}
	if in.Biography != "" {
		film.Biography = in.Biography
	}
	if in.FilmCoverName != "" && len(in.FilmCover) != 0 {
		if film.FilmCover != "" {
			url := strings.Split(film.FilmCover, "/")
			objectName := url[len(url)-1]
			err = l.svcCtx.MinioClient.RemoveObject(l.ctx, l.svcCtx.Config.MinIO.BucketName, objectName, minio.RemoveObjectOptions{})
			if err != nil {
				return nil, errorx.NewDefaultError("服务器删除对象失败")
			}
		}
		objectName := utils.GenerateUUID() + path.Ext(in.FilmCoverName)
		_, err = l.svcCtx.MinioClient.PutObject(context.Background(), l.svcCtx.Config.MinIO.BucketName, objectName, bytes.NewReader(in.FilmCover), int64(len(in.FilmCover)),
			minio.PutObjectOptions{ContentType: "binary/octet-stream"})
		if err != nil {
			//return nil, errorx.NewDefaultError("服务器处理上传失败！")
			return nil, err
		}
		film.FilmCover = l.svcCtx.Config.MinIO.BucketName + "/" + objectName
	}
	err = l.svcCtx.FilmModel.Update(l.ctx, film)
	if err != nil {
		return nil, err
	}
	filmInfo, err := l.svcCtx.FilmInfoModel.FindOneByFilmId(l.ctx, utils.StringToInt64(in.FilmId))
	if in.FilmPreSaleNum != "" {
		filmInfo.FilmPreSaleNum = utils.StringToInt64(in.FilmPreSaleNum)
	}
	err = l.svcCtx.FilmInfoModel.Update(l.ctx, filmInfo)
	if err != nil {
		return nil, err
	}

	score, err := l.svcCtx.FilmScoreModel.FindOneByFilmId(l.ctx, utils.StringToInt64(in.FilmId))
	err = l.PubKqFilmUpdateMessage(film, cateName, utils.Float64ToString(score.FilmScore))
	if err != nil {
		return nil, err
	}
	return &pb.FilmResponse{}, nil
}

func (l *UpdateFilmLogic) PubKqFilmUpdateMessage(film *model.Film, category string, score string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	message := kqueue.FilmUpdateMessage{
		FilmId:       film.Id,
		FilmName:     film.FilmName,
		FilmTime:     utils.TimeToString(film.FilmTime.Time),
		FilmCategory: category,
		FilmArea:     film.FilmArea,
		FilmCoverUrl: film.FilmCover,
		FilmScore:    score,
	}
	if category != "" {
		message.FilmCategory = category
	}
	jsonStr, err := json.MarshalToString(message)
	if err != nil {
		return err
	}
	err = l.svcCtx.KqFilmInsertClient.Push(jsonStr)
	if err != nil {
		return err
	}
	return nil
}
