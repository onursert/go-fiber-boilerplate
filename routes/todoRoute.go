package routes

import (
	"TodoAPI/controllers"

	"github.com/gofiber/fiber/v2"
)

type TodoRoute interface {
	InstallTodoRoute(app *fiber.App)
}

type todoRoute struct {
	todoController controllers.TodoController
}

func NewTodoRoute(todoController controllers.TodoController) TodoRoute {
	return &todoRoute{todoController: todoController}
}

func (r *todoRoute) InstallTodoRoute(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/todos", r.todoController.PostTodo)
	v1.Get("/todos", r.todoController.GetTodo)
	v1.Delete("/todos/:id", r.todoController.DeleteTodo)
}
