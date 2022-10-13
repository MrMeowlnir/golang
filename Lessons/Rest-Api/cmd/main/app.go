package main

import (
	"net"
	"net/http"
	"time"

	"Rest-Api/internal/user"
	"Rest-Api/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create Router")
	router := httprouter.New()

	logger.Info("Register New Handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("Start Application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server listening port 0.0.0.0:1234")
	logger.Fatal(server.Serve(listener))
}
