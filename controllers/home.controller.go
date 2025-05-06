package controllers

import (
	"github.com/gin-gonic/gin"
)

type Demo struct {
	Id      string
	Message string
}

func Home(c *gin.Context) {
	demo := Demo{"1", "Hello World"}

	c.JSON(200, demo)
}
