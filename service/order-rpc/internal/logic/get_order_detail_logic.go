package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/order-rpc/internal/svc"
	"MaoerMovie/service/order-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(in *pb.GetOrderDetailRequest) (*pb.GetOrderDetailResponse, error) {

	order, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.OrderId)
	if err != nil {
		return nil, err
	}
	detail := &pb.OrderDetail{
		OrderId:      order.Uuid,
		CinemaId:     utils.Int64ToString(order.CinemaId),
		FilmId:       utils.Int64ToString(order.FilmId),
		ShowId:       utils.Int64ToString(order.ShowId),
		Price:        utils.Float64ToString(order.Price),
		Status:       utils.Int64ToString(order.Status),
		SeatIds:      order.SeatsIds,
		SeatPosition: order.SeatsPosition,
	}
	return &pb.GetOrderDetailResponse{OrderDetail: detail}, nil
}
