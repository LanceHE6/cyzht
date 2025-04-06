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

func (l *GetImageUrlLogic) GetImageUrl(in *file_server.GetFileUrlRequest) (*file_server.GetFileUrlResponse, error) {
	// todo: add your logic here and delete this line

	return &file_server.GetFileUrlResponse{}, nil
}
