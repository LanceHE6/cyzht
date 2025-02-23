package chat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/pkg/bindparams"
	"server/pkg/jwt"
	"server/pkg/response"
	"strconv"
)

func (c *chatHandler) SendToActivity(ctx *gin.Context) {
	aid, err := strconv.ParseInt(ctx.Param("aid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.FailedResponse(100, err.Error()))
		return
	}
	type SendToActivityRequest struct {
		MsgType  uint8  `json:"msg_type" binding:"required"`
		TextMsg  string `json:"text_msg"`
		FileURL  string `json:"file_url"`
		FileSize int64  `json:"file_size"`
	}
	data := bindparams.BindPostParams[SendToActivityRequest](ctx)
	if data == nil {
		return
	}
	userClaims, _ := jwt.GetClaimsByContext(ctx)

	// 创建消息对象
	msg := model.MsgModel{
		ActivityID:  aid,
		ExhibitorID: 0, // 0 表示大厅聊天
		MetaMsg: model.MetaMsg{
			UserID:   userClaims.ID,
			SendTo:   0, // 0 表示发送给所有人
			MsgType:  model.TEXT,
			TextMsg:  data.TextMsg,
			FileURL:  data.FileURL,
			FileSize: data.FileSize,
		},
	}
	if err := c.MsgRepo.Insert(&msg); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
