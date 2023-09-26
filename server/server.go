package server

import (
	"dvs_api/config"
	"dvs_api/routers"
	"fmt"
	"net/http"
	"time"
)

// Run 启动gin
func Run() {

	cfg := config.Global.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	server := &http.Server{
		Addr:         addr,
		Handler:      routers.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("start http server listening %s", addr)
	server.ListenAndServe()
}
