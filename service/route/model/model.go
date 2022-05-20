package model

// type Conversation struct {
// 	//是否置顶
// 	Pinned bool `gorm:"default:false"`
// 	//复合主键
// 	SessionId string `gorm:"primaryKey"`
// 	UserId    string `gorm:"primaryKey"`
// 	//会话类型，
// 	SessionType int32 `gorm:"index;type:smallint;not null"`
// 	//拉取的最新消息时间戳
// 	AckMsgTime int64 `gorm:"default:0"`
// 	//未读消息数，实时计算
// 	UnreadCount int64 `gorm:"default:0"`
// 	//最新消息时间戳,用户端排序用
// 	Timestamp int64 `gorm:"default:0"`
// }

const (
	ConversationType_c2c = iota + 1
	ConversationType_group
)

//消息发送记录表，群聊读扩散;单聊写扩散
type MsgStorage struct {
	Sender   string `gorm:"not null"`
	FaceUrl  string
	UserId   string `gorm:"index;not null"`
	NickName string
	//会话id/接收id
	ConversationId   string `gorm:"index;not null"`
	ConversationType int32  `gorm:"index;type:smallint;not null"`
	MsgId            string `gorm:"not null"`
	MsgUid           string `gorm:"primaryKey"`
	Status           int32  `gorm:"type:smallint;not null"`
	Type             int32  `gorm:"type:smallint;not null"`
	Content          string
	CloudCustomData  []byte
	//索引，降序
	Timestamp int64 `gorm:"index,sort:desc"`
}

//消息同步表,为写扩散
type MsgSync struct {
	Sender   string `gorm:"not null"`
	FaceURL  string
	NickName string
	UserId   string `gorm:"not null"`
	//会话id/接收id
	ConversationId   string `gorm:"index;not null"`
	ConversationType int32  `gorm:"index;type:smallint;not null"`
	MsgId            string `gorm:"not null"`
	MsgUid           string `gorm:"primaryKey"`
	Status           int32  `gorm:"type:smallint;not null"`
	Type             int32  `gorm:"type:smallint;not null"`
	Content          string
	CloudCustomData  []byte
	//索引，降序
	Timestamp int64 `gorm:"index,sort:desc"`
}
