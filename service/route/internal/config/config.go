package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf   cache.CacheConf // redis缓存
	GateRpcConf zrpc.RpcClientConf
	AuthConf    struct {
		AccessSecret string
		AccessExpire int64
	}
	PostgreSQLConf PostgreSQLConf
}

type PostgreSQLConf struct {
	Tablename  string
	DataSource string
}
