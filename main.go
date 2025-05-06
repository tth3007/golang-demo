package main

import (
	"demo/controllers"
	"demo/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	config := utils.LoadConfigs()
	app := application()
	app.Run(config.ServerAddress)
}

func application() *gin.Engine {
	appRouter := gin.Default()
	client := appRouter.Group("/api")
	{
		client.GET("/home", controllers.Home)
	}

	return appRouter
}
