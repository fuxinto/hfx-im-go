package svc

import (
	"HIMGo/service/gate/gate"
	"HIMGo/service/route/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config  config.Config
	Cache   cache.Cache
	GateRpc gate.Gate
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cache:  cache.New(c.RedisConf, syncx.NewSingleFlight(), cache.NewStat("him"), nil),
		//GateRpc: gate.NewGate(zrpc.MustNewClient(c.GateRpcConf)),
	}
}
