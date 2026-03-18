package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kheya19/task_api/database"
	"github.com/kheya19/task_api/model"
)

func GetTasks(c *fiber.Ctx) error {
	var tasks []model.Task
	query := database.DB.Model(&model.Task{})

	status := c.Query("status")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	createdFrom := c.Query("createdFrom")
	if createdFrom != "" {
		val, err := strconv.ParseInt(createdFrom, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid created from value",
			})
		}
		query = query.Where("createdAt >= to_timestamp(?)", val)
	}
	createdTo := c.Query("createdTo")
	if createdTo != "" {
		val, err := strconv.ParseInt(createdTo, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid created to value",
			})
		}
		query = query.Where("createdAt <= to_timestamp(?)", val)
	}
	expiresFrom := c.Query("expiresFrom")
	if expiresFrom != "" {
		val, err := strconv.ParseInt(expiresFrom, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid expires from value",
			})
		}
		query = query.Where("expiresAt >= to_timestamp(?)", val)
	}
	expiresTo := c.Query("expiresTo")
	if expiresTo != "" {
		val, err := strconv.ParseInt(expiresTo, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid expires to value",
			})
		}
		query = query.Where("expiresAt <= to_timestamp(?)", val)
	}
	result := query.Find(&tasks)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch tasks",
		})
	}

	return c.JSON(fiber.Map{
		"count": len(tasks),
		"tasks": tasks,
	})

}
