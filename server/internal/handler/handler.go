package handler

import (
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"server/internal/config"
	"server/internal/handler/user"
	"server/internal/repo"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

type Handler struct {
	UserHandler      user.UserHandlerInterface
	FileServerClient file_server.FileServiceClient
}

func InitHandler(c *config.Config, repo *repo.Repo) *Handler {
	log.Println("connecting file server: ", c.Server.FileServer.RpcDNS)
	fileServerConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: c.Server.FileServer.RpcDNS,
	})

	fileServer := file_server.NewFileServiceClient(fileServerConn.Conn())
	log.Println("connect file server success")
	return &Handler{
		UserHandler: user.NewUserHandler(c, repo.UserRepo, repo.VerifyCodeRepo, fileServer),
	}
}
