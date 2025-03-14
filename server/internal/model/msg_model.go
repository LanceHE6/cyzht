package model

type MsgType uint8

const (
	TEXT MsgType = iota + 1
	File
)

// MetaMsg 消息元数据
type MetaMsg struct {
	UserID   int64   `gorm:"column:user_id;type:bigint;not null" json:"user_id,string"`
	SendTo   int64   `gorm:"column:send_to;type:bigint;not null" json:"send_to,string"`
	MsgType  MsgType `gorm:"column:msg_type;type:tinyint;not null" json:"msg_type"`
	TextMsg  string  `gorm:"column:text_msg;type:text;not null" json:"text_msg"`
	FileURL  string  `gorm:"column:file_url;type:varchar(255);not null" json:"file_url"`
	FileSize int64   `gorm:"column:file_size;type:bigint;not null" json:"file_size,string"`
}

// MsgModel 聊天消息表
type MsgModel struct {
	BaseModel
	ActivityID  int64 `gorm:"column:activity_id;type:bigint;not null" json:"activity_id"`
	ExhibitorID int64 `gorm:"column:exhibitor_id;type:bigint;not null" json:"exhibitor_id"` // 展商ID, 若为0，则表示是活动大厅聊天消息
	MetaMsg
}

func (*MsgModel) TableName() string {
	return "chat_message"
}
