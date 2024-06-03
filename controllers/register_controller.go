package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/patelajay745/ApplicationStation/models"
	"github.com/patelajay745/ApplicationStation/utils"
	"gorm.io/gorm"
)

func RegisterPutHandler(c *fiber.Ctx, db *gorm.DB) error {
	var newUser models.User

	if err := c.BodyParser(&newUser); err != nil {
		return err
	}

	var exisitingUser models.User
	if err := db.Where("email=?", newUser.Email).First(&exisitingUser).Error; err == nil {
		return c.Redirect("/register?error=email_exists")
	}

	newUser.Password, _ = utils.HashPassword(newUser.Password)

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.Redirect("/login?success=true", fiber.StatusSeeOther)
}

func LoginPutHandler(c *fiber.Ctx, db *gorm.DB) error {
	var userInput models.User
	var existingUser models.User

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse request body")
	}

	// Find user by email
	if err := db.Where("email = ?", userInput.Email).First(&existingUser).Error; err != nil {
		// Redirect to login with error message if user not found
		fmt.Println("email not found")
		return c.Redirect("/login?error=wrong_credentials", fiber.StatusSeeOther)

	}

	// Check if the password matches
	if !utils.CheckPasswordHash(userInput.Password, existingUser.Password) {
		// Redirect to login with error message if password is incorrect
		fmt.Println("password not match")
		return c.Redirect("/login?error=wrong_credentials", fiber.StatusSeeOther)

	}

	return c.Redirect("/dashboard", fiber.StatusSeeOther)
}
