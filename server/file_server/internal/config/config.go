package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type RpcServerConfig struct {
	zrpc.RpcServerConf
	DBPath         string `json:"DBPath"`
	StoragePath    string `json:",default=./storage"` // 文件存储路径
	ThumbMaxWidth  int    `json:",default=200"`       // 默认缩略图最大宽度
	ThumbMaxHeight int    `json:",default=200"`       // 默认缩略图最大高度
	MaxFileSize    int64  `json:",default=10485760"`  // 最大文件大小(10MB)
}

type FileServerConfig struct {
	rest.RestConf
}
