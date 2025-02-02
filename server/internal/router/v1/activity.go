package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/activity"
	"server/internal/middleware"
)

// RegisterActivityRouter
//
//	@Description: 活动路由组
//	@param group *gin.RouterGroup 路由组
func RegisterActivityRouter(group *gin.RouterGroup, activityHandler activity.ActivityHandlerInterface) {
	userGroup := group.Group("/activity")
	userGroup.POST("/add", middleware.Auth(), activityHandler.AddActivity)
	userGroup.DELETE("/del", middleware.Auth(), activityHandler.DeleteActivity)
	userGroup.GET("/search", activityHandler.SearchActivity)
}
