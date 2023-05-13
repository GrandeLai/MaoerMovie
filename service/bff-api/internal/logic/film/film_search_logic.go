package film

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"MaoerMovie/service/film-rpc/types/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilmSearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmSearchLogic {
	return &FilmSearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilmSearchLogic) FilmSearch(req *types.FilmSearchRequest) (resp *types.FilmSearchResponse, err error) {
	resp = new(types.FilmSearchResponse)
	rpcResp, err := l.svcCtx.FilmRPC.SearchFilm(l.ctx, &pb.SearchFilmRequest{
		Keyword:    req.Keyword,
		Page:       req.Page,
		Size:       req.Size,
		Category:   req.Category,
		Area:       req.Area,
		SortedType: req.SortedType,
		TimeStart:  req.TimeStart,
		TimeEnd:    req.TimeEnd,
	})
	if err != nil {
		return resp, err
	}
	for _, film := range rpcResp.Films {
		pre := &types.FilmInfo{
			FilmName:      film.FilmName,
			FilmTime:      film.FilmTime,
			FilmCategory:  film.FilmCategory,
			FilmScore:     film.FilmScore,
			FilmCover:     film.FilmCoverUrl,
			ActorNameList: film.ActorNameList,
		}
		resp.InfoList = append(resp.InfoList, pre)
	}
	resp.Count = utils.Int64ToString(rpcResp.Total)
	return
}
