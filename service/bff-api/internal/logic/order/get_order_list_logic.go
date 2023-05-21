package order

import (
	"MaoerMovie/common/utils"
	pb3 "MaoerMovie/service/cinema-rpc/types/pb"
	pb2 "MaoerMovie/service/film-rpc/types/pb"
	"MaoerMovie/service/order-rpc/types/pb"
	"context"
	"encoding/json"
	"fmt"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListLogic) GetOrderList(req *types.GetOrderListRequest) (resp *types.GetOrderListResponse, err error) {
	resp = new(types.GetOrderListResponse)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	rpcResp, err := l.svcCtx.OrderRPC.GetOrderList(l.ctx, &pb.GetOrderListRequest{
		Page:   req.Page,
		Size:   req.Size,
		UserId: utils.Int64ToString(userId),
	})
	var orderList []types.OrderDetail
	for _, order := range rpcResp.OrderList {
		filmRpcResp, err := l.svcCtx.FilmRPC.GetFilm(l.ctx, &pb2.FilmRequest{Id: order.FilmId})
		if err != nil {
			return nil, err
		}
		cinemaRpcResp, err := l.svcCtx.CinemaRPC.GetCinema(l.ctx, &pb3.GetCinemaRequest{CinemaId: order.CinemaId})
		if err != nil {
			return nil, err
		}
		orderList = append(orderList, types.OrderDetail{
			OrderId:      order.OrderId,
			ShowId:       order.ShowId,
			Price:        order.Price,
			Status:       order.Status,
			SeatIds:      order.SeatIds,
			SeatPosition: order.SeatPosition,
			SeatNum:      utils.IntToString(len(order.SeatIds)),
			FilmName:     filmRpcResp.Film.FilmName,
			CinemaName:   cinemaRpcResp.CinemaName,
		})
	}
	resp.OrderList = orderList
	resp.Total = utils.IntToString(len(orderList))
	return
}
