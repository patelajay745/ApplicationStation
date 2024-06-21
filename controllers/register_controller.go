package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/patelajay745/ApplicationStation/models"
	"github.com/patelajay745/ApplicationStation/utils"
	"gorm.io/gorm"
)

func RegisterPutHandler(c fiber.Ctx, db *gorm.DB) error {
	var newUser models.User

	if err := c.Bind().Body(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
	}

	var exisitingUser models.User
	if err := db.Where("email=?", newUser.Email).First(&exisitingUser).Error; err == nil {
		return c.Redirect().To("/register?error=email_exists")
	}

	newUser.Password, _ = utils.HashPassword(newUser.Password)
	newUser.Status = "notActive"

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.Redirect().To("/login?status=registered")
}

func LoginPutHandler(c fiber.Ctx, db *gorm.DB, store *session.Store) error {
	var userInput models.User
	var existingUser models.User

	if err := c.Bind().Body(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
	}

	if err := db.Where("email = ?", userInput.Email).First(&existingUser).Error; err != nil {
		return c.Redirect().To("/login?status=wrong_email")
	}

	if !utils.CheckPasswordHash(userInput.Password, existingUser.Password) {
		return c.Redirect().To("/login?status=wrong_password")

	}

	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get session")
	}

	sess.Set("authenticated", true)
	sess.Set("userID", existingUser.ID)

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save session")
	}

	fmt.Println("User authenticated successfully")

	return c.Redirect().To("/dashboard")
}

func LogoutHandler(c fiber.Ctx, store *session.Store) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get session")
	}

	// Destroy the session
	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to destroy session")
	}

	// Redirect to login page
	return c.Redirect().To("/")
}
