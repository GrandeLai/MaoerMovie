package order

import (
	"MaoerMovie/common/utils"
	pb3 "MaoerMovie/service/cinema-rpc/types/pb"
	pb2 "MaoerMovie/service/film-rpc/types/pb"
	"MaoerMovie/service/order-rpc/types/pb"
	"context"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(req *types.GetOrderDatilRequest) (resp *types.GetOrderDetailResponse, err error) {
	resp = new(types.GetOrderDetailResponse)
	order, err := l.svcCtx.OrderRPC.GetOrderDetail(l.ctx, &pb.GetOrderDetailRequest{OrderId: req.OrderId})
	if err != nil {
		return nil, err
	}
	filmRpcResp, err := l.svcCtx.FilmRPC.GetFilm(l.ctx, &pb2.FilmRequest{Id: order.OrderDetail.FilmId})
	if err != nil {
		return nil, err
	}
	cinemaRpcResp, err := l.svcCtx.CinemaRPC.GetCinema(l.ctx, &pb3.GetCinemaRequest{CinemaId: order.OrderDetail.CinemaId})
	if err != nil {
		return nil, err
	}
	resp.OrderDetail = types.OrderDetail{
		OrderId:      req.OrderId,
		ShowId:       order.OrderDetail.ShowId,
		Price:        order.OrderDetail.Price,
		Status:       order.OrderDetail.Status,
		SeatIds:      order.OrderDetail.SeatIds,
		SeatPosition: order.OrderDetail.SeatPosition,
		SeatNum:      utils.IntToString(len(order.OrderDetail.SeatIds)),
		FilmName:     filmRpcResp.Film.FilmName,
		CinemaName:   cinemaRpcResp.CinemaName,
	}
	return
}
