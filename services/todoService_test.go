package services

import (
	"TodoAPI/mocks/repositories"
	"TodoAPI/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockTodoRepository *repositories.MockTodoRepository
var ts TodoService

var MockData = []models.TodoModel{
	{Id: primitive.NewObjectID(), Title: "Title 1", Content: "Content 1"},
	{Id: primitive.NewObjectID(), Title: "Title 2", Content: "Content 2"},
	{Id: primitive.NewObjectID(), Title: "Title 3", Content: "Content 3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockTodoRepository = repositories.NewMockTodoRepository(ct)
	ts = NewTodoService(mockTodoRepository)
	return func() {
		defer ct.Finish()
	}
}

func TestTodoService_GetTodo(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockTodoRepository.EXPECT().GetTodo().Return(MockData, nil)
	result, err := ts.GetTodo()
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}
