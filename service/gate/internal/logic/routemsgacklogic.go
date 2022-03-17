package logic

import (
	"context"

	"HIMGo/service/gate/gateClient"
	"HIMGo/service/gate/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RouteMsgAckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRouteMsgAckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RouteMsgAckLogic {
	return &RouteMsgAckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RouteMsgAckLogic) RouteMsgAck(in *gateClient.MessageAckReq) (*gateClient.MessageAckReply, error) {
	// todo: add your logic here and delete this line

	return &gateClient.MessageAckReply{}, nil
}
