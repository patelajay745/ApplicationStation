package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

func AddApplicationHandler(c *fiber.Ctx, db *gorm.DB, store *session.Store) error {
	
	return c.Redirect("/dashboard", fiber.StatusSeeOther)
}
