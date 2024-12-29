package version

import (
	"github.com/gin-gonic/gin"
)

type VersionService interface {
	GetVersion() gin.HandlerFunc // 获取版本号
}

type VersionServiceImpl struct {
}

func NewVersionService() VersionService {
	return &VersionServiceImpl{}
}
