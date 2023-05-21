package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/pay-rpc/internal/svc"
	"MaoerMovie/service/pay-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPayDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPayDetailLogic {
	return &GetPayDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPayDetailLogic) GetPayDetail(in *pb.GetPayDetailRequest) (*pb.GetPayDetailResponse, error) {
	pay, err := l.svcCtx.PayModel.FindByPaySnAndUserId(l.ctx, in.PaySn, utils.StringToInt64(in.UserId))
	if err != nil {
		return nil, err
	}
	resp := &pb.GetPayDetailResponse{
		UserId:  utils.Int64ToString(pay.UserId),
		OrderId: pay.OrderId,
		Price:   utils.Float64ToString(pay.Price),
		Subject: pay.Subject,
		Status:  utils.Int64ToString(pay.Status),
	}
	return resp, nil
}
