package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/film-rpc/model"
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilmLogic {
	return &GetFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFilmLogic) GetFilm(in *pb.FilmRequest) (*pb.GetFilmResponse, error) {
	film := &pb.Film{}
	//先查询缓存有没有该数据
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	cacheFilmKey := utils.CacheFilmKey + strconv.FormatInt(utils.StringToInt64(in.Id), 10)
	success, err := l.svcCtx.RedisClient.Exists(cacheFilmKey)
	if err != nil {
		return nil, err
	}
	//缓存有该数据
	if success {
		redisInfo, err := l.svcCtx.RedisClient.Get(cacheFilmKey)
		//判断该数据是否为空值
		if redisInfo == "" {
			return nil, errorx.NewCodeError(100, "查无此Id的电影！")
		}
		//json反序列化成对象
		var filmModel model.Film
		err = json.UnmarshalFromString(redisInfo, &filmModel)
		if err != nil {
			return nil, err
		}
		film.FilmName = filmModel.FilmName
		film.FilmEnglishName = filmModel.FilmEnName
		film.FilmType = utils.Int64ToString(filmModel.FilmType)
		film.FilmCover = filmModel.FilmCover
		film.FilmCategory = utils.Int64ToString(filmModel.CategoryId)
		film.FilmLength = utils.Int64ToString(filmModel.FilmLength)
		film.FilmArea = filmModel.FilmArea
		film.FilmTime = utils.TimeToString(filmModel.FilmTime.Time)
		film.Biography = filmModel.Biography
		film.Director = utils.Int64ToString(filmModel.DirectorId)
	} else {
		data, err := l.svcCtx.FilmModel.FindOne(l.ctx, utils.StringToInt64(in.Id))
		switch err {
		case nil:
			film.FilmName = data.FilmName
			film.FilmEnglishName = data.FilmEnName
			film.FilmType = utils.Int64ToString(data.FilmType)
			film.FilmCover = data.FilmCover
			film.FilmLength = utils.Int64ToString(data.FilmLength)
			film.FilmCategory = utils.Int64ToString(data.CategoryId)
			film.FilmArea = data.FilmArea
			film.FilmTime = utils.TimeToString(data.FilmTime.Time)
			film.Director = utils.Int64ToString(data.DirectorId)
			film.Biography = data.Biography
			break
		case sqlx.ErrNotFound:
			l.svcCtx.RedisClient.Setex(cacheFilmKey, "", utils.RedisLockExpireSeconds)
			break
		default:
			return nil, err
		}
	}

	cate, err := l.svcCtx.CategoryModel.FindOne(l.ctx, utils.StringToInt64(film.FilmCategory))
	if err == nil {
		film.FilmCategory = cate.CateName
	} else {
		film.FilmCategory = ""
	}

	director, err := l.svcCtx.ActorModel.FindOne(l.ctx, utils.StringToInt64(film.Director))
	if err == nil {
		film.Director = director.ActorName
	} else {
		film.Director = ""
	}
	return &pb.GetFilmResponse{Film: film}, nil
}
