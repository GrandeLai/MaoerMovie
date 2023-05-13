package logic

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/film-rpc/model"
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"path"
	"strings"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFilmInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFilmInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFilmInfoLogic {
	return &UpdateFilmInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFilmInfoLogic) UpdateFilmInfo(in *pb.FilmInfoUpdateRequest) (*pb.FilmResponse, error) {
	var des []string
	for i, byteFile := range in.FilmImgs {
		objectName := utils.GenerateUUID() + path.Ext(in.FilmNameList[i])
		_, err := l.svcCtx.MinioClient.PutObject(context.Background(), l.svcCtx.Config.MinIO.BucketName, objectName, bytes.NewReader(byteFile), int64(len(byteFile)),
			minio.PutObjectOptions{ContentType: "binary/octet-stream"})
		if err != nil {
			//return nil, errorx.NewDefaultError("服务器处理上传失败！")
			return nil, err
		}
		url := l.svcCtx.Config.MinIO.BucketName + "/" + objectName
		des = append(des, url)
	}
	imgs := strings.Join(des, ",")
	filmInfo := &model.FilmInfo{
		FilmId:         utils.StringToInt64(in.FilmId),
		FilmImgs:       imgs,
		FilmPreSaleNum: utils.StringToInt64(in.FilmPreSaleSum),
	}
	_, err := l.svcCtx.FilmInfoModel.FindOneByFilmId(l.ctx, utils.StringToInt64(in.FilmId))
	switch err {
	case nil:
		err := l.svcCtx.FilmInfoModel.Update(l.ctx, filmInfo)
		if err != nil {
			return nil, err
		}
	case sqlc.ErrNotFound:
		filmInfo.Id = utils.GenerateNewId(l.svcCtx.RedisClient, "film_info")
		l.svcCtx.FilmInfoModel.Insert(l.ctx, filmInfo)
	default:
		return nil, err
	}

	return &pb.FilmResponse{}, nil
}
