package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/order-rpc/internal/svc"
	"MaoerMovie/service/order-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderListLogic) GetOrderList(in *pb.GetOrderListRequest) (*pb.GetOrderListResponse, error) {
	page := utils.StringToInt64(in.Page)
	size := utils.StringToInt64(in.Size)
	userId := utils.StringToInt64(in.UserId)
	orderList, err := l.svcCtx.OrderModel.FindAllInPageByUserId(l.ctx, userId, page, size)
	if err != nil {
		return nil, err
	}
	var respList []*pb.OrderDetail
	for _, order := range orderList {
		respList = append(respList, &pb.OrderDetail{
			OrderId:      order.Uuid,
			CinemaId:     utils.Int64ToString(order.CinemaId),
			FilmId:       utils.Int64ToString(order.FilmId),
			ShowId:       utils.Int64ToString(order.ShowId),
			Price:        utils.Float64ToString(order.Price),
			Status:       utils.Int64ToString(order.Status),
			SeatIds:      order.SeatsIds,
			SeatPosition: order.SeatsPosition,
		})
	}
	return &pb.GetOrderListResponse{OrderList: respList}, nil
}
