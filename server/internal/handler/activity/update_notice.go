package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/pkg/bindparams"
	"server/pkg/jwt"
	"server/pkg/response"
	"strconv"
)

func (a *activityHandler) UpdateNotice(ctx *gin.Context) {
	nid, err := strconv.ParseInt(ctx.Param("nid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	type updateNoticeRequest struct {
		Title   string `json:"title" binding:"required" form:"title"`
		Content string `json:"content" binding:"required" form:"content"`
	}
	data := bindparams.BindPostParams[updateNoticeRequest](ctx)
	if data == nil {
		return
	}
	claims, _ := jwt.GetClaimsByContext(ctx)
	if err := a.ActivityNoticeRepo.Update(&model.ActivityNoticeModel{
		BaseModel: model.BaseModel{ID: nid},
		Title:     data.Title,
		Content:   data.Content,
		CreatorID: claims.ID,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
