package logic

import (
	"context"
	"file_server/internal/data/models"
	"fmt"
	"os"
	"strconv"
	"time"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadAvatar 上传文件的RPC方法
func (l *UploadAvatarLogic) UploadAvatar(in *file_server.UploadAvatarRequest) (*file_server.UploadAvatarResponse, error) {
	// 生成唯一的文件ID，这里简单使用时间戳+自增数字（实际可使用更严谨的UUID等方式）
	filename := strconv.FormatInt(time.Now().UnixNano(), 10)
	avatarPath := l.svcCtx.Config.StoragePath + "/avatar/" + filename + in.FileType
	url := "/avatar/" + filename + in.FileType
	// 保存文件到本地磁盘
	err := os.WriteFile(avatarPath, in.FileContent, 0644)
	if err != nil {
		return nil, fmt.Errorf("保存文件到磁盘失败: %v", err.Error())
	}

	// 将文件信息插入到数据库中
	_, err = l.svcCtx.Repo.UserAvatarRepo.InsertOrUpdate(&models.UserAvatarModel{
		BaseModel: models.BaseModel{
			ID: in.Id,
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

	return &file_server.UploadAvatarResponse{
		FileUrl: url,
	}, nil

}
