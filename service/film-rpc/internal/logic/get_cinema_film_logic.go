package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCinemaFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCinemaFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCinemaFilmLogic {
	return &GetCinemaFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCinemaFilmLogic) GetCinemaFilm(in *pb.CinemaFilmRequest) (*pb.CinemaFilmResponse, error) {
	var cinemaFilm []*pb.CinemaFilm
	for _, id := range in.FilmIds {

		filmId := utils.StringToInt64(id)
		film, err := l.svcCtx.FilmModel.FindOne(l.ctx, filmId)
		if err != nil {
			return nil, err
		}
		cate, err := l.svcCtx.CategoryModel.FindOne(l.ctx, film.CategoryId)
		if err != nil {
			return nil, err
		}
		filmActors, err := l.svcCtx.FilmActorModel.FindAllByFilmId(l.ctx, filmId)
		if err != nil {
			return nil, err
		}
		var actorStr string
		for _, filmActor := range filmActors {
			actor, err := l.svcCtx.ActorModel.FindOne(l.ctx, filmActor.ActorId)
			if err != nil {
				return nil, err
			}
			actorStr += actor.ActorName + "/"
		}
		actorStr = actorStr[:len(actorStr)-1]
		cine := &pb.CinemaFilm{
			FilmId:       id,
			FilmName:     film.FilmName,
			FilmLength:   utils.Int64ToString(film.FilmLength),
			FilmCover:    film.FilmCover,
			FilmCategory: cate.CateName,
			ActorList:    actorStr,
		}
		cinemaFilm = append(cinemaFilm, cine)
	}
	return &pb.CinemaFilmResponse{CinemaFilm: cinemaFilm}, nil
}
