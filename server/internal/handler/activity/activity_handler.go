package activity

import (
	"github.com/gin-gonic/gin"
	"server/internal/repo/activity"
)

type HandlerInterface interface {
	AddActivity(ctx *gin.Context)
	DeleteActivity(ctx *gin.Context)
	SearchActivity(ctx *gin.Context)
}

type activityHandler struct {
	ActivityRepo activity.RepoInterface
}

func NewActivityHandler(activityRepo activity.RepoInterface) HandlerInterface {
	return &activityHandler{
		ActivityRepo: activityRepo,
	}
}
