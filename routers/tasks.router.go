package routers

import (
	controller "github.com/a-mohammadalirajab/go-web-test/controllers"
	"github.com/gin-gonic/gin"
)

func InitTasksRouter(router *gin.Engine, controllerInstance controller.TasksController) {
	tasksRouter := router.Group("/task")
	{
		tasksRouter.GET("/", controllerInstance.GetTasks)
		tasksRouter.POST("/", controllerInstance.CreateTask)
	}
}
