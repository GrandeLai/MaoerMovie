package admin

import (
	"MaoerMovie/service/film-rpc/types/pb"
	"context"
	"github.com/jinzhu/copier"
	"strconv"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActorListLogic {
	return &GetActorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActorListLogic) GetActorList(req *types.FilmListRequest) (resp *types.GetActorListResponse, err error) {
	resp = new(types.GetActorListResponse)
	rpcResp, err := l.svcCtx.FilmRPC.GetActorList(l.ctx, &pb.ListRequest{
		Size: req.Size,
		Page: req.Page,
	})
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&resp.ActorList, rpcResp.List)
	if err != nil {
		return nil, err
	}
	resp.Count = strconv.Itoa(len(rpcResp.List))
	return
}
