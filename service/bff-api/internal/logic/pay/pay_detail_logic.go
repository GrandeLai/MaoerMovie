package pay

import (
	"MaoerMovie/common/utils"
	pb2 "MaoerMovie/service/order-rpc/types/pb"
	"MaoerMovie/service/pay-rpc/types/pb"
	"context"
	"encoding/json"
	"fmt"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayDetailLogic {
	return &PayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayDetailLogic) PayDetail(req *types.PayDeatilRequest) (resp *types.PayDeatilResponse, err error) {
	resp = new(types.PayDeatilResponse)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	rpcResp, err := l.svcCtx.PayRPC.GetPayDetail(l.ctx, &pb.GetPayDetailRequest{PaySn: req.PaySn, UserId: utils.Int64ToString(userId)})
	if err != nil {
		return nil, err
	}
	order, err := l.svcCtx.OrderRPC.GetOrderDetail(l.ctx, &pb2.GetOrderDetailRequest{OrderId: rpcResp.OrderId})
	resp.OrderId = rpcResp.OrderId
	resp.Price = rpcResp.Price
	resp.PaySn = req.PaySn
	resp.ShowId = order.OrderDetail.ShowId
	resp.PayStatus = rpcResp.Status
	return
}
