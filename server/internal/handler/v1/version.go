package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/service/version"
)

// VersionRoute
//
//	@Description: 版本路由
//	@param group *gin.RouterGroup
func VersionRoute(group *gin.RouterGroup) {
	versionService := version.NewVersionService()

	group.GET("/ver", func(context *gin.Context) {
		versionService.GetVersion()
	})
	group.GET("/ping", func(context *gin.Context) {
		versionService.GetVersion()
	})
}
