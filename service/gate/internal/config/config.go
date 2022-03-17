package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	RouteRpcConf zrpc.RpcClientConf
	SocketConf   SocketConf
}
type SocketConf struct {
	Port         string
	MaxConnCount int
}
