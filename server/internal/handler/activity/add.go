package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/pkg/bindparams"
	"server/pkg/jwt"
	"server/pkg/response"
	"server/pkg/timeconv"
)

// AddActivity 新增活动
func (a *activityHandler) AddActivity(ctx *gin.Context) {
	type addActivityRequest struct {
		Name      string `json:"name" binding:"required"`
		Introduce string `json:"introduce" binding:"required"`
		StartAt   string `json:"start_at" binding:"required"`
		EndAt     string `json:"end_at" binding:"required"`
		Location  string `json:"location" binding:"required"`
	}
	data := bindparams.BindParams[addActivityRequest](ctx)
	if data == nil {
		return
	}
	startAt, err := timeconv.ParesStrToTime(data.StartAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, "start_at is in the wrong time format"))
		return
	}
	endAt, err := timeconv.ParesStrToTime(data.EndAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, "end_at is in the wrong time format"))
		return
	}

	claims, _ := jwt.GetClaimsByContext(ctx)
	activity := model.ActivityModel{
		ActivityName: data.Name,
		Introduce:    data.Introduce,
		StartAt:      *startAt,
		EndAt:        *endAt,
		Location:     data.Location,
		PromoterID:   claims.ID,
	}
	err = a.ActivityRepo.Insert(&activity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "failed in insert activity", err))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
	return
}
