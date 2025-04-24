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
func RegisterActivityRouter(group *gin.RouterGroup,
	activityHandler activity.HandlerInterface,
) {
	routerGroup := group.Group("/activity")
	routerGroup.POST("/add", middleware.Auth(), activityHandler.AddActivity)
	routerGroup.POST("/del", middleware.Auth(), activityHandler.DeleteActivity)
	routerGroup.GET("/search", activityHandler.SearchActivity)
	routerGroup.POST("/:aid/join", middleware.Auth(), activityHandler.JoinActivity)
	routerGroup.POST("/:aid/withdraw", middleware.Auth(), activityHandler.WithdrawActivity)
	routerGroup.GET("/joined", middleware.Auth(), activityHandler.GetJoinedActivity)
	routerGroup.GET("/:aid/joined", middleware.Auth(), activityHandler.GetJoinedUser)
	routerGroup.POST("/:aid/notice/publish", middleware.Auth(), activityHandler.PublishNotice)
	routerGroup.PUT("/notice/:nid/update", middleware.Auth(), activityHandler.UpdateNotice)
	routerGroup.GET("/:aid/notices", middleware.Auth(), activityHandler.GetNotices)
	routerGroup.DELETE("/notice/:nid/del", middleware.Auth(), activityHandler.DeleteNotice)
}
