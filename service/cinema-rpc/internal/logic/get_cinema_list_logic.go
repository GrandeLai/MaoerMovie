package logic

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/cinema-rpc/internal/svc"
	"MaoerMovie/service/cinema-rpc/types/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCinemaListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCinemaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCinemaListLogic {
	return &GetCinemaListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCinemaListLogic) GetCinemaList(in *pb.CinemaListRequest) (*pb.CinemaListResponse, error) {
	var brandId int64
	if in.BrandId == "" {
		brandId = 0
	} else {
		brandId = utils.StringToInt64(in.BrandId)
	}
	var districtId int64
	if in.DistrictId == "" {
		districtId = 0
	} else {
		districtId = utils.StringToInt64(in.DistrictId)
	}

	var page int64
	if in.Page == "" {
		page = 0
	} else {
		page = utils.StringToInt64(in.Page)
	}
	var size int64
	if in.Size == "" {
		size = 0
	} else {
		size = utils.StringToInt64(in.Size)
	}
	cinemaList, err := l.svcCtx.CinemaModel.FindByFactors(l.ctx, brandId, in.HallType, districtId, page, size)
	if err != nil {
		return nil, err
	}
	var cinemaPreviewList []*pb.CinemaPreview
	for _, cinema := range cinemaList {
		cinemaPreviewList = append(cinemaPreviewList, &pb.CinemaPreview{
			Id:         utils.Int64ToString(cinema.Id),
			CinemaName: cinema.CinemaName,
			Address:    cinema.Address,
			MinPrice:   utils.Int64ToString(cinema.MinPrice),
		})
	}
	return &pb.CinemaListResponse{Cinema: cinemaPreviewList}, nil
}
