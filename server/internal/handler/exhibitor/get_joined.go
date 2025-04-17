package exhibitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/jwt"
	"server/pkg/response"
	"strconv"
)

func (e *exhibitorHandler) GetJoinedExhibitors(ctx *gin.Context) {
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	claims, _ := jwt.GetClaimsByContext(ctx)
	exhibitors, err := e.ExhibitorUserRepo.SelectByUIDAndAID(claims.ID, aid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(map[string]any{
		"rows": *exhibitors,
	}))
}

func (e *exhibitorHandler) GetJoinedUsers(ctx *gin.Context) {
	eid, err := strconv.ParseInt(ctx.Param("eid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	exhibitors, err := e.ExhibitorUserRepo.SelectByEID(eid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(map[string]any{
		"rows": *exhibitors,
	}))
}
