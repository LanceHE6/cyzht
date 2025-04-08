package svc

import (
	"file_server/internal/config"
	"file_server/internal/data/database"
	"file_server/internal/data/repo"
	"file_server/pkg/snowflake"
)

type ServiceContext struct {
	Config          config.RpcServerConfig
	Repo            *repo.Repo
	SnowflakeWorker *snowflake.Worker
}

func NewServiceContext(c config.RpcServerConfig) *ServiceContext {

	db := database.NewSqlite3(c.DBPath)
	worker, err := snowflake.NewWorker(2)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:          c,
		Repo:            repo.NewRepository(db),
		SnowflakeWorker: worker,
	}
}
