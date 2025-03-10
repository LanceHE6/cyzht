package main

import (
	"file_server/api/v1/file_server"
	"file_server/internal/config"
	"file_server/internal/server"
	"file_server/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var rpcConfigFIle = flag.String("rf", "etc/rpc_server.yaml", "the rpc config file")
var fileServerConfigFile = flag.String("ff", "etc/file_server.yaml", "the file server config file")

func main() {
	flag.Parse()

	var c1 config.RpcServerConfig
	var c2 config.FileServerConfig
	conf.MustLoad(*rpcConfigFIle, &c1)
	conf.MustLoad(*fileServerConfigFile, &c2)

	// 文件存储rpc服务
	ctx := svc.NewServiceContext(c1)

	s := zrpc.MustNewServer(c1.RpcServerConf, func(grpcServer *grpc.Server) {
		file_server.RegisterFileServiceServer(grpcServer, server.NewFileServiceServer(ctx))

		if c1.Mode == service.DevMode || c1.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	go func() {
		// 文件访问服务
		//静态文件目录
		fs := http.Dir("./storage")
		//访问url，如"/static/",注意末尾的"/"不能少。
		staticServer := rest.MustNewServer(c2.RestConf, rest.WithFileServer("/static/", fs))
		fmt.Printf("Starting static resources staticServer at %d...\n", c2.Port)
		staticServer.Start()
	}()

	fmt.Printf("file rpc server listening on %s...\n", c1.ListenOn)
	s.Start()
}
