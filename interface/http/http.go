package http

import (
	"fmt"
	"net/http"

	"github.com/VanillaSkys/golang/interface/http/api"
	"github.com/gin-gonic/gin"
)

const (
	ROUTE_GET     = "GET"
	ROUTE_POST    = "POST"
	serverAddress = "localhost:8080"
)

func InitHttpServer() {
	router := gin.Default()

	router.Use(SetRequestId())
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	api.AddRouter(router)
	HttpServer := http.Server{
		Addr:    serverAddress,
		Handler: router,
	}

	if err := HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("HTTP server listen and serves failed")
	}
}
