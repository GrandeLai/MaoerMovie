package pay

import (
	"MaoerMovie/service/order-rpc/types/pb"
	pb2 "MaoerMovie/service/pay-rpc/types/pb"
	"context"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/smartwalle/alipay/v3"
	"net/http"
	"strings"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayCallbackLogic {
	return &PayCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayCallbackLogic) PayCallback(req *types.PayCallBackRequest, r *http.Request) (resp *types.PayCallBackResponse, err error) {
	var noti, _ = l.svcCtx.AlipayClient.GetTradeNotification(r)
	if noti != nil {
		var outTradeNo = r.Form.Get("out_trade_no")
		var p = alipay.TradeQuery{}
		p.OutTradeNo = outTradeNo
		p.TradeNo = r.Form.Get("trade_no")
		rsp, err := l.svcCtx.AlipayClient.TradeQuery(p)

		body := r.FormValue("body")
		orderId := strings.Split(body, ":")[1]
		paySn := rsp.Content.OutTradeNo
		buyerAccount := rsp.Content.BuyerLogonId
		//获取支付RPC和订单RPC各自的BuildTarget
		payRpcBT, err := l.svcCtx.Config.PayRpc.BuildTarget()
		if err != nil {
			return nil, err
		}
		orderRpcBT, err := l.svcCtx.Config.OrderRpc.BuildTarget()
		if err != nil {
			return nil, err
		} //创建saga协议的事务
		dtmServer := l.svcCtx.Config.DtmServer
		gid := dtmgrpc.MustGenGid(dtmServer)
		saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
			Add(orderRpcBT+"/order.orderRpc/SetOrderPaid", orderRpcBT+"/order.orderRpc/SetOrderPaidRollback", &pb.SetOrderPaidRequest{OrderId: orderId}).
			Add(payRpcBT+"/pay.payRpc/SetPayPaid", payRpcBT+"/pay.payRpc/SetPayPaidRollback", &pb2.SetPayStatusRequest{PaySn: paySn, BuyerAccount: buyerAccount})

		//事务提交
		err = saga.Submit()
		dtmimp.FatalIfError(err)
		if err != nil {
			resp.Result = "支付失败"
			return resp, err
		}
		resp.Result = "支付成功"
	} else {
		resp.Result = "支付失败"
	}
	return resp, nil
}
