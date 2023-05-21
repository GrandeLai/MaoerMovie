package order

import (
	"MaoerMovie/common/utils"
	pb2 "MaoerMovie/service/cinema-rpc/types/pb"
	"MaoerMovie/service/order-rpc/types/pb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmgrpc"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	//获取电影院RPC和订单RPC各自的BuildTarget
	cinemaRpcBT, err := l.svcCtx.Config.CinemaRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	orderRpcBT, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, err
	}
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	orderId := utils.GenerateUUID()
	createOrderReq := &pb.CreateOrderRequest{
		OrderId:      orderId,
		CinemaId:     req.CinemaId,
		FilmId:       req.FilmId,
		ShowId:       req.ShowId,
		Price:        req.Price,
		SeatIds:      req.SeatIds,
		SeatPosition: req.SeatPosition,
		SeatNum:      req.SeatNum,
		UserId:       utils.Int64ToString(userId),
	}
	//创建saga协议的事务
	dtmServer := l.svcCtx.Config.DtmServer
	gid := dtmgrpc.MustGenGid(dtmServer)
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcBT+"/order.orderRpc/CreateOrder", orderRpcBT+"/order.orderRpc/CreateOrderRollback", createOrderReq).
		Add(cinemaRpcBT+"/cinema.cinemaRpc/DeductSeats", cinemaRpcBT+"/cinema.cinemaRpc/DeductSeatsRollBack", &pb2.DeductSeatsRequest{
			ShowId: req.ShowId,
			Num:    req.SeatNum,
		})

	//事务提交
	err = saga.Submit()
	dtmimp.FatalIfError(err)
	if err != nil {
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}
	return &types.CreateOrderResponse{
		OrderId: orderId,
	}, nil
}
