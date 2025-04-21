package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/exhibitor"
	"server/internal/middleware"
)

// RegisterExhibitorRouter
//
//	@Description: 参展商路由组
//	@param group *gin.RouterGroup 路由组
func RegisterExhibitorRouter(group *gin.RouterGroup,
	exhibitorHandler exhibitor.HandlerInterface,
) {
	routerGroup := group.Group("/exhibitor")
	routerGroup.POST("/add", middleware.Auth(), exhibitorHandler.AddExhibitor)
	routerGroup.POST("/del", middleware.Auth(), exhibitorHandler.DeleteExhibitor)
	routerGroup.POST("/:eid/join", middleware.Auth(), exhibitorHandler.JoinExhibitor)
	routerGroup.POST("/:eid/withdraw", middleware.Auth(), exhibitorHandler.WithdrawExhibitor)
	routerGroup.GET("/info/:eid", middleware.Auth(), exhibitorHandler.GetExhibitorInfo)
	routerGroup.GET("/:aid/list", middleware.Auth(), exhibitorHandler.ListExhibitors)
	routerGroup.GET("/:aid/joined", middleware.Auth(), exhibitorHandler.GetJoinedExhibitors)
	routerGroup.GET("/joined/users/:eid", middleware.Auth(), exhibitorHandler.GetJoinedUsers)
}
