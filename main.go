package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wangming1993/gohttp/controller"
	"github.com/wangming1993/gohttp/middleware"
)

func main() {
	router := gin.Default()

	router.Use(middleware.AccessLog)

	router.GET("/home", controller.Home)

	router.Run(":9091")
}
