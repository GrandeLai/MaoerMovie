package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/cinema-rpc/internal/svc"
	"MaoerMovie/service/cinema-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShowListLogic {
	return &GetShowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShowListLogic) GetShowList(in *pb.GetShowListRequest) (*pb.GetShowListResponse, error) {
	cinemaId := utils.StringToInt64(in.CinemaId)
	filmId := utils.StringToInt64(in.FilmId)
	resp, err := l.svcCtx.ShowModel.FindByFactors(l.ctx, cinemaId, filmId, in.ShowDate)
	if err != nil {
		return nil, err
	}
	var showList []*pb.CinemaShow
	for _, showResp := range resp {
		hall, err := l.svcCtx.HallModel.FindOne(l.ctx, showResp.HallId)
		if err != nil {
			return nil, err
		}
		showList = append(showList, &pb.CinemaShow{
			Id:        utils.Int64ToString(showResp.Id),
			BeginTime: showResp.BeginTime,
			EndTime:   showResp.EndTime,
			Language:  showResp.FilmLanguage,
			Price:     utils.Int64ToString(showResp.Price),
			HallName:  hall.HallName,
		})
	}
	return &pb.GetShowListResponse{Show: showList}, nil
}
