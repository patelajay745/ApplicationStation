package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/patelajay745/ApplicationStation/config"
	"github.com/patelajay745/ApplicationStation/models"
	"github.com/patelajay745/ApplicationStation/routes"
)

var store *session.Store

func main() {

	// Initialize session store
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: "Lax",
	})

	// Initialize database connection
	db := config.Connect()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Application{})

	app := fiber.New(fiber.Config{
		Views: html.New("templates", ".html"),
	})

	// Middleware to manage session
	app.Use(func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return err
		}
		c.Locals("session", sess)
		return c.Next()
	})

	routes.SetupRoutes(app, db, store)

	app.Static("/", "static")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the Fiber app
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
