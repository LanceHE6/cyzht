package logic

import (
	"context"
	"fmt"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFileInfo 获取文件信息
func (l *GetFileInfoLogic) GetFileInfo(in *file_server.GetFileInfoRequest) (*file_server.GetFileInfoResponse, error) {
	fileInfo, err := l.svcCtx.Repo.FileRepo.FindByID(in.FileId)
	if err != nil {
		return nil, fmt.Errorf("文件不存在: %v", err)
	}

	return &file_server.GetFileInfoResponse{
		FileId:   fileInfo.ID,
		FileUrl:  fileInfo.FileURL,
		ThumbUrl: fileInfo.ThumbURL,
		FileSize: fileInfo.FileSize,
		MimeType: fileInfo.MimeType,
		Width:    int32(fileInfo.Width),
		Height:   int32(fileInfo.Height),
		Duration: int32(fileInfo.Duration),
	}, nil
}
