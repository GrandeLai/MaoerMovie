package admin

import (
	"MaoerMovie/service/film-rpc/types/pb"
	"bytes"
	"context"
	"io"
	"mime/multipart"

	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FilmNewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFilmNewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FilmNewLogic {
	return &FilmNewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FilmNewLogic) FilmNew(req *types.FilmNewRequest, file multipart.File, fileHeader *multipart.FileHeader) (resp *types.FilmCreateResponse, err error) {
	resp = new(types.FilmCreateResponse)

	//将form的文件转换为byte数组
	defer file.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, err
	}
	cover := buf.Bytes()
	if req.FilmId == "" {
		data, err := l.svcCtx.FilmRPC.CreateFilm(l.ctx, &pb.FilmCreateRequest{
			FilmName:        req.FilmName,
			FilmEnglishName: req.FilmEnglishName,
			FilmType:        req.FilmType,
			CoverName:       fileHeader.Filename,
			FilmCover:       cover,
			FilmLength:      req.FilmLength,
			CategoryId:      req.FilmCategory,
			FilmArea:        req.FilmArea,
			FilmTime:        req.FilmTime,
			DirectorId:      req.Director,
			Biography:       req.Biography,
			ActorList:       req.ActorList,
			RoleList:        req.RoleList,
		})
		if err != nil {
			return nil, err
		}
		resp.Id = data.Id
	} else {
		resp.Id = req.FilmId
		_, err := l.svcCtx.FilmRPC.UpdateFilm(l.ctx, &pb.FilmUpdateRequest{
			FilmName:        req.FilmName,
			FilmEnglishName: req.FilmEnglishName,
			FilmType:        req.FilmType,
			FilmCoverName:   fileHeader.Filename,
			FilmCover:       cover,
			FilmLength:      req.FilmLength,
			FilmCategory:    req.FilmCategory,
			FilmArea:        req.FilmArea,
			FilmTime:        req.FilmTime,
			Director:        req.Director,
		})
		if err != nil {
			return nil, err
		}
	}
	return
}
