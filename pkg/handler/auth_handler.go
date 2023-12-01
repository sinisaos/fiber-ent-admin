package handler

import (
	"time"

	"github.com/sinisaos/fiber-ent-admin/pkg/database"
	"github.com/sinisaos/fiber-ent-admin/pkg/model"
	"github.com/sinisaos/fiber-ent-admin/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: service,
	}
}

// Login Handler
func (h AuthHandler) LoginHandler(c *fiber.Ctx) error {
	payload := new(model.LoginUserInput)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error on login request",
			"data":    err,
		})
	}

	// Check username
	newUser, err := h.AuthService.Login(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Incorrect username or password",
		})
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(payload.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Incorrect username or password",
		})
	}

	// Create a token for the user with the correct username and password
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = newUser.Username
	claims["user_id"] = newUser.ID
	claims["superuser"] = newUser.Superuser
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	jwtToken, err := token.SignedString([]byte(database.Config("SECRET_KEY")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Successfully logged in",
		"token":   jwtToken,
	})
}
