package handler

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/v1"
	"server/internal/service"
)

// Route
//
//	@Description: 总路由
//	@param gin *gin.Engine gin框架
func Route(gin *gin.Engine, service *service.Service) {
	api := gin.Group("/api")
	apiV1 := api.Group("/v1")

	v1.UserRoute(apiV1, service.UserService) // 用户路由组
	v1.VersionRoute(apiV1)
}
