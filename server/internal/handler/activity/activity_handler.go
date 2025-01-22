package activity

import (
	"github.com/gin-gonic/gin"
	"server/internal/repo"
)

type ActivityHandlerInterface interface {
	AddActivity(ctx *gin.Context)
	DeleteActivity(ctx *gin.Context)
}

type activityHandler struct {
	ActivityRepo repo.ActivityRepoInterface
}

func NewActivityHandler(activityRepo repo.ActivityRepoInterface) ActivityHandlerInterface {
	return &activityHandler{
		ActivityRepo: activityRepo,
	}
}
