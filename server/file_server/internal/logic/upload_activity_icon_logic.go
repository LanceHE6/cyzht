package logic

import (
	"context"
	"file_server/api/v1/file_server"
	"file_server/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadActivityIconLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadActivityIconLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadActivityIconLogic {
	return &UploadActivityIconLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadActivityIcon 上传活动图标(兼容旧接口)
func (l *UploadActivityIconLogic) UploadActivityIcon(in *file_server.UploadAvatarOrIconRequest) (*file_server.UploadAvatarOrIconResponse, error) {
	fileResp, err := l.svcCtx.FileUploader.UploadFile(l.ctx, &file_server.UploadFileRequest{
		FileContent: in.FileContent,
		FileName:    in.FileName,
		OwnerId:     in.Id,
		OwnerType:   "activity",
	})
	if err != nil {
		return nil, err
	}

	return &file_server.UploadAvatarOrIconResponse{
		FileUrl: fileResp.FileUrl,
	}, nil
}
