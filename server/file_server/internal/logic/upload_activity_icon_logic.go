package logic

import (
	"context"
	"file_server/api/v1/file_server"
	"file_server/internal/data/models"
	"file_server/internal/svc"
	"file_server/pkg/encrypt"
	"fmt"
	"os"

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

// UploadActivityIcon 上传活动图标
func (l *UploadActivityIconLogic) UploadActivityIcon(in *file_server.UploadFileRequest) (*file_server.UploadFileResponse, error) {
	// 计算hash
	hash := encrypt.CalculateFileHash(in.FileContent)
	// 检查是否已存在相同哈希值的文件
	existingIcon, err := l.svcCtx.Repo.ActivityIconRepo.FindByHash(hash)
	if err == nil {
		// 文件已存在，直接更新数据
		_, err := l.svcCtx.Repo.ActivityIconRepo.InsertOrUpdate(&models.ActivityIconModel{
			BaseModel: models.BaseModel{
				ID:   in.Id,
				Hash: hash,
			},
			FileName: existingIcon.FileName,
			FileType: existingIcon.FileType,
			FileSize: existingIcon.FileSize,
			FileURL:  existingIcon.FileURL,
		})
		if err != nil {
			return nil, fmt.Errorf("更新活动图标数据失败: %v", err)
		}

		// 返回已有的 URL
		return &file_server.UploadFileResponse{
			FileUrl: existingIcon.FileURL,
		}, nil
	}
	// 文件名采用hash值
	filename := hash
	dir := "/activity_icon/"
	iconPath := l.svcCtx.Config.StoragePath + dir + filename + in.FileType
	url := dir + filename + in.FileType
	// 保存文件到本地磁盘
	err = os.WriteFile(iconPath, in.FileContent, 0644)
	if err != nil {
		return nil, fmt.Errorf("保存文件到磁盘失败: %v", err.Error())
	}

	// 将文件信息插入到数据库中
	_, err = l.svcCtx.Repo.ActivityIconRepo.InsertOrUpdate(&models.ActivityIconModel{

		BaseModel: models.BaseModel{
			ID:   in.Id,
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

	return &file_server.UploadFileResponse{
		FileUrl: url,
	}, nil
}
