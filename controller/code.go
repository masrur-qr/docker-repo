package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context)  {
	fmt.Println("users list")
	c.JSON(200,"Users list")
}
func Add(c *gin.Context)  {
	fmt.Println("users list")
	c.JSON(200,"Users add")
}