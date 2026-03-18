package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kheya19/task_api/database"
	"github.com/kheya19/task_api/model"
)

func CreateTask(c *fiber.Ctx) error {
	var req model.CreateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if req.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
	}
	if req.ExpiresAt == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ExpiresAt is required"})
	}
	task := model.Task{
		Title:       req.Title,
		Description: req.Description,
		ExpiresAt:   req.ExpiresAt,
	}
	result := database.DB.Create(&task)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create task"})
	}
	return c.Status(fiber.StatusCreated).JSON(task)
}
