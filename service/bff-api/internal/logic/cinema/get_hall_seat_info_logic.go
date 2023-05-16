package cinema

import (
	"MaoerMovie/service/cinema-rpc/types/pb"
	"context"
	"encoding/json"
	"errors"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHallSeatInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHallSeatInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHallSeatInfoLogic {
	return &GetHallSeatInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHallSeatInfoLogic) GetHallSeatInfo(req *types.HallSeatInfoRequest) (resp *types.HallSeatInfoResponse, err error) {
	resp = new(types.HallSeatInfoResponse)
	rpcResp, err := l.svcCtx.CinemaRPC.GetHallSeats(l.ctx, &pb.GetHallSeatsRequest{HallId: req.HallId})
	if err != nil {
		return nil, err
	}
	seatInfo := types.SeatInfo{}
	err = json.Unmarshal(rpcResp.SeatFile, &seatInfo)
	if err != nil {
		return nil, errors.New("json解析时出错")
	}
	resp.SeatInfo = seatInfo
	//TODO:在order服务中查找已经售出的座位，并返回
	return resp, nil
}
