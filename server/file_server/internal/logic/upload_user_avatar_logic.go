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

type UploadUserAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadUserAvatar 上传文件的RPC方法
func (l *UploadUserAvatarLogic) UploadUserAvatar(in *file_server.UploadAvatarOrIconRequest) (*file_server.UploadAvatarOrIconResponse, error) {
	// 计算hash
	hash := encrypt.CalculateFileHash(in.FileContent)
	// 检查是否已存在相同哈希值的文件
	existingIcon, err := l.svcCtx.Repo.UserAvatarRepo.FindByHash(hash)
	if err == nil {
		// 文件已存在，直接更新该用户的头像数据
		_, err := l.svcCtx.Repo.UserAvatarRepo.InsertOrUpdate(&models.UserAvatarModel{
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
			return nil, fmt.Errorf("更新用户头像数据失败: %v", err)
		}

		// 返回已有的 URL
		return &file_server.UploadAvatarOrIconResponse{
			FileUrl: existingIcon.FileURL,
		}, nil
	}
	// 文件名采用hash值
	filename := hash
	dir := "/user_avatar/"
	avatarPath := l.svcCtx.Config.StoragePath + dir + filename + in.FileType
	url := dir + filename + in.FileType
	// 保存文件到本地磁盘
	err = os.WriteFile(avatarPath, in.FileContent, 0644)
	if err != nil {
		return nil, fmt.Errorf("保存文件到磁盘失败: %v", err.Error())
	}

	// 将文件信息插入到数据库中
	_, err = l.svcCtx.Repo.UserAvatarRepo.InsertOrUpdate(&models.UserAvatarModel{

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

	return &file_server.UploadAvatarOrIconResponse{
		FileUrl: url,
	}, nil
}
