package repo

import (
	"server/internal/config"
	"server/internal/db"
	"server/internal/repo/activity"
	"server/internal/repo/activityuser"
	"server/internal/repo/exhibitor"
	"server/internal/repo/msg"
	"server/internal/repo/user"
	"server/internal/repo/verifycode"
	"server/pkg/rpc/file_server/api/v1/file_server"
	"sync"
)

type Repo struct {
	UserRepo         user.RepoInterface
	VerifyCodeRepo   verifycode.RepoInterface
	ActivityRepo     activity.RepoInterface
	ExhibitorRepo    exhibitor.ExhibitorRepoInterface
	MsgRepo          msg.RepoInterface
	ActivityUserRepo activityuser.RepoInterface
}

// repo 全局repo单例
var repo *Repo
var once sync.Once

// InitRepo 初始化Repo
func InitRepo(c *config.Config, conn *db.DBConn, fileRpcServer file_server.FileServiceClient) *Repo {
	if repo == nil {
		once.Do(func() {

			repo = &Repo{
				UserRepo:         user.NewUserRepo(c, conn, fileRpcServer),
				VerifyCodeRepo:   verifycode.NewVerifyCodeRepo(conn),
				ActivityRepo:     activity.NewActivityRepo(conn),
				ExhibitorRepo:    exhibitor.NewExhibitorRepo(conn),
				MsgRepo:          msg.NewMsgRepo(conn),
				ActivityUserRepo: activityuser.NewActivityUserRepo(conn),
			}
		})
	}
	return repo
}

// GetRepo 获取Repo
func GetRepo() *Repo {
	return repo
}
