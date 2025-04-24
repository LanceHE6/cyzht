package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/response"
	"strconv"
)

func (a *activityHandler) DeleteNotice(ctx *gin.Context) {
	nid, err := strconv.ParseInt(ctx.Param("nid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	if err := a.ActivityNoticeRepo.DeleteByID(nid); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
