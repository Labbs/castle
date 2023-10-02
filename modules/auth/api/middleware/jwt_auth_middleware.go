package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/auth/internal/tokenutil"
)

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == "/api/auth/login" || c.Path() == "/api/auth/refresh" {
			return c.Next()
		}
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing authorization header",
			})
		}
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authorized, err := tokenutil.IsAuthorized(t[1])
			if authorized {
				username, err := tokenutil.GetUsernameFromToken(t[1])
				if err != nil {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"message": "Invalid authorization token",
					})
				}

				c.Context().SetUserValue("username", username)
				return c.Next()
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid authorization header",
		})
	}
}
