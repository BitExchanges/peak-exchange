package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"peak-exchange/config"
	"peak-exchange/routes"
	"peak-exchange/utils"
	"time"
)

func main() {

	initialize()
	gin.ForceConsoleColor()
	//记录日志
	//f, _ := os.Create("peak.log")
	//仅将日志写入文件
	//gin.DefaultWriter = io.MultiWriter(f)

	//同时将日志写入文件和控制台
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("static/*")
	routes.SetInterfaces(router)
	srv := &http.Server{
		Addr:    "192.168.0.114:8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器 (超时时间设置为5秒)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("关闭服务:", err)
	}
	log.Println("服务退出")
}

func initialize() {
	config.InitEnv()
	utils.InitMainDB()
	utils.InitRedisPools()
}
