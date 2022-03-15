package main

import (
	"context"
	"route/kitex_gen/routeRpc"
	"route/kitex_gen/rpc"
)

// RouteRpcImpl implements the last service interface defined in the IDL.
type RouteRpcImpl struct{}

// Login implements the RouteRpcImpl interface.
func (s *RouteRpcImpl) Login(ctx context.Context, req *routeRpc.LoginReq) (resp *routeRpc.LoginReply, err error) {
	// TODO: Your code here...
	return
}

// GatePushMsg implements the RouteRpcImpl interface.
func (s *RouteRpcImpl) GatePushMsg(ctx context.Context, req *rpc.MessagePushReq) (resp *rpc.MessagePushReply, err error) {
	// TODO: Your code here...
	return
}
