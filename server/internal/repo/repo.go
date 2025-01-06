package repo

import (
	"server/internal/db"
	"server/internal/repo/user"
	"server/internal/repo/verify_code"
)

type Repo struct {
	UserRepo       user.UserRepoInterface
	VerifyCodeRepo verify_code.VerifyCodeRepoInterface
}

func InitRepo(conn *db.DBConn) *Repo {
	return &Repo{
		UserRepo:       user.NewUserRepo(conn.MySQLConn, conn.RedisConn),
		VerifyCodeRepo: verify_code.NewVerifyCodeRepo(conn.RedisConn),
	}
}
