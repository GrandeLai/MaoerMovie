package pay

import (
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"context"
	"net/http"
	"net/url"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayReturnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer http.ResponseWriter
}

func NewPayReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayReturnLogic {
	return &PayReturnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayReturnLogic) PayReturn(req *types.PayReturnRequest, r *http.Request) (resp *types.PayReturnResponse, err error) {
	//获取url并转成*URL
	x, _ := url.Parse(r.URL.String())
	//验证是否成功支付
	ok, err := l.svcCtx.AlipayClient.VerifySign(x.Query())
	if err == nil && ok {
		resp.Result = "支付成功"
		return resp, nil
	} else {
		resp.Result = "支付失败"
		return resp, nil
	}
}
