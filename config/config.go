package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	Global *Config
)

// Config 配置参数
type Config struct {
	Swagger string
	HTTP    HTTP  `yaml:"http"`
	Log     Log   `yaml:"log"`
	Redis   Redis `yaml:"redis"`
	Gorm    Gorm  `yaml:"gorm"`
	MySQL   MySQL `yaml:"mysql"`
}

// HTTP 配置参数, 配置启动server监听端口，是否启用https
type HTTP struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Log 日志配置参数
type Log struct {
	Level   int    `yaml:"level"`
	Format  string `yaml:"format"`
	Output  string `yaml:"output"`
	LogFile string `yaml:"log_file"`
}

// Redis 配置参数
type Redis struct {
	Host     string `yaml:"host"`
	db       string `yaml:"db"`
	password string `yaml:"password"`
}

// Gorm gorm配置参数
type Gorm struct {
	Debug             bool   `yaml:"debug"`
	DBType            string `yaml:"db_type"`
	MaxLifetime       int    `yaml:"max_lifetime"`
	MaxOpenConns      int    `yaml:"max_open_conns"`
	MaxIdleConns      int    `yaml:"max_idle_conns"`
	TablePrefix       string `yaml:"table_prefix"`
	EnableAutoMigrate bool   `yaml:"enable_auto_migrate"`
}

// MySQL 配置参数
type MySQL struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"db_name"`
	Parameters string `yaml:"parameters"`
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

// Init 解析配置文件，初始化变量
func Init(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name "config" with no extension.
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/dvs")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		err := viper.Unmarshal(&Global)
		if err != nil {
			fmt.Printf("parse error, %v", err)
			os.Exit(1)
		}
	}
}
