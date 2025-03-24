package model

type MsgType uint8

const (
	TEXT MsgType = iota + 1
	File
)

// MetaMsg 消息元数据
type MetaMsg struct {
	FromUID  int64      `gorm:"column:from_uid;type:bigint;not null" json:"-,string"`
	FromUser *UserModel `gorm:"foreignKey:FromUID;references:ID" json:"from_user"`
	ToUID    int64      `gorm:"column:to_uid;type:bigint;not null" json:"-,string"`
	ToUser   *UserModel `gorm:"foreignKey:ToUID;references:ID" json:"to_user"`
	MsgType  MsgType    `gorm:"column:msg_type;type:tinyint;not null" json:"msg_type"`
	TextMsg  string     `gorm:"column:text_msg;type:text;not null" json:"text_msg"`
	FileURL  string     `gorm:"column:file_url;type:varchar(255);not null" json:"file_url"`
	FileSize int64      `gorm:"column:file_size;type:bigint;not null" json:"file_size,string"`
}

// MsgModel 聊天消息表
type MsgModel struct {
	BaseModel
	ActivityID  int64           `gorm:"column:activity_id;type:bigint;not null" json:"-"`
	Activity    *ActivityModel  `gorm:"foreignKey:ActivityID;references:ID" json:"activity"`
	ExhibitorID int64           `gorm:"column:exhibitor_id;type:bigint;not null" json:"-"` // 展商ID, 若为0，则表示是活动大厅聊天消息
	Exhibitor   *ExhibitorModel `gorm:"foreignKey:ExhibitorID;references:ID" json:"exhibitor"`
	MetaMsg
}

func (*MsgModel) TableName() string {
	return "chat_message"
}
