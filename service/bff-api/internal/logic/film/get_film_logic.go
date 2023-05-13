package film

import (
	"MaoerMovie/service/film-rpc/types/pb"
	"context"
	"github.com/jinzhu/copier"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilmLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilmLogic {
	return &GetFilmLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFilmLogic) GetFilm(req *types.FilmRequest) (resp *types.GetFilmResponse, err error) {
	resp = new(types.GetFilmResponse)
	rpcResp, err := l.svcCtx.FilmRPC.GetFilm(l.ctx, &pb.FilmRequest{Id: req.Id})
	err = copier.Copy(&resp, rpcResp.Film)
	if err != nil {
		return nil, err
	}
	return
}
