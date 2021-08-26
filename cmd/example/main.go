package main

import (
	"context"
	"errors"
	"flag"
	"github.com/zs368/gin-example/internal/router"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zs368/gin-example/pkg/log"
)

var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs", "config path, eg: -conf config.yaml")
}

// @title gin-example
// @version 0.2.x
func main() {
	flag.Parse()
	logger, err := log.CustomLogger()
	if err != nil {
		panic(err)
	}

	logger.Info("test")
	if err != nil {
		panic(err)
	}

	_, err = initApp()
	if err != nil {
		panic(err)
	}



	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		os.Exit(1)
	}

	logger.Info("Server exiting")
}
