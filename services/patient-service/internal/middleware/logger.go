// internal/middleware/logger.go
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} ${latency}\n",
		CustomTags: map[string]logger.LogFunc{
			"user_id": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				if userID := c.Locals("userID"); userID != nil {
					return output.WriteString(userID.(string))
				}
				return output.WriteString("-")
			},
		},
	})
}
