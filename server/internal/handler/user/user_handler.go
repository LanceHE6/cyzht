package user

import (
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/internal/repo/user"
	"server/internal/repo/verifycode"
	"server/pkg/rpc/file_server/api/v1/file_server"
)

// UserHandlerInterface
//
//	@Description: 用户服务接口
type UserHandlerInterface interface {
	RegisterAndLoginSendCode(ctx *gin.Context)   // 发送登录注册验证码
	RegisterAndLoginVerifyCode(ctx *gin.Context) // 验证登录注册验证码
	Login(ctx *gin.Context)                      // 用户登录
	OnlineHeartbeat(ctx *gin.Context)            // 用户在线心跳
	UpdatePassword(ctx *gin.Context)             // 修改密码
	UpdateAvatar(ctx *gin.Context)               // 修改头像
	UpdateProfile(ctx *gin.Context)              // 修改个人资料
	GetUserInfo(ctx *gin.Context)                // 获取用户信息
}

// userHandler
//
//	@Description: 用户服务实现
type userHandler struct {
	C              *config.Config
	UserRepo       user.UserRepoInterface
	VerifyCodeRepo verifycode.VerifyCodeRepoInterface
	FileRpcServer  file_server.FileServiceClient
}

// NewUserHandler
//
//	@Description: 创建用户服务实例
//	@return UserHandlerInterface 用户服务实例
func NewUserHandler(
	c *config.Config,
	userRepo user.UserRepoInterface,
	verifyCodeRepo verifycode.VerifyCodeRepoInterface,
	fileRpcServer file_server.FileServiceClient,
) UserHandlerInterface {
	return &userHandler{
		C:              c,
		UserRepo:       userRepo,
		VerifyCodeRepo: verifyCodeRepo,
		FileRpcServer:  fileRpcServer,
	}
}
