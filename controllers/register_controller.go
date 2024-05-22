package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/patelajay745/ApplicationStation/models"
	"gorm.io/gorm"
)

func RegisterPutHandler(c *fiber.Ctx, db *gorm.DB) error {
	var newUser models.User

	if err := c.BodyParser(&newUser); err != nil {
		return err
	}

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.Redirect("/login")
}
