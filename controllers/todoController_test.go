package controllers

import (
	"TodoAPI/mocks/services"
	"TodoAPI/models"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockService *services.MockTodoService
var tc TodoController

var MockData = []models.TodoModel{
	{Id: primitive.NewObjectID(), Title: "Title 1", Content: "Content 1"},
	{Id: primitive.NewObjectID(), Title: "Title 2", Content: "Content 2"},
	{Id: primitive.NewObjectID(), Title: "Title 3", Content: "Content 3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)

	mockService = services.NewMockTodoService(ct)
	tc = NewTodoController(mockService)
	return func() {
		defer ct.Finish()
	}
}

func TestTodoRepository_GetTodo(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	router := fiber.New()
	router.Get("/api/v1/todos", tc.GetTodo)

	mockService.EXPECT().GetTodo().Return(MockData, nil)
	req := httptest.NewRequest("GET", "/api/v1/todos", nil)
	res, _ := router.Test(req, 1)

	assert.Equal(t, 200, res.StatusCode)
}
