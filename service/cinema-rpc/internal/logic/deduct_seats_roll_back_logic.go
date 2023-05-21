package logic

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/cinema-rpc/internal/svc"
	"MaoerMovie/service/cinema-rpc/types/pb"
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductSeatsRollBackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSeatsRollBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSeatsRollBackLogic {
	return &DeductSeatsRollBackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSeatsRollBackLogic) DeductSeatsRollBack(in *pb.DeductSeatsRequest) (*pb.DeductSeatsResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		_, err := l.svcCtx.ShowModel.TxUpdateSurplusNumWithLock(tx, utils.StringToInt64(in.Num), utils.StringToInt64(in.ShowId))
		redisQueryKey := utils.CacheShowStockKey + in.ShowId
		_, err = l.svcCtx.RedisClient.Incrby(redisQueryKey, 1)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.DeductSeatsResponse{}, nil
}
