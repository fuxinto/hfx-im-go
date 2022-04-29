package logic

import (
	"HIMGo/pkg/pb"
	"HIMGo/service/route/model"
	"HIMGo/service/route/routeClient"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) SessionPullHandler(body []byte, channelId string) (*routeClient.MessagePushReply, error) {
	return &routeClient.MessagePushReply{}, nil
	pull := pb.SessionPull{}
	err := proto.Unmarshal(body, &pull)
	if err != nil {
		logx.Error("会话列表拉body解析失败")
	}
	push := pb.SessionPush{}
	push.UnreadCount = 1
	sessions := []model.Session{}
	err1 := l.svcCtx.Db.Where("timestamp > ? AND user_id = ?", pull.Timestamp, pull.UserId).Limit(20).Find(&sessions).Error
	if err1 != nil {
		return &routeClient.MessagePushReply{}, nil
	}
	for _, v := range sessions {
		var msgs = []model.Message{}
		err2 := l.svcCtx.Db.Where("target_id = ?", v.UserId).Order("timestamp DESC").Limit(20).Find(&msgs).Error
		
	}
}
