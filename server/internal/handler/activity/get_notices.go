package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/response"
	"strconv"
)

func (a *activityHandler) GetNotices(ctx *gin.Context) {
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	result, err := a.ActivityNoticeRepo.SelectByAID(aid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(map[string]any{
		"rows": *result,
	}))
}
