package svc

import (
	"HIMGo/pkg/fxerror"
	"HIMGo/pkg/fxgorm"
	"HIMGo/service/gate/gate"
	"HIMGo/service/route/internal/config"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	Cache   cache.Cache
	GateRpc gate.Gate
	Db      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		Cache:  cache.New(c.RedisConf, syncx.NewSingleFlight(), cache.NewStat("him"), fxerror.RedisNotFound),
		//GateRpc: gate.NewGate(zrpc.MustNewClient(c.GateRpcConf)),
		Db: fxgorm.GormPgSql(c.PostgreSQLConf.DataSource),
	}
}
