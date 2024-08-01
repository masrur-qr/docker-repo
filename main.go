package main

import (
	"test/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/list",controller.List)
	r.POST("/add",controller.Add)

	r.Run(":6765")
}