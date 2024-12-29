package v1

import (
	"github.com/gin-gonic/gin"
	"server/internal/service/user"
)

// UserRoute
//
//	@Description: 用户路由组
//	@param group *gin.RouterGroup 路由组
func UserRoute(group *gin.RouterGroup, service user.UserServiceInterface) {
	userGroup := group.Group("/user")
	userGroup.POST("/login", service.Login())
	userGroup.POST("/register&login_send_code", service.RegisterAndLoginSendCode())
	userGroup.POST("/register&login_verify_code", service.RegisterAndLoginVerifyCode())
	//userGroup.GET("/ws/online", service.UserService.OnlineHeartbeat())
	userGroup.PUT("/update/psw", service.AuthMiddleware(), service.UpdatePassword())
	userGroup.PUT("/update/avatar", service.AuthMiddleware(), service.UpdateAvatar())
	userGroup.PUT("/update/profile", service.AuthMiddleware(), service.UpdateProfile())
	userGroup.GET("/info", service.AuthMiddleware(), service.GetUserInfo())
}
