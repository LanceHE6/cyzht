package logic

import (
	"context"
	"fmt"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetImageUrlsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetImageUrlsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetImageUrlsLogic {
	return &BatchGetImageUrlsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BatchGetImageUrls 批量获取图片URL
func (l *BatchGetImageUrlsLogic) BatchGetImageUrls(in *file_server.BatchGetImageUrlsRequest) (*file_server.BatchGetImageUrlsResponse, error) {
	fileInfos, err := l.svcCtx.Repo.FileRepo.FindByIDs(in.FileIds)
	if err != nil {
		return nil, fmt.Errorf("获取文件信息失败: %v", err)
	}

	resp := &file_server.BatchGetImageUrlsResponse{}
	for _, file := range fileInfos {
		imageInfo := &file_server.BatchGetImageUrlsResponse_ImageInfo{
			FileId:  file.ID,
			FileUrl: file.FileURL,
			Width:   int32(file.Width),
			Height:  int32(file.Height),
		}

		if in.NeedThumb && file.ThumbURL != "" {
			imageInfo.ThumbUrl = file.ThumbURL
		}

		resp.Images = append(resp.Images, imageInfo)
	}

	return resp, nil
}
