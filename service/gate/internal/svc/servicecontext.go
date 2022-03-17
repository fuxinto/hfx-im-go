package svc

import (
	"HIMGo/service/gate/internal/config"
	"HIMGo/service/route/route"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	RouteRpc route.Route
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		RouteRpc: route.NewRoute(zrpc.MustNewClient(c.RouteRpcConf)),
	}
}
