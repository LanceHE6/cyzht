package chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/response"
	"strconv"
	"time"
)

func (c *chatHandler) WithdrawMsg(ctx *gin.Context) {
	mid, err := strconv.ParseInt(ctx.Param("mid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	msg, err := c.MsgRepo.SelectByID(mid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	// 判断发送的消息是否已超过5分钟
	if msg.CreatedAt.Add(5 * time.Minute).Before(time.Now()) {
		ctx.JSON(http.StatusOK, response.FailedResponse(11, "消息已超过5分钟，无法撤回"))
		return
	}
	if err := c.MsgRepo.UpdateStatus(mid, 2); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-2, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
