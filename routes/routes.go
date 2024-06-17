package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/patelajay745/ApplicationStation/controllers"
	"gorm.io/gorm"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App, db *gorm.DB, store *session.Store) {
	app.Get("/register", func(c fiber.Ctx) error {
		errorMessage := c.Query(("error"))
		var errorMsg string
		if errorMessage == "email_exists" {
			errorMsg = "This email is already registered. Please sign in."
		}
		return c.Render("layout/register", fiber.Map{
			"ErrorMsg": errorMsg,
		})
	})

	app.Get("/login", func(c fiber.Ctx) error {

		message := c.Redirect().Message("status")

		if len(message) > 0 {
			fmt.Println("Sucees")
			return c.Render("layout/login", fiber.Map{
				"Success": message,
			})
		} else {
			fmt.Println("error")
			return c.Render("layout/login", fiber.Map{
				"Error": "An error occurred. Please try again.",
			})
		}
	})

	app.Post("/login", func(c fiber.Ctx) error {

		return controllers.LoginPutHandler(c, db, store)

	})

	app.Post("/register", func(c fiber.Ctx) error {
		return controllers.RegisterPutHandler(c, db)
	})

	app.Get("/logout", func(c fiber.Ctx) error {
		return controllers.LogoutHandler(c, store)
	})

	// Middleware to check if the user is logged in
	//app.Use("/", middleware.AuthRequired(store))

	app.Get("/", func(c fiber.Ctx) error {
		sess := c.Locals("session").(*session.Session)
		if sess.Get("authenticated") != true {
			return c.Redirect().To("/login")
		}
		return c.Redirect().To("/dashboard")
	})

	app.Get("/dashboard", func(c fiber.Ctx) error {
		// sess := c.Locals("session").(*session.Session)
		// if sess.Get("authenticated") != true {
		// 	return c.Redirect("/login")
		// }

		message := c.Redirect().Message("status")

		if len(message) > 0 {
			return c.Render("layout/dashboard", fiber.Map{"message": message}, "layout/main")
		}

		return c.Render("layout/dashboard", fiber.Map{}, "layout/main")
	})
	// Add Application form route
	app.Get("/add_application", func(c fiber.Ctx) error {
		return c.Render("layout/add_application", fiber.Map{
			"Title": "Add Application | Job Application Tracker",
		}, "layout/main")
	})

	app.Post("/add_application", func(c fiber.Ctx) error {

		return controllers.AddApplicationHandler(c, db, store)
	})
}
