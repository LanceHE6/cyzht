package database

import (
	"file_server/internal/data/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite3(dbPath string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect to SQLite database successfully")
	fmt.Println("AutoMigrate...")
	err = db.AutoMigrate(
		&models.FileModel{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	fmt.Println("AutoMigrate successfully")
	return db
}
