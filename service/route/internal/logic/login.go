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

func (l *GatePushMsgLogic) newPack(code int32, msg, userId string, err error) (*routeClient.RouteReply, error) {
	loginAck := pb.LoginAck{Code: code, Msg: msg, UserId: userId}
	data, err1 := pb.NewFrom(pb.PackType_loginAck, &loginAck)
	if err1 != nil {
		return &routeClient.RouteReply{}, err1
	}
	return &routeClient.RouteReply{Body: data}, err
}

func (l *GatePushMsgLogic) LoginHandler(body []byte, channelId string) (*routeClient.RouteReply, error) {
	signin := pb.LoginReq{}
	err := proto.Unmarshal(body, &signin)
	if err != nil {
		errx := errors.New("登录请求proto解析失败" + err.Error())
		logx.Error(errx.Error())
		return l.newPack(20002, "其它错误", "", errx)
	}
	//解析token获取uid
	uid, err := jwtx.GetUidWithToken(signin.Token, l.svcCtx.Config.AuthConf.AccessSecret)
	if err != nil {

		return l.newPack(20001, "token错误", "", err)
	}
	var cid string
	err = l.svcCtx.Cache.Get(uid, &cid)
	if err != nil {
		//uid key不存在，登录操作
		if errors.Is(err, fxerror.RedisNotFound) {
			l.svcCtx.Cache.Set(uid, channelId)
			return l.newPack(20000, "登录成功", uid, nil)
		}
		return l.newPack(20002, "其它错误", "", err)
	}

	//已有登录，挤下线
	// go l.logout(uid, channelId)
	//if err != nil {
	//	return &routeClient.MessagePushReply{}, err
	//}
	err = l.svcCtx.Cache.Set(uid, channelId)
	if err != nil {
		errx := errors.New("redis设置uid key失败：" + err.Error())
		logx.Error(errx.Error())
		return l.newPack(20002, "其它错误", "", errx)
	}
	return l.newPack(20000, "登录成功", uid, nil)
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
