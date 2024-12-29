package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

// @Description: viper配置对象
var v *viper.Viper

var ConfigData = &Config{}

type (
	Config struct {
		Server   ServerConf   `json:"server"`
		DataBase DataBaseConf `json:"database"`
	}
	ServerConf struct {
		Port       int            `json:"port"`
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
		StaticDNS string `json:"static_dns"`
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

// LoadConfig
//
//	@Description: 初始化配置
func LoadConfig() *Config {
	v = viper.New()
	v.SetConfigName("config") // 配置文件的名称（不需要带后缀）
	v.SetConfigType("yaml")   // 配置文件的类型
	v.AddConfigPath("./etc/") // 查找配置文件所在的路径

	// 尝试加载配置文件
	if err := v.ReadInConfig(); err != nil {
		// 如果配置文件不存在，警告用户并使用默认值
		fmt.Println("Using default config settings...")

		v.SetDefault("server.port", "8080")  // 服务器端口
		v.SetDefault("server.mode", "debug") // 服务器日志模式
		//v.SetDefault("server.logger.path", "logs")                  // 服务器日志路径
		//v.SetDefault("server.logger.max_files", 15)                 // 服务器日志文件最大数量
		v.SetDefault("server.secret_key", "net_chat_secret_key") // 服务器密钥

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
	err := v.Unmarshal(&ConfigData)
	if err != nil {
		panic("error Unmarshal config file to config struct: " + err.Error())
	}

	fmt.Println(ConfigData.string())
	return ConfigData
}

// 获取配置,如果环境变量存在则使用环境变量，否则使用配置文件

func GetServerPort() string {
	envValue := os.Getenv("SERVER_PORT")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("server.port")
	}
}
func GetServerMode() string {
	envValue := os.Getenv("SERVER_MODE")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("server.mode")
	}
}
func GetServerLogPath() string {
	envValue := os.Getenv("SERVER_LOG_PATH")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("server.logger.path")
	}
}
func GetServerLogMaxFiles() int {
	envValue := os.Getenv("SERVER_LOG_MAX_FILES")
	if envValue != "" {
		value, _ := strconv.Atoi(envValue)
		return value
	} else {
		return v.GetInt("server.logger.max_files")
	}
}
func GetServerSecretKey() string {
	envValue := os.Getenv("SERVER_SECRET_KEY")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("server.secretKey")
	}
}
func GetDBMySQLPort() string {
	envValue := os.Getenv("MYSQL_PORT")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("database.mysql.port")
	}
}
func GetDBMySQLAccount() string {
	envValue := os.Getenv("MYSQL_ACCOUNT")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("database.mysql.account")
	}
}
func GetDBMySQLPassword() string {
	envValue := os.Getenv("MYSQL_PASSWORD")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("database.mysql.password")
	}
}
func GetDBMySQLDBName() string {
	envValue := os.Getenv("MYSQL_DBNAME")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("database.mysql.databasename")
	}
}

func GetDBMySQLHost() string {
	envValue := os.Getenv("MYSQL_HOST")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("database.mysql.host")
	}
}

func GetRedisHost() string {
	envValue := os.Getenv("REDIS_HOST")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("database.redis.host")
	}
}
func GetRedisPort() string {
	envValue := os.Getenv("REDIS_PORT")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("database.redis.port")
	}
}
func GetRedisPassword() string {
	envValue := os.Getenv("REDIS_PASSWORD")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("database.redis.password")
	}
}
func GetRedisDBName() int {
	envValue := os.Getenv("REDIS_DBNAME")
	if envValue != "" {
		value, _ := strconv.Atoi(envValue)
		return value
	} else {
		value, _ := strconv.Atoi(v.GetString("database.redis.databasename"))
		return value
	}
}

//func GetMongoDBHost() string {
//	envValue := os.Getenv("MONGODB_HOST")
//	if envValue != "" {
//		return envValue
//	} else {
//		return v.GetString("database.mongodatabase.host")
//	}
//}
//func GetMongoDBPort() string {
//	envValue := os.Getenv("MONGODB_PORT")
//	if envValue != "" {
//		return envValue
//	} else {
//		return v.GetString("database.mongodatabase.port")
//	}
//}
//func GetMongoDBAccount() string {
//	envValue := os.Getenv("MONGODB_ACCOUNT")
//	if envValue != "" {
//		return envValue
//	} else {
//		return v.GetString("database.mongodatabase.account")
//	}
//}
//func GetMongoDBPassword() string {
//	envValue := os.Getenv("MONGODB_PASSWORD")
//	if envValue != "" {
//		return envValue
//	} else {
//		return v.GetString("database.mongodatabase.password")
//	}
//}
//func GetMongoDBName() string {
//	envValue := os.Getenv("MONGODB_DBNAME")
//	if envValue != "" {
//		return envValue
//	} else {
//		return v.GetString("database.mongodatabase.databasename")
//	}
//}

func GetSMTPHost() string {
	envValue := os.Getenv("SMTP_HOST")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("smtp.host")
	}
}

func GetSMTPPort() int {
	envValue := os.Getenv("SMTP_PORT")
	if envValue != "" {
		value, _ := strconv.Atoi(envValue)
		return value
	} else {
		return v.GetInt("smtp.port")
	}
}

func GetSMTPAccount() string {
	envValue := os.Getenv("SMTP_ACCOUNT")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("smtp.account")
	}
}

func GetSMTPPassword() string {
	envValue := os.Getenv("SMTP_PASSWORD")
	if envValue != "" {
		return envValue
	} else {
		return v.GetString("smtp.password")
	}
}
