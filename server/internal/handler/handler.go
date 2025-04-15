package handler

import (
	"server/internal/config"
	"server/internal/handler/activity"
	"server/internal/handler/chat"
	"server/internal/handler/exhibitor"
	"server/internal/handler/upload"
	"server/internal/handler/user"
	"server/internal/handler/version"
	"server/internal/repo"
	"sync"
)

type Handler struct {
	VersionHandler version.HandlerInterface
	UserHandler    user.HandlerInterface

	ActivityHandler  activity.HandlerInterface
	ExhibitorHandler exhibitor.HandlerInterface
	ChatHandler      chat.HandlerInterface

	UploadHandler upload.HandlerInterface
}

// handler 全局单例
var handler *Handler
var once sync.Once

// InitHandler 初始化并获取handler
func InitHandler(c *config.Config, repo *repo.Repo) *Handler {
	if handler == nil {
		once.Do(func() {

			handler = &Handler{
				VersionHandler: version.NewVersionHandler(),
				UserHandler: user.NewUserHandler(c,
					repo.UserRepo,
					repo.VerifyCodeRepo,
					repo.ActivityUserRepo,
				),
				ActivityHandler:  activity.NewActivityHandler(repo.ActivityRepo, repo.ActivityUserRepo),
				ExhibitorHandler: exhibitor.NewExhibitorHandler(repo.ExhibitorRepo, repo.ExhibitorUserRepo),
				ChatHandler:      chat.NewChatHandler(repo.MsgRepo),
				UploadHandler:    upload.NewUploadHandler(c, repo.FileRpcServer),
			}
		})
	}
	return handler
}

// GetHandler 获取handler
func GetHandler() *Handler {
	return handler
}
