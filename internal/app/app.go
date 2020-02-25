package app

import (
	"context"
	"os"

	"github.com/souxls/dvs_api/internal/app/bll/impl"
	"github.com/souxls/dvs_api/internal/app/config"
	"github.com/souxls/dvs_api/pkg/logger"
	"go.uber.org/dig"
)

type options struct {
	ConfigFile string
	Version    string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Init 初始化应用
func Init(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobal(o.ConfigFile)
	handleError(err)

	//	cfg := config.Global()

	logger.Printf(ctx, "服务启动，运行模式：%s，版本号：%s，进程号：%d", o.Version, os.Getpid())

	handleError(err)

	loggerCall, err := InitLogger()
	handleError(err)

	// 创建依赖注入容器
	container, containerCall := BuildContainer()

	httpCall := InitHTTPServer(ctx, container)
	return func() {
		if httpCall != nil {
			httpCall()
		}
		if containerCall != nil {
			containerCall()
		}
		if loggerCall != nil {
			loggerCall()
		}
	}
}

// BuildContainer 创建依赖注入容器
func BuildContainer() (*dig.Container, func()) {
	// 创建依赖注入容器
	container := dig.New()

	// 注入存储模块
	storeCall, err := InitStore(container)
	handleError(err)

	// 注入bll
	err = impl.Inject(container)
	handleError(err)

	return container, func() {
		// 释放资源
		//		ReleaseCasbinEnforcer(container)

		if storeCall != nil {
			storeCall()
		}
	}
}
