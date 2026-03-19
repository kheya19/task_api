package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kheya19/task_api/database"
	"github.com/kheya19/task_api/model"
)

func GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var task model.Task
	result := database.DB.Where("id = ?", id).First(&task)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}
	return c.JSON(task)
}
