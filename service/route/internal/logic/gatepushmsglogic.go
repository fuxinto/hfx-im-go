package logic

import (
	"context"

	"HIMGo/service/route/internal/svc"
	"HIMGo/service/route/routeClient"

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

func (l *GatePushMsgLogic) GatePushMsg(in *routeClient.MessagePushReq) (*routeClient.MessagePushReply, error) {
	// todo: add your logic here and delete this line

	return &routeClient.MessagePushReply{}, nil
}
