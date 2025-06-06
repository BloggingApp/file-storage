package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BloggingApp/file-storage/internal/config"
	"github.com/BloggingApp/file-storage/internal/handler"
	"github.com/BloggingApp/file-storage/internal/server"
	"github.com/BloggingApp/file-storage/internal/service"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if err := initConfig(); err != nil {
		log.Panicf("failed to initialize yaml config: %s", err.Error())
	}

	logger, err := newLogger()
	if err != nil {
		log.Panicf("failed to create new zap logger: %s", err.Error())
	}

	services := service.New(logger)
	handlers := handler.New(services)

	mux := handlers.Init()

	srv := server.New()
	cfg := config.ServerConfig{
		Port: viper.GetString("app.port"),
		Handler: mux,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	go func(srv *server.Server, cfg config.ServerConfig) {
		if err := srv.Run(cfg); err != nil {
			log.Panicf("failed to run server: %s", err.Error())
		}
	}(srv, cfg)

	log.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Server shutting down")
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")
	return viper.ReadInConfig()
}

func newLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./app.log",
	}
	return cfg.Build()
}
