package main

import (
	"github.com/alokkumarv/Vyapaar/internal/registration"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName:           "Vyapaar",
		EnablePrintRoutes: true,
	})
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/register", registration.HandleRegistration).Name("ShopKeeper Registration")
	// auth.Post("/login")
	// auth.Post("/logout")
	// auth.Get("/profile")
	// auth.Put("/profile")
	// auth.Get("/users/:id")

	app.Listen(":3000")
}

// /register	Register a new user
// POST	/login	User login, returns JWT token
// POST	/logout	Logout user
// GET	/profile	Get current user profile
// PUT	/profile	Update user profile
// GET	/users/:id
