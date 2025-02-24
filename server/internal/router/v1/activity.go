package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/activity"
	"server/internal/handler/chat"
	"server/internal/middleware"
)

// RegisterActivityRouter
//
//	@Description: 活动路由组
//	@param group *gin.RouterGroup 路由组
func RegisterActivityRouter(group *gin.RouterGroup,
	activityHandler activity.HandlerInterface,
	chatHandler chat.HandlerInterface,
) {
	routerGroup := group.Group("/activity")
	routerGroup.POST("/add", middleware.Auth(), activityHandler.AddActivity)
	routerGroup.DELETE("/del", middleware.Auth(), activityHandler.DeleteActivity)
	routerGroup.GET("/search", activityHandler.SearchActivity)
	routerGroup.POST("/:aid/send", middleware.Auth(), chatHandler.SendToActivity)
	routerGroup.GET("/:aid/msg", middleware.Auth(), chatHandler.GetActivityMsg)
	routerGroup.POST("/:aid/join", middleware.Auth(), activityHandler.JoinActivity)
	routerGroup.POST("/:aid/exit", middleware.Auth(), activityHandler.ExitActivity)
}
