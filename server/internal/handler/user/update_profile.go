package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/jwt"
	"server/pkg/response"
)

// UpdateProfile
//
//	@Description: 更新用户信息
//	@receiver s userHandler
//	@return gin.HandlerFunc
func (s userHandler) UpdateProfile(ctx *gin.Context) {
	type updateProfileReq struct {
		Nickname string `json:"nickname" form:"nickname"`
		Sex      int    `json:"sex" form:"sex"`
	}
	var data updateProfileReq
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	if data.Sex == 0 && data.Nickname == "" {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, "无修改参数"))
		return
	}
	userInfo, _ := jwt.GetClaimsByContext(ctx)
	err := s.UserRepo.UpdateProfile(userInfo.ID, data.Nickname, data.Sex)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
