package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"Rest-Api/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Create Router")
	router := httprouter.New()

	log.Println("Register New Handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("Start Application")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}
