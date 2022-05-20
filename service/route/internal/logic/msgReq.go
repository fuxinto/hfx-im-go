package logic

import (
	"strings"

	"HIMGo/pkg/pb"
	"HIMGo/service/route/model"
	"HIMGo/service/route/msgSeq"
	"HIMGo/service/route/routeClient"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/internal/errors"
)

func (l *GatePushMsgLogic) MsgReqHandler(body []byte, channelId string) error {
	msg := &pb.Message{}
	err := proto.Unmarshal(body, msg)
	if err != nil {
		return err
	}
	msg.MsgUid, msg.Timestamp = msgSeq.GetMsgSeq(msg.SessionId, int(msg.Type))
	msgAck := pb.MessageAck{MsgId: msg.MsgID, Code: 2000, MsgUid: msg.MsgUid}
	//消息保存
	err = l.saveMsgStorage(msg)
	if err != nil {
		msgAck.Code = 2002
	} else {
		//消息推送
		err = l.msgSendHander(msg)
	}
	return err
}

func (l *GatePushMsgLogic) saveMsgStorage(msg *pb.Message) error {
	var err errors
	switch msg.ConversationType {
	case pb.ConversationType_c2c:
		msgs := [2]model.MsgStorage{}
		msgs[0] := &model.MsgStorage{ConversationId: msg.TargetId, UserId: msg.Sender, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), Sender: msg.Sender,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		
		msgs[1] := &model.MsgStorage{ConversationId: msg.Sender, UserId: msg.TargetId, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), Sender: msg.Sender,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		err = l.svcCtx.Db.Create(&msgs).Error 
		
	case pb.SessionType_group:
		groupMsg := model.MsgStorage{SessionId: msg.SessionId, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), SenderId: msg.SenderId,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		err = l.svcCtx.Db.Create(&groupMsg).Error
	}
	if err != nil {
		return errors.Wrap(err,"msg：%+v",msg)
	}
	return err
}
func createRedisSessionKey(userId string) string {
	var b strings.Builder
	b.WriteString(userId)
	b.WriteString(":sessionList")
	return b.String()
}

func (l *GatePushMsgLogic) msgAddSync(msg *pb.Message) error {
	var err errors
	switch msg.SessionType {
	case pb.SessionType_c2c:
		msgs := [2]model.MsgStorage{}
		msgs[0] := &model.MsgStorage{ConversationId: msg.TargetId, UserId: msg.Sender, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), Sender: msg.Sender,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		
		msgs[1] := &model.MsgStorage{ConversationId: msg.Sender, UserId: msg.TargetId, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), Sender: msg.Sender,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		err = l.svcCtx.Db.Create(&msgs).Error 
		

	case pb.SessionType_group:
		//循序查询群成员，
		// l.svcCtx.Cache.Get()
	}
	if err != nil {
		return errors.Wrap(err,"msg：%+v",msg)
	}
	return err
}

func (l *GatePushMsgLogic) msgSendHander(msg *pb.Message) {
	var cid string
	err := l.svcCtx.Cache.Get(msg.SenderId, &cid)
	if err != nil {
		if errors.Is(err, fxerror.RedisNotFound) {
			//对方不在线，离场缓存+设备推送
			err = l.msgAddSync(msg)
			if err == nil {
				//离线消息添加成功，会话消息更新成功,推送push

			}
		}
		
		//其它错误
		//日志记录
		logx.Errorf("用户redis key:", msg.SenderId, "错误")
	} else {

		msg.SessionId = msg.SenderId
		body, err := pb.NewFrom(pb.PackType_msgAck, msg)
		if err != nil {
			return err
		}
		l.pushGata(cid, body)
	}

}
