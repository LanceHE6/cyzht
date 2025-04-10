package logic

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/disintegration/imaging"
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

func (l *UploadImageLogic) UploadImage(in *file_server.UploadImageRequest) (*file_server.UploadImageResponse, error) {
	// 1. 使用文件上传器上传文件
	fileResp, err := l.svcCtx.FileUploader.UploadFile(l.ctx, &file_server.UploadFileRequest{
		FileContent: in.FileContent,
		FileName:    in.FileName,
		OwnerId:     in.OwnerId,
		OwnerType:   in.OwnerType,
	})
	if err != nil {
		return nil, err
	}

	resp := &file_server.UploadImageResponse{
		FileId:  fileResp.FileId,
		FileUrl: fileResp.FileUrl,
	}

	// 2. 如果不是图片，直接返回
	if !strings.HasPrefix(fileResp.MimeType, "image/") {
		return resp, nil
	}

	// 3. 获取文件信息
	fileInfo, err := l.svcCtx.Repo.FileRepo.FindByID(fileResp.FileId)
	if err != nil {
		return nil, fmt.Errorf("获取文件信息失败: %v", err)
	}

	resp.Width = int32(fileInfo.Width)
	resp.Height = int32(fileInfo.Height)

	// 4. 如果需要缩略图
	if in.NeedThumb {
		thumbWidth, thumbHeight := l.calculateThumbSize(fileInfo.Width, fileInfo.Height,
			int(in.ThumbWidth), int(in.ThumbHeight))

		// 读取原始图片
		filePath := filepath.Join(l.svcCtx.Config.StoragePath, strings.TrimPrefix(fileInfo.FileURL, "/files/"))
		img, err := imaging.Open(filePath)
		if err != nil {
			return nil, fmt.Errorf("打开图片失败: %v", err)
		}

		// 生成缩略图
		thumbImg := imaging.Resize(img, thumbWidth, thumbHeight, imaging.Lanczos)
		thumbFilename := "thumb_" + filepath.Base(filePath)
		thumbPath := filepath.Join(filepath.Dir(filePath), thumbFilename)
		thumbUrl := filepath.Join(filepath.Dir(fileInfo.FileURL), thumbFilename)

		// 保存缩略图
		if err := imaging.Save(thumbImg, thumbPath); err != nil {
			return nil, fmt.Errorf("保存缩略图失败: %v", err)
		}

		// 更新数据库
		fileInfo.ThumbURL = thumbUrl
		if _, err := l.svcCtx.Repo.FileRepo.Upsert(fileInfo); err != nil {
			os.Remove(thumbPath)
			return nil, fmt.Errorf("更新缩略图信息失败: %v", err)
		}

		resp.ThumbUrl = thumbUrl
		resp.ThumbWidth = int32(thumbWidth)
		resp.ThumbHeight = int32(thumbHeight)
	}

	return resp, nil
}

func (l *UploadImageLogic) calculateThumbSize(origWidth, origHeight, reqWidth, reqHeight int) (int, int) {
	maxWidth := l.svcCtx.Config.ThumbMaxWidth
	maxHeight := l.svcCtx.Config.ThumbMaxHeight

	// 如果请求指定了尺寸，使用请求尺寸
	if reqWidth > 0 && reqHeight > 0 {
		return reqWidth, reqHeight
	} else if reqWidth > 0 {
		return reqWidth, origHeight * reqWidth / origWidth
	} else if reqHeight > 0 {
		return origWidth * reqHeight / origHeight, reqHeight
	}

	// 否则使用配置的默认尺寸
	if origWidth <= maxWidth && origHeight <= maxHeight {
		return origWidth, origHeight
	}

	ratio := float64(origWidth) / float64(origHeight)
	if origWidth > origHeight {
		return maxWidth, int(float64(maxWidth) / ratio)
	} else {
		return int(float64(maxHeight) * ratio), maxHeight
	}
}
