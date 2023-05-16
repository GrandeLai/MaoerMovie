package cinema

import (
	"MaoerMovie/service/cinema-rpc/types/pb"
	pb2 "MaoerMovie/service/film-rpc/types/pb"
	"context"
	"github.com/jinzhu/copier"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCinemaFilmListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCinemaFilmListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCinemaFilmListLogic {
	return &GetCinemaFilmListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCinemaFilmListLogic) GetCinemaFilmList(req *types.CinemaFilmListRequest) (resp *types.CinemaFilmListResponse, err error) {
	resp = new(types.CinemaFilmListResponse)
	rpcResp, err := l.svcCtx.CinemaRPC.GetCinema(l.ctx, &pb.GetCinemaRequest{CinemaId: req.CinemaId})
	if err != nil {
		return nil, err
	}
	resp.CinemaName = rpcResp.CinemaName
	resp.CinemaPhone = rpcResp.CinemaPhone
	resp.CinemaAddress = rpcResp.CinemaAddress
	resp.CinemaImgs = rpcResp.CinemaAddress
	filmList, err := l.svcCtx.FilmRPC.GetCinemaFilm(l.ctx, &pb2.CinemaFilmRequest{FilmIds: rpcResp.FilmIds})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp.FilmList, filmList.CinemaFilm)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
