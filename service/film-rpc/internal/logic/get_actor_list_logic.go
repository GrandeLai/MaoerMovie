package logic

import (
	"MaoerMovie/common/utils"
	"context"

	"MaoerMovie/service/film-rpc/internal/svc"
	"MaoerMovie/service/film-rpc/types/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActorListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActorListLogic {
	return &GetActorListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActorListLogic) GetActorList(in *pb.ListRequest) (*pb.ActorListResponse, error) {
	page := utils.StringToInt64(in.Page)
	size := utils.StringToInt64(in.Size)
	resp, err := l.svcCtx.ActorModel.FindAllInPage(l.ctx, page, size)
	if err != nil {
		return nil, err
	}
	var list []*pb.Actor
	for _, actor := range resp {
		list = append(list, &pb.Actor{
			ActorId:     utils.Int64ToString(actor.Id),
			ActorName:   actor.ActorName,
			ActorAvatar: actor.ActorImg,
		})
	}
	return &pb.ActorListResponse{List: list, Count: string(len(list))}, nil
}
