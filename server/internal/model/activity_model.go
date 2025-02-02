package model

import "time"

// ActivityModel
//
//	@Description: 活动建表模型
type ActivityModel struct {
	BaseModel
	Name      string    `gorm:"column:name;type:varchar(255);not null;" json:"name"`           // 活动名
	Introduce string    `gorm:"column:introduce;type:varchar(255);not null" json:"introduce"`  // 活动介绍
	StartAt   time.Time `gorm:"column:start_at;type:datetime;not null" json:"start_at"`        // 活动开始时间
	EndAt     time.Time `gorm:"column:end_at;type:datetime;not null" json:"end_at"`            // 活动结束时间
	Location  string    `gorm:"column:location;type:varchar(255);not null" json:"location"`    // 活动地点
	CreatorID int64     `gorm:"column:creator_id;type:bigint;not null" json:"creator_id"`      // 活动发起人ID
	Creator   UserModel `gorm:"foreignKey:CreatorID;references:ID" json:"creator"`             // 活动发起人
	IsDeleted bool      `gorm:"column:is_deleted;type:tinyint(1);default 0" json:"is_deleted"` // 是否被标记删除
}

func (*ActivityModel) TableName() string {
	return "activity"
}
