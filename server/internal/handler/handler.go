package handler

import (
	"github.com/zeromicro/go-zero/zrpc"
	"server/internal/config"
	"server/internal/handler/user"
	"server/internal/repo"
	"server/pkg/logger"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

type Handler struct {
	UserHandler      user.UserHandlerInterface
	FileServerClient file_server.FileServiceClient
}

func InitHandler(c *config.Config, repo *repo.Repo) *Handler {
	logger.Logger.Infof("connecting file server: %s", c.Server.FileServer.RpcDNS)
	fileServerConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: c.Server.FileServer.RpcDNS,
	})

	fileServer := file_server.NewFileServiceClient(fileServerConn.Conn())
	logger.Logger.Info("connect file server success")
	return &Handler{
		UserHandler: user.NewUserHandler(c, repo.UserRepo, repo.VerifyCodeRepo, fileServer),
	}
}
