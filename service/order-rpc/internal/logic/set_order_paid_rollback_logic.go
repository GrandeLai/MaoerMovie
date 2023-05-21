package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"MaoerMovie/service/order-rpc/internal/svc"
	"MaoerMovie/service/order-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetOrderPaidRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetOrderPaidRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetOrderPaidRollbackLogic {
	return &SetOrderPaidRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetOrderPaidRollbackLogic) SetOrderPaidRollback(in *pb.SetOrderPaidRequest) (*pb.SetOrderPaidResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		order, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.OrderId)
		if err != nil {
			return err
		}
		order.Status = 0
		_, err = l.svcCtx.OrderModel.TxUpdate(tx, order)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.SetOrderPaidResponse{}, nil
}
