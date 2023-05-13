package admin

import (
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"MaoerMovie/service/film-rpc/types/pb"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type FilmDetailAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmDetailAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmDetailAddLogic {
	return &FilmDetailAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilmDetailAddLogic) FilmDetailAdd(req *types.FilmDetailAddRequest, filmByte [][]byte, filmNameList []string) (resp *types.CommonResponse, err error) {

	_, err = l.svcCtx.FilmRPC.UpdateFilmInfo(l.ctx, &pb.FilmInfoUpdateRequest{FilmImgs: filmByte, FilmNameList: filmNameList, FilmPreSaleSum: req.FilmPreSaleNum, FilmId: req.FilmId})
	if err != nil {
		return nil, err
	}
	return &types.CommonResponse{}, nil
}
