package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lemoba/go-sweet/framework/gin"
	"github.com/lemoba/go-sweet/route"
)

func main() {
	//core := framework.NewCore()
	core := gin.New()

	//core.Use(middleware.Recovery())
	core.Use(gin.Recovery())

	route.RegisterRouter(core)

	server := &http.Server{
		Addr:    ":8888",
		Handler: core,
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)

	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatalln("Server Shutdown: ", err)
	}
}
