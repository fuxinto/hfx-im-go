package logic

import (
	"HIMGo/pkg/fxerror"

	"HIMGo/pkg/pb"
	"HIMGo/service/route/msgSeq"
	"HIMGo/service/route/routeClient"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) MsgReqHandler(body []byte, channelId string) (*routeClient.MessagePushReply, error) {
	msg := pb.Message{}
	err := proto.Unmarshal(body, &msg)
	if err != nil {
		logx.Error("msg请求解析失败")
	}
	l.msgSendHander(msg)
	return &routeClient.MessagePushReply{}, nil

	//ack := pb.LoginAck{}
	//ack.Code = int32(200)
	//ack.UserId = "str"
	//ack.Msg = "登录成功"
}

func (l *GatePushMsgLogic) msgSendHander(msg pb.Message) {
	var cid string
	err := l.svcCtx.Cache.Get(msg.TargetId, &cid)
	if err != nil {
		if !errors.Is(err, fxerror.RedisNotFound) {
			//对方不在线，离场缓存+设备推送
		}
		//其它错误
		//日志记录
		logx.Errorf("用户redis key:", msg.TargetId, "错误")
	} else {
		uid := msgSeq.GetMsgSeq(msg.MessageId, 15)
		msgAck := &pb.MessageAck{}
		msgAck.MessageId = msg.MessageId
		msgAck.MessagedUid = uid

		body, err := pb.NewFrom(pb.PackType_msgAck, msgAck)
		if err != nil {
			logx.Error("PackType_loginAck,proto编码失败")
		}
		l.pushGata(cid, body)
	}

}
