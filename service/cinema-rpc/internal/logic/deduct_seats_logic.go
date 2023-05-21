package logic

import (
	"MaoerMovie/common/utils"
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"MaoerMovie/service/cinema-rpc/internal/svc"
	"MaoerMovie/service/cinema-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeductSeatsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeductSeatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeductSeatsLogic {
	return &DeductSeatsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeductSeatsLogic) DeductSeats(in *pb.DeductSeatsRequest) (*pb.DeductSeatsResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		res, err := l.svcCtx.ShowModel.TxUpdateSurplusNumWithLock(tx, -utils.StringToInt64(in.Num), utils.StringToInt64(in.ShowId))
		if err != nil {
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if rows == 0 {
			//库存不足，走回滚
			return dtmcli.ErrFailure
		}
		return nil
	})
	if err == dtmcli.ErrFailure {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}
	if err != nil {
		return nil, err
	}
	return &pb.DeductSeatsResponse{}, nil
}
