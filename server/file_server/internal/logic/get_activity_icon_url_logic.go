package logic

import (
	"context"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityIconUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityIconUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityIconUrlLogic {
	return &GetActivityIconUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityIconUrlLogic) GetActivityIconUrl(in *file_server.GetFileUrlRequest) (*file_server.GetFileUrlResponse, error) {
	icon, err := l.svcCtx.Repo.ActivityIconRepo.FindByID(in.Id)
	if icon == nil {
		return &file_server.GetFileUrlResponse{}, err
	}

	return &file_server.GetFileUrlResponse{
		FileUrl: icon.FileURL,
	}, nil
}
