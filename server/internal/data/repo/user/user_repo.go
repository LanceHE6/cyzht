package user

import (
	"github.com/jinzhu/gorm"
	mysql2 "server/internal/data/db/mysql"
	"server/internal/data/model"
	"server/pkg"
	"server/pkg/hash"
	"strconv"
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
	SearchUsers(params pkg.SearchUsersParams) ([]model.UserModel, int)
}

// NewUserRepo
//
//	@Description: 创建用户仓库实例
//	@return UserRepoInterface 用户仓库实例
func NewUserRepo(mysqlConn *gorm.DB) UserRepoInterface {
	return &userRepo{
		MyDB: mysqlConn,
	}
}

// userRepo
//
//	@Description: 用户仓库实现
type userRepo struct {
	MyDB *gorm.DB
}

// modelMyDB
//
//	@Description: 获取用户表
//	@receiver u userRepo
//	@return *gorm.db 用户表
func (u userRepo) modelMyDB() *gorm.DB {
	return u.MyDB.Model(model.UserModel{})
}

// SelectByAccount
//
//	@Description: 根据账号查询用户
//	@receiver u userRepo
//	@param account string 账号
//	@return *model.UserModel 用户数据
func (u userRepo) SelectByAccount(account string) *model.UserModel {
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
func (u userRepo) SelectByID(id int64) *model.UserModel {
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
func (u userRepo) SelectByAccountAndPsw(account, password string) *model.UserModel {
	var user model.UserModel
	err := u.modelMyDB().Where("account = ? and password = ?", account, hash.HashPsw(password)).First(&user)
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
func (u userRepo) SelectAll() []model.UserModel {
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
func (u userRepo) Insert(user model.UserModel) error {
	// 密码加密
	user.Password = hash.HashPsw(user.Password)
	return u.modelMyDB().Create(&user).Error
}

// UpdateSessionID
//
//	@Description: 更新用户sessionID
//	@receiver u userRepo
//	@param id int64 用户id
//	@param sessionID string sessionID
//	@return error 错误信息
func (u userRepo) UpdateSessionID(id int64, sessionID string) error {
	return u.modelMyDB().Where("id = ?", id).Update("session_id", sessionID).Error
}

// UpdatePassword
//
//	@Description: 更新用户密码
//	@receiver u userRepo
//	@param id int64 用户id
//	@param newPsw string 新密码
//	@return error 错误信息
func (u userRepo) UpdatePassword(id int64, newPsw string) error {
	return u.modelMyDB().Where("id = ?", id).Update("password", hash.HashPsw(newPsw)).Error
}

func (u userRepo) SearchUsers(params pkg.SearchUsersParams) ([]model.UserModel, int) {
	query := mysql2.GetMySQLConnection().Table("users")
	if params.Account != nil {
		query = query.Where("account like ?", "%"+*params.Account+"%")
	}
	if params.Name != nil {
		query = query.Where("name like ?", "%"+*params.Name+"%")
	}
	if params.Keyword != nil {
		query = query.Where("account like ? OR name like ? ",
			"%"+*params.Keyword+"%",
			"%"+*params.Keyword+"%",
		)
	}
	if params.Direction != nil {
		// 未填写方向的
		if *params.Direction == 0 {
			query = query.Where("JSON_LENGTH(direction) = 0")
		} else // 填写了两个方向的
		if *params.Direction == 3 {
			query = query.Where("JSON_CONTAINS(direction, '1') AND JSON_CONTAINS(direction, '2')")
		} else {
			// 查询单个方向的
			query = query.Where("JSON_CONTAINS(direction, ?)", strconv.Itoa(*params.Direction))
		}
	}
	// 统计总数
	var count int
	query.Count(&count)
	// 分页
	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	} else {
		params.Limit = new(int)
		*params.Limit = 10
	}
	if params.Page != nil {
		query = query.Offset((*params.Page - 1) * *params.Limit)
	}
	var users []model.UserModel
	query.Find(&users)
	return users, count
}

// UpdateUserInfo
//
//	@Description: 更新用户信息
//	@receiver u userRepo
//	@param id int64
//	@param user model.UserModel
//	@return error 错误信息
func (u userRepo) UpdateUserInfo(id int64, user model.UserModel) error {
	return u.modelMyDB().Where("id = ?", id).Updates(user).Error
}

// UpdateOnlineStatus
//
//	@Description: 更新用户在线状态
//	@receiver u userRepo
//	@param id int64 用户id
//	@param onlineStatus pkg.OnlineStatus 在线状态
//	@return error 错误信息
func (u userRepo) UpdateOnlineStatus(id int64, onlineStatus int) error {
	return u.modelMyDB().Where("id = ?", id).Update("online_status", onlineStatus).Error
}

// UpdateAvatar
//
//	@Description: 更新用户头像
//	@receiver u userRepo
//	@param avatar mongodb.UserAvatarModel
//	@return error 错误信息
func (u userRepo) UpdateAvatar(id int64, avatarUrl string) error {
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
func (u userRepo) UpdateProfile(id int64, nickname string, sex int) error {
	return u.modelMyDB().Where("id = ?", id).Updates(&model.UserModel{
		Nickname: nickname,
		Sex:      sex,
	}).Error
}
