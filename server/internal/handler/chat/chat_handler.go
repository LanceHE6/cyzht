package chat

import (
	"github.com/gin-gonic/gin"
	"server/internal/repo/msg"
)

type HandlerInterface interface {
	SendToActivity(ctx *gin.Context)
	GetActivityMsg(ctx *gin.Context)
}

type chatHandler struct {
	MsgRepo msg.RepoInterface
}

func NewChatHandler(msgRepo msg.RepoInterface) HandlerInterface {
	return &chatHandler{
		MsgRepo: msgRepo,
	}
}
