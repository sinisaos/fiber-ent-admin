package handler

import (
	"strconv"

	"github.com/sinisaos/fiber-ent-admin/pkg/model"
	"github.com/sinisaos/fiber-ent-admin/pkg/service"

	"github.com/gofiber/fiber/v2"
)

type QuestionHandler struct {
	QuestionService service.QuestionService
}

func NewQuestionHandler(service service.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		QuestionService: service,
	}
}

type QParams struct {
	Title   string `query:"title"`
	Content string `query:"content"`
	Sort    string `query:"_sort"`
	Order   string `query:"_order"`
	Start   int    `query:"_start"`
	End     int    `query:"_end"`
}

// All Questions
func (h QuestionHandler) GetAllQuestionsHandler(c *fiber.Ctx) error {
	p := new(QParams)
	if err := c.QueryParser(p); err != nil {
		return err
	}

	questions, err := h.QuestionService.GetAllQuestions(p.Start, p.End, p.Sort, p.Order, p.Title, p.Content)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// total number of filtered records for React Admin pagination
	if p.Title != "" || p.Content != "" {
		c.Set("X-Total-Count", strconv.Itoa(len(questions)))
	} else {
		// total number of records for React Admin pagination
		count, _ := h.QuestionService.CountQuestions()
		c.Set("X-Total-Count", strconv.Itoa(count))
	}
	return c.Status(200).JSON(questions)
}

// Single Question
func (h QuestionHandler) GetQuestionHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	question, err := h.QuestionService.GetQuestion(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(question)
}

// New Question
func (h QuestionHandler) CreateQuestionHandler(c *fiber.Ctx) error {
	payload := new(model.NewQuestionInput)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Payload error",
			"data":    err.Error(),
		})
	}

	question, err := h.QuestionService.CreateQuestion(payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"data": question,
	})
}

// Update Question
func (h QuestionHandler) UpdateQuestionHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	payload := new(model.UpdateQuestionInput)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Payload error",
			"data":    err.Error(),
		})
	}

	question, err := h.QuestionService.UpdateQuestion(id, payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Question successfully updated",
		"data":    question,
	})
}

// Delete Question
func (h QuestionHandler) DeleteQuestionHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	// Check if the record exists
	err := h.QuestionService.DeleteQuestion(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(204).JSON(fiber.Map{
		"message": "Question successfully deleted",
	})
}

// Question Answers
func (h QuestionHandler) GetQuestionAnswersHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	question, err := h.QuestionService.GetQuestionAnswers(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": question,
	})
}

// Question Author
func (h QuestionHandler) GetQuestionAuthorHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	question, err := h.QuestionService.GetQuestionAuthor(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": question,
	})
}

// Question Tags
func (h QuestionHandler) GetQuestionTagsHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	question, err := h.QuestionService.GetQuestionTags(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Record not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": question,
	})
}
