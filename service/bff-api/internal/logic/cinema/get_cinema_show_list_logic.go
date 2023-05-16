package cinema

import (
	"MaoerMovie/service/cinema-rpc/types/pb"
	"context"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCinemaShowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCinemaShowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCinemaShowListLogic {
	return &GetCinemaShowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCinemaShowListLogic) GetCinemaShowList(req *types.CinemaShowListRequest) (resp *types.CinemaShowListResponse, err error) {
	resp = new(types.CinemaShowListResponse)
	rpcResp, err := l.svcCtx.CinemaRPC.GetShowList(l.ctx, &pb.GetShowListRequest{
		CinemaId: req.CinemaId,
		FilmId:   req.FilmId,
		ShowDate: req.Date,
	})
	if err != nil {
		return nil, err
	}
	var cinemaShows []types.CinemaShow
	for _, show := range rpcResp.Show {
		cinemaShows = append(cinemaShows, types.CinemaShow{
			ShowId:    show.Id,
			BeginTime: show.BeginTime,
			EndTime:   show.EndTime,
			Language:  show.Language,
			Price:     show.Price,
			HallName:  show.HallName,
		})
	}
	resp.ShowList = cinemaShows
	return resp, nil
}
