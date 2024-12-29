package database

import (
	"file_server/internal/data/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite3(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.UserAvatarModel{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}
