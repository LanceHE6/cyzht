package repo

import (
	"server/internal/db"
	"sync"
)

type Repo struct {
	UserRepo       UserRepoInterface
	VerifyCodeRepo VerifyCodeRepoInterface
	ActivityRepo   ActivityRepoInterface
}

// repo 全局repo单例
var repo *Repo
var once sync.Once

// InitRepo 初始化Repo
func InitRepo(conn *db.DBConn) *Repo {
	if repo == nil {
		once.Do(func() {
			repo = &Repo{
				UserRepo:       NewUserRepo(conn.MySQLConn, conn.RedisConn),
				VerifyCodeRepo: NewVerifyCodeRepo(conn.RedisConn),
				ActivityRepo:   NewActivityRepo(conn.MySQLConn),
			}
		})
	}
	return repo
}

// GetRepo 获取Repo
func GetRepo() *Repo {
	return repo
}
