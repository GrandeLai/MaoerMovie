package admin

import (
	"MaoerMovie/service/film-rpc/types/pb"
	"context"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilmRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmRemoveLogic {
	return &FilmRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilmRemoveLogic) FilmRemove(req *types.FilmRequest) (resp *types.FilmResponse, err error) {
	resp = new(types.FilmResponse)
	_, err = l.svcCtx.FilmRPC.DeleteFilm(l.ctx, &pb.FilmRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
