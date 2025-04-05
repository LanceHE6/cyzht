package db

import (
	goredis "github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"server/internal/config"
	"server/internal/db/mysql"
	"server/internal/db/redis"
	"sync"
)

type DBConn struct {
	MySQLConn *gorm.DB        // mysql连接
	RedisConn *goredis.Client // redis连接
}

// dbConn 单例
var dbConn *DBConn
var once sync.Once

// InitDBConn 初始化数据库连接
func InitDBConn(c *config.Config) *DBConn {
	if dbConn == nil {
		once.Do(func() {
			// 显式初始化顺序,确保在mysql初始化之前redis已经初始化
			// 用于确保gorm钩子函数中redis操作不会报错
			r := redis.InitRedisConn(c)
			m := mysql.InitMySQLConn(c)
			dbConn = &DBConn{
				MySQLConn: m,
				RedisConn: r,
			}
		})
	}
	return dbConn
}

// GetDBConn 获取数据库连接
func GetDBConn() *DBConn {
	return dbConn
}
