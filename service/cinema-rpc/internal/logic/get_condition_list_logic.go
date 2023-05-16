package logic

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/cinema-rpc/internal/svc"
	"MaoerMovie/service/cinema-rpc/types/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConditionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConditionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConditionListLogic {
	return &GetConditionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetConditionListLogic) GetConditionList(in *pb.ConditionListRequest) (*pb.ConditionListResponse, error) {
	districtList, err := l.svcCtx.DistrictModel.FindAllByCityName(l.ctx, in.CityName)
	if err != nil {
		return nil, err
	}
	brandList, err := l.svcCtx.BrandModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	hallList, err := l.svcCtx.HallModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	var hallType []*pb.HallType
	for _, hall := range hallList {
		hallType = append(hallType, &pb.HallType{
			Id:           utils.Int64ToString(hall.Id),
			HallTypeName: hall.HallName,
		})
	}
	var districtCondition []*pb.DistrictCondition
	for _, district := range districtList {
		districtCondition = append(districtCondition, &pb.DistrictCondition{
			Id:           utils.Int64ToString(district.Id),
			DistrictName: district.DistrictName,
		})
	}
	var brandCondition []*pb.BrandCondition
	for _, brand := range brandList {
		brandCondition = append(brandCondition, &pb.BrandCondition{
			Id:        utils.Int64ToString(brand.Id),
			BrandName: brand.BrandName,
		})
	}
	return &pb.ConditionListResponse{BrandList: brandCondition, HallList: hallType, DistrictList: districtCondition}, nil
}
