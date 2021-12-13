package service_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	dtos "github.com/a-mohammadalirajab/go-web-test/DTOs"
	"github.com/a-mohammadalirajab/go-web-test/models"
	service "github.com/a-mohammadalirajab/go-web-test/services"
)

var dbMock *gorm.DB
var taskServiceInstance service.TasksService

var _ = Describe("GoWebTest", func() {
	BeforeSuite(func() {
		connection, err := gorm.Open(sqlite.Open("mock.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		connection.AutoMigrate(&models.TaskModel{})
		dbMock = connection
		taskServiceInstance = service.InitTasksService(dbMock)
	})
	Describe("Create a task and read all of them.", func() {
		sampleTask := dtos.TasksDtoRequest{Title: "test", Description: "test"}
		Context("Add a new task when it does not exist.", func() {
			It("Check if it Exists (Should be not)", func() {
				result, err := taskServiceInstance.FindTaskByTitle(sampleTask.Title)
				var dbdata dtos.TasksDtoResponse
				if err := dbMock.Model(&models.TaskModel{}).First(&dbdata, "Title = ?", sampleTask.Title).
					Error; err != nil {
					Expect(err.Error()).To(Equal("record not found"))
				}
				Expect(err).To(BeNil())
				Expect(result.ID).To(BeZero())
			})
			It("Should add new task.", func() {
				err := taskServiceInstance.CreateTask(sampleTask)
				Expect(err).To(BeNil())
			})
		})
		Context("Get tasks list with that task.", func() {
			It("Check if it Exists (Should be)", func() {
				result, err := taskServiceInstance.FindTaskByTitle(sampleTask.Title)
				var dbdata dtos.TasksDtoResponse
				if err := dbMock.Model(&models.TaskModel{}).First(&dbdata, "Title = ?", sampleTask.Title).
					Error; err != nil {
					Expect(err.Error()).NotTo(Equal("record not found"))
				}
				Expect(err).To(BeNil())
				Expect(result.ID).NotTo(BeZero())
			})
			It("Check tasks list that contains only one task.", func() {
				result, err := taskServiceInstance.GetAllTasks()
				var dbdata []dtos.TasksDtoResponse
				if err := dbMock.Model(&[]models.TaskModel{}).Find(&dbdata).Error; err != nil {
					Expect(err).To(BeNil())
				}
				Expect(err).To(BeNil())
				Expect(len(dbdata)).To(Equal(1))
				Expect(reflect.DeepEqual(result, dbdata)).To(Equal(true))
			})
		})
	})
})
