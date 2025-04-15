package exhibitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/jwt"
	"server/pkg/response"
	"strconv"
)

func (e *exhibitorHandler) WithdrawExhibitor(ctx *gin.Context) {
	eid, err := strconv.ParseInt(ctx.Param("eid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}

	claims, _ := jwt.GetClaimsByContext(ctx)
	err = e.ExhibitorUserRepo.Delete(claims.ID, eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
