package model

// ActivityUserModel 用户参与活动绑定模型
type ActivityUserModel struct {
	BaseModel
	ActivityID int64         `gorm:"column:activity_id;type:bigint;not null" json:"activity_id"`
	Activity   ActivityModel `gorm:"foreignKey:ActivityID;references:ID" json:"-"`
	UserID     int64         `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	User       UserModel     `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

func (*ActivityUserModel) TableName() string {
	return "activity_user"
}
