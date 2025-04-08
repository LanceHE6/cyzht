package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/upload"
	"server/internal/middleware"
)

// RegisterUploadRouter
//
//	@Description: 全局上传路由组
//	@param group *gin.RouterGroup 路由组
func RegisterUploadRouter(group *gin.RouterGroup,
	uploadHandler upload.HandlerInterface,
) {
	routerGroup := group.Group("/upload")
	routerGroup.POST("/image", middleware.Auth(), uploadHandler.UploadImage)
	routerGroup.GET("/image/:id", middleware.Auth(), uploadHandler.GetImageUrl)
}
