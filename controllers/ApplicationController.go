package controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/patelajay745/ApplicationStation/models"
	"gorm.io/gorm"
)

func AddApplicationHandler(c *fiber.Ctx, db *gorm.DB, store *session.Store) error {

	// Log raw body
	log.Printf("Raw request body: %s\n", c.Body())

	var application models.Application

	if err := c.BodyParser(&application); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(("Failed to parse request body"))
	}

	fmt.Println(application.WorkType)
	application.AppliedDate = time.Now()
	application.Status = "Pending"

	if err := db.Create(&application).Error; err != nil {
		return err
	}

	return c.Redirect("/dashboard", fiber.StatusSeeOther)
}
