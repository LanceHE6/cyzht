package service

import (
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"server/internal/config"
	"server/internal/data/repo"
	"server/internal/service/user"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

type Service struct {
	UserService      user.UserServiceInterface
	FileServerClient file_server.FileServiceClient
}

func InitService(c *config.Config, repo *repo.Repo) *Service {
	log.Println("connecting file server: ", c.Server.FileServer.RpcDNS)
	fileServerConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: c.Server.FileServer.RpcDNS,
	})

	fileServer := file_server.NewFileServiceClient(fileServerConn.Conn())
	log.Println("connect file server success")
	return &Service{
		UserService: user.NewUserService(c, repo.UserRepo, repo.VerifyCodeRepo, fileServer),
	}
}
