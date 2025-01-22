package user

import (
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/internal/repo"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

// UserHandlerInterface
//
//	@Description: 用户服务接口
type UserHandlerInterface interface {
	RegisterAndLoginSendCode() gin.HandlerFunc   // 发送登录注册验证码
	RegisterAndLoginVerifyCode() gin.HandlerFunc // 验证登录注册验证码
	Login() gin.HandlerFunc                      // 用户登录
	OnlineHeartbeat() gin.HandlerFunc            // 用户在线心跳
	UpdatePassword() gin.HandlerFunc             // 修改密码
	UpdateAvatar() gin.HandlerFunc               // 修改头像
	UpdateProfile() gin.HandlerFunc              // 修改个人资料
	GetUserInfo() gin.HandlerFunc                // 获取用户信息
}

// userHandler
//
//	@Description: 用户服务实现
type userHandler struct {
	C              *config.Config
	UserRepo       repo.UserRepoInterface
	VerifyCodeRepo repo.VerifyCodeRepoInterface
	FileRpcServer  file_server.FileServiceClient
}

// NewUserHandler
//
//	@Description: 创建用户服务实例
//	@return UserHandlerInterface 用户服务实例
func NewUserHandler(
	c *config.Config,
	userRepo repo.UserRepoInterface,
	verifyCodeRepo repo.VerifyCodeRepoInterface,
	fileRpcServer file_server.FileServiceClient,
) UserHandlerInterface {
	return &userHandler{
		C:              c,
		UserRepo:       userRepo,
		VerifyCodeRepo: verifyCodeRepo,
		FileRpcServer:  fileRpcServer,
	}
}
