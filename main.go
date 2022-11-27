package main

import (
	"TodoAPI/configs"
	"TodoAPI/controllers"
	"TodoAPI/db"
	"TodoAPI/repositories"
	"TodoAPI/routes"
	"TodoAPI/services"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "Todo API"})
	})

	dbConnection := db.NewDBConnection()
	todoRepository := repositories.NewTodoRepository(dbConnection.Database.Collection("todos"))
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)
	todoRoute := routes.NewTodoRoute(todoController)
	todoRoute.InstallTodoRoute(app)

	log.Fatalln(app.Listen(":" + configs.GetEnvPortNumber()))
}
