package handler

import (
	"strconv"

	"github.com/sinisaos/fiber-ent-admin/model"
	"github.com/sinisaos/fiber-ent-admin/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

type UParams struct {
	UserName  string `query:"username"`
	Email     string `query:"email"`
	CreatedAt string `query:"created_at"`
	Sort      string `query:"_sort"`
	Order     string `query:"_order"`
	Start     int    `query:"_start"`
	End       int    `query:"_end"`
}

// All users
func (h UserHandler) GetAllUsersHandler(c *fiber.Ctx) error {
	p := new(UParams)
	if err := c.QueryParser(p); err != nil {
		return err
	}

	users, err := h.UserService.GetAllUsers(p.Start, p.End, p.Sort, p.Order, p.UserName, p.Email, p.CreatedAt)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	// total number of filtered records for React Admin pagination
	if p.UserName != "" || p.Email != "" || p.CreatedAt != "" {
		c.Set("X-Total-Count", strconv.Itoa(len(users)))
	} else {
		// total number of records for React Admin pagination
		count, _ := h.UserService.CountUsers()
		c.Set("X-Total-Count", strconv.Itoa(count))
	}
	return c.Status(200).JSON(users)
}

// Single user
func (h UserHandler) GetUserHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := h.UserService.GetUser(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(user)
}

// New user
func (h UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	payload := new(model.NewUserInput)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Payload error",
			"data":    err,
		})

	}

	user, err := h.UserService.CreateUser(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "User already exists",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": user,
	})
}

// Update user
func (h UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	payload := new(model.UpdateUserInput)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Payload error",
			"data":    err,
		})
	}

	user, err := h.UserService.UpdateUser(id, payload)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}
	return c.Status(200).JSON(user)

}

// Delete user
func (h UserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	// Check if the record exists
	err := h.UserService.DeleteUser(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(204).JSON(fiber.Map{
		"message": "User successfully deleted",
	})
}

// User questions
func (h UserHandler) GetUserQuestionsHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := h.UserService.GetUserQuestions(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": user,
	})
}

// User answers
func (h UserHandler) GetUserAnswersHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user, err := h.UserService.GetUserAnswers(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": user,
	})
}
