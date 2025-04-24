package activity

import (
	"github.com/gin-gonic/gin"
	"server/internal/repo/activity"
	"server/internal/repo/activitynotice"
	"server/internal/repo/activityuser"
)

type HandlerInterface interface {
	AddActivity(ctx *gin.Context)
	DeleteActivity(ctx *gin.Context)
	SearchActivity(ctx *gin.Context)
	JoinActivity(ctx *gin.Context)
	WithdrawActivity(ctx *gin.Context)
	GetJoinedActivity(ctx *gin.Context)
	GetJoinedUser(ctx *gin.Context)
	PublishNotice(ctx *gin.Context)
	UpdateNotice(ctx *gin.Context)
	GetNotices(ctx *gin.Context)
	DeleteNotice(ctx *gin.Context)
}

type activityHandler struct {
	ActivityRepo       activity.RepoInterface
	ActivityUserRepo   activityuser.RepoInterface
	ActivityNoticeRepo activitynotice.RepoInterface
}

func NewActivityHandler(activityRepo activity.RepoInterface,
	activityUserRepo activityuser.RepoInterface,
	activityNoticeRepo activitynotice.RepoInterface,
) HandlerInterface {
	return &activityHandler{
		ActivityRepo:       activityRepo,
		ActivityUserRepo:   activityUserRepo,
		ActivityNoticeRepo: activityNoticeRepo,
	}
}
