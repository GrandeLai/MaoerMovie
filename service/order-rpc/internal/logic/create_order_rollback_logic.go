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

type CreateOrderRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderRollbackLogic {
	return &CreateOrderRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderRollbackLogic) CreateOrderRollback(in *pb.CreateOrderRequest) (*pb.CreateOrderRequest, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		err = l.svcCtx.OrderModel.TxDelete(tx, in.OrderId)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.CreateOrderRequest{OrderId: in.OrderId}, nil
}
