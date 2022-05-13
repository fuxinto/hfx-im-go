package logic

import (
	"HIMGo/pkg/pb"
	"HIMGo/service/gate/gateClient"
	"HIMGo/service/route/routeClient"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) Handler(in *routeClient.GateReq) (*routeClient.RouteReply, error) {
	////未登录拦截
	//if c.IsSignIn == false && pack.Type != pb.PackType_LOGIN_REQ {
	//	c.Release()
	//	return
	//}

	pack := pb.Pack{}
	err1 := proto.Unmarshal(in.Message, &pack)
	if err1 != nil {
		logx.Error("proto解析失败")
	}

	switch pack.Type {
	case pb.PackType_loginReq:
		l.LoginHandler(pack.Body, in.ChannelId)
	case pb.PackType_msgReq:
		l.MsgReqHandler(pack.Body, in.ChannelId)
	// case pb.PackType_sessionPull:
	// l.SessionPullHandler(pack.Body, in.ChannelId)
	case pb.PackType_msgPullReq:
		l.MsgPullHandler(pack.Body, in.ChannelId)

	default:
		str := fmt.Sprintf("%v未设置处理handler", pack.Type)
		logx.Error(str)
		return &routeClient.RouteReply{}, fmt.Errorf(str)
	}
	return &routeClient.RouteReply{}, nil
}

func (l *GatePushMsgLogic) pushGata(channelId string, body []byte) error {
	req := &gateClient.MessagePushReq{
		ChannelId: channelId,
		Body:      body,
	}
	_, err := l.svcCtx.GateRpc.RoutePushMsg(l.ctx, req)
	if err != nil {
		logx.Errorf("pushGate失败，channelId:", channelId)
		return err
	}
	return nil
}
