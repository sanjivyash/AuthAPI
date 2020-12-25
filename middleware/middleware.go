package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sanjivyash/AuthAPI/models/token"
)

// token authentication middleware
func Auth(c *fiber.Ctx) error {
	msg := c.Query("token")

	if msg == "" {
		return c.Status(400).JSON(map[string]string{"error": "Query parameter token is missing"})
	}

	token := token.Token{
		CreatedAt: 0,
		Message:   msg,
	}
	
	if err := token.Authenticate(); err != nil {
		return c.Status(401).JSON(map[string]string{"error": "Invalid token provided"})
	}

	return c.Next()
}
