package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/patelajay745/ApplicationStation/config"
	"github.com/patelajay745/ApplicationStation/routes"

	"github.com/patelajay745/ApplicationStation/models"
)

func main() {

	// Initialize database connection
	db := config.Connect()
	db.AutoMigrate(&models.User{})

	app := fiber.New(fiber.Config{
		Views: html.New("templates", ".html"),
	})

	routes.SetupRoutes(app, db)

	app.Static("/", "static")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start the Fiber app
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
