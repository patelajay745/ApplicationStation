package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("../templates", ".html"),
	})

	app.Static("/", "../static")

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("layout/register", fiber.Map{})
	})

	// Start the Fiber app
	err := app.Listen(":3005")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
