package admin

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/film-rpc/types/pb"
	"context"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminFilmListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminFilmListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminFilmListLogic {
	return &AdminFilmListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminFilmListLogic) AdminFilmList(req *types.FilmListRequest) (resp *types.AdminFilmListResponse, err error) {
	resp = new(types.AdminFilmListResponse)
	pageSize := req.Size
	if req.Size == "0" {
		pageSize = "8"
	}
	startPage := req.Page
	if startPage == "0" {
		startPage = "1"
	}
	rpcResp, err := l.svcCtx.FilmRPC.AdminFilmList(l.ctx, &pb.ListRequest{
		Size: pageSize,
		Page: startPage,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.AdminFilm
	for _, film := range rpcResp.List {
		fList := &types.AdminFilm{
			FilmName:        film.FilmName,
			FilmEnglishName: film.FilmEnglishName,
			FilmType:        film.FilmType,
			FilmCover:       film.FilmCover,
			FilmLength:      film.FilmLength,
			FilmCategory:    film.FilmCategory,
			FilmArea:        film.FilmArea,
			FilmTime:        film.FilmTime,
			Director:        film.Director,
		}
		list = append(list, fList)
	}
	resp.List = list
	resp.Count = utils.IntToString(len(list))
	return
}
