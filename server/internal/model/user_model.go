package model

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"server/internal/config"
	myredis "server/internal/db/redis"
)

// UserModel
//
//	@Description: 用户表结构
type UserModel struct {
	BaseModel
	Account  string `gorm:"column:account;type:varchar(255);not null;unique" json:"account"` // 账号
	Password string `gorm:"column:password;type:varchar(255);not null" json:"-"`             // 密码
	Nickname string `gorm:"column:nickname;type:varchar(255);not null" json:"nickname"`      // 姓名
	Avatar   string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`                  // 头像url
	Sex      int    `gorm:"column:sex;type:int;default -1" json:"sex"`                       // 性别1:男 2:女 -1:保密 避免零值
	//Signature    string `gorm:"column:signature;type:varchar(255);" json:"signature"`            // 个性签名
	OnlineStatus uint   `gorm:"-" json:"online_status"`                                // 在线状态(不映射进数据库) 0为离线1为在线
	SessionID    string `gorm:"column:session_id;type:varchar(255)" json:"session_id"` // session_id
}

func (*UserModel) TableName() string {
	return "user"
}

// AfterFind 钩子函数, 在查询到数据后执行数据补充
func (u *UserModel) AfterFind(*gorm.DB) (err error) {
	if u.Avatar != "" {
		u.Avatar = config.GetConfig().Server.FileServer.StaticURL + u.Avatar
	}
	//  从 Redis 中获取用户的在线状态
	onlineStatus, err := getOnlineStatus(u.ID)
	if err != nil {
		return err
	}

	// 写入 OnlineStatus 字段
	u.OnlineStatus = onlineStatus
	return nil
}

// 从 Redis 中获取用户的在线状态
func getOnlineStatus(userID int64) (uint, error) {
	// Redis 中存储用户在线状态的键为 user:{userID}:status
	userStatusKey := fmt.Sprintf("user:%d:status", userID)

	// 从 Redis 中获取用户状态信息
	status, err := myredis.GetRedisConnection().HGet(userStatusKey, "status").Int()
	if err != nil {
		// 如果 Redis 中没有该用户的状态信息，默认返回离线状态
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, err
	}

	return uint(status), nil
}
