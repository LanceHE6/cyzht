package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/bindparams"
	"server/pkg/response"
)

// DeleteActivity 删除活动
func (a *activityHandler) DeleteActivity(ctx *gin.Context) {
	type delActivityRequest struct {
		ID int64 `json:"id" binding:"required"`
	}
	data := bindparams.BindParams[delActivityRequest](ctx)
	if data == nil {
		return
	}
	if err := a.ActivityRepo.DeleteByID(data.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "failed to delete activity", err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
