package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/hash"
	"server/pkg/response"
)

// UpdatePassword
//
//	@Description: 修改密码
//	@receiver s userHandler
//	@param id 用户id
//	@param oldPsw 原密码
//	@param newPsw 新密码
//	@return *pkg.Response 返回结果
func (s userHandler) UpdatePassword() gin.HandlerFunc {
	return func(context *gin.Context) {
		// updatePasswordRequest
		// @Description: 修改密码请求参数结构体
		type updatePasswordRequest struct {
			OldPassword string `json:"old_password" form:"old_password" binding:"required"`
			NewPassword string `json:"new_password" form:"new_password" binding:"required"`
		}
		var data updatePasswordRequest
		if err := context.ShouldBind(&data); err != nil {
			context.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
			return
		}
		claims, err := GetUserInfoByContext(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.FailedResponse(10, err.Error()))
		}
		user := s.UserRepo.SelectByID(claims.ID)
		if user == nil {
			context.JSON(http.StatusBadRequest, response.FailedResponse(11, "用户不存在"))
			return
		}

		// 修改自己密码
		if !hash.CheckPsw(user.Password, data.OldPassword) {
			context.JSON(http.StatusOK, response.FailedResponse(1, "原密码错误"))
			return
		}
		err = s.UserRepo.UpdatePassword(claims.ID, data.NewPassword)
		if err != nil {
			context.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "修改密码失败", err))
			return
		} else {
			context.JSON(http.StatusOK, response.SuccessResponse(nil))
			return
		}
	}
}
