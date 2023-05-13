package film

import (
	"MaoerMovie/service/film-rpc/types/pb"
	"context"
	"github.com/jinzhu/copier"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilmDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmDetailLogic {
	return &FilmDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilmDetailLogic) FilmDetail(req *types.FilmRequest) (resp *types.FilmDetailResponse, err error) {
	resp = new(types.FilmDetailResponse)
	rpcResp, err := l.svcCtx.FilmRPC.GetFilmDetail(l.ctx, &pb.FilmRequest{Id: req.Id})
	if err != nil {
		return resp, err
	}
	err = copier.Copy(&resp, rpcResp.FilmDetail)
	if err != nil {
		return nil, err
	}
	return
}
