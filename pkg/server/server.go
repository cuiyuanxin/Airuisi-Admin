package server

import (
	"context"
	"github.com/cuiyuanxin/airuisi-admin/global"
	"github.com/gin-gonic/gin"
	"github.com/zztroot/color"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(route *gin.Engine) {
	gin.SetMode(global.ServerSetting.Mode)

	s := &http.Server{
		Addr:           global.ServerSetting.Addr + ":" + global.ServerSetting.Port,
		Handler:        route,
		ReadTimeout:    global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err：%s", color.Coat("requires at least one arg", color.Red))
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal)
	// 接受syscall.SIGINT和syscall.SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting tdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", color.Coat("requires at least one arg", color.Red))
	}
	log.Println("Server exiting")
}
