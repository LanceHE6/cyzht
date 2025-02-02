package model

import (
	"github.com/jinzhu/gorm"
)

// CreateTable
//
//	@Description: 创建表
//	@param db *gorm.db 数据库连接
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&ActivityModel{})
	db.AutoMigrate(&MsgModel{})
	db.AutoMigrate(&ExhibitorModel{})
}
