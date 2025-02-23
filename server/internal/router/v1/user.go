package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler/user"
	"server/internal/middleware"
)

// RegisterUserRouter
//
//	@Description: 用户路由组
//	@param group *gin.RouterGroup 路由组
func RegisterUserRouter(group *gin.RouterGroup, userHandler user.HandlerInterface) {
	routerGroup := group.Group("/user")
	routerGroup.POST("/login", userHandler.Login)
	routerGroup.POST("/register&login_send_code", userHandler.RegisterAndLoginSendCode)
	routerGroup.POST("/register&login_verify_code", userHandler.RegisterAndLoginVerifyCode)
	routerGroup.GET("/ws/online", userHandler.OnlineHeartbeat)
	routerGroup.PUT("/update/psw", middleware.Auth(), userHandler.UpdatePassword)
	routerGroup.PUT("/update/avatar", middleware.Auth(), userHandler.UpdateAvatar)
	routerGroup.PUT("/update/profile", middleware.Auth(), userHandler.UpdateProfile)
	routerGroup.GET("/info", middleware.Auth(), userHandler.GetUserInfo)
}
