package logic

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/cinema-rpc/internal/svc"
	"MaoerMovie/service/cinema-rpc/types/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCinemaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCinemaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCinemaLogic {
	return &GetCinemaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCinemaLogic) GetCinema(in *pb.GetCinemaRequest) (*pb.GetCinemaResponse, error) {
	cinemaId := utils.StringToInt64(in.CinemaId)
	cinema, err := l.svcCtx.CinemaModel.FindOne(l.ctx, cinemaId)
	if err != nil {
		return nil, err
	}
	var filmIds []string
	films, err := l.svcCtx.CinemaFilmModel.FindAllByCinemaId(l.ctx, cinemaId)
	if err != nil {
		return nil, err
	}
	for _, film := range films {
		filmIds = append(filmIds, utils.Int64ToString(film.FilmId))
	}
	return &pb.GetCinemaResponse{
		CinemaName:    cinema.CinemaName,
		CinemaAddress: cinema.Address,
		CinemaPhone:   cinema.CinemaPhone,
		CinemaImgs:    cinema.CinemaImgs,
		FilmIds:       filmIds,
	}, nil
}
