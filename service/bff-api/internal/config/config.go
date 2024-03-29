package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRPC    zrpc.RpcClientConf
	FilmRpc    zrpc.RpcClientConf
	CinemaRpc  zrpc.RpcClientConf
	CommentRpc zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf
	DtmServer  string
	Salt       string
	//AlipayClient alipay.Client
	PayRpc zrpc.RpcClientConf
}
