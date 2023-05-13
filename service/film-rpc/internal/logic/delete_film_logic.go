package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/utils"
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFilmLogic {
	return &DeleteFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFilmLogic) DeleteFilm(in *pb.FilmRequest) (*pb.FilmResponse, error) {
	_, err := l.svcCtx.FilmModel.FindOne(l.ctx, utils.StringToInt64(in.Id))
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的电影！")
	default:
		return nil, err
	}
	err = l.svcCtx.FilmModel.Delete(l.ctx, utils.StringToInt64(in.Id))
	if err != nil {
		return nil, err
	}
	filmRedisKey := utils.CacheFilmKey + strconv.FormatInt(utils.StringToInt64(in.Id), 10)
	success, err := l.svcCtx.RedisClient.Exists(filmRedisKey)
	if err != nil {
		return nil, err
	}
	if success {
		_, err = l.svcCtx.RedisClient.Del(filmRedisKey)
		if err != nil {
			return nil, err
		}
	}

	_, err = l.svcCtx.FilmInfoModel.FindOneByFilmId(l.ctx, utils.StringToInt64(in.Id))
	switch err {
	case nil:
		_ = l.svcCtx.FilmInfoModel.DeleteByFilmId(l.ctx, utils.StringToInt64(in.Id))
		filmInfoRedisKey := utils.CacheFilmInfoKey + strconv.FormatInt(utils.StringToInt64(in.Id), 10)
		success, err := l.svcCtx.RedisClient.Exists(filmInfoRedisKey)
		if err != nil {
			return nil, err
		}
		if success {
			_, err = l.svcCtx.RedisClient.Del(filmInfoRedisKey)
			if err != nil {
				return nil, err
			}
		}
		break
	case sqlx.ErrNotFound:
		break
	default:
		return nil, err
	}

	_, err = l.svcCtx.FilmScoreModel.FindOneByFilmId(l.ctx, utils.StringToInt64(in.Id))
	switch err {
	case nil:
		_ = l.svcCtx.FilmScoreModel.DeleteByFilmId(l.ctx, utils.StringToInt64(in.Id))
		filmScoreRedisKey := utils.CacheFilmScoreKey + strconv.FormatInt(utils.StringToInt64(in.Id), 10)
		success, err := l.svcCtx.RedisClient.Exists(filmScoreRedisKey)
		if err != nil {
			return nil, err
		}
		if success {
			_, err = l.svcCtx.RedisClient.Del(filmScoreRedisKey)
			if err != nil {
				return nil, err
			}
		}
		break
	case sqlx.ErrNotFound:
		break
	default:
		return nil, err
	}

	_, err = l.svcCtx.FilmActorModel.FindOneByFilmId(l.ctx, utils.StringToInt64(in.Id))
	switch err {
	case nil:
		_ = l.svcCtx.FilmActorModel.DeleteByFilmId(l.ctx, utils.StringToInt64(in.Id))
		filmActorRedisKey := utils.CacheFilmActorKey + strconv.FormatInt(utils.StringToInt64(in.Id), 10)
		success, err := l.svcCtx.RedisClient.Exists(filmActorRedisKey)
		if err != nil {
			return nil, err
		}
		if success {
			_, err = l.svcCtx.RedisClient.Del(filmActorRedisKey)
			if err != nil {
				return nil, err
			}
		}
		break
	case sqlx.ErrNotFound:
		break
	default:
		return nil, err
	}

	return &pb.FilmResponse{}, nil
}
