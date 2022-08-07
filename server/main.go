package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/hexiaopi/blog-service/global"
	"github.com/hexiaopi/blog-service/internal/config"
	"github.com/hexiaopi/blog-service/internal/routers"
)

func init() {
	var err error
	global.AppConfig, err = config.NewAppConfig("./configs/app.yaml")
	if err != nil {
		log.Fatalf("init config fail %v", err)
	}
	log.Infof("config:%v", global.AppConfig)
	if err = initLogger(global.AppConfig.LogLevel); err != nil {
		log.Fatalf("init log fail err:%v", err)
	}
	global.DBEngine, err = config.NewDBEngine(&global.AppConfig.DataBase, global.AppConfig.LogLevel)
	if err != nil {
		log.Fatalf("init db engine fail err:%v", err)
	}

	global.Tracer, _, err = config.NewJaegerTracer(global.AppConfig.ServiceName, global.AppConfig.TraceAgent)
	if err != nil {
		log.Fatal("init jaeger client fail.")
	}
}

// @title Blog Service API
// @version 1.0
// @description This is a blog server restful api docs.
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 30,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Starting Server...")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exit.")
}

//日志打印初始化
func initLogger(logLevel string) error {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
