package logic

import (
	"HIMGo/pkg/fxerror"
	"strings"

	"HIMGo/pkg/pb"
	"HIMGo/service/route/model"
	"HIMGo/service/route/msgSeq"
	"HIMGo/service/route/routeClient"
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *GatePushMsgLogic) MsgReqHandler(body []byte, channelId string) (*routeClient.RouteReply, error) {
	msg := &pb.Message{}
	err := proto.Unmarshal(body, msg)
	if err != nil {
		logx.Error("msg请求解析失败")
	}
	msg.MsgUid, msg.Timestamp = msgSeq.GetMsgSeq(msg.SessionId, int(msg.Type))
	msgAck := pb.MessageAck{MsgId: msg.MsgID, Code: 2000, MsgUid: msg.MsgUid}
	//消息保存
	err = l.saveMsgStorage(msg)
	if err != nil {
		logx.Errorf("l.saveMsgStorage(msg) err:", err.Error())
		msgAck.Code = 2002
	} else {
		//消息推送
		l.msgSendHander(msg)
	}
	return &routeClient.RouteReply{}, nil
}

func (l *GatePushMsgLogic) saveMsgStorage(msg *pb.Message) error {
	switch msg.SessionType {
	case pb.SessionType_c2c:
		// 开始事务
		tx := l.svcCtx.Db.Begin()
		msgStorage1 := model.MsgStorage{SessionId: msg.SessionId, UserId: msg.SenderId, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), SenderId: msg.SenderId,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&msgStorage1).Error; err != nil {
			// 返回任何错误都会回滚事务
			// 遇到错误时回滚事务
			tx.Rollback()
			return err
		}
		msgStorage2 := model.MsgStorage{SessionId: msg.SenderId, UserId: msg.SessionId, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), SenderId: msg.SenderId,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&msgStorage2).Error; err != nil {
			// 返回任何错误都会回滚事务
			// 遇到错误时回滚事务
			tx.Rollback()
			return err
		}
		// 否则，提交事务
		tx.Commit()
	case pb.SessionType_group:
		groupMsg := model.MsgStorage{SessionId: msg.SessionId, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), SenderId: msg.SenderId,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		if err := l.svcCtx.Db.Create(&groupMsg).Error; err != nil {
			logx.Errorf("创建群消息历史失败；error:", err.Error())
		}
	}
	return nil
}
func createRedisSessionKey(userId string) string {
	var b strings.Builder
	b.WriteString(userId)
	b.WriteString(":sessionList")
	return b.String()
}

func (l *GatePushMsgLogic) msgAddSync(msg *pb.Message) error {
	// 开始事务
	tx := l.svcCtx.Db.Begin()
	switch msg.SessionType {
	case pb.SessionType_c2c:
		msgStorage1 := model.MsgSync{SessionId: msg.SessionId, UserId: msg.SenderId, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), SenderId: msg.SenderId,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&msgStorage1).Error; err != nil {
			// 返回任何错误都会回滚事务
			// 遇到错误时回滚事务
			tx.Rollback()
			return err
		}
		msgStorage2 := model.MsgSync{SessionId: msg.SenderId, UserId: msg.SessionId, Type: int32(msg.Type),
			SessionType: int32(msg.SessionType), Content: msg.Content, CloudCustomData: msg.CloudCustomData,
			MsgId: msg.MsgID, MsgUid: msg.MsgUid, Status: int32(msg.Status), SenderId: msg.SenderId,
			Timestamp: msg.Timestamp, FaceURL: msg.FaceURL, NickName: msg.NickName,
		}
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&msgStorage2).Error; err != nil {
			// 返回任何错误都会回滚事务
			// 遇到错误时回滚事务
			tx.Rollback()
			return err
		}
	case pb.SessionType_group:
		//循序查询群成员，
		// l.svcCtx.Cache.Get()
	}
	// 否则，提交事务
	tx.Commit()
	return nil
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
			logx.Error("PackType_loginAck,proto编码失败")
		}
		l.pushGata(cid, body)
	}

}
