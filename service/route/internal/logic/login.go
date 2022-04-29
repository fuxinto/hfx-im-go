package logic

import (
	"HIMGo/pkg/fxerror"
	"HIMGo/pkg/jwtx"
	"HIMGo/pkg/pb"

	"HIMGo/service/route/routeClient"

	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) LoginHandler(body []byte, channelId string) (*routeClient.MessagePushReply, error) {
	signin := pb.LoginReq{}
	err := proto.Unmarshal(body, &signin)
	if err != nil {
		logx.Error("登录请求解析失败")
	}

	//解析token获取uid
	uid, err := jwtx.GetUidWithToken(signin.Token, l.svcCtx.Config.AuthConf.AccessSecret)
	if err != nil {
		logx.Errorf("token:", signin.Token, "解析错误:", err)
		return &routeClient.MessagePushReply{}, err
	}
	var cid string
	err = l.svcCtx.Cache.Get(uid, &cid)
	if err != nil {
		//uid key不存在，登录操作
		if errors.Is(err, fxerror.RedisNotFound) {
			l.svcCtx.Cache.Set(uid, channelId)
			return &routeClient.MessagePushReply{}, nil

		}
		return &routeClient.MessagePushReply{}, err
	}

	//已有登录，挤下线
	// go l.logout(uid, channelId)
	//if err != nil {
	//	return &routeClient.MessagePushReply{}, err
	//}
	err = l.svcCtx.Cache.Set(uid, channelId)
	if err != nil {
		logx.Errorf("redis设置uid key失败：", err)
		return &routeClient.MessagePushReply{}, err
	}
	return &routeClient.MessagePushReply{}, nil
	//ack := pb.LoginAck{}
	//ack.Code = int32(200)
	//ack.UserId = "str"
	//ack.Msg = "登录成功"
}

func (l *GatePushMsgLogic) logout(uid, channelId string) {
	loginAck := &pb.LoginAck{
		Code:   1001,
		Msg:    "在别地登录",
		UserId: uid,
	}
	body, err := pb.NewFrom(pb.PackType_loginAck, loginAck)
	if err != nil {
		logx.Error("loginAck,proto编码失败")
	} else {
		l.pushGata(channelId, body)
	}
}
