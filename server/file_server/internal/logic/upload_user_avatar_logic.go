package logic

import (
	"context"
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

// UploadUserAvatar 上传用户头像(兼容旧接口)
func (l *UploadUserAvatarLogic) UploadUserAvatar(in *file_server.UploadAvatarOrIconRequest) (*file_server.UploadAvatarOrIconResponse, error) {
	fileResp, err := l.svcCtx.FileUploader.UploadFile(l.ctx, &file_server.UploadFileRequest{
		FileContent: in.FileContent,
		FileName:    in.FileName,
		OwnerId:     in.Id,
		OwnerType:   "user",
	})
	if err != nil {
		return nil, err
	}

	return &file_server.UploadAvatarOrIconResponse{
		FileUrl: fileResp.FileUrl,
	}, nil
}
