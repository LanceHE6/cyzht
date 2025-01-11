package config

import (
	"fmt"
	"github.com/spf13/viper"
	"server/pkg/logger"
	"strings"
	"sync"
)

// @Description: viper配置对象
var v *viper.Viper

var c *Config      // 全局配置对象
var once sync.Once // 只执行一次配置加载

type (
	Config struct {
		Server   ServerConf   `json:"server"`
		DataBase DataBaseConf `json:"database"`
	}
	ServerConf struct {
		Port       string         `json:"port"`
		Mode       string         `json:"mode"`
		JWTSecret  string         `json:"jwt_secret"` // jwt加密密钥
		SMTP       SMTPConf       `json:"smtp"`
		FileServer FileServerConf `json:"file_server"`
	}

	SMTPConf struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Account  string `json:"account"`
		Password string `json:"password"`
	}
	FileServerConf struct {
		RpcDNS    string `json:"rpc_dns"`
		StaticURL string `json:"static_url"`
	}

	DataBaseConf struct {
		MySQL MySQLConf `json:"mysql"`
		Redis RedisConf `json:"redis"`
		//MongoDB MongoDBConf `json:"mongodb"`
	}

	MySQLConf struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		UserName string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
	}
	RedisConf struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
		Database int    `json:"database"`
	}
	//MongoDBConf struct {
	//	Host     string `json:"host"`
	//	Port     string `json:"port"`
	//	UserName string `json:"username"`
	//	Password string `json:"password"`
	//	Database string `json:"database"`
	//}
)

func (c *Config) string() string {
	return fmt.Sprintf(
		"ServerConf: %+v\n", c,
	)
}

// GetConfig 获取全局配置变量
func GetConfig() *Config {
	once.Do(func() {
		initConfig()
	})
	return c
}

// initConfig
//
//	@Description: 初始化配置
func initConfig() *Config {
	v = viper.New()
	v.SetConfigName("config")                 // 配置文件的名称（不需要带后缀）
	v.SetConfigType("yaml")                   // 配置文件的类型
	v.AddConfigPath("./etc/")                 // 查找配置文件所在的路径
	v.AutomaticEnv()                          // 绑定环境变量
	replacer := strings.NewReplacer(".", "_") // 替换环境变量中的点为下划线
	v.SetEnvKeyReplacer(replacer)             // 设置环境变量替换器

	// 尝试加载配置文件
	if err := v.ReadInConfig(); err != nil {
		// 如果配置文件不存在，警告用户并使用默认值
		fmt.Println("Using default config settings...")

		v.SetDefault("server.port", "8080")  // 服务器端口
		v.SetDefault("server.mode", "debug") // 服务器日志模式
		//v.SetDefault("server.logger.path", "logs")                  // 服务器日志路径
		//v.SetDefault("server.logger.max_files", 15)                 // 服务器日志文件最大数量
		v.SetDefault("server.secret_key", "secret_key")                        // 服务器密钥
		v.SetDefault("server.file_server.rpc_dns", "dns:///127.0.0.1:5173")    // 文件rpc服务器地址
		v.SetDefault("server.file_server.static_url", "127.0.0.1:5174/static") // 文件静态资源地址

		v.SetDefault("database.mysql.host", "localhost")    // mysql数据库地址
		v.SetDefault("database.mysql.port", "3306")         // mysql数据库端口
		v.SetDefault("database.mysql.account", "root")      // mysql数据库账号
		v.SetDefault("database.mysql.password", "root")     // mysql数据库密码
		v.SetDefault("database.mysql.database", "net_chat") // mysql数据库名称

		v.SetDefault("database.redis.host", "localhost")  // redis数据库地址
		v.SetDefault("database.redis.port", "6379")       // redis数据库端口
		v.SetDefault("database.redis.password", "123456") // redis数据库密码
		v.SetDefault("database.redis.database", 0)        // redis数据库名称

		//v.SetDefault("database.mongodatabase.host", "localhost")  // mongodatabase数据库地址
		//v.SetDefault("database.mongodatabase.port", "27017")      // mongodatabase数据库端口
		//v.SetDefault("database.mongodatabase.account", "root")    // mongodatabase数据库账号
		//v.SetDefault("database.mongodatabase.password", "123456") // mongodatabase数据库密码
		//v.SetDefault("database.mongodatabase.database", 0)          // mongodatabase数据库名称

		v.SetDefault("smtp.host", "smtp.qq.com")      // 邮箱服务器地址
		v.SetDefault("smtp.port", "465")              // 邮箱服务器端口
		v.SetDefault("smtp.account", "123456@qq.com") // 邮箱服务器账号
		v.SetDefault("smtp.password", "123456")       // 邮箱服务器密码

		// 选择在此处写入默认配置文件或者不写
		err := v.WriteConfigAs("config.yaml")
		if err != nil {
			fmt.Println("error writing default config file: " + err.Error())
		}
	}
	err := v.Unmarshal(&c)
	if err != nil {
		panic("error Unmarshal config file to config struct: " + err.Error())
	}

	logger.Logger.Debugf("server with config: %+v", c)
	return c
}
