package logic

import (
	"context"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAvatarUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAvatarUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAvatarUrlLogic {
	return &GetUserAvatarUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAvatarUrlLogic) GetUserAvatarUrl(in *file_server.GetAvatarUrlRequest) (*file_server.GetAvatarUrlResponse, error) {
	avatar, err := l.svcCtx.Repo.UserAvatarRepo.FindByID(in.Id)
	if avatar == nil {
		return &file_server.GetAvatarUrlResponse{}, err
	}

	return &file_server.GetAvatarUrlResponse{
		FileUrl: avatar.FileURL,
	}, nil
}
