package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RoleMiddleware(allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleInterface := c.Locals("role")
		role, ok := roleInterface.(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Invalid role format"})
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}
}
