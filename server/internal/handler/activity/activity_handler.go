package activity

import (
	"github.com/gin-gonic/gin"
	"server/internal/repo/activity"
	"server/internal/repo/activityuser"
)

type HandlerInterface interface {
	AddActivity(ctx *gin.Context)
	DeleteActivity(ctx *gin.Context)
	SearchActivity(ctx *gin.Context)
	JoinActivity(ctx *gin.Context)
	ExitActivity(ctx *gin.Context)
}

type activityHandler struct {
	ActivityRepo     activity.RepoInterface
	ActivityUserRepo activityuser.RepoInterface
}

func NewActivityHandler(activityRepo activity.RepoInterface,
	activityUserRepo activityuser.RepoInterface,
) HandlerInterface {
	return &activityHandler{
		ActivityRepo:     activityRepo,
		ActivityUserRepo: activityUserRepo,
	}
}
