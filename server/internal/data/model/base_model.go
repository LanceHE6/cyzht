package model

import "time"

// BaseModel 基础模型，所有模型都继承该模型
type BaseModel struct {
	ID        int64     `gorm:"primary_key;" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
