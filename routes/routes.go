package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patelajay745/ApplicationStation/controllers"
	"gorm.io/gorm"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/register", func(c *fiber.Ctx) error {
		errorMessage := c.Query(("error"))
		var errorMsg string
		if errorMessage == "email_exists" {
			errorMsg = "This email is already registered. Please sign in."
		}
		return c.Render("layout/register", fiber.Map{
			"ErrorMsg": errorMsg,
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		success := c.Query("success")
		error := c.Query("error")
		var errorMsg string
		if error == "wrong_credentials" {
			errorMsg = "Email or Password is wrong"
		}

		return c.Render("layout/login", fiber.Map{
			"Success": success, "error": errorMsg,
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {

		return controllers.LoginPutHandler(c, db)

	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.RegisterPutHandler(c, db)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/login")
	})

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Render("layout/dashboard", fiber.Map{})
	})
}
