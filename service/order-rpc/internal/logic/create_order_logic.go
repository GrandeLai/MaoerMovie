package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/order-rpc/internal/svc"
	"MaoerMovie/service/order-rpc/model"
	"MaoerMovie/service/order-rpc/types/pb"
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"os"
	"reflect"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	//调用lua脚本，从缓存中判断影票库存是否足够并删减
	client := redis.NewClient(&redis.Options{
		Addr:     l.svcCtx.RedisClient.Addr,
		Password: l.svcCtx.RedisClient.Pass,
	})
	file, err := os.ReadFile("common/script/decStock.lua")
	if err != nil {
		return nil, err
	}
	script := redis.NewScript(string(file))
	redisQueryKey := utils.CacheShowStockKey + in.ShowId
	//l.svcCtx.RedisClient.Set(redisQueryKey, utils.IntToString(24))
	res, err := script.Run(context.Background(), client, []string{}, redisQueryKey).Result()
	if err != nil {
		panic(err)
	}
	value := reflect.ValueOf(res).Int()
	//不存在该电影
	if value == 1 {
		return nil, errorx.NewCodeError(100, "查无此Id的电影！")
	}
	//库存不足
	if value == 2 {
		return nil, errorx.NewCodeError(100, "影票库存不足！")
	}

	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, err
	}
	// 获取子事务屏障对象，来源于GRPC
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, err
	}
	//开启该子事务屏障
	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		_, err := l.svcCtx.OrderModel.TxInsert(tx, &model.Order{
			Uuid:          in.OrderId,
			CinemaId:      utils.StringToInt64(in.CinemaId),
			ShowId:        utils.StringToInt64(in.ShowId),
			FilmId:        utils.StringToInt64(in.FilmId),
			SeatsIds:      in.SeatIds,
			SeatsPosition: in.SeatPosition,
			Price:         utils.StringToFloat64(in.Price),
			UserId:        utils.StringToInt64(in.UserId),
			Status:        0,
		})
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{OrderId: in.OrderId}, nil
}
