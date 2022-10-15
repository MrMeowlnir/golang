package main

import (
	"Rest-Api/internal/config"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"Rest-Api/internal/user"
	"Rest-Api/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create Router")
	router := httprouter.New()

	cfg := config.GetConfig()

	logger.Info("Register New Handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("Start Application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		logger.Info("Detect App Path")
		cwd, err := os.Getwd()
		if err != nil{
			logger.Fatal(err)
		}
		appDir, err := filepath.Abs(cwd)
		if err != nil{
			logger.Fatal(err)
		}
		logger.Info("Create Socket")
		socketPath := path.Join(appDir, "app.sock")
		logger.Debugf("Socket Path: %s", socketPath)

		logger.Info("Listen Unix Socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("Server is Listening Unix Socket: %s", socketPath)

	} else {
		logger.Info("Listen TCP")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("Server Listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
