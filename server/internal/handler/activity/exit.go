package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/jwt"
	"server/pkg/response"
	"strconv"
)

func (a *activityHandler) ExitActivity(ctx *gin.Context) {
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}

	claims, _ := jwt.GetClaimsByContext(ctx)
	err = a.ActivityUserRepo.Delete(claims.ID, aid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
