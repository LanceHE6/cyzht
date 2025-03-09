package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/jwt"
	"server/pkg/response"
)

func (a *activityHandler) GetJoinedActivity(ctx *gin.Context) {
	userInfo, _ := jwt.GetClaimsByContext(ctx)
	aus, err := a.ActivityUserRepo.SelectByUID(userInfo.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(http.StatusInternalServerError, err.Error()))
		return
	}
	data := map[string]any{
		"rows": aus,
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(data))
}
