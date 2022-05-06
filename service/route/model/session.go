package model

import (
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	//是否有同步完的离线消息
	OfflineMsg bool `gorm:"default:0"`
	//是否置顶
	Pinned bool `gorm:"default:0"`
	//联合索引
	SessionId string
	UserId    string `gorm:"index:idx_session;not null"`
	//唯一索引，降序
	Timestamp int64 `gorm:"index:idx_session;not null"`
	//未读消息数，实时计算
	// UnreadCount int64 `gorm:"default:0"`
	//最新的20条消息
	LatestMsg []Message `gorm:"-"`
}

type SessionModel struct {
	Db *gorm.DB
}

type Message struct {
	Content      string
	ExpansionDic string
	MsgId        string
	MsgUid       string `gorm:"uniqueIndex"`
	MsgType      int32  `gorm:"type:smallint;not null"`
	SessionType  int32  `gorm:"type:smallint;not null"`
	Status       int32  `gorm:"type:smallint;not null"`
	//联合索引
	SendId   string `gorm:"index:idx_msg;not null"`
	TargetId string `gorm:"index:idx_msg;not null"`
	//唯一索引，降序
	Timestamp int64 `gorm:"uniqueIndex:idx_time,sort:desc"`
}
