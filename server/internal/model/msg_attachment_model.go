package model

import (
	"github.com/jinzhu/gorm"
	"server/internal/config"
)

// Attachment 消息附件
type Attachment struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id,string"`
	MessageID int64  `gorm:"column:message_id;type:bigint;not null" json:"message_id,string"`
	Type      uint8  `gorm:"column:type;type:tinyint;not null" json:"type"` // 1-图片, 2-文件等
	FileURL   string `gorm:"column:file_url;type:varchar(512);not null" json:"file_url"`
	ThumbURL  string `gorm:"column:thumb_url;type:varchar(512)" json:"thumb_url"` // 缩略图URL
	FileName  string `gorm:"column:file_name;type:varchar(255)" json:"file_name"`
	FileSize  int64  `gorm:"column:file_size;type:bigint" json:"file_size,string"`
	Width     int    `gorm:"column:width;type:int" json:"width"`                     // 图片宽度
	Height    int    `gorm:"column:height;type:int" json:"height"`                   // 图片高度
	Duration  int    `gorm:"column:duration;type:int" json:"duration"`               // 音视频时长(秒)
	ExtraInfo string `gorm:"column:extra_info;type:varchar(512)" json:"extra_info"`  // 额外信息
	SortOrder int    `gorm:"column:sort_order;type:int;default:0" json:"sort_order"` // 排序
}

func (*Attachment) TableName() string {
	return "message_attachment"
}

func (a *Attachment) AfterFind(*gorm.DB) (err error) {
	if a.FileURL != "" {
		a.FileURL = config.GetConfig().Server.FileServer.StaticURL + a.FileURL
	}
	if a.ThumbURL != "" {
		a.ThumbURL = config.GetConfig().Server.FileServer.StaticURL + a.ThumbURL
	}
	return nil
}
