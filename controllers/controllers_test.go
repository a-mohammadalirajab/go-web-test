package controller_test

import (
	"reflect"

	dtos "github.com/a-mohammadalirajab/go-web-test/DTOs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Controllers", func() {
	Describe("Try to create new tasks and delete or get list of them.", func() {
		BeforeEach(func() {
			createTaskSpy = false
			getAllTasksSpy = false
			findTaskByTitleSpy = false
		})
		Context("Try to create new tasks.", func() {
			It("Should check duplicated task and create a new task when it does not exists.", func() {
				err := serviceInstance.CreateTask(sampleTaskRequest)
				Expect(err).To(BeNil())
				Expect(findTaskByTitleSpy).To(Equal(true))
				Expect(createTaskSpy).To(Equal(true))
			})
		})
		Context("Should return already exist error.", func() {
			It("Should check duplicated task and create a new task when it does not exists.", func() {
				err := serviceInstance.CreateTask(sampleTaskRequest)
				Expect(err).NotTo(BeNil())
				Expect(findTaskByTitleSpy).To(Equal(true))
				Expect(createTaskSpy).To(Equal(false))
			})
		})
		Context("Get list of tasks.", func() {
			It("GetTasks service have to returns tasks list.", func() {
				tasks, err := serviceInstance.GetAllTasks()
				Expect(err).To(BeNil())
				Expect(findTaskByTitleSpy).To(Equal(false))
				Expect(getAllTasksSpy).To(Equal(true))
				Expect(len(tasks)).To(Equal(1))
				var outputInstance []dtos.TasksDtoResponse = append(
					[]dtos.TasksDtoResponse{}, sampleTaskResponse,
				)
				Expect(reflect.DeepEqual(outputInstance, tasks)).To(Equal(true))
			})
		})
	})
})
