package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"yukboard/utils"
)

func Auth(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{"error": "Tidak terautentikasi"})
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	userID, err := utils.ParseToken(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Token tidak valid"})
	}

	c.Locals("userID", userID)
	return c.Next()
}
