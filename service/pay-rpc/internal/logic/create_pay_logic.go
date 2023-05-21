package logic

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/pay-rpc/internal/svc"
	"MaoerMovie/service/pay-rpc/model"
	"MaoerMovie/service/pay-rpc/types/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePayLogic {
	return &CreatePayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePayLogic) CreatePay(in *pb.CreatePayRequest) (*pb.CreatePayResponse, error) {
	userId := utils.StringToInt64(in.UserId)
	price := utils.StringToFloat64(in.Price)

	_, err := l.svcCtx.PayModel.Insert(l.ctx, &model.Pay{
		PaySn:        in.PaySn,
		UserId:       userId,
		OrderId:      in.OrderId,
		Price:        price,
		Subject:      in.Subject,
		Status:       0,
		BuyerAccount: "",
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreatePayResponse{}, nil
}
