package repo

import (
	"server/internal/db"
	"server/internal/repo/activity"
	"server/internal/repo/exhibitor"
	"server/internal/repo/msg"
	"server/internal/repo/user"
	"server/internal/repo/verifycode"
	"sync"
)

type Repo struct {
	UserRepo       user.UserRepoInterface
	VerifyCodeRepo verifycode.VerifyCodeRepoInterface
	ActivityRepo   activity.ActivityRepoInterface
	ExhibitorRepo  exhibitor.ExhibitorRepoInterface
	MsgRepo        msg.MsgRepoInterface
}

// repo 全局repo单例
var repo *Repo
var once sync.Once

// InitRepo 初始化Repo
func InitRepo(conn *db.DBConn) *Repo {
	if repo == nil {
		once.Do(func() {
			repo = &Repo{
				UserRepo:       user.NewUserRepo(conn.MySQLConn, conn.RedisConn),
				VerifyCodeRepo: verifycode.NewVerifyCodeRepo(conn.RedisConn),
				ActivityRepo:   activity.NewActivityRepo(conn.MySQLConn),
				ExhibitorRepo:  exhibitor.NewExhibitorRepo(conn.MySQLConn),
				MsgRepo:        msg.NewMsgRepo(conn.MySQLConn),
			}
		})
	}
	return repo
}

// GetRepo 获取Repo
func GetRepo() *Repo {
	return repo
}
