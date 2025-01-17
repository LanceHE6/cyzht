package repo

import (
	"server/internal/db"
)

type Repo struct {
	UserRepo       UserRepoInterface
	VerifyCodeRepo VerifyCodeRepoInterface
}

func InitRepo(conn *db.DBConn) *Repo {
	return &Repo{
		UserRepo:       NewUserRepo(conn.MySQLConn, conn.RedisConn),
		VerifyCodeRepo: NewVerifyCodeRepo(conn.RedisConn),
	}
}
