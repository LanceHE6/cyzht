package user

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"server/internal/model"
	"server/pkg/encrypt"
	"strconv"
	"time"
)

// 定义 Redis 键名
const (
	onlineUsersKey = "online_users"
	userStatusKey  = "user:%d:status"
)

// UserRepoInterface
// @Description: 用户仓库接口
type UserRepoInterface interface {
	SelectByID(id int64) *model.UserModel
	SelectAll() []model.UserModel
	SelectByAccount(account string) *model.UserModel
	SelectByAccountAndPsw(account, password string) *model.UserModel
	Insert(user model.UserModel) error
	UpdateSessionID(id int64, sessionID string) error
	UpdatePassword(id int64, newPsw string) error
	UpdateOnlineStatus(id int64, onlineStatus int) error
	UpdateAvatar(id int64, avatarUrl string) error
	UpdateProfile(id int64, nickname string, sex int) error

	// SetUserOnline 设置用户在线状态
	SetUserOnline(userID int64, sessionID string, activityID uint) error
	// SetUserOffline 设置用户离线状态
	SetUserOffline(userID int64) error
	// IsUserOnline 判断用户是否在线
	IsUserOnline(userID int64) (bool, error)
	// GetOnlineUsers 获取所有在线用户
	GetOnlineUsers() ([]int64, error)
	// GetUserStatus 获取用户状态信息
	GetUserStatus(userID int64) (map[string]string, error)
	// UpdateUserHeartbeat 更新用户心跳
	UpdateUserHeartbeat(userID int64) error
}

// NewUserRepo
//
//	@Description: 创建用户仓库实例
//	@return UserRepoInterface 用户仓库实例
func NewUserRepo(mysqlConn *gorm.DB, redisConn *redis.Client) UserRepoInterface {
	return &userRepo{
		MyDB:  mysqlConn,
		Redis: redisConn,
	}
}

// userRepo
//
//	@Description: 用户仓库实现
type userRepo struct {
	MyDB  *gorm.DB
	Redis *redis.Client
}

// modelMyDB
//
//	@Description: 获取用户表
//	@receiver u userRepo
//	@return *gorm.db 用户表
func (u *userRepo) modelMyDB() *gorm.DB {
	return u.MyDB.Model(model.UserModel{})
}

// SelectByAccount
//
//	@Description: 根据账号查询用户
//	@receiver u userRepo
//	@param account string 账号
//	@return *model.UserModel 用户数据
func (u *userRepo) SelectByAccount(account string) *model.UserModel {
	var user model.UserModel
	err := u.modelMyDB().Where("account = ?", account).First(&user)
	if err.Error != nil {
		return nil
	}
	return &user
}

// SelectByID
//
//	@Description: 根据id查询用户
//	@receiver u userRepo
//	@param id int64 用户id
//	@return *model.UserModel 用户数据
func (u *userRepo) SelectByID(id int64) *model.UserModel {
	var user model.UserModel
	err := u.modelMyDB().Where("id = ?", id).First(&user)
	if err.Error != nil {
		return nil
	}
	return &user
}

// SelectByAccountAndPsw
//
//	@Description: 根据账号密码查询用户
//	@receiver u userRepo
//	@param account string 账号
//	@param password string 密码
//	@return *model.UserModel 用户数据
func (u *userRepo) SelectByAccountAndPsw(account, password string) *model.UserModel {
	var user model.UserModel
	err := u.modelMyDB().Where("account = ? and password = ?", account, encrypt.HashPsw(password)).First(&user)
	if err.Error != nil {
		return nil
	}
	return &user
}

// SelectAll
//
//	@Description: 获取所有用户
//	@receiver u userRepo
//	@return []model.UserModel 用户列表
func (u *userRepo) SelectAll() []model.UserModel {
	var users []model.UserModel
	u.modelMyDB().Find(&users)
	return users
}

// Insert
//
//	@Description: 插入用户(密码加密)
//	@receiver u userRepo
//	@param user *model.UserModel 用户数据
//	@return error 错误信息
func (u *userRepo) Insert(user model.UserModel) error {
	// 密码加密
	user.Password = encrypt.HashPsw(user.Password)
	return u.modelMyDB().Create(&user).Error
}

// UpdateSessionID
//
//	@Description: 更新用户sessionID
//	@receiver u userRepo
//	@param id int64 用户id
//	@param sessionID string sessionID
//	@return error 错误信息
func (u *userRepo) UpdateSessionID(id int64, sessionID string) error {
	return u.modelMyDB().Where("id = ?", id).Update("session_id", sessionID).Error
}

