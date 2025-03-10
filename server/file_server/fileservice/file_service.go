// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: file_server.proto

package fileservice

import (
	"context"

	"file_server/api/v1/file_server"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetAvatarUrlRequest  = file_server.GetAvatarUrlRequest
	GetAvatarUrlResponse = file_server.GetAvatarUrlResponse
	UploadAvatarRequest  = file_server.UploadAvatarRequest
	UploadAvatarResponse = file_server.UploadAvatarResponse

	FileService interface {
		// 上传文件的RPC方法
		UploadUserAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error)
		GetUserAvatarUrl(ctx context.Context, in *GetAvatarUrlRequest, opts ...grpc.CallOption) (*GetAvatarUrlResponse, error)
		UploadActivityAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error)
		GetActivityAvatarUrl(ctx context.Context, in *GetAvatarUrlRequest, opts ...grpc.CallOption) (*GetAvatarUrlResponse, error)
	}

	defaultFileService struct {
		cli zrpc.Client
	}
)

func NewFileService(cli zrpc.Client) FileService {
	return &defaultFileService{
		cli: cli,
	}
}

// 上传文件的RPC方法
func (m *defaultFileService) UploadUserAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error) {
	client := file_server.NewFileServiceClient(m.cli.Conn())
	return client.UploadUserAvatar(ctx, in, opts...)
}

func (m *defaultFileService) GetUserAvatarUrl(ctx context.Context, in *GetAvatarUrlRequest, opts ...grpc.CallOption) (*GetAvatarUrlResponse, error) {
	client := file_server.NewFileServiceClient(m.cli.Conn())
	return client.GetUserAvatarUrl(ctx, in, opts...)
}

func (m *defaultFileService) UploadActivityAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error) {
	client := file_server.NewFileServiceClient(m.cli.Conn())
	return client.UploadActivityAvatar(ctx, in, opts...)
}

func (m *defaultFileService) GetActivityAvatarUrl(ctx context.Context, in *GetAvatarUrlRequest, opts ...grpc.CallOption) (*GetAvatarUrlResponse, error) {
	client := file_server.NewFileServiceClient(m.cli.Conn())
	return client.GetActivityAvatarUrl(ctx, in, opts...)
}
