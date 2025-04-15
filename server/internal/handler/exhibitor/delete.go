package exhibitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/bindparams"
	"server/pkg/response"
)

// DeleteExhibitor 删除参展商
func (e *exhibitorHandler) DeleteExhibitor(ctx *gin.Context) {
	type delExhibitorRequest struct {
		ID int64 `json:"id,string" binding:"required"`
	}
	data := bindparams.BindPostParams[delExhibitorRequest](ctx)
	if data == nil {
		return
	}
	if err := e.ExhibitorRepo.DeleteByID(data.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "删除参展商失败", err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
