package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/version"
)

// RegisterVersionRouter
//
//	@Description: 版本路由
//	@param group *gin.RouterGroup
func RegisterVersionRouter(group *gin.RouterGroup) {
	versionService := version.NewVersionHandler()

	group.GET("/ver", func(context *gin.Context) {
		versionService.GetVersion()
	})
	group.GET("/ping", func(context *gin.Context) {
		versionService.GetVersion()
	})
}
