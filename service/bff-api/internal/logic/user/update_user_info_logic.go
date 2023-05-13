package user

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/bff-api/internal/svc"
	"MaoerMovie/service/bff-api/internal/types"
	"MaoerMovie/service/user-rpc/types/user"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoRequest, file multipart.File, fileHeader *multipart.FileHeader) (resp *types.UpdateUserInfoResponse, err error) {
	resp = new(types.UpdateUserInfoResponse)
	userId, err := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))).Int64()
	if err != nil {
		return nil, err
	}

	//将form的文件转换为byte数组
	defer file.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return nil, err
	}
	avatar := buf.Bytes()

	_, err = l.svcCtx.UserRPC.UpdateUserInfo(l.ctx, &user.UpdateUserInfoRequest{
		Id:         utils.Int64ToString(userId),
		Name:       req.Name,
		Gender:     req.Gender,
		Phone:      req.Phone,
		AvatarName: fileHeader.Filename,
		Avatar:     avatar,
	})
	if err != nil {
		return
	}

	return
}
