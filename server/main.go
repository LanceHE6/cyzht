package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/zrpc"
	"server/internal/config"
	"server/internal/db"
	"server/internal/handler"
	"server/internal/middleware"
	"server/internal/repo"
	"server/internal/router"
	"server/internal/ws"
	"server/pkg/logger"
	"server/pkg/logo"
	"server/pkg/rpc/file_server/api/v1/file_server"
	"server/pkg/smtp"
)

func main() {
	logo.PrintLogo()

	ginServer := gin.Default()
	// 跨域
	ginServer.Use(middleware.Cors())
	ginServer.Use(middleware.LoggerToFile())

	// 获取配置
	c := config.InitConfig()
	// 连接文件服务器
	logger.Logger.Infof("connecting file server: %s", c.Server.FileServer.RpcDNS)
	fileServerConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: c.Server.FileServer.RpcDNS,
	})

	fileServer := file_server.NewFileServiceClient(fileServerConn.Conn())
	logger.Logger.Info("connect file server success")

	// 初始化数据库连接
	dbConn := db.InitDBConn(c)
	// 初始化repo
	repository := repo.InitRepo(c, dbConn, fileServer)
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

	// 初始化消息广播
	ws.InitBroadcast(repository.ActivityUserRepo)

	err := ginServer.Run(":" + c.Server.Port)
	if err != nil {
		return
	}
}
