package main

import (
	"fmt"
	"go.gin/configs"
	"go.gin/router"
	"log"
	"net/http"
	"time"
)

func main() {
	server, _ := NewHTTPServer()
	log.Printf("[info] start http server listening %s", configs.Get().Server.HttpPort)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("[error] start http server error")
	}
}

func NewHTTPServer() (*http.Server, error) {
	serverConfig := configs.Get().Server
	maxHeaderBytes := 1 << 20
	_router := router.InitRouter()
	endPoint := fmt.Sprintf(":%d", serverConfig.HttpPort)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        _router,
		ReadTimeout:    serverConfig.ReadTimeout * time.Second,
		WriteTimeout:   serverConfig.WriteTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	return server, nil
}
