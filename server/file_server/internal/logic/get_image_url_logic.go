package logic

import (
	"context"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImageUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetImageUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImageUrlLogic {
	return &GetImageUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetImageUrlLogic) GetImageUrl(in *file_server.GetImageUrlRequest) (*file_server.GetImageUrlResponse, error) {
	icon, err := l.svcCtx.Repo.ImageRepo.FindByID(in.Id)
	if icon == nil {
		return &file_server.GetImageUrlResponse{}, err
	}

	return &file_server.GetImageUrlResponse{
		FileUrl: icon.FileURL,
	}, nil
}
