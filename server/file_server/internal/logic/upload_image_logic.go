package logic

import (
	"context"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传图片(消息)
func (l *UploadImageLogic) UploadImage(in *file_server.UploadFileRequest) (*file_server.UploadFileResponse, error) {
	// todo: add your logic here and delete this line

	return &file_server.UploadFileResponse{}, nil
}
