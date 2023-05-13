package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/film-rpc/model"
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"strings"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilmDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFilmDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilmDetailLogic {
	return &GetFilmDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFilmDetailLogic) GetFilmDetail(in *pb.FilmRequest) (*pb.GetFilmDetailResponse, error) {
	detail := &pb.FilmDetail{}
	//先查询缓存有没有该数据
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	filmInfoKey := utils.CacheFilmInfoKey + strconv.FormatInt(utils.StringToInt64(in.Id), 10)
	success, err := l.svcCtx.RedisClient.Exists(filmInfoKey)
	if err != nil {
		return nil, err
	}
	//缓存有该数据
	if success {
		redisFilmInfo, err := l.svcCtx.RedisClient.Get(filmInfoKey)
		//判断该数据是否为空值
		if redisFilmInfo == "" {
			return nil, errorx.NewCodeError(100, "查无此Id的电影！")
		}
		//json反序列化成对象
		var filmInfo model.FilmInfo
		err = json.UnmarshalFromString(redisFilmInfo, &filmInfo)
		if err != nil {
			return nil, err
		}
		detail.FilmPreSaleNum = utils.Int64ToString(filmInfo.FilmPreSaleNum)
		detail.FilmBoxOffice = utils.Int64ToString(filmInfo.FilmBoxOffice)
		detail.FilmImgs = filmInfo.FilmImgs
	} else {
		filmInfo, err := l.svcCtx.FilmInfoModel.FindOneByFilmId(l.ctx, utils.StringToInt64(in.Id))
		switch err {
		case nil:
			break
		case sqlx.ErrNotFound:
			l.svcCtx.RedisClient.Setex(filmInfoKey, "", utils.RedisLockExpireSeconds)
			return nil, errorx.NewCodeError(100, "查无此Id的电影！")
		default:
			return nil, err
		}
		//查到该数据，存入json序列化后的对象到缓存中
		jsonStr, err := json.MarshalToString(filmInfo)
		detail.FilmPreSaleNum = utils.Int64ToString(filmInfo.FilmPreSaleNum)
		detail.FilmBoxOffice = utils.Int64ToString(filmInfo.FilmBoxOffice)
		detail.FilmImgs = filmInfo.FilmImgs
		l.svcCtx.RedisClient.Setex(filmInfoKey, jsonStr, utils.RedisLockExpireSeconds)
	}

	filmScoreKey := utils.CacheFilmScoreKey + strconv.FormatInt(utils.StringToInt64(in.Id), 10)
	success, err = l.svcCtx.RedisClient.Exists(filmScoreKey)
	if err != nil {
		return nil, err
	}
	//缓存有该数据
	if success {
		redisFilmScore, err := l.svcCtx.RedisClient.Get(filmScoreKey)
		//判断该数据是否为空值
		if redisFilmScore == "" {
			detail.FilmScore = ""
			detail.FilmScoreNum = "0"
		} else {
			//json反序列化成对象
			var filmScore model.FilmScore
			err = json.UnmarshalFromString(redisFilmScore, &filmScore)
			if err != nil {
				return nil, err
			}
			detail.FilmScore = utils.Float64ToString(filmScore.FilmScore)
			detail.FilmScoreNum = utils.Int64ToString(filmScore.FilmScoreNum)
		}
	} else {
		filmScore, err := l.svcCtx.FilmScoreModel.FindOneByFilmId(l.ctx, utils.StringToInt64(in.Id))
		switch err {
		case nil:
			//查到该数据，存入json序列化后的对象到缓存中
			jsonStr, err := json.MarshalToString(filmScore)
			if err != nil {
				return nil, err
			}
			l.svcCtx.RedisClient.Setex(filmScoreKey, jsonStr, utils.RedisLockExpireSeconds)
			detail.FilmScore = utils.Float64ToString(filmScore.FilmScore)
			detail.FilmScoreNum = utils.Int64ToString(filmScore.FilmScoreNum)
			break
		case sqlx.ErrNotFound:
			l.svcCtx.RedisClient.Setex(filmScoreKey, "", utils.RedisLockExpireSeconds)
			break
		default:
			return nil, err
		}

	}

	filmActorKey := utils.CacheFilmActorKey + strconv.FormatInt(utils.StringToInt64(in.Id), 10)
	success, err = l.svcCtx.RedisClient.Exists(filmActorKey)
	if err != nil {
		return nil, err
	}
	//缓存有该数据
	if success {
		redisFilmActor, _ := l.svcCtx.RedisClient.Get(filmActorKey)
		//判断该数据是否为空值
		if redisFilmActor == "" {
			detail.ActorList = ""
		} else {
			detail.ActorList = redisFilmActor
		}
	} else {
		filmActorList, err := l.svcCtx.FilmActorModel.FindAllByFilmId(l.ctx, utils.StringToInt64(in.Id))
		switch err {
		case nil:
			var arr []string
			for _, filmActor := range filmActorList {
				actor, err := l.svcCtx.ActorModel.FindOne(l.ctx, filmActor.ActorId)
				if err != nil {
					return nil, err
				}
				arr = append(arr, actor.ActorName)
			}
			actorListStr := strings.Join(arr, "/")
			l.svcCtx.RedisClient.Setex(filmActorKey, actorListStr, utils.RedisLockExpireSeconds)
			detail.ActorList = actorListStr
			break
		case sqlx.ErrNotFound:
			l.svcCtx.RedisClient.Setex(filmActorKey, "", utils.RedisLockExpireSeconds)
			break
		default:
			return nil, err
		}
	}

	return &pb.GetFilmDetailResponse{FilmDetail: detail}, nil
}
