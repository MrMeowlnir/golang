package main

import (
	"Rest-Api/internal/config"
	"Rest-Api/internal/user/db"
	"Rest-Api/pkg/client/mongodb"
	"context"
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

	mongoDBClient, err := mongodb.NewCient(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port,
		cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Collection, logger)


	users, err := storage.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	user1 := user.User{
		ID:           "",
		Email:        "mrmeowlnir@valhall.la",
		Username:     "MrMeowlnir",
		PasswordHash: "123456",
	}
	user1ID, err := storage.Create(context.Background(), user1)
	if err != nil {
		panic(err)
	}
	logger.Info(user1ID)

	user2 := user.User{
		ID:           "",
		Email:        "2mrmeowlnir@valhall.la",
		Username:     "2MrMeowlnir",
		PasswordHash: "222222",
		}
	user2ID, err := storage.Create(context.Background(), user2)
	if err != nil {
		panic(err)
	}
	logger.Info(user2ID)

	users, err = storage.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	user2Found, err := storage.FindOne(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(user2Found)

	user2Found.Email = "newEmailuser2@valhall.la"
	err = storage.Update(context.Background(), user2Found)
	if err != nil {
		panic(err)
	}

	user2Found, err = storage.FindOne(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(user2Found)

	err = storage.Delete(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}

	users, err = storage.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	_, err = storage.FindOne(context.Background(), user2ID)
	if err != nil {
		panic(err)
	}

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
