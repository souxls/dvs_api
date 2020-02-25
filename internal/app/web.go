package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/souxls/dvs_api/internal/app/config"
	"github.com/souxls/dvs_api/internal/app/middleware"
	"github.com/souxls/dvs_api/internal/app/routers/api"
	"github.com/souxls/dvs_api/pkg/logger"
	"go.uber.org/dig"
)

// InitWeb 初始化web引擎
func InitWeb(container *dig.Container) *gin.Engine {
	cfg := config.Global()

	app := gin.New()

	// 跨域请求
	if cfg.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	// 注册/api路由
	err := api.RegisterRouter(app, container)
	handleError(err)

	return app
}

// InitHTTPServer 初始化http服务
func InitHTTPServer(ctx context.Context, container *dig.Container) func() {
	cfg := config.Global().HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      InitWeb(container),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.Printf(ctx, "HTTP服务开始启动，地址监听在：[%s]", addr)
		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			logger.Errorf(ctx, err.Error())
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf(ctx, err.Error())
		}
	}
}
