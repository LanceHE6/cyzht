package repo

import (
	"server/internal/db"
)

type Repo struct {
	UserRepo       UserRepoInterface
	VerifyCodeRepo VerifyCodeRepoInterface
	ActivityRepo   ActivityRepoInterface
}

func InitRepo(conn *db.DBConn) *Repo {
	return &Repo{
		UserRepo:       NewUserRepo(conn.MySQLConn, conn.RedisConn),
		VerifyCodeRepo: NewVerifyCodeRepo(conn.RedisConn),
		ActivityRepo:   NewActivityRepo(conn.MySQLConn),
	}
}
