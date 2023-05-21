package logic

import (
	"MaoerMovie/service/order-rpc/internal/svc"
	"MaoerMovie/service/order-rpc/types/pb"
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetOrderPaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetOrderPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetOrderPaidLogic {
	return &SetOrderPaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetOrderPaidLogic) SetOrderPaid(in *pb.SetOrderPaidRequest) (*pb.SetOrderPaidResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	order, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.OrderId)
	if err != nil {
		return nil, err
	}
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		order.Status = 1
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
