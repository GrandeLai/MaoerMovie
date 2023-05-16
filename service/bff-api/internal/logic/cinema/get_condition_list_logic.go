package cinema

import (
	"MaoerMovie/service/cinema-rpc/types/pb"
	"context"
	"github.com/jinzhu/copier"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConditionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConditionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConditionListLogic {
	return &GetConditionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConditionListLogic) GetConditionList(req *types.ConditionListRequest) (resp *types.ConditionListResponse, err error) {
	resp = new(types.ConditionListResponse)
	rpcResp, err := l.svcCtx.CinemaRPC.GetConditionList(l.ctx, &pb.ConditionListRequest{CityName: req.CityName})
	if err != nil {
		return nil, err
	}

	var brandCondition []types.BrandCondition
	err = copier.Copy(&brandCondition, rpcResp.BrandList)
	if err != nil {
		return nil, err
	}
	resp.BrandList = brandCondition

	var districtCondition []types.DistrictCondition
	err = copier.Copy(&districtCondition, rpcResp.DistrictList)
	if err != nil {
		return nil, err
	}
	resp.DistrictList = districtCondition

	var hallType []types.HallType
	err = copier.Copy(&hallType, rpcResp.HallList)
	if err != nil {
		return nil, err
	}
	resp.HallTypeList = hallType
	return resp, nil
}
