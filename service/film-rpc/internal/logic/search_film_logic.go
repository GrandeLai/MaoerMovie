package logic

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/model/esClient"
	"MaoerMovie/service/film-rpc/types/pb"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFilmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFilmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFilmLogic {
	return &SearchFilmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFilmLogic) SearchFilm(in *pb.SearchFilmRequest) (*pb.SearchFilmResponse, error) {
	var resp []*pb.FilmPreview
	filmList, num, err := l.svcCtx.EsModel.SearchFilm(l.ctx, &esClient.SearchFilterFactor{
		Keyword:    in.Keyword,
		Category:   in.Category,
		Page:       utils.StringToInt64(in.Page),
		Size:       utils.StringToInt64(in.Size),
		Area:       in.Area,
		SortedType: utils.StringToInt64(in.SortedType),
		TimeStart:  utils.StringToInt64(in.TimeStart),
		TimeEnd:    utils.StringToInt64(in.TimeEnd),
	})
	if err != nil {
		return nil, err
	}
	for _, film := range filmList {
		pre := &pb.FilmPreview{
			FilmName:      film.FilmName,
			FilmTime:      utils.TimeToString(film.FilmTime),
			FilmCategory:  film.FilmCategory,
			FilmScore:     utils.Float64ToString(film.FilmScore),
			FilmCoverUrl:  film.FilmCoverUrl,
			ActorNameList: film.ActorNameList,
			FilmArea:      film.FilmArea,
		}
		resp = append(resp, pre)
	}
	return &pb.SearchFilmResponse{Films: resp, Total: num}, nil
}
