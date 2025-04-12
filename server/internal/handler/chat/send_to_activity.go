package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/internal/model"
	"server/internal/ws"
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

	type AttachmentRequest struct {
		Type     uint8  `json:"type" binding:"required"` // 1-图片, 2-文件等
		FileURL  string `json:"file_url" binding:"required"`
		ThumbURL string `json:"thumb_url"` // 缩略图URL
		FileName string `json:"file_name"`
		FileSize int64  `json:"file_size"`
		Width    int    `json:"width"`    // 图片宽度
		Height   int    `json:"height"`   // 图片高度
		Duration int    `json:"duration"` // 音视频时长(秒)
	}

	type SendToActivityRequest struct {
		MsgType     model.MsgType       `json:"msg_type" binding:"required"`
		Content     string              `json:"content"`     // 文本内容
		Attachments []AttachmentRequest `json:"attachments"` // 消息附件
	}

	data := bindparams.BindPostParams[SendToActivityRequest](ctx)
	if data == nil {
		return
	}
	fmt.Printf("%v\n", data)

	userClaims, _ := jwt.GetClaimsByContext(ctx)

	// 创建消息对象
	msg := model.MsgModel{
		ActivityID:  aid,
		ExhibitorID: 0, // 0 表示大厅聊天
		MetaMsg: model.MetaMsg{
			FromUID: userClaims.ID,
			ToUID:   0, // 0 表示发送给所有人
			MsgType: data.MsgType,
			Content: data.Content,
		},
		Status: 1, // 默认正常状态
	}

	// 处理附件
	if len(data.Attachments) > 0 {
		msg.Attachments = make([]model.Attachment, len(data.Attachments))
		for i, att := range data.Attachments {
			msg.Attachments[i] = model.Attachment{
				Type:     att.Type,
				FileURL:  att.FileURL,
				ThumbURL: att.ThumbURL,
				FileName: att.FileName,
				FileSize: att.FileSize,
				Width:    att.Width,
				Height:   att.Height,
				Duration: att.Duration,
			}
		}
	}

	insertMsg, err := c.MsgRepo.Insert(&msg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.FailedResponse(-1, err.Error()))
		return
	}

	// 将消息写入消息推送队列
	ws.ActivityMQ <- insertMsg

	ctx.JSON(http.StatusOK, response.SuccessResponse(nil))
}
