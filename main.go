package main

import (
	controller "github.com/a-mohammadalirajab/go-web-test/controllers"
	"github.com/a-mohammadalirajab/go-web-test/lib"
	"github.com/a-mohammadalirajab/go-web-test/routers"
	service "github.com/a-mohammadalirajab/go-web-test/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	databaseConnection *gorm.DB                   = lib.InitDatabase()
	tasksService       service.TasksService       = service.InitTasksService(databaseConnection)
	tasksController    controller.TasksController = controller.InitTasksController(&tasksService)
)

func setupRouter() *gin.Engine {
	mainRouter := gin.Default()
	routers.InitTasksRouter(mainRouter, tasksController)
	return mainRouter
}

func main() {
	server := setupRouter()
	server.Run(":8080")
}
