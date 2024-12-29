package svc

import (
	"file_server/internal/config"
	"file_server/internal/data/database"
	"file_server/internal/data/repo"
)

type ServiceContext struct {
	Config config.RpcServerConfig
	Repo   *repo.Repo
}

func NewServiceContext(c config.RpcServerConfig) *ServiceContext {

	db := database.NewSqlite3(c.DBPath)

	return &ServiceContext{
		Config: c,
		Repo:   repo.NewRepository(db),
	}
}
