package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type RpcServerConfig struct {
	zrpc.RpcServerConf
	DBPath      string `json:"DBPath"`
	StoragePath string `json:"StoragePath"`
}

type FileServerConfig struct {
	rest.RestConf
}
