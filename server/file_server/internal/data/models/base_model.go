package models

import (
	"time"
)

type BaseModel struct {
	ID        int64     `gorm:"primary_key;" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
}
