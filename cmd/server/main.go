package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/souxls/dvs_api/internal/app"
	"github.com/souxls/dvs_api/pkg/logger"
	"github.com/souxls/dvs_api/pkg/util"
)

// VERSION 设定版本号
var VERSION = "0.0.1"

var (
	help       bool
	version    bool
	configFile string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.BoolVar(&version, "v", false, "show version and exit")
	flag.StringVar(&configFile, "c", "", "config file")

	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `dvs version: dvs/%s
Usage: dvs [-h] [-c filename]

Options:
`, VERSION)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if configFile == "" {
		panic("配置文件不能为空，请使用 -c 选项")
	}

	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := logger.NewTraceIDContext(context.Background(), util.NewTraceID())
	span := logger.StartSpanWithCall(ctx)

	call := app.Init(ctx,
		app.SetConfigFile(configFile),
		app.SetVersion(VERSION))

EXIT:
	for {
		sig := <-sc
		span().Printf("获取到信号[%s]", sig.String())

		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	if call != nil {
		call()
	}

	span().Printf("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}
