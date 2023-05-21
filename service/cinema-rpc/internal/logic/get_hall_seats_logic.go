package logic

import (
	"MaoerMovie/common/utils"
	"context"
	"encoding/json"
	"errors"
	"github.com/minio/minio-go/v7"
	"io/ioutil"
	"strings"

	"MaoerMovie/service/cinema-rpc/internal/svc"
	"MaoerMovie/service/cinema-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHallSeatsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHallSeatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHallSeatsLogic {
	return &GetHallSeatsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type Seat struct {
	AllSeats string
}

func (l *GetHallSeatsLogic) GetHallSeats(in *pb.GetHallSeatsRequest) (*pb.GetHallSeatsResponse, error) {
	hallId := utils.StringToInt64(in.HallId)
	hall, err := l.svcCtx.HallModel.FindOne(l.ctx, hallId)
	if err != nil {
		return nil, err
	}
	arr := strings.Split(hall.SeatAddress, "/")
	object, err := l.svcCtx.MinioClient.GetObject(l.ctx, l.svcCtx.Config.MinIO.BucketName, arr[len(arr)-1], minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.New("无法获取影厅下座位")
	}
	defer object.Close()
	data, err := ioutil.ReadAll(object)
	if err != nil {
		return nil, err
	}
	var seats Seat
	err = json.Unmarshal(data, &seats)
	if err != nil {
		return nil, err
	}
	redisQueryKey := utils.CacheHallKey + in.HallId
	err = l.svcCtx.RedisClient.Set(redisQueryKey, seats.AllSeats)
	if err != nil {
		return nil, err
	}
	return &pb.GetHallSeatsResponse{SeatFile: data}, nil
}
