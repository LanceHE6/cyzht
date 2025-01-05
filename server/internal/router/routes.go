package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler"
	"server/internal/router/v1"
)

// InitRouter
//
//	@Description: 总路由
//	@param gin *gin.Engine gin框架
func InitRouter(gin *gin.Engine, handler *handler.Handler) {
	api := gin.Group("/api")
	apiV1 := api.Group("/v1")

	v1.RegisterUserRouter(apiV1, handler.UserHandler) // 用户路由组
	v1.RegisterVersionRouter(apiV1)
}
