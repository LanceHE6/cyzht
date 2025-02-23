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
	ActivityRepo activity.ActivityRepoInterface
}

func NewActivityHandler(activityRepo activity.ActivityRepoInterface) HandlerInterface {
	return &activityHandler{
		ActivityRepo: activityRepo,
	}
}
