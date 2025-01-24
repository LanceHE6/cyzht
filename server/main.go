package main

import (
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/internal/db"
	"server/internal/handler"
	"server/internal/middleware"
	"server/internal/repo"
	"server/internal/router"
	"server/pkg/smtp"
)

func main() {
	ginServer := gin.Default()
	// 跨域
	ginServer.Use(middleware.Cors())
	ginServer.Use(middleware.LoggerToFile())

	// 获取配置
	c := config.InitConfig()
	// 初始化数据库连接
	dbConn := db.InitDBConn(c)
	// 初始化repo
	repository := repo.InitRepo(dbConn)
	// 初始化服务类
	svc := handler.InitHandler(c, repository)

	// 初始化smtp服务
	smtp.InitSMTPService(
		c.Server.SMTP.Host,
		c.Server.SMTP.Port,
		c.Server.SMTP.Account,
		c.Server.SMTP.Password,
	)
	// 加载路由
	router.InitRouter(ginServer, svc)

	err := ginServer.Run(":" + c.Server.Port)
	if err != nil {
		return
	}
}
