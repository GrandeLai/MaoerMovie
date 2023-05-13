package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminFilmListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminFilmListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminFilmListLogic {
	return &AdminFilmListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminFilmListLogic) AdminFilmList(in *pb.ListRequest) (*pb.AdminFilmListResponse, error) {
	page := utils.StringToInt64(in.Page)
	size := utils.StringToInt64(in.Size)
	resp, err := l.svcCtx.FilmModel.FindAllInPage(l.ctx, page, size)
	if err != nil {
		return nil, err
	}
	var list []*pb.Film
	for _, film := range resp {
		category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, film.CategoryId)
		if err != nil {
			return nil, err
		}
		director, err := l.svcCtx.ActorModel.FindOne(l.ctx, film.DirectorId)
		list = append(list, &pb.Film{
			FilmName:        film.FilmName,
			FilmEnglishName: film.FilmEnName,
			FilmType:        utils.Int64ToString(film.FilmType),
			FilmCover:       film.FilmCover,
			FilmLength:      utils.Int64ToString(film.FilmLength),
			FilmCategory:    category.CateName,
			FilmArea:        film.FilmArea,
			FilmTime:        film.FilmTime.Time.Format("2003-06-21"),
			Director:        director.ActorName,
		})
	}
	return &pb.AdminFilmListResponse{List: list}, nil
}
