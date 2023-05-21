package logic

import (
	"MaoerMovie/common/utils"
	"context"
	"sort"
	"strconv"
	"strings"

	"MaoerMovie/service/order-rpc/internal/svc"
	"MaoerMovie/service/order-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSoldSeatsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSoldSeatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSoldSeatsLogic {
	return &GetSoldSeatsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSoldSeatsLogic) GetSoldSeats(in *pb.GetSoldSeatsRequest) (*pb.GetSoldSeatsResponse, error) {
	showId := utils.StringToInt64(in.ShowId)
	orderList, err := l.svcCtx.OrderModel.FindAllByShowId(l.ctx, showId)
	if err != nil {
		return nil, err
	}
	var seatIds string
	for _, order := range orderList {
		seatIds += order.SeatsIds
	}
	//redisSoldSeatsKey := utils.CacheSoldSeatsKey + in.ShowId
	//
	split := strings.Split(seatIds, ",")

	var seats []int64
	for _, str := range split {
		seat, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			seats = append(seats, seat)
		}
	}
	sort.SliceStable(seats, func(i, j int) bool {
		return seats[i] < seats[j] //从小到大排列
	})
	return &pb.GetSoldSeatsResponse{SoldSeats: seats}, nil
}
