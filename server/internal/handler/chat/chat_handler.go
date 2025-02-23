package chat

import (
	"github.com/gin-gonic/gin"
	"server/internal/repo/msg"
)

type HandlerInterface interface {
	SendToActivity(ctx *gin.Context)
}

type chatHandler struct {
	MsgRepo msg.MsgRepoInterface
}

func NewChatHandler(msgRepo msg.MsgRepoInterface) HandlerInterface {
	return &chatHandler{
		MsgRepo: msgRepo,
	}
}
