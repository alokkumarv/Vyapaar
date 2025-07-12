package user

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserHandler(c *fiber.Ctx) error {
	user := c.Body()
	log.Printf("Got user *v", string(user))
	return c.JSON(string(user))
}
