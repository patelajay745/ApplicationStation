package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/patelajay745/ApplicationStation/controllers"
	"gorm.io/gorm"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App, db *gorm.DB, store *session.Store) {
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

		return controllers.LoginPutHandler(c, db, store)

	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return controllers.RegisterPutHandler(c, db)
	})

	app.Get("/logout", func(c *fiber.Ctx) error {
		return controllers.LogoutHandler(c, store)
	})

	// Middleware to check if the user is logged in
	//app.Use("/", middleware.AuthRequired(store))

	app.Get("/", func(c *fiber.Ctx) error {
		sess := c.Locals("session").(*session.Session)
		if sess.Get("authenticated") != true {
			return c.Redirect("/login")
		}
		return c.Redirect("/dashboard")
	})

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		// sess := c.Locals("session").(*session.Session)
		// if sess.Get("authenticated") != true {
		// 	return c.Redirect("/login")
		// }
		return c.Render("layout/dashboard", fiber.Map{}, "layout/main")
	})
	// Add Application form route
	app.Get("/add_application", func(c *fiber.Ctx) error {
		return c.Render("layout/add_application", fiber.Map{
			"Title": "Add Application | Job Application Tracker",
		}, "layout/main")
	})
}
