package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"server/internal/model"
	"server/pkg/hash"
)

// InitData
//
//	@Description: 初始化数据
func InitData() {
	// 初始化用户数据
	createIfNotExists(db, &model.UserModel{
		Account:  "admin",
		Password: hash.HashPsw("123456"),
		Nickname: "admin",
	}, "admin", "account")
}

// createIfNotExists
//
//	@Description: 创建数据，如果不存在
//	@param db 数据库
//	@param value 数据
//	@param id id
//	@param idFieldName id字段名
func createIfNotExists(db *gorm.DB, value interface{}, id interface{}, idFieldName string) {
	// 检查数据是否存在
	if err := db.Where(idFieldName+" = ?", id).First(value).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 数据不存在，插入新数据
			if err = db.Create(value).Error; err != nil {
				// 插入错误
				log.Println(err)
				os.Exit(-100)
			}
		} else {
			// 查询错误
			log.Println(err)
			os.Exit(-200)
		}
	}
}
