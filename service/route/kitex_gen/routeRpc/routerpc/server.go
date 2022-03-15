// Code generated by Kitex v0.2.0. DO NOT EDIT.
package routerpc

import (
	"github.com/cloudwego/kitex/server"
	"route/kitex_gen/routeRpc"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler routeRpc.RouteRpc, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}