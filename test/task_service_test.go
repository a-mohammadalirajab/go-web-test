package test2

import (
	"reflect"
	"testing"

	dtos "github.com/a-mohammadalirajab/go-web-test/DTOs"
	"github.com/a-mohammadalirajab/go-web-test/models"
	service "github.com/a-mohammadalirajab/go-web-test/services"
	. "github.com/franela/goblin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbMock *gorm.DB
var taskServiceInstance service.TasksService

func init() {
	connection, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	connection.AutoMigrate(&models.TaskModel{})
	dbMock = connection
	taskServiceInstance = service.InitTasksService(dbMock)
}

func TestTasksService(t *testing.T) {
	g := Goblin(t)
	g.Describe("TaskService", func() {
		g.It("Should tasks empty.", func() {
			result, err := taskServiceInstance.GetAllTasks()
			var dbdata []models.TaskModel
			if err := dbMock.Model(&models.TaskModel{}).Find(&dbdata).Error; err != nil {
				t.Errorf("[-] Bad database condig: %s", err.Error())
			}
			g.Assert(err).IsNil()
			g.Assert(len(result)).IsZero()
			g.Assert(len(dbdata)).IsZero()
		})
		g.It("Should add new task.", func() {
			body := dtos.TasksDtoRequest{Title: "test", Description: "test"}
			err := taskServiceInstance.CreateTask(body)
			var dbdata []models.TaskModel
			if err := dbMock.Model(&models.TaskModel{}).Find(&dbdata).Error; err != nil {
				t.Errorf("[-] Database error: %s", err.Error())
			}
			g.Assert(err).IsNil()
			g.Assert(len(dbdata)).Equal(1)
			g.Assert(dbdata[0].Title).Equal(body.Title)
			g.Assert(dbdata[0].Description).Equal(body.Description)
		})
		g.It("Should find the task with title 'test'.", func() {
			result, err := taskServiceInstance.FindTaskByTitle("test")
			var dbdata dtos.TasksDtoResponse
			if err := dbMock.Model(&models.TaskModel{}).First(&dbdata, "Title = ?", "test").
				Error; err != nil {
				t.Errorf("[-] Database error: %s", err.Error())
			}
			g.Assert(err).IsNil()
			g.Assert(result.ID).IsNotZero()
			g.Assert(result.ID).IsNotZero()
			g.Assert(reflect.DeepEqual(dbdata, result)).IsTrue()
		})
	})
}
