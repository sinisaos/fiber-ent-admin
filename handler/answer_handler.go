package handler

import (
	"strconv"

	"github.com/sinisaos/fiber-ent-admin/model"
	"github.com/sinisaos/fiber-ent-admin/service"

	"github.com/gofiber/fiber/v2"
)

type AnswerHandler struct {
	AnswerService service.AnswerService
}

func NewAnswerHandler(service service.AnswerService) *AnswerHandler {
	return &AnswerHandler{
		AnswerService: service,
	}
}

type AParams struct {
	Content string `query:"content"`
	Sort    string `query:"_sort"`
	Order   string `query:"_order"`
	Start   int    `query:"_start"`
	End     int    `query:"_end"`
}

// All Answers
func (h AnswerHandler) GetAllAnswersHandler(c *fiber.Ctx) error {
	p := new(AParams)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	// filtered result
	answers, err := h.AnswerService.GetAllAnswers(p.Start, p.End, p.Sort, p.Order, p.Content)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// total number of filtered records for React Admin pagination
	if p.Content != "" {
		c.Set("X-Total-Count", strconv.Itoa(len(answers)))
	} else {
		// total number of records for React Admin pagination
		count, _ := h.AnswerService.CountAnswers()
		c.Set("X-Total-Count", strconv.Itoa(count))
	}

	return c.Status(200).JSON(answers)
}

// Single Answer
func (h AnswerHandler) GetAnswerHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	answer, err := h.AnswerService.GetAnswer(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(answer)
}

// New Answer
func (h AnswerHandler) CreateAnswerHandler(c *fiber.Ctx) error {
	payload := new(model.NewAnswerInput)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Payload error",
			"data":    err.Error(),
		})
	}

	answer, err := h.AnswerService.CreateAnswer(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": answer,
	})
}

// Update Answer
func (h AnswerHandler) UpdateAnswerHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	payload := new(model.UpdateAnswerInput)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Payload error",
			"data":    err.Error(),
		})
	}

	answer, err := h.AnswerService.UpdateAnswer(id, payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Answer successfully updated",
		"data":    answer,
	})

}

// Delete Answer
func (h AnswerHandler) DeleteAnswerHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	// Check if the record exists
	err := h.AnswerService.DeleteAnswer(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(204).JSON(fiber.Map{
		"message": "Answer successfully deleted",
	})
}

// Answer Questions
func (h AnswerHandler) GetAnswerQuestionHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	answer, err := h.AnswerService.GetAnswerQuestion(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": answer,
	})
}

// Answer Author
func (h AnswerHandler) GetAnswerAuthorHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	answer, err := h.AnswerService.GetAnswerAuthor(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": answer,
	})
}
