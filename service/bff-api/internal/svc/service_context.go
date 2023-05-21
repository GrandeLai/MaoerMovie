package svc

import (
	"MaoerMovie/common/utils"
	"MaoerMovie/service/bff-api/internal/config"
	"MaoerMovie/service/cinema-rpc/cinemarpc"
	"MaoerMovie/service/comment-rpc/commentrpc"
	"MaoerMovie/service/film-rpc/filmrpc"
	"MaoerMovie/service/order-rpc/orderrpc"
	"MaoerMovie/service/pay-rpc/payrpc"
	"MaoerMovie/service/user-rpc/userrpc"
	"github.com/smartwalle/alipay/v3"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	UserRPC      userrpc.UserRpc
	FilmRPC      filmrpc.FilmRpc
	CinemaRPC    cinemarpc.CinemaRpc
	CommentRPC   commentrpc.CommentRpc
	OrderRPC     orderrpc.OrderRpc
	AlipayClient *alipay.Client
	PayRPC       payrpc.PayRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := alipay.New(utils.AppId, utils.PrivateKey, false)
	if err != nil {
		panic(err)
	}
	client.LoadAliPayPublicKey(utils.AliPublicKey)
	return &ServiceContext{
		Config:       c,
		UserRPC:      userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRPC)),
		FilmRPC:      filmrpc.NewFilmRpc(zrpc.MustNewClient(c.FilmRpc)),
		CinemaRPC:    cinemarpc.NewCinemaRpc(zrpc.MustNewClient(c.CinemaRpc)),
		CommentRPC:   commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpc)),
		OrderRPC:     orderrpc.NewOrderRpc(zrpc.MustNewClient(c.OrderRpc)),
		AlipayClient: client,
		PayRPC:       payrpc.NewPayRpc(zrpc.MustNewClient(c.PayRpc)),
	}
}
