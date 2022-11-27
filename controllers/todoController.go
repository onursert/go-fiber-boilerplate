package controllers

import (
	"TodoAPI/models"
	"TodoAPI/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoController interface {
	PostTodo(ctx *fiber.Ctx) error
	GetTodo(ctx *fiber.Ctx) error
	DeleteTodo(ctx *fiber.Ctx) error
}

type todoController struct {
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) TodoController {
	return &todoController{todoService: todoService}
}

func (c todoController) PostTodo(ctx *fiber.Ctx) error {
	var todoModel models.TodoModel

	err := ctx.BodyParser(&todoModel)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	err = c.todoService.PostTodo(todoModel)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"state": true})
}

func (c todoController) GetTodo(ctx *fiber.Ctx) error {
	result, err := c.todoService.GetTodo()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(result)
}

func (c todoController) DeleteTodo(ctx *fiber.Ctx) error {
	query := ctx.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := c.todoService.DeleteTodo(cnv)
	if err != nil || !result {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"state": false})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"state": true})
}