// UpdatePassword
//
//	@Description: 更新用户密码
//	@receiver u userRepo
//	@param id int64 用户id
//	@param newPsw string 新密码
//	@return error 错误信息
func (u *userRepo) UpdatePassword(id int64, newPsw string) error {
	return u.modelMyDB().Where("id = ?", id).Update("password", encrypt.HashPsw(newPsw)).Error
}

// UpdateOnlineStatus
//
//	@Description: 更新用户在线状态
//	@receiver u userRepo
//	@param id int64 用户id
//	@param onlineStatus pkg.OnlineStatus 在线状态
//	@return error 错误信息
func (u *userRepo) UpdateOnlineStatus(id int64, onlineStatus int) error {
	return u.modelMyDB().Where("id = ?", id).Update("online_status", onlineStatus).Error
}

// UpdateAvatar
//
//	@Description: 更新用户头像
//	@receiver u userRepo
//	@param avatar mongodb.UserAvatarModel
//	@return error 错误信息
func (u *userRepo) UpdateAvatar(id int64, avatarUrl string) error {
	return u.modelMyDB().Where("id = ?", id).Update("avatar", avatarUrl).Error
}

// UpdateProfile
//
//	@Description: 更新用户资料
//	@receiver u userRepo
//	@param id 用户id
//	@param nickname 昵称
//	@param sex 性别
//	@param signature 签名
//	@return error 错误信息
func (u *userRepo) UpdateProfile(id int64, nickname string, sex int) error {
	return u.modelMyDB().Where("id = ?", id).Updates(&model.UserModel{
		Nickname: nickname,
		Sex:      sex,
	}).Error
}

// SetUserOnline 设置用户在线状态
func (u *userRepo) SetUserOnline(userID int64, sessionID string, activityID uint) error {
	// 将用户 ID 添加到 online_users 集合中
	if err := u.Redis.SAdd(onlineUsersKey, userID).Err(); err != nil {
		return err
	}

	// 构建用户状态信息的键名
	userStatusKey := fmt.Sprintf(userStatusKey, userID)

	// 在 user:{userID}:status 中存储用户的详细信息
	userStatus := map[string]interface{}{
		"user_id":     userID,
		"session_id":  sessionID,
		"login_at":    time.Now().Format("2006-01-02 15:04:05"),
		"activity_id": activityID,
	}
	if err := u.Redis.HMSet(userStatusKey, userStatus).Err(); err != nil {
		return err
	}

	// 设置过期时间为 10s
	if err := u.Redis.Expire(userStatusKey, 10*time.Second).Err(); err != nil {
		return err
	}

	return nil
}

// SetUserOffline 设置用户离线状态
func (u *userRepo) SetUserOffline(userID int64) error {
	// 将用户 ID 从 online_users 集合中移除
	if err := u.Redis.SRem(onlineUsersKey, userID).Err(); err != nil {
		return err
	}

	// 删除 user:{userID}:status
	userStatusKey := fmt.Sprintf(userStatusKey, userID)
	if err := u.Redis.Del(userStatusKey).Err(); err != nil {
		return err
	}

	return nil
}

// IsUserOnline 判断用户是否在线
func (u *userRepo) IsUserOnline(userID int64) (bool, error) {
	// 检查用户 ID 是否在 online_users 集合中
	isMember, err := u.Redis.SIsMember(onlineUsersKey, userID).Result()
	if err != nil {
		return false, err
	}

	return isMember, nil
}

// GetOnlineUsers 获取所有在线用户
func (u *userRepo) GetOnlineUsers() ([]int64, error) {
	// 获取 online_users 集合中的所有用户 ID
	userIDs, err := u.Redis.SMembers(onlineUsersKey).Result()
	if err != nil {
		return nil, err
	}

	// 将字符串类型的用户 ID 转换为 uint 类型
	var onlineUserIDs []int64
	for _, userIDStr := range userIDs {
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			return nil, err
		}
		onlineUserIDs = append(onlineUserIDs, userID)
	}

	return onlineUserIDs, nil
}

// GetUserStatus 获取用户状态信息
func (u *userRepo) GetUserStatus(userID int64) (map[string]string, error) {
	// 获取 user:{userID}:status 中的用户详细信息
	userStatusKey := fmt.Sprintf(userStatusKey, userID)
	userStatus, err := u.Redis.HGetAll(userStatusKey).Result()
	if err != nil {
		return nil, err
	}

	return userStatus, nil
}

// UpdateUserHeartbeat 更新用户心跳(延长用户在线状态)
func (u *userRepo) UpdateUserHeartbeat(userID int64) error {
	// 更新 user:{userID}:status 的过期时间
	userStatusKey := fmt.Sprintf(userStatusKey, userID)
	if err := u.Redis.Expire(userStatusKey, 5*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}
