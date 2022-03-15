// Code generated by Kitex v0.2.0. DO NOT EDIT.

package routerpc

import (
	"github.com/cloudwego/kitex/server"
	"route/kitex_gen/routeRpc"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler routeRpc.RouteRpc, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
