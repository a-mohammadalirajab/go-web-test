package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	dtos "github.com/a-mohammadalirajab/go-web-test/DTOs"
	"github.com/a-mohammadalirajab/go-web-test/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTasks(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/task/", nil)
	router.ServeHTTP(w, req)

	var responseResultBody []dtos.TasksDtoResponse
	var databaseResult []dtos.TasksDtoResponse

	json.Unmarshal(w.Body.Bytes(), &responseResultBody)

	// Database fetching
	if err := databaseConnection.Model(&models.TaskModel{}).Find(&databaseResult).Error; err != nil {
		t.Error(err.Error())
	}

	// this line generate failed assertion
	// responseResultBody = append(responseResultBody, dtos.TasksDtoResponse{})

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Result().Header["Content-Type"][0])
	if reflect.DeepEqual(responseResultBody, databaseResult) == false {
		t.Errorf("Want response body: %v, Got: %v", databaseResult, responseResultBody)
	}
}
