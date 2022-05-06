package logic

import (
	"HIMGo/pkg/pb"
	"HIMGo/service/route/model"
	"HIMGo/service/route/routeClient"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) SessionPullHandler(body []byte, channelId string) (*routeClient.RouteReply, error) {
	pull := pb.SessionPull{}
	err := proto.Unmarshal(body, &pull)
	if err != nil {
		logx.Error("会话列表拉body解析失败")
	}
	push := pb.SessionPush{}

	sessions := []model.Session{}
	err1 := l.svcCtx.Db.Where("timestamp > ? AND user_id = ?", pull.Timestamp, pull.UserId).Limit(20).Find(&sessions).Error
	if err1 != nil {
		return &routeClient.RouteReply{}, nil
	}
	for _, v := range sessions {
		session := pb.Session{}

		var count int64 = 0
		l.svcCtx.Db.Where("target_id = ?", v.UserId).Count(&count)

		session.UnreadCount = count
		session.IsPinned = v.Pinned
		session.SessionId = v.SessionId

		session.IsOfflineMsg = count > 20

		var msgs = []model.Message{}
		err2 := l.svcCtx.Db.Where("target_id = ?", v.UserId).Order("timestamp DESC").Limit(20).Find(&msgs).Error
		if err2 == nil {
			v.LatestMsg = msgs
		}
		push.SessionList = append(push.SessionList, &session)
	}

	data, err4 := proto.Marshal(&push)
	if err4 != nil {
		return nil, err4
	}

	return &routeClient.RouteReply{Body: data}, nil
}
