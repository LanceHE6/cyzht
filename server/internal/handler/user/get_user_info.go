package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/jwt"
	"server/pkg/response"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

// GetUserInfo
//
//	@Description: 获取用户信息
//	@receiver s userHandler
//	@return gin.HandlerFunc 返回用户信息
func (s userHandler) GetUserInfo(ctx *gin.Context) {
	claims, _ := jwt.GetClaimsByContext(ctx)
	user := s.UserRepo.SelectByID(claims.ID)
	// 获取文件服务器中的头像地址
	rsp, err := s.FileRpcServer.GetAvatarUrl(context.Background(), &file_server.GetAvatarUrlRequest{
		Id: user.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, response.ErrorResponse(-1, "获取用户头像失败", err))
		return
	}
	// 如果文件服务器和本地数据库的头像地址不一致，更新本地数据库数据库
	if rsp.FileUrl != user.Avatar {
		user.Avatar = rsp.FileUrl
		err = s.UserRepo.UpdateAvatar(user.ID, rsp.FileUrl)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-2, "更新头像失败", err))
			return
		}
	}
	// 拼接头像地址
	if user.Avatar != "" {
		user.Avatar = s.C.Server.FileServer.StaticURL + user.Avatar
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(user))
}
