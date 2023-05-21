package pay

import (
	"MaoerMovie/common/uniqueid"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	pb2 "MaoerMovie/service/pay-rpc/types/pb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/smartwalle/alipay/v3"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePayLogic {
	return &CreatePayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePayLogic) CreatePay(req *types.CreatePayRequest) (resp *types.CreatePayResponse, err error) {
	resp = new(types.CreatePayResponse)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	pay := alipay.TradePagePay{}
	// 支付宝回调地址（需要在支付宝后台配置）
	// 支付成功后，支付宝会发送一个POST消息到该地址
	pay.NotifyURL = "http://" + l.svcCtx.Config.Host + ":" + utils.IntToString(l.svcCtx.Config.Port) + "/pay/call_back"
	// 支付成功之后，浏览器将会重定向到该 URL
	pay.ReturnURL = "http://" + l.svcCtx.Config.Host + ":" + utils.IntToString(l.svcCtx.Config.Port) + "/pay/return"
	//支付标题
	pay.Subject = "猫耳电影线上订座"
	//订单号，一个订单号只能支付一次
	sn := uniqueid.GenSn(uniqueid.SN_PREFIX_THIRD_PAYMENT)
	pay.OutTradeNo = sn
	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//金额
	pay.TotalAmount = req.Price
	pay.Body = req.OrderId
	//pay.BusinessParams = req.OrderId
	url, err := l.svcCtx.AlipayClient.TradePagePay(pay)
	if err != nil {
		return nil, err
	}

	//_, err = l.svcCtx.OrderRPC.SetOrderPaid(l.ctx, &pb.SetOrderPaidRequest{OrderId: req.OrderId})
	//if err != nil {
	//	return nil, errors.New("修改订单信息失败")
	//}
	pageURL := url.String()
	_, err = l.svcCtx.PayRPC.CreatePay(l.ctx, &pb2.CreatePayRequest{
		OrderId: req.OrderId,
		UserId:  utils.Int64ToString(userId),
		PaySn:   sn,
		Price:   req.Price,
		Subject: pay.Subject,
	})
	if err != nil {
		return nil, err
	}
	//这个 payURL 是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	resp.PaySn = sn
	resp.PayUrl = pageURL
	return resp, nil
}
