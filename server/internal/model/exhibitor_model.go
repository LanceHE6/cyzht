package model

// ExhibitorModel 参展商数据库模型
type ExhibitorModel struct {
	BaseModel
	Name       string        `gorm:"column:name;type:varchar(255);not null" json:"name"`                // 展商名称
	Introduce  string        `gorm:"column:introduce;type:varchar(255);not null" json:"introduce"`      // 展商介绍
	CreatorID  int64         `gorm:"column:creator_id;type:bigint;not null" json:"creator_id,string"`   // 创建者ID
	Creator    UserModel     `gorm:"foreignKey:CreatorID;references:ID" json:"-"`                       // 创建者
	ActivityID int64         `gorm:"column:activity_id;type:bigint;not null" json:"activity_id,string"` // 所属活动ID
	Activity   ActivityModel `gorm:"foreignKey:ActivityID;references:ID" json:"-"`                      // 所属活动
}

func (*ExhibitorModel) TableName() string {
	return "exhibitor"
}
