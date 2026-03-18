package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kheya19/task_api/database"
	"github.com/kheya19/task_api/model"
)

func AutoExpire(c *fiber.Ctx) error {
	now := time.Now().Unix()

	database.DB.Model(&model.Task{}).Where("expires_at <= ? AND status != ?", now, model.StatusExpired).Update("status", model.StatusExpired)
	return c.Next()
}
