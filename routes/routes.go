package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patelajay745/ApplicationStation/controllers"
	"gorm.io/gorm"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("layout/register", fiber.Map{})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("layout/login", fiber.Map{})
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.RegisterPutHandler(c, db)
	})
}
