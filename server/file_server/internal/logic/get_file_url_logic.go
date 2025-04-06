package logic

import (
	"context"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileUrlLogic {
	return &GetFileUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileUrlLogic) GetFileUrl(in *file_server.GetFileUrlRequest) (*file_server.GetFileUrlResponse, error) {
	// todo: add your logic here and delete this line

	return &file_server.GetFileUrlResponse{}, nil
}
