package cinema

import (
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"MaoerMovie/service/cinema-rpc/types/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCinemaListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCinemaListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCinemaListLogic {
	return &GetCinemaListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCinemaListLogic) GetCinemaList(req *types.CinemaListRequest) (resp *types.CinemaListResponse, err error) {
	resp = new(types.CinemaListResponse)
	rpcResp, err := l.svcCtx.CinemaRPC.GetCinemaList(l.ctx, &pb.CinemaListRequest{
		BrandId:    req.BrandId,
		HallType:   req.HallType,
		DistrictId: req.DistrictId,
		Page:       req.Page,
		Size:       req.Size,
	})
	if err != nil {
		return nil, err
	}
	for _, cinema := range rpcResp.Cinema {
		resp.List = append(resp.List, types.CinemaPreview{
			CinemaId:   cinema.Id,
			CinemaName: cinema.CinemaName,
			Address:    cinema.Address,
			MinPrice:   cinema.MinPrice,
		})
	}
	return resp, nil
}
