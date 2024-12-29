package db

import (
	goredis "github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"server/internal/config"
	"server/internal/data/db/mysql"
	"server/internal/data/db/redis"
)

type DBConn struct {
	MySQLConn *gorm.DB
	RedisConn *goredis.Client
}

// InitDBConn 初始化数据库连接
func InitDBConn(c *config.Config) *DBConn {
	return &DBConn{
		MySQLConn: mysql.InitMySQLConn(c),
		RedisConn: redis.InitRedisConn(c),
	}
}
