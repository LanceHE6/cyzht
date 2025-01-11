package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/jwt"
	"server/pkg/response"
)

// GetUserInfo
//
//	@Description: 获取用户信息
//	@receiver s userHandler
//	@return gin.HandlerFunc 返回用户信息
func (s userHandler) GetUserInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		claims, _ := jwt.GetClaimsByContext(context)
		user := s.UserRepo.SelectByID(claims.ID)
		// 拼接头像地址
		user.Avatar = s.C.Server.FileServer.StaticURL + user.Avatar
		context.JSON(http.StatusOK, response.SuccessResponse(user))
	}
}
