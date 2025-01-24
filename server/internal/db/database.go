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
			dbConn = &DBConn{
				MySQLConn: mysql.InitMySQLConn(c),
				RedisConn: redis.InitRedisConn(c),
			}
		})
	}
	return dbConn
}

// GetDBConn 获取数据库连接
func GetDBConn() *DBConn {
	return dbConn
}
