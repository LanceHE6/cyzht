package version

import (
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	GetVersion(ctx *gin.Context) // 获取版本号
}

type versionHandler struct {
}

func NewVersionHandler() HandlerInterface {
	return &versionHandler{}
}
