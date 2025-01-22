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
func RegisterUserRouter(group *gin.RouterGroup, userHandler user.UserHandlerInterface) {
	userGroup := group.Group("/user")
	userGroup.POST("/login", userHandler.Login())
	userGroup.POST("/register&login_send_code", userHandler.RegisterAndLoginSendCode())
	userGroup.POST("/register&login_verify_code", userHandler.RegisterAndLoginVerifyCode())
	userGroup.GET("/ws/online", userHandler.OnlineHeartbeat())
	userGroup.PUT("/update/psw", middleware.Auth(), userHandler.UpdatePassword())
	userGroup.PUT("/update/avatar", middleware.Auth(), userHandler.UpdateAvatar())
	userGroup.PUT("/update/profile", middleware.Auth(), userHandler.UpdateProfile())
	userGroup.GET("/info", middleware.Auth(), userHandler.GetUserInfo())
}
