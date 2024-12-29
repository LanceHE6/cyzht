package main

import (
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/internal/data/db"
	"server/internal/data/repo"
	"server/internal/handler"
	"server/internal/middleware"
	"server/internal/service"
	"server/pkg/smtp"
)

func main() {
	ginServer := gin.Default()
	// 跨域
	ginServer.Use(middleware.Cors())

	// 加载配置
	c := config.LoadConfig()
	// 初始化数据库连接
	dbConn := db.InitDBConn(c)
	// 初始化repo
	repository := repo.InitRepo(dbConn)
	// 初始化服务类
	svc := service.InitService(c, repository)
	// 初始化smtp服务
	smtp.InitSMTPService(
		config.ConfigData.Server.SMTP.Host,
		config.ConfigData.Server.SMTP.Port,
		config.ConfigData.Server.SMTP.Account,
		config.ConfigData.Server.SMTP.Password,
	)
	// 加载路由
	handler.Route(ginServer, svc)

	err := ginServer.Run(":" + config.GetServerPort())
	if err != nil {
		return
	}
}
