package logic

import (
	"context"
	"file_server/internal/data/models"
	"file_server/pkg/encrypt"
	"fmt"
	"os"

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

// UploadImage 上传图片(消息)
func (l *UploadImageLogic) UploadImage(in *file_server.UploadImageRequest) (*file_server.UploadImageResponse, error) {
	// 计算hash
	hash := encrypt.CalculateFileHash(in.FileContent)
	// 检查是否已存在相同哈希值的文件
	existingIcon, err := l.svcCtx.Repo.ImageRepo.FindByHash(hash)
	if err == nil {
		// 文件已存在，直接返回已有的 URL
		return &file_server.UploadImageResponse{
			Id:      existingIcon.ID,
			FileUrl: existingIcon.FileURL,
		}, nil
	}
	// 文件名采用hash值
	filename := hash
	dir := "/image/"
	iconPath := l.svcCtx.Config.StoragePath + dir + filename + in.FileType
	url := dir + filename + in.FileType
	// 保存文件到本地磁盘
	err = os.WriteFile(iconPath, in.FileContent, 0644)
	if err != nil {
		return nil, fmt.Errorf("保存文件到磁盘失败: %v", err.Error())
	}

	// 将文件信息插入到数据库中
	id := l.svcCtx.SnowflakeWorker.NextId() // 雪花算法生成id
	_, err = l.svcCtx.Repo.ImageRepo.InsertOrUpdate(&models.ImageModel{
		BaseModel: models.BaseModel{
			ID:   id,
			Hash: hash,
		},
		FileName: filename,
		FileType: in.FileType,
		FileSize: int64(len(in.FileContent)),
		FileURL:  url,
	},
	)
	if err != nil {
		return nil, fmt.Errorf("插入文件信息到数据库失败: %v", err)
	}

	return &file_server.UploadImageResponse{
		Id:      id,
		FileUrl: url,
	}, nil
}
