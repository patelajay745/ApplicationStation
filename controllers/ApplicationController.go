package controllers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/patelajay745/ApplicationStation/models"
	"gorm.io/gorm"
)

func AddApplicationHandler(c fiber.Ctx, db *gorm.DB, store *session.Store) error {

	var application models.Application

	if err := c.Bind().Body(&application); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(("Failed to parse request body"))
	}
	application.AppliedDate = time.Now()
	application.Status = "Pending"

	if err := db.Create(&application).Error; err != nil {
		return err
	}

	// return c.Redirect("/dashboard", fiber.StatusSeeOther)
	return c.Redirect().With("status","added").To("/dashboard")
}
