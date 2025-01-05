package main

import (
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/internal/db"
	"server/internal/handler"
	"server/internal/middleware"
	"server/internal/repo"
	"server/internal/service"
	"server/pkg/smtp"
)

func main() {
	ginServer := gin.Default()
	// 跨域
	ginServer.Use(middleware.Cors())
	ginServer.Use(middleware.LoggerToFile())

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
		c.Server.SMTP.Host,
		c.Server.SMTP.Port,
		c.Server.SMTP.Account,
		c.Server.SMTP.Password,
	)
	// 加载路由
	handler.Route(ginServer, svc)

	err := ginServer.Run(":" + config.GetServerPort())
	if err != nil {
		return
	}
}
