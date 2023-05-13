package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCategoryLogic {
	return &GetAllCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllCategoryLogic) GetAllCategory(in *pb.CommonRequest) (*pb.GetCategoryListResponse, error) {
	resp, err := l.svcCtx.CategoryModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	var list []*pb.Category
	for _, cate := range resp {
		list = append(list, &pb.Category{
			CategoryId:   utils.Int64ToString(cate.Id),
			CategoryName: cate.CateName,
		})
	}
	return &pb.GetCategoryListResponse{Category: list}, nil
}
