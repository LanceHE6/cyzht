package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/version"
)

// RegisterVersionRouter
//
//	@Description: 版本路由
//	@param group *gin.RouterGroup
func RegisterVersionRouter(group *gin.RouterGroup, handler version.HandlerInterface) {
	group.GET("/ver", handler.GetVersion)
	group.GET("/ping", handler.GetVersion)
}
