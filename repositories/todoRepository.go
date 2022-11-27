package repositories

import (
	"TodoAPI/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// to create mock for repository
//go:generate mockgen -destination=../mocks/repositories/mock_todoRepository.go -package=repositories TodoAPI/repositories TodoRepository

type TodoRepository interface {
	PostTodo(todo models.TodoModel) (bool, error)
	GetTodo() ([]models.TodoModel, error)
	DeleteTodo(id primitive.ObjectID) (bool, error)
}

type todoRepository struct {
	todoCollection *mongo.Collection
}

func NewTodoRepository(todoCollection *mongo.Collection) TodoRepository {
	return &todoRepository{todoCollection: todoCollection}
}

func (r todoRepository) PostTodo(todoModel models.TodoModel) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todoModel.Id = primitive.NewObjectID()
	result, err := r.todoCollection.InsertOne(ctx, todoModel)
	if result.InsertedID == nil || err != nil {
		return false, err
	}

	return true, nil
}

func (r todoRepository) GetTodo() ([]models.TodoModel, error) {
	var todoModel models.TodoModel
	var todoModels []models.TodoModel

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.todoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todoModel); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		todoModels = append(todoModels, todoModel)
	}
	return todoModels, nil
}

func (r todoRepository) DeleteTodo(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.todoCollection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}

	return true, nil
}
