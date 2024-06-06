package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/patelajay745/ApplicationStation/config"
	"github.com/patelajay745/ApplicationStation/routes"

	"github.com/patelajay745/ApplicationStation/models"
)

var sessionManager *scs.SessionManager

func main() {

	// Initialize database connection
	db := config.Connect()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Application{})

	// Initialize session manager
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = true

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
