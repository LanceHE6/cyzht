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
	FileUploader    FileUploader // 通用文件上传
}

func NewServiceContext(c config.RpcServerConfig) *ServiceContext {

	db := database.NewSqlite3(c.DBPath)
	worker, err := snowflake.NewWorker(2)
	if err != nil {
		panic(err)
	}

	r := repo.NewRepository(db)

	return &ServiceContext{
		Config:          c,
		Repo:            r,
		SnowflakeWorker: worker,
		FileUploader:    NewFileUploader(r, c, worker),
	}
}
