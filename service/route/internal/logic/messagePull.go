package logic

import (
	"HIMGo/pkg/pb"
	"HIMGo/service/route/model"
	"HIMGo/service/route/routeClient"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) MsgPullHandler(body []byte, channelId string) (*routeClient.RouteReply, error) {
	logx.Infof("心跳拉取消息")
	return &routeClient.RouteReply{}, nil

	pull := &pb.MessagePullReq{}
	err := proto.Unmarshal(body, pull)
	if err != nil {
		logx.Error("会话列表拉body解析失败")
		return &routeClient.RouteReply{}, nil
	}
	if pull.Timestamp > 0 {
		err = l.svcCtx.Db.Where("user_id = ? AND timestamp <= ? ", pull.UserId, pull.Timestamp).Delete(model.MsgSync{}).Error
		if err != nil {
			logx.Error("离线消息表删除匹配时间戳失败;error:", err.Error())
			return &routeClient.RouteReply{}, nil
		}
	}
	msgs := []model.MsgSync{}
	var count int64 = 0
	err1 := l.svcCtx.Db.Where("user_id = ? AND timestamp > ?", pull.Timestamp, pull.UserId).Count(&count).Order("timestamp ASC").Limit(20).Find(&msgs).Error
	push := pb.MessagePullAck{}
	if err1 == nil {
		pbMsgs := make([]*pb.Message, 20, 20)
		for _, v := range msgs {
			msg := &pb.Message{}
			msg.CloudCustomData = v.CloudCustomData
			msg.Content = v.Content
			msg.FaceURL = v.FaceURL
			msg.MsgID = v.MsgId
			msg.MsgUid = v.MsgUid
			msg.NickName = v.NickName
			msg.SenderId = v.SenderId
			msg.SessionId = v.SessionId
			msg.SessionType = pb.SessionType(v.SessionType)
			msg.Status = pb.MessageStatus(v.Status)
			msg.Timestamp = v.Timestamp
			msg.Type = pb.ElemType(v.Type)
			pbMsgs = append(pbMsgs, msg)
		}
		push.Msglist = pbMsgs
	}
	push.NLeft = count
	data, err4 := proto.Marshal(&push)
	if err4 != nil {
		logx.Errorf("data, err4 := proto.Marshal(&push),", err4.Error())
	}
	return &routeClient.RouteReply{Body: data}, nil
}
