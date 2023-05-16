package svc

import (
	"MaoerMovie/service/bff-api/internal/config"
	"MaoerMovie/service/cinema-rpc/cinemarpc"
	"MaoerMovie/service/film-rpc/filmrpc"
	"MaoerMovie/service/user-rpc/userrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserRPC   userrpc.UserRpc
	FilmRPC   filmrpc.FilmRpc
	CinemaRPC cinemarpc.CinemaRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserRPC:   userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRPC)),
		FilmRPC:   filmrpc.NewFilmRpc(zrpc.MustNewClient(c.FilmRpc)),
		CinemaRPC: cinemarpc.NewCinemaRpc(zrpc.MustNewClient(c.CinemaRpc)),
	}
}
