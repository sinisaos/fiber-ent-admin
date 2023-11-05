package middleware

import (
	"github.com/sinisaos/fiber-ent-admin/database"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// Protect specific routes
func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(database.Config("SECRET_KEY"))},
		ErrorHandler: JWTError,
	})
}

func JWTError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{
				"message": "Missing or malformed JWT",
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{
			"message": "Invalid or expired JWT",
		})
}
