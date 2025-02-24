package chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/repo/msg"
	"server/pkg/bindparams"
	"server/pkg/response"
	"strconv"
)

func (c *chatHandler) GetActivityMsg(ctx *gin.Context) {
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}

	type searchParams struct {
		// form tag绑定url参数
		Page  *int `form:"page"`      // 页码
		Limit *int `form:"page_size"` // 每页条数
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

	result, total, err := c.MsgRepo.GetActivityMsg(aid,
		msg.WithPage(data.Page, data.Limit),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(-1, "failed to get activity msg", err))
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
