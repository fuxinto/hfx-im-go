package logic

import (
	"HIMGo/service/route/routeClient"

	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) MsgPullHandler(body []byte, channelId string) (*routeClient.RouteReply, error) {
	logx.Infof("心跳拉取消息")
	return &routeClient.RouteReply{}, nil
}
