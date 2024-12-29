package repo

import (
	"server/internal/data/db"
	"server/internal/data/repo/user"
	"server/internal/data/repo/verify_code"
)

type Repo struct {
	UserRepo       user.UserRepoInterface
	VerifyCodeRepo verify_code.VerifyCodeRepoInterface
}

func InitRepo(conn *db.DBConn) *Repo {
	return &Repo{
		UserRepo:       user.NewUserRepo(conn.MySQLConn),
		VerifyCodeRepo: verify_code.NewVerifyCodeRepo(conn.RedisConn),
	}
}
