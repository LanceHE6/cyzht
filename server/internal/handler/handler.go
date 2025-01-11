package handler

import (
	"github.com/zeromicro/go-zero/zrpc"
	"server/internal/config"
	"server/internal/handler/user"
	"server/internal/handler/version"
	"server/internal/repo"
	"server/pkg/logger"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

type Handler struct {
	VersionHandler   version.VersionHandlerInterface
	UserHandler      user.UserHandlerInterface
	FileServerClient file_server.FileServiceClient
}

// InitHandler 初始化handler
func InitHandler(c *config.Config, repo *repo.Repo) *Handler {
	// 连接文件服务器
	logger.Logger.Infof("connecting file server: %s", c.Server.FileServer.RpcDNS)
	fileServerConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: c.Server.FileServer.RpcDNS,
	})

	fileServer := file_server.NewFileServiceClient(fileServerConn.Conn())
	logger.Logger.Info("connect file server success")

	return &Handler{
		VersionHandler:   version.NewVersionHandler(),
		UserHandler:      user.NewUserHandler(c, repo.UserRepo, repo.VerifyCodeRepo, fileServer),
		FileServerClient: fileServer,
	}
}
