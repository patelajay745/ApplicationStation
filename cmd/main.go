package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"column:fullname"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func main() {

	connStr := "postgresql://postgres:0745@localhost/applicationTracker?sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	app := fiber.New(fiber.Config{
		Views: html.New("../templates", ".html"),
	})

	app.Static("/", "../static")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("layout/register", fiber.Map{})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		registerMessage := c.Locals("registerMessage").(string)
		fmt.Println("Login route accessed")
		fmt.Println("Register message:", registerMessage)
		return c.Render("layout/login", fiber.Map{"Message": registerMessage})
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return registerPutHandler(c, db)
	})

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func registerPutHandler(c *fiber.Ctx, db *gorm.DB) error {
	var newUser User

	if err := c.BodyParser(&newUser); err != nil {
		return nil
	}

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	// Set a success message to be displayed on the login page
	c.Locals("registerMessage", "Registration successful! You can now login.")

	fmt.Println("Register message set:", c.Locals("registerMessage"))
	return c.Redirect("/login")

}
