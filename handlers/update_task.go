package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kheya19/task_api/database"
	"github.com/kheya19/task_api/model"
)

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task model.Task
	result := database.DB.Where("id = ?", id).First(&task)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}
	var req model.UpdateTaskRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.Status != "" {
		task.Status = req.Status
	}
	if req.ExpiresAt != 0 {
		task.ExpiresAt = req.ExpiresAt
	}
	result = database.DB.Save(&task)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to update task",
		})
	}
	return c.JSON(task)
}
