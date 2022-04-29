package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"HIMGo/service/route/internal/config"
	"HIMGo/service/route/internal/server"
	"HIMGo/service/route/internal/svc"
	"HIMGo/service/route/routeClient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/route.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewRouteServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		routeClient.RegisterRouteServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	logx.Error("测试")
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
