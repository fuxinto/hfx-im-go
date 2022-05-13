package logic

import (
	"HIMGo/pkg/pb"

	"HIMGo/service/route/model"
	"HIMGo/service/route/routeClient"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"

	"gorm.io/gorm"
)

func (l *GatePushMsgLogic) SessionPullHandler(body []byte, channelId string) (*routeClient.RouteReply, error) {
	return &routeClient.RouteReply{}, nil

	pull := pb.SessionPull{}
	err := proto.Unmarshal(body, &pull)
	if err != nil {
		logx.Error("会话列表拉body解析失败")
	}
	push := pb.SessionPush{}
	var count int64 = 0
	sessions := []model.Session{}
	if pull.Timestamp > 0 {
		l.svcCtx.Db.Where("user_id = ? AND timestamp > ?", pull.UserId)
	}
	err = l.svcCtx.Db.Where("unread_count > 0 AND user_id = ?", pull.UserId).Count(&count).Order("timestamp DESC").Limit(20).Find(&sessions).Error
	if err == nil || err == gorm.ErrRecordNotFound {
		for _, v := range sessions {
			//查询离线消息表
			msgs := []model.MsgSync{}
			var count int64 = 0
			pbSession := pb.Session{}
			//查询未读消息条数
			// switch int(v.SessionType) {
			// case model.SessionType_c2c:
			// 	//查询单聊
			// 	err = l.svcCtx.Db.Model(&model.MsgSync{}).Where("sender_id = ?", v.SessionId).Order("timestamp DESC").Limit(20).Find(&msgs).Error
			// case model.SessionType_group:

			// }
			err = l.svcCtx.Db.Model(&model.MsgSync{}).Where("session_id = ?", v.SessionId).Order("timestamp DESC").Limit(20).Find(&msgs).Error
			if err != nil || err != gorm.ErrRecordNotFound {
				logx.Errorf("查询消息同步列表失败;error:", err.Error())
				return nil, err
			}
			if len(msgs) > 0 {
				pbMsgs := make([]*pb.Message, 20, 20)
				for _, msg := range msgs {
					pbMsg := pb.Message{}
					pbMsg.CloudCustomData = msg.CloudCustomData
					pbMsg.Content = msg.Content
					pbMsg.FaceURL = msg.FaceURL
					pbMsg.MsgID = msg.MsgId
					pbMsg.MsgUid = msg.MsgUid
					pbMsg.NickName = msg.NickName
					pbMsg.SenderId = msg.SenderId
					pbMsg.SessionId = msg.SessionId
					pbMsg.SessionType = pb.SessionType(msg.SessionType)
					pbMsg.Status = pb.MessageStatus(msg.Status)
					pbMsg.Timestamp = msg.Timestamp
					pbMsg.Type = pb.ElemType(msg.Type)
					pbMsgs = append(pbMsgs, &pbMsg)
				}
				pbSession.Msglist = pbMsgs
			}
			pbSession.UnreadCount = count
			pbSession.IsPinned = v.Pinned
			pbSession.SessionId = v.SessionId
			pbSession.IsOfflineMsg = count == 0
			push.SessionList = append(push.SessionList, &pbSession)
		}
		data, err4 := proto.Marshal(&push)
		if err4 != nil {
			logx.Errorf("SessionPullHandler;error:", err4.Error())
			return nil, err4
		}
		return &routeClient.RouteReply{Body: data}, nil
	}
	logx.Errorf("查询会话列表失败;error:", err.Error())
	return nil, err
}
