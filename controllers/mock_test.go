package controller_test

import (
	"errors"

	dtos "github.com/a-mohammadalirajab/go-web-test/DTOs"
	"github.com/a-mohammadalirajab/go-web-test/models"
	service "github.com/a-mohammadalirajab/go-web-test/services"
)

var sampleTaskResponse dtos.TasksDtoResponse = dtos.TasksDtoResponse{
	ID:          1,
	Title:       "test",
	Description: "test",
	IsCompleted: false,
}

var sampleTaskRequest dtos.TasksDtoRequest = dtos.TasksDtoRequest{
	Title:       "test",
	Description: "test",
}

type mockTasksService struct {
	db []models.TaskModel
}

func InitMockTasksService() service.TasksService {
	return &mockTasksService{[]models.TaskModel{}}
}

var serviceInstance service.TasksService = &mockTasksService{}

// Service Spies and Mocks

var idCounter uint = 1

var createTaskSpy bool = false

func (r *mockTasksService) CreateTask(body dtos.TasksDtoRequest) error {
	if _, err := r.FindTaskByTitle(body.Title); err != nil {
		return err
	}
	createTaskSpy = true
	r.db = append(r.db, models.TaskModel{
		ID:          idCounter,
		Title:       body.Title,
		Description: body.Description,
		IsCompleted: false,
	})
	idCounter += 1
	return nil
}

var getAllTasksSpy bool = false

func (r *mockTasksService) GetAllTasks() ([]dtos.TasksDtoResponse, error) {
	getAllTasksSpy = true
	var result []dtos.TasksDtoResponse = []dtos.TasksDtoResponse{}
	for _, v := range r.db {
		result = append(result, dtos.TasksDtoResponse{
			ID:          uint8(v.ID),
			Title:       v.Title,
			Description: v.Description,
			IsCompleted: v.IsCompleted,
		})
	}
	return result, nil
}

var findTaskByTitleSpy bool = false

func (r *mockTasksService) FindTaskByTitle(title string) (dtos.TasksDtoResponse, error) {
	findTaskByTitleSpy = true
	var targetTask dtos.TasksDtoResponse = dtos.TasksDtoResponse{}
	for _, v := range r.db {
		if v.Title == title {
			targetTask = dtos.TasksDtoResponse{
				ID:          uint8(v.ID),
				Title:       v.Title,
				Description: v.Description,
				IsCompleted: v.IsCompleted,
			}
			break
		}
	}
	if len(targetTask.Title) != 0 {
		return targetTask, errors.New("10000002")
	}
	return targetTask, nil
}
