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

type SetPayPaidRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetPayPaidRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPayPaidRollbackLogic {
	return &SetPayPaidRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetPayPaidRollbackLogic) SetPayPaidRollback(in *pb.SetPayStatusRequest) (*pb.SetPayStatusResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		pay, err := l.svcCtx.PayModel.FindByPaySn(l.ctx, in.PaySn)
		if err != nil {
			return err
		}
		pay.Status = 0
		pay.BuyerAccount = in.BuyerAccount
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
