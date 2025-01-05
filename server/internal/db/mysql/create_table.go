package mysql

import (
	"github.com/jinzhu/gorm"
	"server/internal/model"
)

// CreateTable
//
//	@Description: 创建表
//	@param db *gorm.db 数据库连接
func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&model.UserModel{})
}
