package logic

import (
	"context"

	"HIMGo/service/gate/gateClient"
	"HIMGo/service/gate/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoutePushMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoutePushMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoutePushMsgLogic {
	return &RoutePushMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoutePushMsgLogic) RoutePushMsg(in *gateClient.MessagePushReq) (*gateClient.MessagePushReply, error) {
	// todo: add your logic here and delete this line

	return &gateClient.MessagePushReply{}, nil
}
