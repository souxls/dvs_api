package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	global *Config
)

// LoadGlobal 加载全局配置
func LoadGlobal(fpath string) error {
	c, err := Parse(fpath)
	if err != nil {
		return err
	}
	global = c
	return nil
}

// Global 获取全局配置
func Global() *Config {
	if global == nil {
		return &Config{}
	}
	return global
}

// Parse 解析配置文件
func Parse(fpath string) (*Config, error) {
	YamlFile, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	var c Config
	err = yaml.Unmarshal(YamlFile, &c)
	fmt.Println(c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Config 配置参数
type Config struct {
	Swagger string `yaml:"swagger"`
	HTTP    HTTP   `yaml:"http"`
	Log     Log    `yaml:"log"`
	CORS    CORS   `yaml:"cors"`
	Redis   Redis  `yaml:"redis"`
	Gorm    Gorm   `yaml:"gorm"`
	MySQL   MySQL  `yaml:"mysql"`
}

// HTTP 配置参数, 配置启动server监听端口，是否启用https
type HTTP struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	HTTPS           bool   `yaml:"https"`
	CertFile        string `yaml:"cert_file"`
	KeyFile         string `yaml:"key_file"`
	ShutdownTimeout int    `yaml:"shutdown_timeout"`
}

// Log 日志配置参数
type Log struct {
	Level   int    `yaml:"level"`
	Format  string `yaml:"format"`
	Output  string `yaml:"output"`
	LogFile string `yaml:"log_file"`
}

// CORS 跨域请求配置参数
type CORS struct {
	Enable           bool     `yaml:"enable"`
	AllowOrigins     []string `yaml:"allow_origins"`
	AllowMethods     []string `yaml:"allow_methods"`
	AllowHeaders     []string `yaml:"allow_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	MaxAge           int      `yaml:"max_age"`
}

// Redis 配置参数
type Redis struct {
	Host     string `"yaml:host"`
	db       string `"yaml:db"`
	password string `"yaml:password"`
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
	Port       string `yaml:"port"`
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
