package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/response"
)

// GetUserInfo
//
//	@Description: 获取用户信息
//	@receiver s userService
//	@return gin.HandlerFunc 返回用户信息
func (s userService) GetUserInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		claims, _ := GetUserInfoByContext(context)
		user := s.UserRepo.SelectByID(claims.ID)
		// 拼接头像地址
		user.Avatar = s.C.Server.FileServer.StaticDNS + user.Avatar
		context.JSON(http.StatusOK, response.SuccessResponse(user))
	}
}
