package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/chat"
	"server/internal/middleware"
)

// RegisterChatRouter
//
//	@Description: 聊天路由组
//	@param group *gin.RouterGroup 路由组
func RegisterChatRouter(group *gin.RouterGroup,
	chatHandler chat.HandlerInterface,
) {
	routerGroup := group.Group("/chat")
	routerGroup.POST("/:aid/:eid/send", middleware.Auth(), chatHandler.SendMsg)
	routerGroup.GET("/:aid/:eid/msg", middleware.Auth(), chatHandler.GetActivityMsg)
}
