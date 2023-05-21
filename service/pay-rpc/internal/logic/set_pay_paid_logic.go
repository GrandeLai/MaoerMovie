package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"MaoerMovie/service/pay-rpc/internal/svc"
	"MaoerMovie/service/pay-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetPayPaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetPayPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPayPaidLogic {
	return &SetPayPaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetPayPaidLogic) SetPayPaid(in *pb.SetPayStatusRequest) (*pb.SetPayStatusResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	pay, err := l.svcCtx.PayModel.FindByPaySn(l.ctx, in.PaySn)
	if err != nil {
		return nil, err
	}
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		pay.BuyerAccount = in.BuyerAccount
		pay.Status = 1
		_, err = l.svcCtx.PayModel.TxUpdate(tx, pay)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.SetPayStatusResponse{}, nil
}
