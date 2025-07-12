package registration

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleRegistration(c *fiber.Ctx) error {
	var registrationData ShopkeeperRegisterRequest
	err := c.BodyParser(&registrationData)
	if err != nil {
		log.Printf("Error while parsing body : %v", err)
	}
	log.Printf("Data received %v", registrationData)
	return c.JSON(string("Registers"))
}
