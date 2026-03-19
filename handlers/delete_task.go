package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kheya19/task_api/database"
	"github.com/kheya19/task_api/model"
)

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	var task model.Task
	result := database.DB.Where("id = ?", id).First(&task)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "task not found",
		})
	}
	result = database.DB.Delete(&task)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete task",
		})
	}
	return c.JSON(fiber.Map{
		"message": "task deleted successfully",
	})
}
