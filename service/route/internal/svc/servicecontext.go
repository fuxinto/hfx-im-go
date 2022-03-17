package svc

import (
	"HIMGo/service/route/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cache:  cache.New(c.RedisConf, syncx.NewSingleFlight(), cache.NewStat("him"), nil),
	}
}
