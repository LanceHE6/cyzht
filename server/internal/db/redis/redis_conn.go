package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"server/internal/config"
	"server/pkg/logger"
)

var rdb *redis.Client

func InitRedisConn(c *config.Config) *redis.Client {
	// 连接Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.DataBase.Redis.Host, c.DataBase.Redis.Port),
		Password: c.DataBase.Redis.Password,
		DB:       c.DataBase.Redis.Database, // 默认数据库
		PoolSize: 20000,                     // 连接池大小
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		logger.Logger.Error("can not connect redis: " + err.Error())
		os.Exit(-3)
	}
	logger.Logger.Info("connect redis successfully")
	return rdb
}

func GetRedisConnection() *redis.Client {
	return rdb
}
