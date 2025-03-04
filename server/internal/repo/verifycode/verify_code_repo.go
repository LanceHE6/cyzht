package verifycode

import (
	"github.com/go-redis/redis"
	"server/internal/db"
	"time"
)

type RepoInterface interface {
	SetVerifyCode(account string, code string)
	GetVerifyCode(account string) string // 获取验证码
	DeleteVerifyCode(account string)     // 删除验证码
}

// NewVerifyCodeRepo
//
//	@Description: 初始化验证码repo
//	@return RepoInterface 验证码repo
func NewVerifyCodeRepo(dbConn *db.DBConn) RepoInterface {
	return &verifyCodeRepo{RDB: dbConn.RedisConn}
}

// verifyCodeRepo
//
//	@Description: 验证码仓库实现
type verifyCodeRepo struct {
	RDB *redis.Client
}

// DeleteVerifyCode
//
//	@Description: 删除验证码
//	@receiver v verifyCodeRepo
//	@param account 账号
//	@return error 错误信息
func (v verifyCodeRepo) DeleteVerifyCode(account string) {
	v.RDB.Del(account)
}

// GetVerifyCode
//
//	@Description: 获取验证码
//	@receiver v verifyCodeRepo
//	@param account 账号
//	@return string 验证码
//	@return error 错误信息
func (v verifyCodeRepo) GetVerifyCode(account string) string {
	code, _ := v.RDB.Get(account).Result()
	return code
}

// SetVerifyCode
//
//	@Description: 设置验证码
//	@receiver v verifyCodeRepo
//	@param account 账号
//	@param code 验证码
//	@return error 错误信息
func (v verifyCodeRepo) SetVerifyCode(account string, code string) {
	// 创建一个键值对用于存放验证码, key为account, value为code, 过期时间为5分钟
	v.RDB.Set(account, code, 5*time.Minute)
}
