package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

func AuthRequired(store *session.Store) fiber.Handler {
	return func(c fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to get session")
		}

		// Check if user is authenticated
		if sess.Get("authenticated") != true {
			return c.Redirect().To("/login")
		}

		return c.Next()
	}
}
