package logic

import (
	"MaoerMovie/common/errorx"
	"MaoerMovie/common/kqueue"
	"MaoerMovie/common/utils"
	"MaoerMovie/service/user-rpc/internal/svc"
	"MaoerMovie/service/user-rpc/types/user"
	"bytes"
	"context"
	"database/sql"
	jsoniter "github.com/json-iterator/go"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"path"
	"strings"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error) {
	userId := in.Id
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, utils.StringToInt64(userId))
	switch err {
	case nil:
		break
	case sqlx.ErrNotFound:
		return nil, errorx.NewCodeError(100, "查无此Id的用户！")
	default:
		return nil, err
	}
	if in.Name != "" {
		userInfo.Name = in.Name
	}
	gender := utils.StringToInt64(in.Gender)
	if gender != 0 {
		userInfo.Gender = gender
	}
	if in.Phone != "" {
		userInfo.Phone = in.Phone
	}
	var avatarUrl string
	if in.AvatarName != "" && len(in.Avatar) != 0 {
		if userInfo.AvatarUrl.String != "" {
			url := strings.Split(userInfo.AvatarUrl.String, "/")
			objectName := url[len(url)-1]
			err = l.svcCtx.MinioClient.RemoveObject(l.ctx, l.svcCtx.Config.MinIO.BucketName, objectName, minio.RemoveObjectOptions{})
			if err != nil {
				return nil, errorx.NewDefaultError("服务器删除对象失败")
			}
		}
		objectName := utils.GenerateUUID() + path.Ext(in.AvatarName)
		_, err = l.svcCtx.MinioClient.PutObject(context.Background(), l.svcCtx.Config.MinIO.BucketName, objectName, bytes.NewReader(in.Avatar), int64(len(in.Avatar)),
			minio.PutObjectOptions{ContentType: "binary/octet-stream"})
		if err != nil {
			//return nil, errorx.NewDefaultError("服务器处理上传失败！")
			return nil, err
		}
		avatarUrl = l.svcCtx.Config.MinIO.BucketName + "/" + objectName
		userInfo.AvatarUrl = sql.NullString{String: avatarUrl, Valid: true}
	}
	err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
	if err != nil {
		return nil, err
	}
	err = l.PubKqUserUpdateMessage(utils.StringToInt64(userId), in.Name, avatarUrl)
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserInfoResponse{}, nil
}

func (l *UpdateUserInfoLogic) PubKqUserUpdateMessage(id int64, Name string, AvatarUrl string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonStr, err := json.MarshalToString(kqueue.UserUpdateMessage{
		Id:        id,
		Name:      Name,
		AvatarUrl: AvatarUrl,
	})
	if err != nil {
		return err
	}
	err = l.svcCtx.KqUserUpdateClient.Push(jsonStr)
	if err != nil {
		return err
	}
	return nil
}
