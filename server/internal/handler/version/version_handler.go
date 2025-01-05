package version

import (
	"github.com/gin-gonic/gin"
)

type VersionHandlerInterface interface {
	GetVersion() gin.HandlerFunc // 获取版本号
}

type versionHandler struct {
}

func NewVersionHandler() VersionHandlerInterface {
	return &versionHandler{}
}
