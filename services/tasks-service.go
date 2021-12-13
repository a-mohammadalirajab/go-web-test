package service

import (
	"errors"

	dtos "github.com/a-mohammadalirajab/go-web-test/DTOs"
	"github.com/a-mohammadalirajab/go-web-test/models"
	"gorm.io/gorm"
)

type TasksService interface {
	CreateTask(body dtos.TasksDtoRequest) error
	GetAllTasks() ([]dtos.TasksDtoResponse, error)
	FindTaskByTitle(title string) (dtos.TasksDtoResponse, error)
}

type tasksService struct {
	db *gorm.DB
}

func InitTasksService(db *gorm.DB) TasksService {
	return &tasksService{db}
}

func (r *tasksService) GetAllTasks() ([]dtos.TasksDtoResponse, error) {
	var tasksList []dtos.TasksDtoResponse
	if err := r.db.Model(&models.TaskModel{}).Find(&tasksList).Error; err != nil {
		return nil, err
	}
	return tasksList, nil
}

func (r *tasksService) FindTaskByTitle(title string) (dtos.TasksDtoResponse, error) {
	var task dtos.TasksDtoResponse
	if err := r.db.Model(&models.TaskModel{}).First(&task, "Title = ?", title).Error; err != nil {
		if err.Error() != "record not found" {
			return dtos.TasksDtoResponse{}, errors.New("90000001")
		}
	}
	return task, nil
}

func (r *tasksService) CreateTask(body dtos.TasksDtoRequest) error {
	if err := r.db.Create(&models.TaskModel{
		Title: body.Title, Description: body.Description, IsCompleted: false,
	}).Error; err != nil {
		return errors.New("90000001")
	}
	return nil
}
