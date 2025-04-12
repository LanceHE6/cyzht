package model

type MsgType uint8

const (
	TEXT      MsgType = iota + 1 // 纯文本消息
	IMAGE                        // 纯图片消息
	COMPOSITE                    // 复合消息(文字+图片)
	// 可以继续扩展其他类型
)

// MetaMsg 消息元数据
type MetaMsg struct {
	FromUID     int64        `gorm:"column:from_uid;type:bigint;not null" json:"-,string"`
	FromUser    *UserModel   `gorm:"foreignKey:FromUID;references:ID" json:"from_user"`
	ToUID       int64        `gorm:"column:to_uid;type:bigint;not null" json:"-,string"` // 单聊时使用
	ToUser      *UserModel   `gorm:"foreignKey:ToUID;references:ID" json:"to_user"`
	MsgType     MsgType      `gorm:"column:msg_type;type:tinyint;not null" json:"msg_type"`
	Content     string       `gorm:"column:content;type:text" json:"content"` // 文本内容
	Attachments []Attachment `gorm:"foreignKey:MessageID" json:"attachments"` // 消息附件
}

// MsgModel 聊天消息表
type MsgModel struct {
	BaseModel
	ActivityID  int64           `gorm:"column:activity_id;type:bigint;not null" json:"-"`
	Activity    *ActivityModel  `gorm:"foreignKey:ActivityID;references:ID" json:"activity"`
	ExhibitorID int64           `gorm:"column:exhibitor_id;type:bigint;not null" json:"-"` // 展商ID, 若为0，则表示是活动大厅聊天消息
	Exhibitor   *ExhibitorModel `gorm:"foreignKey:ExhibitorID;references:ID" json:"exhibitor"`
	MetaMsg
	Status uint8 `gorm:"column:status;type:tinyint;default:1" json:"status"` // 消息状态:1-正常,2-已撤回等
}

func (*MsgModel) TableName() string {
	return "chat_message"
}
