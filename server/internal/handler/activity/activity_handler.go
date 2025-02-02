package activity

import (
	"github.com/gin-gonic/gin"
	"server/internal/repo/activity"
)

type ActivityHandlerInterface interface {
	AddActivity(ctx *gin.Context)
	DeleteActivity(ctx *gin.Context)
	SearchActivity(ctx *gin.Context)
}

type activityHandler struct {
	ActivityRepo activity.ActivityRepoInterface
}

func NewActivityHandler(activityRepo activity.ActivityRepoInterface) ActivityHandlerInterface {
	return &activityHandler{
		ActivityRepo: activityRepo,
	}
}
