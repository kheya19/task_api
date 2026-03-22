package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kheya19/task_api/database"
	"github.com/kheya19/task_api/handlers"
	"github.com/kheya19/task_api/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	database.ConnectDatabase()

	app := fiber.New()

	app.Get("/helath", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	api := app.Group("/api")
	tasks := api.Group("/tasks")
	tasks.Use(middleware.AutoExpire)
	tasks.Post("/", handlers.CreateTask)
	tasks.Get("/", handlers.GetTasks)
	tasks.Get("/:id", handlers.GetTaskByID)
	tasks.Put("/:id", handlers.UpdateTask)
	tasks.Delete("/:id", handlers.DeleteTask)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
