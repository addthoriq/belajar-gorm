package main

import (
	"belajar-gorm/models"
	"belajar-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDB()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	routes.UserRoutes(app)
	app.Listen(":3000")
}
