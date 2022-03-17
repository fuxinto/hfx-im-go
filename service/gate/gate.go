package main

import (
	"HIMGo/service/gate/gateClient"
	"HIMGo/service/gate/internal/config"
	"HIMGo/service/gate/internal/server"
	"HIMGo/service/gate/internal/svc"
	"HIMGo/service/route/route"
	"context"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/gate.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewGateServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		gateClient.RegisterGateServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	req := route.LoginReq{ChannelId: "channelid333", Uid: "uid123"}
	_, err := ctx.RouteRpc.Login(context.TODO(), &req)
	if err != nil {
		fmt.Println("rpc失败")
	}
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
