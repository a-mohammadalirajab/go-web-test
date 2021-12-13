package controller

import (
	"errors"
	"net/http"

	dtos "github.com/a-mohammadalirajab/go-web-test/DTOs"
	"github.com/a-mohammadalirajab/go-web-test/lib"
	service "github.com/a-mohammadalirajab/go-web-test/services"
	"github.com/gin-gonic/gin"
)

type TasksController interface {
	GetTasks(c *gin.Context)
	CreateTask(c *gin.Context)
}

type tasksController struct {
	s service.TasksService
}

func InitTasksController(tasksService *service.TasksService) TasksController {
	return &tasksController{s: *tasksService}
}

func (tc *tasksController) GetTasks(c *gin.Context) {
	result, err := tc.s.GetAllTasks()
	if err != nil {
		code, message := lib.ErrorMaker(err)
		c.JSON(code, message)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func (tc *tasksController) CreateTask(c *gin.Context) {
	var body dtos.TasksDtoRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		code, message := lib.ErrorMaker(err)
		c.JSON(code, message)
		return
	}
	targetTask, err := tc.s.FindTaskByTitle(body.Title)
	if err != nil {
		code, message := lib.ErrorMaker(err)
		c.JSON(code, message)
		return
	}
	//already exist
	if len(targetTask.Title) != 0 {
		code, message := lib.ErrorMaker(errors.New("10000002"))
		c.JSON(code, message)
		return
	}
	if err := tc.s.CreateTask(body); err != nil {
		code, message := lib.ErrorMaker(err)
		c.JSON(code, message)
	} else {
		c.JSON(http.StatusCreated, lib.SuccessfulMessage)
	}
}
