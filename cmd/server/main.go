package main

import (
	"github.com/alokkumarv/Vyapaar/internal/user"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName:           "Vyapaar",
		EnablePrintRoutes: true,
	})
	app.Post("/login", user.UserHandler)
	app.Listen(":3000")
}
