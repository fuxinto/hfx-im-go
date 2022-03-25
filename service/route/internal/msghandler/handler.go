package msghandler

import (
	"HIMGo/pkg/pb"
	"HIMGo/service/route/routeClient"
	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessageHandler interface {
	handler(msg []byte) (*routeClient.MessagePushReply, error)
}

var loginHandler = LoginHandler{}

func Handler(in *routeClient.MessagePushReq) (*routeClient.MessagePushReply, error) {
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
		//登录
		return loginHandler.Handler(pack.Body)
	}
	return nil, nil
}
