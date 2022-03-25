// Code generated by goctl. DO NOT EDIT!
// Source: route.proto

package server

import (
	"context"

	"HIMGo/service/route/internal/logic"
	"HIMGo/service/route/internal/svc"
	"HIMGo/service/route/routeClient"
)

type RouteServer struct {
	svcCtx *svc.ServiceContext
	routeClient.UnimplementedRouteServer
}

func NewRouteServer(svcCtx *svc.ServiceContext) *RouteServer {
	return &RouteServer{
		svcCtx: svcCtx,
	}
}

func (s *RouteServer) GatePushMsg(ctx context.Context, in *routeClient.MessagePushReq) (*routeClient.MessagePushReply, error) {
	l := logic.NewGatePushMsgLogic(ctx, s.svcCtx)
	return l.GatePushMsg(in)
}
