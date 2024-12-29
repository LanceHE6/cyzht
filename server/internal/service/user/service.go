package user

import (
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/internal/data/repo/user"
	"server/internal/data/repo/verify_code"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

// UserServiceInterface
//
//	@Description: 用户服务接口
type UserServiceInterface interface {
	AuthMiddleware() gin.HandlerFunc             // 鉴权
	RegisterAndLoginSendCode() gin.HandlerFunc   // 发送登录注册验证码
	RegisterAndLoginVerifyCode() gin.HandlerFunc // 验证登录注册验证码
	Login() gin.HandlerFunc                      // 用户登录
	//OnlineHeartbeat() gin.HandlerFunc            // 用户在线心跳
	UpdatePassword() gin.HandlerFunc // 修改密码
	UpdateAvatar() gin.HandlerFunc   // 修改头像
	UpdateProfile() gin.HandlerFunc  // 修改个人资料
	GetUserInfo() gin.HandlerFunc    // 获取用户信息
}

// userService
//
//	@Description: 用户服务实现
type userService struct {
	C              *config.Config
	UserRepo       user.UserRepoInterface
	VerifyCodeRepo verify_code.VerifyCodeRepoInterface
	FileRpcServer  file_server.FileServiceClient
}

// NewUserService
//
//	@Description: 创建用户服务实例
//	@return UserServiceInterface 用户服务实例
func NewUserService(
	c *config.Config,
	userRepo user.UserRepoInterface,
	verifyCodeRepo verify_code.VerifyCodeRepoInterface,
	fileRpcServer file_server.FileServiceClient,
) UserServiceInterface {
	return &userService{
		C:              c,
		UserRepo:       userRepo,
		VerifyCodeRepo: verifyCodeRepo,
		FileRpcServer:  fileRpcServer,
	}
}
