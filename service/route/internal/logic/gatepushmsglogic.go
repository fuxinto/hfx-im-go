package logic

import (
	"HIMGo/service/route/internal/svc"
	"HIMGo/service/route/routeClient"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GatePushMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGatePushMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GatePushMsgLogic {
	return &GatePushMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GatePushMsgLogic) GatePushMsg(in *routeClient.GateReq) (*routeClient.RouteReply, error) {
	// todo: add your logic here and delete this line

	return l.Handler(in)
}
