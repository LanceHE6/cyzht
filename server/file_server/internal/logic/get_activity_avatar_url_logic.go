package logic

import (
	"context"

	"file_server/api/v1/file_server"
	"file_server/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityAvatarUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityAvatarUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityAvatarUrlLogic {
	return &GetActivityAvatarUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityAvatarUrlLogic) GetActivityAvatarUrl(in *file_server.GetAvatarUrlRequest) (*file_server.GetAvatarUrlResponse, error) {
	avatar, err := l.svcCtx.Repo.ActivityAvatarRepo.FindByID(in.Id)
	if avatar == nil {
		return &file_server.GetAvatarUrlResponse{}, err
	}

	return &file_server.GetAvatarUrlResponse{
		FileUrl: avatar.FileURL,
	}, nil
}
