package model

// UserModel
//
//	@Description: 用户表结构
type UserModel struct {
	BaseModel
	Account  string `gorm:"column:account;type:varchar(255);not null;unique" json:"account"` // 账号
	Password string `gorm:"column:password;type:varchar(255);not null" json:"-"`             // 密码
	Nickname string `gorm:"column:nickname;type:varchar(255);not null" json:"nickname"`      // 姓名
	Avatar   string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`                  // mongodb头像ID
	Sex      int    `gorm:"column:sex;type:int;default -1" json:"sex"`                       // 性别0:男 1:女 -1:未知
	//Signature    string `gorm:"column:signature;type:varchar(255);" json:"signature"`            // 个性签名
	//OnlineStatus int    `gorm:"column:online_status;type:int;default '0'" json:"online_status"`  // 在线状态0为离线1为在线
	SessionID string `gorm:"column:session_id;type:varchar(255)" json:"session_id"` // session_id
}

func (*UserModel) TableName() string {
	return "user"
}
