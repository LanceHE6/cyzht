package mysql

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"server/internal/config"
	"server/internal/model"
	"server/pkg/logger"
	"server/pkg/snowflake"
)

var db *gorm.DB

// InitMySQLConn
//
//	@Description: 初始化数据库连接
func InitMySQLConn(c *config.Config) *gorm.DB {
	var err error
	account := c.DataBase.MySQL.UserName
	password := c.DataBase.MySQL.Password
	host := c.DataBase.MySQL.Host
	port := c.DataBase.MySQL.Port
	dbname := c.DataBase.MySQL.Database

	// 创建MySQL连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		account,
		password,
		host,
		port,
	)

	// 连接到MySQL
	tdb, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Logger.Error("can not connect to database: " + err.Error())
		os.Exit(-1)
		return nil
	}

	// 创建数据库
	_, err = tdb.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		logger.Logger.Error("can not create database: " + err.Error())
		os.Exit(-2)
		return nil
	}
	// 关闭数据库连接
	err = tdb.Close()
	if err != nil {
		logger.Logger.Error("can not close database: " + err.Error())
		os.Exit(-3)
		return nil
	}

	// 创建MySQL连接字符串

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		account,
		password,
		host,
		port,
		dbname,
	)

	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		logger.Logger.Error("Cannot connect to MYSQL database: " + err.Error())
		os.Exit(-4)
	}
	logger.Logger.Info("Connect to MYSQL database successfully")
	// 初始化雪花算法生成worker
	worker, err := snowflake.NewWorker(1)
	if err != nil {
		logger.Logger.Error("Snow Flake NewWorker error: " + err.Error())
		return nil
	}
	// 设置雪花算法生成id
	db.Callback().Create().Before("gorm:create").Replace("id", func(scope *gorm.Scope) {
		id := worker.NextId()
		err := scope.SetColumn("id", id)
		if err != nil {
			logger.Logger.Error("Cannot set snowflake id: " + err.Error())
			return
		}
	})

	// 自动建表
	model.CreateTable(db)

	// 初始化数据
	logger.Logger.Info("Init data...")
	InitData()
	logger.Logger.Info("Init data done.")

	return db
}

// CloseMyDb
//
//	@Description: 关闭数据库连接
func CloseMyDb() {
	if db != nil {
		if err := db.Close(); err != nil {
			logger.Logger.Error("Close Db error: " + err.Error())
		}
	}
}

// GetMySQLConnection
//
//	@Description: 获取数据库连接
//	@return *gorm.db 数据库连接
func GetMySQLConnection() *gorm.DB {
	return db
}
