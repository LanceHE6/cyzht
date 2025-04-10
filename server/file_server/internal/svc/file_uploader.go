package svc

import (
	"context"
	"file_server/api/v1/file_server"
	"file_server/internal/config"
	"file_server/internal/data/models"
	"file_server/internal/data/repo"
	"file_server/pkg/encrypt"
	"file_server/pkg/snowflake"
	"fmt"
	"image"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileUploader interface {
	UploadFile(ctx context.Context, in *file_server.UploadFileRequest) (*file_server.UploadFileResponse, error)
}

type fileUploader struct {
	repo   *repo.Repo
	config config.RpcServerConfig
	worker *snowflake.Worker
}

func NewFileUploader(repo *repo.Repo, cfg config.RpcServerConfig, worker *snowflake.Worker) FileUploader {
	return &fileUploader{
		repo:   repo,
		config: cfg,
		worker: worker,
	}
}

func (u *fileUploader) UploadFile(ctx context.Context, in *file_server.UploadFileRequest) (*file_server.UploadFileResponse, error) {
	// 1. 验证文件类型
	mimeType := in.MimeType
	if mimeType == "" {
		mimeType = http.DetectContentType(in.FileContent[:512])
	}

	// 2. 计算文件哈希
	hash := encrypt.CalculateFileHash(in.FileContent)

	// 3. 检查文件是否已存在
	existingFile, err := u.repo.FileRepo.FindByHash(hash)
	if err == nil && existingFile != nil {
		return &file_server.UploadFileResponse{
			FileId:   existingFile.ID,
			FileUrl:  existingFile.FileURL,
			FileSize: existingFile.FileSize,
			MimeType: mimeType,
		}, nil
	}

	// 4. 确定文件存储路径
	fileExt := strings.ToLower(filepath.Ext(in.FileName))
	if fileExt == "" {
		exts, _ := mime.ExtensionsByType(mimeType)
		if len(exts) > 0 {
			fileExt = exts[0]
		} else {
			fileExt = ".bin"
		}
	}

	filename := hash + fileExt
	storagePath := u.config.StoragePath + "/" + in.OwnerType
	filePath := filepath.Join(storagePath, filename)
	fileUrl := "/" + in.OwnerType + "/" + filename

	// 5. 保存文件
	if err := u.saveFile(filePath, in.FileContent); err != nil {
		return nil, fmt.Errorf("保存文件失败: %v", err)
	}

	// 6. 保存到数据库
	fileType := u.determineFileType(mimeType)
	fileModel := &models.FileModel{
		BaseModel: models.BaseModel{
			ID:   u.worker.NextId(),
			Hash: hash,
		},
		FileName:  in.FileName,
		FileType:  fileType,
		MimeType:  mimeType,
		FileSize:  int64(len(in.FileContent)),
		FileURL:   fileUrl,
		OwnerID:   in.OwnerId,
		OwnerType: in.OwnerType,
	}

	// 如果是图片，提取宽高信息
	if strings.HasPrefix(mimeType, "image/") {
		if img, _, err := image.Decode(strings.NewReader(string(in.FileContent))); err == nil {
			fileModel.Width = img.Bounds().Dx()
			fileModel.Height = img.Bounds().Dy()
		}
	}

	if _, err := u.repo.FileRepo.Upsert(fileModel); err != nil {
		os.Remove(filePath)
		return nil, fmt.Errorf("保存文件信息失败: %v", err)
	}

	return &file_server.UploadFileResponse{
		FileId:   fileModel.ID,
		FileUrl:  fileModel.FileURL,
		FileSize: fileModel.FileSize,
		MimeType: fileModel.MimeType,
	}, nil
}

func (u *fileUploader) saveFile(path string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, content, 0644)
}

func (u *fileUploader) determineFileType(mimeType string) models.FileType {
	switch {
	case strings.HasPrefix(mimeType, "image/"):
		return models.FileTypeImage
	case strings.HasPrefix(mimeType, "video/"):
		return models.FileTypeVideo
	case strings.HasPrefix(mimeType, "audio/"):
		return models.FileTypeAudio
	default:
		return models.FileTypeDocument
	}
}
