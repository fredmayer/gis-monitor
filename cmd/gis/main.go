package main

import (
	"fmt"
	"gis-crawler/internal/app"
	"gis-crawler/internal/config"
	"gis-crawler/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	run()
}

func run() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	logging.Init(cfg.LogLevel)

	App := app.New(cfg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	//todo gracefully shutdown
	fmt.Printf("shutdown... \n")
	App.Stop()
}
