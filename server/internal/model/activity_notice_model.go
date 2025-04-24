package model

type ActivityNoticeModel struct {
	BaseModel
	Activity   *ActivityModel `gorm:"foreignKey:ActivityID;references:ID" json:"activity"`
	ActivityID int64          `gorm:"column:activity_id" json:"-"`
	Title      string         `gorm:"column:title" json:"title"`
	Content    string         `gorm:"column:content" json:"content"`
	Creator    *UserModel     `gorm:"foreignKey:CreatorID;references:ID" json:"creator"`
	CreatorID  int64          `gorm:"column:creator_id;type:bigint;not null" json:"-"`
}

func (*ActivityNoticeModel) TableName() string {
	return "activity_notice"
}
