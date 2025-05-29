package routes

import (
	"belajar-gorm/service"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")

	// user.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{"message": "Anda di /user"})
	// })

	user.Get("/", service.GetAllUser)
}
