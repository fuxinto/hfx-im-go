package logic

import (
	"HIMGo/pkg/pb"
	"HIMGo/service/route/model"
	"HIMGo/service/route/routeClient"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) MsgPullHandler(body []byte, channelId string) error {
	logx.Infof("心跳拉取消息")
	return &routeClient.RouteReply{}, nil

	pull := &pb.MessagePullReq{}
	err := proto.Unmarshal(body, pull)
	if err != nil {
		return err
	}
	if pull.Timestamp > 0 {
		err = l.svcCtx.Db.Where("user_id = ? AND timestamp <= ? ", pull.UserId, pull.Timestamp).Delete(model.MsgSync{}).Error
		if err != nil {
			return err
		}
	}
	msgs := []model.MsgSync{}
	//每页同步数据页数
	limit := 200
	err1 := l.svcCtx.Db.Where("user_id = ? AND timestamp > ?", pull.Timestamp, pull.UserId).Order("timestamp ASC").Limit(limit).Find(&msgs).Error
	push := pb.MessagePullAck{}
	if err1 == nil {
		num := len(msgs)
		pbMsgs := make([]*pb.Message, num, num)
		for _, v := range msgs {
			msg := &pb.Message{}
			msg.CloudCustomData = v.CloudCustomData
			msg.Content = v.Content
			msg.FaceURL = v.FaceURL
			msg.MsgID = v.MsgId
			msg.MsgUid = v.MsgUid
			msg.NickName = v.NickName
			msg.Sender = v.Sender
			msg.ConversationId = v.ConversationId
			msg.ConversationType = pb.ConversationType(v.ConversationType)
			msg.Status = pb.MessageStatus(v.Status)
			msg.Timestamp = v.Timestamp
			msg.Type = pb.ElemType(v.Type)
			pbMsgs = append(pbMsgs, msg)
		}
		push.Msglist = pbMsgs
	}
	data, err4 := proto.Marshal(&push)
	if err4 != nil {
		return err4
	}
	return nil
}
