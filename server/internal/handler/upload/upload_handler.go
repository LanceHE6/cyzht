package upload

import (
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

type HandlerInterface interface {
	UploadFile(ctx *gin.Context)
	GetFileUrl(ctx *gin.Context)
	UploadImage(ctx *gin.Context)
	GetImageUrl(ctx *gin.Context)
}

type uploadHandler struct {
	C             *config.Config
	FileRpcServer file_server.FileServiceClient
}

func NewUploadHandler(c *config.Config, fileRpcServer file_server.FileServiceClient) HandlerInterface {
	return &uploadHandler{
		C:             c,
		FileRpcServer: fileRpcServer,
	}
}
