package activity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	activityRepo "server/internal/repo/activity"
	"server/pkg/bindparams"
	"server/pkg/response"
)

func (a *activityHandler) SearchActivity(ctx *gin.Context) {
	type searchParams struct {
		// form tag绑定url参数
		Page  *int `form:"page"`      // 页码
		Limit *int `form:"page_size"` // 每页条数

		ID           *int64  `form:"id,string"`
		Name         *string `form:"name"`
		Creator      *string `form:"creator"`
		Location     *string `form:"location"`
		IsInProgress *bool   `form:"is_in_progress"`
		Keyword      *string `form:"keyword"`
	}
	var data = bindparams.BindQueryParams[searchParams](ctx)
	if data.Page == nil {
		data.Page = new(int)
		*data.Page = 1
	}
	if data.Limit == nil || *data.Limit < 1 {
		data.Limit = new(int)
		*data.Limit = 10
	}
	result, total, err := a.ActivityRepo.Search(
		activityRepo.WithID(data.ID),
		activityRepo.WithName(data.Name),
		activityRepo.WithCreator(data.Creator),
		activityRepo.WithLocation(data.Location),
		activityRepo.WithIsInProgress(data.IsInProgress),
		activityRepo.WithKeyword(data.Keyword),
		activityRepo.WithPage(data.Page, data.Limit),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "failed to search activity", err))
		return
	}
	respData := map[string]any{
		"page":      *data.Page,
		"page_size": *data.Limit,
		"rows":      *result,
		"total":     total,
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(respData))
	return
}
