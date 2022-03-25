package svc

import (
	"HIMGo/service/gate/internal/config"
	"HIMGo/service/gate/socket"
	"HIMGo/service/gate/socket/tcp"
	"HIMGo/service/route/route"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config     config.Config
	HIMService socket.Server
}

func NewServiceContext(c config.Config) *ServiceContext {

	handler := &socket.Handler{
		ServiceID: c.Name,
		RouteRpc:  route.NewRoute(zrpc.MustNewClient(c.RouteRpcConf)),
	}

	ser := tcp.NewServer(c.SocketConf.Port)
	ser.SetReadWait(time.Minute)
	ser.SetAcceptor(handler)
	ser.SetMessageListener(handler)
	ser.SetStateListener(handler)

	go startHIMService(ser)
	return &ServiceContext{
		Config:     c,
		HIMService: ser,
	}
}

func startHIMService(ser socket.Server) {
	err := ser.Start()
	if err != nil {
		panic(err)
	}
}
