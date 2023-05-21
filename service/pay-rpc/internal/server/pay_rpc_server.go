// Code generated by goctl. DO NOT EDIT!
// Source: pay.proto

package server

import (
	"context"

	"MaoerMovie/service/pay-rpc/internal/logic"
	"MaoerMovie/service/pay-rpc/internal/svc"
	"MaoerMovie/service/pay-rpc/types/pb"
)

type PayRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPayRpcServer
}

func NewPayRpcServer(svcCtx *svc.ServiceContext) *PayRpcServer {
	return &PayRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *PayRpcServer) CreatePay(ctx context.Context, in *pb.CreatePayRequest) (*pb.CreatePayResponse, error) {
	l := logic.NewCreatePayLogic(ctx, s.svcCtx)
	return l.CreatePay(in)
}

func (s *PayRpcServer) SetPayPaid(ctx context.Context, in *pb.SetPayStatusRequest) (*pb.SetPayStatusResponse, error) {
	l := logic.NewSetPayPaidLogic(ctx, s.svcCtx)
	return l.SetPayPaid(in)
}

func (s *PayRpcServer) SetPayPaidRollback(ctx context.Context, in *pb.SetPayStatusRequest) (*pb.SetPayStatusResponse, error) {
	l := logic.NewSetPayPaidRollbackLogic(ctx, s.svcCtx)
	return l.SetPayPaidRollback(in)
}

func (s *PayRpcServer) GetPayDetail(ctx context.Context, in *pb.GetPayDetailRequest) (*pb.GetPayDetailResponse, error) {
	l := logic.NewGetPayDetailLogic(ctx, s.svcCtx)
	return l.GetPayDetail(in)
}