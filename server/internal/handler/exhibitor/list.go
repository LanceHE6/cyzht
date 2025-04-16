package exhibitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/response"
	"strconv"
)

func (e *exhibitorHandler) ListExhibitors(ctx *gin.Context) {
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	exhibitors, err := e.ExhibitorRepo.SelectByAID(aid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(map[string]any{
		"rows": *exhibitors,
	}))
}
