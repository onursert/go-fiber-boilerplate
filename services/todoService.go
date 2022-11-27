package services

import (
	"TodoAPI/models"
	"TodoAPI/repositories"
	"TodoAPI/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// to create mock for service
//go:generate mockgen -destination=../mocks/services/mock_todoService.go -package=services TodoAPI/services TodoService

type TodoService interface {
	PostTodo(todo models.TodoModel) error
	GetTodo() ([]models.TodoModel, error)
	DeleteTodo(id primitive.ObjectID) (bool, error)
}

type todoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &todoService{todoRepository: todoRepository}
}

func (s todoService) PostTodo(todoModel models.TodoModel) error {
	if len(todoModel.Title) <= 2 || len(todoModel.Content) <= 2 {
		return utils.ErrTooFewChar
	}

	result, err := s.todoRepository.PostTodo(todoModel)
	if err != nil || result == false {
		return err
	}

	return nil
}

func (s todoService) GetTodo() ([]models.TodoModel, error) {
	result, err := s.todoRepository.GetTodo()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s todoService) DeleteTodo(id primitive.ObjectID) (bool, error) {
	result, err := s.todoRepository.DeleteTodo(id)
	if err != nil || result == false {
		return false, err
	}

	return true, nil
}
